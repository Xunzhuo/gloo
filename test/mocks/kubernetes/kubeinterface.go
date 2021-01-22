// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes (interfaces: Interface)

// Package mock_kubernetes is a generated GoMock package.
package mock_kubernetes

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	discovery "k8s.io/client-go/discovery"
	v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	v1beta1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1beta1"
	v10 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v1beta10 "k8s.io/client-go/kubernetes/typed/apps/v1beta1"
	v1beta2 "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	v11 "k8s.io/client-go/kubernetes/typed/authentication/v1"
	v1beta11 "k8s.io/client-go/kubernetes/typed/authentication/v1beta1"
	v12 "k8s.io/client-go/kubernetes/typed/authorization/v1"
	v1beta12 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	v13 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	v2beta1 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta1"
	v2beta2 "k8s.io/client-go/kubernetes/typed/autoscaling/v2beta2"
	v14 "k8s.io/client-go/kubernetes/typed/batch/v1"
	v1beta13 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	v2alpha1 "k8s.io/client-go/kubernetes/typed/batch/v2alpha1"
	v15 "k8s.io/client-go/kubernetes/typed/certificates/v1"
	v1beta14 "k8s.io/client-go/kubernetes/typed/certificates/v1beta1"
	v16 "k8s.io/client-go/kubernetes/typed/coordination/v1"
	v1beta15 "k8s.io/client-go/kubernetes/typed/coordination/v1beta1"
	v17 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1alpha1 "k8s.io/client-go/kubernetes/typed/discovery/v1alpha1"
	v1beta16 "k8s.io/client-go/kubernetes/typed/discovery/v1beta1"
	v18 "k8s.io/client-go/kubernetes/typed/events/v1"
	v1beta17 "k8s.io/client-go/kubernetes/typed/events/v1beta1"
	v1beta18 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	v1alpha10 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1alpha1"
	v19 "k8s.io/client-go/kubernetes/typed/networking/v1"
	v1beta19 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	v1alpha11 "k8s.io/client-go/kubernetes/typed/node/v1alpha1"
	v1beta110 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	v1beta111 "k8s.io/client-go/kubernetes/typed/policy/v1beta1"
	v110 "k8s.io/client-go/kubernetes/typed/rbac/v1"
	v1alpha12 "k8s.io/client-go/kubernetes/typed/rbac/v1alpha1"
	v1beta112 "k8s.io/client-go/kubernetes/typed/rbac/v1beta1"
	v111 "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	v1alpha13 "k8s.io/client-go/kubernetes/typed/scheduling/v1alpha1"
	v1beta113 "k8s.io/client-go/kubernetes/typed/scheduling/v1beta1"
	v1alpha14 "k8s.io/client-go/kubernetes/typed/settings/v1alpha1"
	v112 "k8s.io/client-go/kubernetes/typed/storage/v1"
	v1alpha15 "k8s.io/client-go/kubernetes/typed/storage/v1alpha1"
	v1beta114 "k8s.io/client-go/kubernetes/typed/storage/v1beta1"
)

// MockInterface is a mock of Interface interface
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// AdmissionregistrationV1 mocks base method
func (m *MockInterface) AdmissionregistrationV1() v1.AdmissionregistrationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdmissionregistrationV1")
	ret0, _ := ret[0].(v1.AdmissionregistrationV1Interface)
	return ret0
}

// AdmissionregistrationV1 indicates an expected call of AdmissionregistrationV1
func (mr *MockInterfaceMockRecorder) AdmissionregistrationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdmissionregistrationV1", reflect.TypeOf((*MockInterface)(nil).AdmissionregistrationV1))
}

// AdmissionregistrationV1beta1 mocks base method
func (m *MockInterface) AdmissionregistrationV1beta1() v1beta1.AdmissionregistrationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdmissionregistrationV1beta1")
	ret0, _ := ret[0].(v1beta1.AdmissionregistrationV1beta1Interface)
	return ret0
}

