// Code generated by MockGen. DO NOT EDIT.
// Source: registry.go

// Package mock_registry is a generated GoMock package.
package mock_registry

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-broker/interoperator/pkg/apis/resource/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockClusterRegistry is a mock of ClusterRegistry interface
type MockClusterRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockClusterRegistryMockRecorder
}

// MockClusterRegistryMockRecorder is the mock recorder for MockClusterRegistry
type MockClusterRegistryMockRecorder struct {
	mock *MockClusterRegistry
}

// NewMockClusterRegistry creates a new mock instance
func NewMockClusterRegistry(ctrl *gomock.Controller) *MockClusterRegistry {
	mock := &MockClusterRegistry{ctrl: ctrl}
	mock.recorder = &MockClusterRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterRegistry) EXPECT() *MockClusterRegistryMockRecorder {
	return m.recorder
}

// GetClient mocks base method
func (m *MockClusterRegistry) GetClient(clusterID string) (client.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", clusterID)
	ret0, _ := ret[0].(client.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient
func (mr *MockClusterRegistryMockRecorder) GetClient(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockClusterRegistry)(nil).GetClient), clusterID)
}

// GetCluster mocks base method
func (m *MockClusterRegistry) GetCluster(clusterID string) (v1alpha1.SFClusterInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCluster", clusterID)
	ret0, _ := ret[0].(v1alpha1.SFClusterInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockClusterRegistryMockRecorder) GetCluster(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterRegistry)(nil).GetCluster), clusterID)
}

// ListClusters mocks base method
func (m *MockClusterRegistry) ListClusters(options *client.ListOptions) (*v1alpha1.SFClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", options)
	ret0, _ := ret[0].(*v1alpha1.SFClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters
func (mr *MockClusterRegistryMockRecorder) ListClusters(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockClusterRegistry)(nil).ListClusters), options)
}
