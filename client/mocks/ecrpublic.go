// Code generated by MockGen. DO NOT EDIT.
// Source: ecrpublic.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	ecrpublic "github.com/aws/aws-sdk-go-v2/service/ecrpublic"
	gomock "github.com/golang/mock/gomock"
)

// MockEcrpublicClient is a mock of EcrpublicClient interface.
type MockEcrpublicClient struct {
	ctrl     *gomock.Controller
	recorder *MockEcrpublicClientMockRecorder
}

// MockEcrpublicClientMockRecorder is the mock recorder for MockEcrpublicClient.
type MockEcrpublicClientMockRecorder struct {
	mock *MockEcrpublicClient
}

// NewMockEcrpublicClient creates a new mock instance.
func NewMockEcrpublicClient(ctrl *gomock.Controller) *MockEcrpublicClient {
	mock := &MockEcrpublicClient{ctrl: ctrl}
	mock.recorder = &MockEcrpublicClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEcrpublicClient) EXPECT() *MockEcrpublicClientMockRecorder {
	return m.recorder
}

// DescribeImageTags mocks base method.
func (m *MockEcrpublicClient) DescribeImageTags(arg0 context.Context, arg1 *ecrpublic.DescribeImageTagsInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImageTagsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImageTags", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeImageTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImageTags indicates an expected call of DescribeImageTags.
func (mr *MockEcrpublicClientMockRecorder) DescribeImageTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImageTags", reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeImageTags), varargs...)
}

// DescribeImages mocks base method.
func (m *MockEcrpublicClient) DescribeImages(arg0 context.Context, arg1 *ecrpublic.DescribeImagesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImages", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImages indicates an expected call of DescribeImages.
func (mr *MockEcrpublicClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeImages), varargs...)
}

// DescribeRegistries mocks base method.
func (m *MockEcrpublicClient) DescribeRegistries(arg0 context.Context, arg1 *ecrpublic.DescribeRegistriesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRegistriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRegistries", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeRegistriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRegistries indicates an expected call of DescribeRegistries.
func (mr *MockEcrpublicClientMockRecorder) DescribeRegistries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRegistries", reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeRegistries), varargs...)
}

// DescribeRepositories mocks base method.
func (m *MockEcrpublicClient) DescribeRepositories(arg0 context.Context, arg1 *ecrpublic.DescribeRepositoriesInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.DescribeRepositoriesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRepositories", varargs...)
	ret0, _ := ret[0].(*ecrpublic.DescribeRepositoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRepositories indicates an expected call of DescribeRepositories.
func (mr *MockEcrpublicClientMockRecorder) DescribeRepositories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRepositories", reflect.TypeOf((*MockEcrpublicClient)(nil).DescribeRepositories), varargs...)
}

// GetAuthorizationToken mocks base method.
func (m *MockEcrpublicClient) GetAuthorizationToken(arg0 context.Context, arg1 *ecrpublic.GetAuthorizationTokenInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetAuthorizationTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAuthorizationToken", varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetAuthorizationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAuthorizationToken indicates an expected call of GetAuthorizationToken.
func (mr *MockEcrpublicClientMockRecorder) GetAuthorizationToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthorizationToken", reflect.TypeOf((*MockEcrpublicClient)(nil).GetAuthorizationToken), varargs...)
}

// GetRegistryCatalogData mocks base method.
func (m *MockEcrpublicClient) GetRegistryCatalogData(arg0 context.Context, arg1 *ecrpublic.GetRegistryCatalogDataInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetRegistryCatalogDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRegistryCatalogData", varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetRegistryCatalogDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRegistryCatalogData indicates an expected call of GetRegistryCatalogData.
func (mr *MockEcrpublicClientMockRecorder) GetRegistryCatalogData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRegistryCatalogData", reflect.TypeOf((*MockEcrpublicClient)(nil).GetRegistryCatalogData), varargs...)
}

// GetRepositoryCatalogData mocks base method.
func (m *MockEcrpublicClient) GetRepositoryCatalogData(arg0 context.Context, arg1 *ecrpublic.GetRepositoryCatalogDataInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryCatalogDataOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRepositoryCatalogData", varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetRepositoryCatalogDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryCatalogData indicates an expected call of GetRepositoryCatalogData.
func (mr *MockEcrpublicClientMockRecorder) GetRepositoryCatalogData(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryCatalogData", reflect.TypeOf((*MockEcrpublicClient)(nil).GetRepositoryCatalogData), varargs...)
}

// GetRepositoryPolicy mocks base method.
func (m *MockEcrpublicClient) GetRepositoryPolicy(arg0 context.Context, arg1 *ecrpublic.GetRepositoryPolicyInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.GetRepositoryPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRepositoryPolicy", varargs...)
	ret0, _ := ret[0].(*ecrpublic.GetRepositoryPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRepositoryPolicy indicates an expected call of GetRepositoryPolicy.
func (mr *MockEcrpublicClientMockRecorder) GetRepositoryPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRepositoryPolicy", reflect.TypeOf((*MockEcrpublicClient)(nil).GetRepositoryPolicy), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockEcrpublicClient) ListTagsForResource(arg0 context.Context, arg1 *ecrpublic.ListTagsForResourceInput, arg2 ...func(*ecrpublic.Options)) (*ecrpublic.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*ecrpublic.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockEcrpublicClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockEcrpublicClient)(nil).ListTagsForResource), varargs...)
}