// AdmissionregistrationV1beta1 indicates an expected call of AdmissionregistrationV1beta1
func (mr *MockInterfaceMockRecorder) AdmissionregistrationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdmissionregistrationV1beta1", reflect.TypeOf((*MockInterface)(nil).AdmissionregistrationV1beta1))
}

// AppsV1 mocks base method
func (m *MockInterface) AppsV1() v10.AppsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1")
	ret0, _ := ret[0].(v10.AppsV1Interface)
	return ret0
}

// AppsV1 indicates an expected call of AppsV1
func (mr *MockInterfaceMockRecorder) AppsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1", reflect.TypeOf((*MockInterface)(nil).AppsV1))
}

// AppsV1beta1 mocks base method
func (m *MockInterface) AppsV1beta1() v1beta10.AppsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1beta1")
	ret0, _ := ret[0].(v1beta10.AppsV1beta1Interface)
	return ret0
}

// AppsV1beta1 indicates an expected call of AppsV1beta1
func (mr *MockInterfaceMockRecorder) AppsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1beta1", reflect.TypeOf((*MockInterface)(nil).AppsV1beta1))
}

// AppsV1beta2 mocks base method
func (m *MockInterface) AppsV1beta2() v1beta2.AppsV1beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppsV1beta2")
	ret0, _ := ret[0].(v1beta2.AppsV1beta2Interface)
	return ret0
}

// AppsV1beta2 indicates an expected call of AppsV1beta2
func (mr *MockInterfaceMockRecorder) AppsV1beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppsV1beta2", reflect.TypeOf((*MockInterface)(nil).AppsV1beta2))
}

// AuthenticationV1 mocks base method
func (m *MockInterface) AuthenticationV1() v11.AuthenticationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationV1")
	ret0, _ := ret[0].(v11.AuthenticationV1Interface)
	return ret0
}

// AuthenticationV1 indicates an expected call of AuthenticationV1
func (mr *MockInterfaceMockRecorder) AuthenticationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationV1", reflect.TypeOf((*MockInterface)(nil).AuthenticationV1))
}

// AuthenticationV1beta1 mocks base method
func (m *MockInterface) AuthenticationV1beta1() v1beta11.AuthenticationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthenticationV1beta1")
	ret0, _ := ret[0].(v1beta11.AuthenticationV1beta1Interface)
	return ret0
}

// AuthenticationV1beta1 indicates an expected call of AuthenticationV1beta1
func (mr *MockInterfaceMockRecorder) AuthenticationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthenticationV1beta1", reflect.TypeOf((*MockInterface)(nil).AuthenticationV1beta1))
}

// AuthorizationV1 mocks base method
func (m *MockInterface) AuthorizationV1() v12.AuthorizationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationV1")
	ret0, _ := ret[0].(v12.AuthorizationV1Interface)
	return ret0
}

// AuthorizationV1 indicates an expected call of AuthorizationV1
func (mr *MockInterfaceMockRecorder) AuthorizationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationV1", reflect.TypeOf((*MockInterface)(nil).AuthorizationV1))
}

// AuthorizationV1beta1 mocks base method
func (m *MockInterface) AuthorizationV1beta1() v1beta12.AuthorizationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationV1beta1")
	ret0, _ := ret[0].(v1beta12.AuthorizationV1beta1Interface)
	return ret0
}

// AuthorizationV1beta1 indicates an expected call of AuthorizationV1beta1
func (mr *MockInterfaceMockRecorder) AuthorizationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationV1beta1", reflect.TypeOf((*MockInterface)(nil).AuthorizationV1beta1))
}

// AutoscalingV1 mocks base method
func (m *MockInterface) AutoscalingV1() v13.AutoscalingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV1")
	ret0, _ := ret[0].(v13.AutoscalingV1Interface)
	return ret0
}

