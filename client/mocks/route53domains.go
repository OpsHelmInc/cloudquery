// Code generated by MockGen. DO NOT EDIT.
// Source: route53domains.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	route53domains "github.com/aws/aws-sdk-go-v2/service/route53domains"
	gomock "github.com/golang/mock/gomock"
)

// MockRoute53domainsClient is a mock of Route53domainsClient interface.
type MockRoute53domainsClient struct {
	ctrl     *gomock.Controller
	recorder *MockRoute53domainsClientMockRecorder
}

// MockRoute53domainsClientMockRecorder is the mock recorder for MockRoute53domainsClient.
type MockRoute53domainsClientMockRecorder struct {
	mock *MockRoute53domainsClient
}

// NewMockRoute53domainsClient creates a new mock instance.
func NewMockRoute53domainsClient(ctrl *gomock.Controller) *MockRoute53domainsClient {
	mock := &MockRoute53domainsClient{ctrl: ctrl}
	mock.recorder = &MockRoute53domainsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoute53domainsClient) EXPECT() *MockRoute53domainsClientMockRecorder {
	return m.recorder
}

// GetContactReachabilityStatus mocks base method.
func (m *MockRoute53domainsClient) GetContactReachabilityStatus(arg0 context.Context, arg1 *route53domains.GetContactReachabilityStatusInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetContactReachabilityStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContactReachabilityStatus", varargs...)
	ret0, _ := ret[0].(*route53domains.GetContactReachabilityStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactReachabilityStatus indicates an expected call of GetContactReachabilityStatus.
func (mr *MockRoute53domainsClientMockRecorder) GetContactReachabilityStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactReachabilityStatus", reflect.TypeOf((*MockRoute53domainsClient)(nil).GetContactReachabilityStatus), varargs...)
}

// GetDomainDetail mocks base method.
func (m *MockRoute53domainsClient) GetDomainDetail(arg0 context.Context, arg1 *route53domains.GetDomainDetailInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetDomainDetailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainDetail", varargs...)
	ret0, _ := ret[0].(*route53domains.GetDomainDetailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainDetail indicates an expected call of GetDomainDetail.
func (mr *MockRoute53domainsClientMockRecorder) GetDomainDetail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainDetail", reflect.TypeOf((*MockRoute53domainsClient)(nil).GetDomainDetail), varargs...)
}

// GetDomainSuggestions mocks base method.
func (m *MockRoute53domainsClient) GetDomainSuggestions(arg0 context.Context, arg1 *route53domains.GetDomainSuggestionsInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetDomainSuggestionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainSuggestions", varargs...)
	ret0, _ := ret[0].(*route53domains.GetDomainSuggestionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainSuggestions indicates an expected call of GetDomainSuggestions.
func (mr *MockRoute53domainsClientMockRecorder) GetDomainSuggestions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainSuggestions", reflect.TypeOf((*MockRoute53domainsClient)(nil).GetDomainSuggestions), varargs...)
}

// GetOperationDetail mocks base method.
func (m *MockRoute53domainsClient) GetOperationDetail(arg0 context.Context, arg1 *route53domains.GetOperationDetailInput, arg2 ...func(*route53domains.Options)) (*route53domains.GetOperationDetailOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOperationDetail", varargs...)
	ret0, _ := ret[0].(*route53domains.GetOperationDetailOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperationDetail indicates an expected call of GetOperationDetail.
func (mr *MockRoute53domainsClientMockRecorder) GetOperationDetail(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationDetail", reflect.TypeOf((*MockRoute53domainsClient)(nil).GetOperationDetail), varargs...)
}

// ListDomains mocks base method.
func (m *MockRoute53domainsClient) ListDomains(arg0 context.Context, arg1 *route53domains.ListDomainsInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListDomainsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDomains", varargs...)
	ret0, _ := ret[0].(*route53domains.ListDomainsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomains indicates an expected call of ListDomains.
func (mr *MockRoute53domainsClientMockRecorder) ListDomains(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomains", reflect.TypeOf((*MockRoute53domainsClient)(nil).ListDomains), varargs...)
}

// ListOperations mocks base method.
func (m *MockRoute53domainsClient) ListOperations(arg0 context.Context, arg1 *route53domains.ListOperationsInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListOperationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOperations", varargs...)
	ret0, _ := ret[0].(*route53domains.ListOperationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOperations indicates an expected call of ListOperations.
func (mr *MockRoute53domainsClientMockRecorder) ListOperations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperations", reflect.TypeOf((*MockRoute53domainsClient)(nil).ListOperations), varargs...)
}

// ListPrices mocks base method.
func (m *MockRoute53domainsClient) ListPrices(arg0 context.Context, arg1 *route53domains.ListPricesInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListPricesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPrices", varargs...)
	ret0, _ := ret[0].(*route53domains.ListPricesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPrices indicates an expected call of ListPrices.
func (mr *MockRoute53domainsClientMockRecorder) ListPrices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPrices", reflect.TypeOf((*MockRoute53domainsClient)(nil).ListPrices), varargs...)
}

// ListTagsForDomain mocks base method.
func (m *MockRoute53domainsClient) ListTagsForDomain(arg0 context.Context, arg1 *route53domains.ListTagsForDomainInput, arg2 ...func(*route53domains.Options)) (*route53domains.ListTagsForDomainOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForDomain", varargs...)
	ret0, _ := ret[0].(*route53domains.ListTagsForDomainOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForDomain indicates an expected call of ListTagsForDomain.
func (mr *MockRoute53domainsClientMockRecorder) ListTagsForDomain(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForDomain", reflect.TypeOf((*MockRoute53domainsClient)(nil).ListTagsForDomain), varargs...)
}
