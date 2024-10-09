// Code generated by MockGen. DO NOT EDIT.
// Source: servicediscovery.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	servicediscovery "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	gomock "github.com/golang/mock/gomock"
)

// MockServicediscoveryClient is a mock of ServicediscoveryClient interface.
type MockServicediscoveryClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicediscoveryClientMockRecorder
}

// MockServicediscoveryClientMockRecorder is the mock recorder for MockServicediscoveryClient.
type MockServicediscoveryClientMockRecorder struct {
	mock *MockServicediscoveryClient
}

// NewMockServicediscoveryClient creates a new mock instance.
func NewMockServicediscoveryClient(ctrl *gomock.Controller) *MockServicediscoveryClient {
	mock := &MockServicediscoveryClient{ctrl: ctrl}
	mock.recorder = &MockServicediscoveryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicediscoveryClient) EXPECT() *MockServicediscoveryClientMockRecorder {
	return m.recorder
}

// GetInstance mocks base method.
func (m *MockServicediscoveryClient) GetInstance(arg0 context.Context, arg1 *servicediscovery.GetInstanceInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.GetInstanceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetInstance")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstance", varargs...)
	ret0, _ := ret[0].(*servicediscovery.GetInstanceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstance indicates an expected call of GetInstance.
func (mr *MockServicediscoveryClientMockRecorder) GetInstance(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstance", reflect.TypeOf((*MockServicediscoveryClient)(nil).GetInstance), varargs...)
}

// GetInstancesHealthStatus mocks base method.
func (m *MockServicediscoveryClient) GetInstancesHealthStatus(arg0 context.Context, arg1 *servicediscovery.GetInstancesHealthStatusInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.GetInstancesHealthStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetInstancesHealthStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstancesHealthStatus", varargs...)
	ret0, _ := ret[0].(*servicediscovery.GetInstancesHealthStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstancesHealthStatus indicates an expected call of GetInstancesHealthStatus.
func (mr *MockServicediscoveryClientMockRecorder) GetInstancesHealthStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstancesHealthStatus", reflect.TypeOf((*MockServicediscoveryClient)(nil).GetInstancesHealthStatus), varargs...)
}

// GetNamespace mocks base method.
func (m *MockServicediscoveryClient) GetNamespace(arg0 context.Context, arg1 *servicediscovery.GetNamespaceInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.GetNamespaceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetNamespace")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNamespace", varargs...)
	ret0, _ := ret[0].(*servicediscovery.GetNamespaceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNamespace indicates an expected call of GetNamespace.
func (mr *MockServicediscoveryClientMockRecorder) GetNamespace(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockServicediscoveryClient)(nil).GetNamespace), varargs...)
}

// GetOperation mocks base method.
func (m *MockServicediscoveryClient) GetOperation(arg0 context.Context, arg1 *servicediscovery.GetOperationInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.GetOperationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetOperation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOperation", varargs...)
	ret0, _ := ret[0].(*servicediscovery.GetOperationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperation indicates an expected call of GetOperation.
func (mr *MockServicediscoveryClientMockRecorder) GetOperation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperation", reflect.TypeOf((*MockServicediscoveryClient)(nil).GetOperation), varargs...)
}

// GetService mocks base method.
func (m *MockServicediscoveryClient) GetService(arg0 context.Context, arg1 *servicediscovery.GetServiceInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.GetServiceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetService")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetService", varargs...)
	ret0, _ := ret[0].(*servicediscovery.GetServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockServicediscoveryClientMockRecorder) GetService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockServicediscoveryClient)(nil).GetService), varargs...)
}

// ListInstances mocks base method.
func (m *MockServicediscoveryClient) ListInstances(arg0 context.Context, arg1 *servicediscovery.ListInstancesInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.ListInstancesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListInstances")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInstances", varargs...)
	ret0, _ := ret[0].(*servicediscovery.ListInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInstances indicates an expected call of ListInstances.
func (mr *MockServicediscoveryClientMockRecorder) ListInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInstances", reflect.TypeOf((*MockServicediscoveryClient)(nil).ListInstances), varargs...)
}

// ListNamespaces mocks base method.
func (m *MockServicediscoveryClient) ListNamespaces(arg0 context.Context, arg1 *servicediscovery.ListNamespacesInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.ListNamespacesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListNamespaces")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNamespaces", varargs...)
	ret0, _ := ret[0].(*servicediscovery.ListNamespacesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockServicediscoveryClientMockRecorder) ListNamespaces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockServicediscoveryClient)(nil).ListNamespaces), varargs...)
}

// ListOperations mocks base method.
func (m *MockServicediscoveryClient) ListOperations(arg0 context.Context, arg1 *servicediscovery.ListOperationsInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.ListOperationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListOperations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperations", varargs...)
	ret0, _ := ret[0].(*servicediscovery.ListOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockServicediscoveryClientMockRecorder) ListOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockServicediscoveryClient)(nil).ListOperations), varargs...)
}

// ListServices mocks base method.
func (m *MockServicediscoveryClient) ListServices(arg0 context.Context, arg1 *servicediscovery.ListServicesInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.ListServicesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListServices")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListServices", varargs...)
	ret0, _ := ret[0].(*servicediscovery.ListServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServices indicates an expected call of ListServices.
func (mr *MockServicediscoveryClientMockRecorder) ListServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServices", reflect.TypeOf((*MockServicediscoveryClient)(nil).ListServices), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockServicediscoveryClient) ListTagsForResource(arg0 context.Context, arg1 *servicediscovery.ListTagsForResourceInput, arg2 ...func(*servicediscovery.Options)) (*servicediscovery.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &servicediscovery.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTagsForResource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*servicediscovery.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockServicediscoveryClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockServicediscoveryClient)(nil).ListTagsForResource), varargs...)
}