// AutoscalingV1 indicates an expected call of AutoscalingV1
func (mr *MockInterfaceMockRecorder) AutoscalingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV1", reflect.TypeOf((*MockInterface)(nil).AutoscalingV1))
}

// AutoscalingV2beta1 mocks base method
func (m *MockInterface) AutoscalingV2beta1() v2beta1.AutoscalingV2beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2beta1")
	ret0, _ := ret[0].(v2beta1.AutoscalingV2beta1Interface)
	return ret0
}

// AutoscalingV2beta1 indicates an expected call of AutoscalingV2beta1
func (mr *MockInterfaceMockRecorder) AutoscalingV2beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2beta1", reflect.TypeOf((*MockInterface)(nil).AutoscalingV2beta1))
}

// AutoscalingV2beta2 mocks base method
func (m *MockInterface) AutoscalingV2beta2() v2beta2.AutoscalingV2beta2Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoscalingV2beta2")
	ret0, _ := ret[0].(v2beta2.AutoscalingV2beta2Interface)
	return ret0
}

// AutoscalingV2beta2 indicates an expected call of AutoscalingV2beta2
func (mr *MockInterfaceMockRecorder) AutoscalingV2beta2() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoscalingV2beta2", reflect.TypeOf((*MockInterface)(nil).AutoscalingV2beta2))
}

// BatchV1 mocks base method
func (m *MockInterface) BatchV1() v14.BatchV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV1")
	ret0, _ := ret[0].(v14.BatchV1Interface)
	return ret0
}

// BatchV1 indicates an expected call of BatchV1
func (mr *MockInterfaceMockRecorder) BatchV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV1", reflect.TypeOf((*MockInterface)(nil).BatchV1))
}

// BatchV1beta1 mocks base method
func (m *MockInterface) BatchV1beta1() v1beta13.BatchV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV1beta1")
	ret0, _ := ret[0].(v1beta13.BatchV1beta1Interface)
	return ret0
}

// BatchV1beta1 indicates an expected call of BatchV1beta1
func (mr *MockInterfaceMockRecorder) BatchV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV1beta1", reflect.TypeOf((*MockInterface)(nil).BatchV1beta1))
}

// BatchV2alpha1 mocks base method
func (m *MockInterface) BatchV2alpha1() v2alpha1.BatchV2alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchV2alpha1")
	ret0, _ := ret[0].(v2alpha1.BatchV2alpha1Interface)
	return ret0
}

// BatchV2alpha1 indicates an expected call of BatchV2alpha1
func (mr *MockInterfaceMockRecorder) BatchV2alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchV2alpha1", reflect.TypeOf((*MockInterface)(nil).BatchV2alpha1))
}

// CertificatesV1 mocks base method
func (m *MockInterface) CertificatesV1() v15.CertificatesV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificatesV1")
	ret0, _ := ret[0].(v15.CertificatesV1Interface)
	return ret0
}

// CertificatesV1 indicates an expected call of CertificatesV1
func (mr *MockInterfaceMockRecorder) CertificatesV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificatesV1", reflect.TypeOf((*MockInterface)(nil).CertificatesV1))
}

// CertificatesV1beta1 mocks base method
func (m *MockInterface) CertificatesV1beta1() v1beta14.CertificatesV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CertificatesV1beta1")
	ret0, _ := ret[0].(v1beta14.CertificatesV1beta1Interface)
	return ret0
}

// CertificatesV1beta1 indicates an expected call of CertificatesV1beta1
func (mr *MockInterfaceMockRecorder) CertificatesV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CertificatesV1beta1", reflect.TypeOf((*MockInterface)(nil).CertificatesV1beta1))
}

// CoordinationV1 mocks base method
func (m *MockInterface) CoordinationV1() v16.CoordinationV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoordinationV1")
	ret0, _ := ret[0].(v16.CoordinationV1Interface)
	return ret0
}

