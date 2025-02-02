/*
Copyright 2018 The Service Fabrik Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sfservicesreplicator

import (
	"context"

	osbv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/apis/osb/v1alpha1"
	resourcev1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/apis/resource/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/cluster/registry"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kubernetes "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("service.replicator")

// Add creates a new SFServicesReplicator Controller and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	clusterRegistry, err := registry.New(mgr.GetConfig(), mgr.GetScheme(), mgr.GetRESTMapper())
	if err != nil {
		return err
	}

	return add(mgr, newReconciler(mgr, clusterRegistry))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager, clusterRegistry registry.ClusterRegistry) reconcile.Reconciler {
	return &ReconcileSFServices{
		Client:          mgr.GetClient(),
		scheme:          mgr.GetScheme(),
		clusterRegistry: clusterRegistry,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("sfservices-mcd-controller", mgr, controller.Options{
		Reconciler:              r,
		MaxConcurrentReconciles: 1,
	})
	if err != nil {
		return err
	}

	// Watch for changes to SFCluster
	err = c.Watch(&source.Kind{Type: &resourcev1alpha1.SFCluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	clusterRegistry, err := registry.New(mgr.GetConfig(), mgr.GetScheme(), mgr.GetRESTMapper())
	if err != nil {
		return err
	}

	// Define a mapping from the object in the event(sfservice/sfplan) to
	// list of sfclusters to reconcile
	mapFn := handler.ToRequestsFunc(
		func(a handler.MapObject) []reconcile.Request {
			return enqueueRequestForAllClusters(clusterRegistry)
		})

	err = c.Watch(&source.Kind{Type: &osbv1alpha1.SFService{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests: mapFn,
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &osbv1alpha1.SFPlan{}}, &handler.EnqueueRequestsFromMapFunc{
		ToRequests: mapFn,
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileSFServices{}

// ReconcileSFServices reconciles SFServices state across clusters
type ReconcileSFServices struct {
	client.Client
	scheme          *runtime.Scheme
	clusterRegistry registry.ClusterRegistry
}

// Reconcile is called for a SFCluster. It replicates all SFServices and all SFPlans to
// the SFCluster
func (r *ReconcileSFServices) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// Fetch the SFCluster
	clusterInstance := &resourcev1alpha1.SFCluster{}
	err := r.Get(context.TODO(), request.NamespacedName, clusterInstance)
	if err != nil {
		if apiErrors.IsNotFound(err) {
			// Object not found, return.
			return reconcile.Result{}, nil
		}
		log.Error(err, "Failed to get SFCluster")
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	clusterID := clusterInstance.GetName()
	log.Info("Reconcile started for cluster", "clusterID", clusterID)
	targetClient, err := r.clusterRegistry.GetClient(clusterID)
	if err != nil {
		log.Error(err, "Following error occurred while getting client for cluster ", "clusterID", clusterID)
		return reconcile.Result{}, err
	}

	log.Info("Trying to list all the services", "namespace", request.NamespacedName.Namespace)
	options := kubernetes.InNamespace(request.NamespacedName.Namespace)
	services := &osbv1alpha1.SFServiceList{}
	err = r.List(context.TODO(), options, services)
	if err != nil {
		log.Error(err, "error while fetching services while processing cluster id ", "clusterID", clusterID)
		return reconcile.Result{}, err
	}
	log.Info("services fetched ", "count", len(services.Items), "clusterID", clusterID)
	for _, obj := range services.Items {
		log.Info("Service is fetched from master cluster", "serviceID", obj.Spec.ID)
		service := &osbv1alpha1.SFService{}
		serviceKey := types.NamespacedName{
			Name:      obj.GetName(),
			Namespace: obj.GetNamespace(),
		}
		log.Info("Checking if service already exists on target cluster", "serviceID", obj.Spec.ID, "clusterID", clusterID)
		err = targetClient.Get(context.TODO(), serviceKey, service)
		if err != nil {
			if apiErrors.IsNotFound(err) {
				replicateSFServiceResourceData(&obj, service)
				err = targetClient.Create(context.TODO(), service)
				if err != nil {
					log.Error(err, "Creating new service on sister cluster failed due to following error: ")
					return reconcile.Result{}, err
				}
				log.Info("Created service on cluster", "serviceName", service.Spec.Name, "clusterID", clusterID)
				err := r.handleServicePlans(service, clusterID, &targetClient)
				if err != nil {
					log.Error(err, "Error while replicating plans for service ", "serviceName", service.Spec.Name)
					return reconcile.Result{}, err
				}
			} else {
				log.Error(err, "Getting the service from sister cluster ", "clusterID", clusterID)
				return reconcile.Result{}, err
			}
		} else {
			replicateSFServiceResourceData(&obj, service)
			err = targetClient.Update(context.TODO(), service)
			if err != nil {
				log.Error(err, "Updating service on sister cluster failed due to following error: ")
				return reconcile.Result{}, err
			}
			log.Info("Updated service on cluster", "serviceName", service.Spec.Name, "clusterID", clusterID)
			err = r.handleServicePlans(service, clusterID, &targetClient)
			if err != nil {
				log.Error(err, "Error while replicating plans for service ", "serviceName", service.Spec.Name)
				return reconcile.Result{}, err
			}
		}
	}
	return reconcile.Result{}, nil
}

func (r *ReconcileSFServices) handleServicePlans(service *osbv1alpha1.SFService, clusterID string, targetClient *kubernetes.Client) error {
	log.Info("Trying  to list all the plans for service in the master cluster", "serviceName", service.Spec.Name)
	plans := &osbv1alpha1.SFPlanList{}
	searchLabels := make(map[string]string)
	searchLabels["serviceId"] = service.Spec.ID
	options := kubernetes.MatchingLabels(searchLabels)
	options.Namespace = service.GetNamespace()
	err := r.List(context.TODO(), options, plans)
	if err != nil {
		log.Error(err, "error while fetching plans while processing cluster id ", "clusterID", clusterID)
		return err
	}
	log.Info("plans fetched for cluster", "count", len(plans.Items), "clusterID", clusterID)
	for _, obj := range plans.Items {
		log.Info("Plan is fetched from master cluster", "planID", obj.Spec.ID)
		plan := &osbv1alpha1.SFPlan{}
		planKey := types.NamespacedName{
			Name:      obj.GetName(),
			Namespace: obj.GetNamespace(),
		}
		log.Info("Checking if plan already exists on target cluster", "clusterID", clusterID, "planID", obj.Spec.ID)
		err = (*targetClient).Get(context.TODO(), planKey, plan)
		if err != nil {
			if apiErrors.IsNotFound(err) {
				replicateSFPlanResourceData(&obj, plan)
				err = controllerutil.SetControllerReference(service, plan, r.scheme)
				if err != nil {
					return err
				}
				err = (*targetClient).Create(context.TODO(), plan)
				if err != nil {
					log.Error(err, "Creating new plan on sister cluster failed")
					return err
				}
				log.Info("Created plan on cluster", "clusterID", clusterID, "planName", plan.Spec.Name)
			}
		} else {
			replicateSFPlanResourceData(&obj, plan)
			err = controllerutil.SetControllerReference(service, plan, r.scheme)
			if err != nil {
				return err
			}
			err = (*targetClient).Update(context.TODO(), plan)
			if err != nil {
				log.Error(err, "Updating plan on sister cluster failed")
				return err
			}
			log.Info("Updated plan on cluster ", "clusterID", clusterID, "planName", plan.Spec.Name)
		}
	}
	return nil
}

func enqueueRequestForAllClusters(clusterRegistry registry.ClusterRegistry) []reconcile.Request {
	clusterList, err := clusterRegistry.ListClusters(nil)
	if err != nil {
		return nil
	}
	reconcileRequests := make([]reconcile.Request, len(clusterList.Items))
	for i, cluster := range clusterList.Items {
		reconcileRequests[i] = reconcile.Request{
			NamespacedName: types.NamespacedName{
				Name:      cluster.GetName(),
				Namespace: cluster.GetNamespace(),
			},
		}
	}
	return reconcileRequests
}

func replicateSFServiceResourceData(source *osbv1alpha1.SFService, dest *osbv1alpha1.SFService) {
	source.Spec.DeepCopyInto(&dest.Spec)
	dest.SetName(source.GetName())
	dest.SetNamespace(source.GetNamespace())
	dest.SetLabels(source.GetLabels())
}

func replicateSFPlanResourceData(source *osbv1alpha1.SFPlan, dest *osbv1alpha1.SFPlan) {
	source.Spec.DeepCopyInto(&dest.Spec)
	dest.SetName(source.GetName())
	dest.SetNamespace(source.GetNamespace())
	dest.SetLabels(source.GetLabels())
}
