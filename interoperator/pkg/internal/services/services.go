package services

import (
	"context"

	osbv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/apis/osb/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/errors"
	kubernetes "sigs.k8s.io/controller-runtime/pkg/client"
)

// FindServiceInfo fetches the details of a service
// from the services path
func FindServiceInfo(client kubernetes.Client, serviceID string, planID string, namespace string) (*osbv1alpha1.SFService, *osbv1alpha1.SFPlan, error) {
	services := &osbv1alpha1.SFServiceList{}
	labels := make(map[string]string)
	labels["serviceId"] = serviceID
	options := kubernetes.MatchingLabels(labels)
	options.Namespace = namespace

	err := client.List(context.TODO(), options, services)
	if err != nil {
		return nil, nil, err
	}
	var service *osbv1alpha1.SFService
	for _, obj := range services.Items {
		if obj.Spec.ID == serviceID {
			service = &obj
		}
	}
	if service == nil {
		return nil, nil, errors.NewSFServiceNotFound(serviceID, nil)
	}

	plans := &osbv1alpha1.SFPlanList{}
	labels = make(map[string]string)
	labels["serviceId"] = serviceID
	labels["planId"] = planID
	options = kubernetes.MatchingLabels(labels)
	options.Namespace = namespace

	err = client.List(context.TODO(), options, plans)
	if err != nil {
		return nil, nil, err
	}

	for _, plan := range plans.Items {
		if plan.Spec.ID == planID {
			return service, &plan, nil
		}
	}
	return nil, nil, errors.NewSFPlanNotFound(planID, nil)
}