// CoordinationV1 indicates an expected call of CoordinationV1
func (mr *MockInterfaceMockRecorder) CoordinationV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoordinationV1", reflect.TypeOf((*MockInterface)(nil).CoordinationV1))
}

// CoordinationV1beta1 mocks base method
func (m *MockInterface) CoordinationV1beta1() v1beta15.CoordinationV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoordinationV1beta1")
	ret0, _ := ret[0].(v1beta15.CoordinationV1beta1Interface)
	return ret0
}

// CoordinationV1beta1 indicates an expected call of CoordinationV1beta1
func (mr *MockInterfaceMockRecorder) CoordinationV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoordinationV1beta1", reflect.TypeOf((*MockInterface)(nil).CoordinationV1beta1))
}

// CoreV1 mocks base method
func (m *MockInterface) CoreV1() v17.CoreV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CoreV1")
	ret0, _ := ret[0].(v17.CoreV1Interface)
	return ret0
}

// CoreV1 indicates an expected call of CoreV1
func (mr *MockInterfaceMockRecorder) CoreV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CoreV1", reflect.TypeOf((*MockInterface)(nil).CoreV1))
}

// Discovery mocks base method
func (m *MockInterface) Discovery() discovery.DiscoveryInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Discovery")
	ret0, _ := ret[0].(discovery.DiscoveryInterface)
	return ret0
}

// Discovery indicates an expected call of Discovery
func (mr *MockInterfaceMockRecorder) Discovery() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discovery", reflect.TypeOf((*MockInterface)(nil).Discovery))
}

// DiscoveryV1alpha1 mocks base method
func (m *MockInterface) DiscoveryV1alpha1() v1alpha1.DiscoveryV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoveryV1alpha1")
	ret0, _ := ret[0].(v1alpha1.DiscoveryV1alpha1Interface)
	return ret0
}

// DiscoveryV1alpha1 indicates an expected call of DiscoveryV1alpha1
func (mr *MockInterfaceMockRecorder) DiscoveryV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryV1alpha1", reflect.TypeOf((*MockInterface)(nil).DiscoveryV1alpha1))
}

// DiscoveryV1beta1 mocks base method
func (m *MockInterface) DiscoveryV1beta1() v1beta16.DiscoveryV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscoveryV1beta1")
	ret0, _ := ret[0].(v1beta16.DiscoveryV1beta1Interface)
	return ret0
}

// DiscoveryV1beta1 indicates an expected call of DiscoveryV1beta1
func (mr *MockInterfaceMockRecorder) DiscoveryV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscoveryV1beta1", reflect.TypeOf((*MockInterface)(nil).DiscoveryV1beta1))
}

// EventsV1 mocks base method
func (m *MockInterface) EventsV1() v18.EventsV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsV1")
	ret0, _ := ret[0].(v18.EventsV1Interface)
	return ret0
}

// EventsV1 indicates an expected call of EventsV1
func (mr *MockInterfaceMockRecorder) EventsV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsV1", reflect.TypeOf((*MockInterface)(nil).EventsV1))
}

// EventsV1beta1 mocks base method
func (m *MockInterface) EventsV1beta1() v1beta17.EventsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsV1beta1")
	ret0, _ := ret[0].(v1beta17.EventsV1beta1Interface)
	return ret0
}

// EventsV1beta1 indicates an expected call of EventsV1beta1
func (mr *MockInterfaceMockRecorder) EventsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsV1beta1", reflect.TypeOf((*MockInterface)(nil).EventsV1beta1))
}

// ExtensionsV1beta1 mocks base method
func (m *MockInterface) ExtensionsV1beta1() v1beta18.ExtensionsV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtensionsV1beta1")
	ret0, _ := ret[0].(v1beta18.ExtensionsV1beta1Interface)
	return ret0
}

// ExtensionsV1beta1 indicates an expected call of ExtensionsV1beta1
func (mr *MockInterfaceMockRecorder) ExtensionsV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionsV1beta1", reflect.TypeOf((*MockInterface)(nil).ExtensionsV1beta1))
}

// FlowcontrolV1alpha1 mocks base method
func (m *MockInterface) FlowcontrolV1alpha1() v1alpha10.FlowcontrolV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowcontrolV1alpha1")
	ret0, _ := ret[0].(v1alpha10.FlowcontrolV1alpha1Interface)
	return ret0
}

// FlowcontrolV1alpha1 indicates an expected call of FlowcontrolV1alpha1
func (mr *MockInterfaceMockRecorder) FlowcontrolV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowcontrolV1alpha1", reflect.TypeOf((*MockInterface)(nil).FlowcontrolV1alpha1))
}

// NetworkingV1 mocks base method
func (m *MockInterface) NetworkingV1() v19.NetworkingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkingV1")
	ret0, _ := ret[0].(v19.NetworkingV1Interface)
	return ret0
}

// NetworkingV1 indicates an expected call of NetworkingV1
func (mr *MockInterfaceMockRecorder) NetworkingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkingV1", reflect.TypeOf((*MockInterface)(nil).NetworkingV1))
}

// NetworkingV1beta1 mocks base method
func (m *MockInterface) NetworkingV1beta1() v1beta19.NetworkingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkingV1beta1")
	ret0, _ := ret[0].(v1beta19.NetworkingV1beta1Interface)
	return ret0
}

// NetworkingV1beta1 indicates an expected call of NetworkingV1beta1
func (mr *MockInterfaceMockRecorder) NetworkingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkingV1beta1", reflect.TypeOf((*MockInterface)(nil).NetworkingV1beta1))
}

// NodeV1alpha1 mocks base method
func (m *MockInterface) NodeV1alpha1() v1alpha11.NodeV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeV1alpha1")
	ret0, _ := ret[0].(v1alpha11.NodeV1alpha1Interface)
	return ret0
}

// NodeV1alpha1 indicates an expected call of NodeV1alpha1
func (mr *MockInterfaceMockRecorder) NodeV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeV1alpha1", reflect.TypeOf((*MockInterface)(nil).NodeV1alpha1))
}

// NodeV1beta1 mocks base method
func (m *MockInterface) NodeV1beta1() v1beta110.NodeV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeV1beta1")
	ret0, _ := ret[0].(v1beta110.NodeV1beta1Interface)
	return ret0
}

// NodeV1beta1 indicates an expected call of NodeV1beta1
func (mr *MockInterfaceMockRecorder) NodeV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeV1beta1", reflect.TypeOf((*MockInterface)(nil).NodeV1beta1))
}

// PolicyV1beta1 mocks base method
func (m *MockInterface) PolicyV1beta1() v1beta111.PolicyV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PolicyV1beta1")
	ret0, _ := ret[0].(v1beta111.PolicyV1beta1Interface)
	return ret0
}

// PolicyV1beta1 indicates an expected call of PolicyV1beta1
func (mr *MockInterfaceMockRecorder) PolicyV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PolicyV1beta1", reflect.TypeOf((*MockInterface)(nil).PolicyV1beta1))
}

// RbacV1 mocks base method
func (m *MockInterface) RbacV1() v110.RbacV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1")
	ret0, _ := ret[0].(v110.RbacV1Interface)
	return ret0
}

// RbacV1 indicates an expected call of RbacV1
func (mr *MockInterfaceMockRecorder) RbacV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1", reflect.TypeOf((*MockInterface)(nil).RbacV1))
}

// RbacV1alpha1 mocks base method
func (m *MockInterface) RbacV1alpha1() v1alpha12.RbacV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1alpha1")
	ret0, _ := ret[0].(v1alpha12.RbacV1alpha1Interface)
	return ret0
}

// RbacV1alpha1 indicates an expected call of RbacV1alpha1
func (mr *MockInterfaceMockRecorder) RbacV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1alpha1", reflect.TypeOf((*MockInterface)(nil).RbacV1alpha1))
}

// RbacV1beta1 mocks base method
func (m *MockInterface) RbacV1beta1() v1beta112.RbacV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RbacV1beta1")
	ret0, _ := ret[0].(v1beta112.RbacV1beta1Interface)
	return ret0
}

// RbacV1beta1 indicates an expected call of RbacV1beta1
func (mr *MockInterfaceMockRecorder) RbacV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RbacV1beta1", reflect.TypeOf((*MockInterface)(nil).RbacV1beta1))
}

// SchedulingV1 mocks base method
func (m *MockInterface) SchedulingV1() v111.SchedulingV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1")
	ret0, _ := ret[0].(v111.SchedulingV1Interface)
	return ret0
}

// SchedulingV1 indicates an expected call of SchedulingV1
func (mr *MockInterfaceMockRecorder) SchedulingV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1", reflect.TypeOf((*MockInterface)(nil).SchedulingV1))
}

// SchedulingV1alpha1 mocks base method
func (m *MockInterface) SchedulingV1alpha1() v1alpha13.SchedulingV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1alpha1")
	ret0, _ := ret[0].(v1alpha13.SchedulingV1alpha1Interface)
	return ret0
}

// SchedulingV1alpha1 indicates an expected call of SchedulingV1alpha1
func (mr *MockInterfaceMockRecorder) SchedulingV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1alpha1", reflect.TypeOf((*MockInterface)(nil).SchedulingV1alpha1))
}

// SchedulingV1beta1 mocks base method
func (m *MockInterface) SchedulingV1beta1() v1beta113.SchedulingV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SchedulingV1beta1")
	ret0, _ := ret[0].(v1beta113.SchedulingV1beta1Interface)
	return ret0
}

// SchedulingV1beta1 indicates an expected call of SchedulingV1beta1
func (mr *MockInterfaceMockRecorder) SchedulingV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SchedulingV1beta1", reflect.TypeOf((*MockInterface)(nil).SchedulingV1beta1))
}

// SettingsV1alpha1 mocks base method
func (m *MockInterface) SettingsV1alpha1() v1alpha14.SettingsV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettingsV1alpha1")
	ret0, _ := ret[0].(v1alpha14.SettingsV1alpha1Interface)
	return ret0
}

// SettingsV1alpha1 indicates an expected call of SettingsV1alpha1
func (mr *MockInterfaceMockRecorder) SettingsV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettingsV1alpha1", reflect.TypeOf((*MockInterface)(nil).SettingsV1alpha1))
}

// StorageV1 mocks base method
func (m *MockInterface) StorageV1() v112.StorageV1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1")
	ret0, _ := ret[0].(v112.StorageV1Interface)
	return ret0
}

// StorageV1 indicates an expected call of StorageV1
func (mr *MockInterfaceMockRecorder) StorageV1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1", reflect.TypeOf((*MockInterface)(nil).StorageV1))
}

// StorageV1alpha1 mocks base method
func (m *MockInterface) StorageV1alpha1() v1alpha15.StorageV1alpha1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1alpha1")
	ret0, _ := ret[0].(v1alpha15.StorageV1alpha1Interface)
	return ret0
}

// StorageV1alpha1 indicates an expected call of StorageV1alpha1
func (mr *MockInterfaceMockRecorder) StorageV1alpha1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1alpha1", reflect.TypeOf((*MockInterface)(nil).StorageV1alpha1))
}

// StorageV1beta1 mocks base method
func (m *MockInterface) StorageV1beta1() v1beta114.StorageV1beta1Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageV1beta1")
	ret0, _ := ret[0].(v1beta114.StorageV1beta1Interface)
	return ret0
}

// StorageV1beta1 indicates an expected call of StorageV1beta1
func (mr *MockInterfaceMockRecorder) StorageV1beta1() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageV1beta1", reflect.TypeOf((*MockInterface)(nil).StorageV1beta1))
}
