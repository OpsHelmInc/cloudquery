// Code generated by MockGen. DO NOT EDIT.
// Source: appsync.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	appsync "github.com/aws/aws-sdk-go-v2/service/appsync"
	gomock "github.com/golang/mock/gomock"
)

// MockAppsyncClient is a mock of AppsyncClient interface.
type MockAppsyncClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppsyncClientMockRecorder
}

// MockAppsyncClientMockRecorder is the mock recorder for MockAppsyncClient.
type MockAppsyncClientMockRecorder struct {
	mock *MockAppsyncClient
}

// NewMockAppsyncClient creates a new mock instance.
func NewMockAppsyncClient(ctrl *gomock.Controller) *MockAppsyncClient {
	mock := &MockAppsyncClient{ctrl: ctrl}
	mock.recorder = &MockAppsyncClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppsyncClient) EXPECT() *MockAppsyncClientMockRecorder {
	return m.recorder
}

// GetApiAssociation mocks base method.
func (m *MockAppsyncClient) GetApiAssociation(arg0 context.Context, arg1 *appsync.GetApiAssociationInput, arg2 ...func(*appsync.Options)) (*appsync.GetApiAssociationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetApiAssociation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApiAssociation", varargs...)
	ret0, _ := ret[0].(*appsync.GetApiAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApiAssociation indicates an expected call of GetApiAssociation.
func (mr *MockAppsyncClientMockRecorder) GetApiAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiAssociation", reflect.TypeOf((*MockAppsyncClient)(nil).GetApiAssociation), varargs...)
}

// GetApiCache mocks base method.
func (m *MockAppsyncClient) GetApiCache(arg0 context.Context, arg1 *appsync.GetApiCacheInput, arg2 ...func(*appsync.Options)) (*appsync.GetApiCacheOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetApiCache")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApiCache", varargs...)
	ret0, _ := ret[0].(*appsync.GetApiCacheOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApiCache indicates an expected call of GetApiCache.
func (mr *MockAppsyncClientMockRecorder) GetApiCache(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiCache", reflect.TypeOf((*MockAppsyncClient)(nil).GetApiCache), varargs...)
}

// GetDataSource mocks base method.
func (m *MockAppsyncClient) GetDataSource(arg0 context.Context, arg1 *appsync.GetDataSourceInput, arg2 ...func(*appsync.Options)) (*appsync.GetDataSourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDataSource")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataSource", varargs...)
	ret0, _ := ret[0].(*appsync.GetDataSourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataSource indicates an expected call of GetDataSource.
func (mr *MockAppsyncClientMockRecorder) GetDataSource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataSource", reflect.TypeOf((*MockAppsyncClient)(nil).GetDataSource), varargs...)
}

// GetDataSourceIntrospection mocks base method.
func (m *MockAppsyncClient) GetDataSourceIntrospection(arg0 context.Context, arg1 *appsync.GetDataSourceIntrospectionInput, arg2 ...func(*appsync.Options)) (*appsync.GetDataSourceIntrospectionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDataSourceIntrospection")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDataSourceIntrospection", varargs...)
	ret0, _ := ret[0].(*appsync.GetDataSourceIntrospectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataSourceIntrospection indicates an expected call of GetDataSourceIntrospection.
func (mr *MockAppsyncClientMockRecorder) GetDataSourceIntrospection(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataSourceIntrospection", reflect.TypeOf((*MockAppsyncClient)(nil).GetDataSourceIntrospection), varargs...)
}

// GetDomainName mocks base method.
func (m *MockAppsyncClient) GetDomainName(arg0 context.Context, arg1 *appsync.GetDomainNameInput, arg2 ...func(*appsync.Options)) (*appsync.GetDomainNameOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDomainName")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDomainName", varargs...)
	ret0, _ := ret[0].(*appsync.GetDomainNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDomainName indicates an expected call of GetDomainName.
func (mr *MockAppsyncClientMockRecorder) GetDomainName(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainName", reflect.TypeOf((*MockAppsyncClient)(nil).GetDomainName), varargs...)
}

// GetFunction mocks base method.
func (m *MockAppsyncClient) GetFunction(arg0 context.Context, arg1 *appsync.GetFunctionInput, arg2 ...func(*appsync.Options)) (*appsync.GetFunctionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetFunction")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFunction", varargs...)
	ret0, _ := ret[0].(*appsync.GetFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFunction indicates an expected call of GetFunction.
func (mr *MockAppsyncClientMockRecorder) GetFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFunction", reflect.TypeOf((*MockAppsyncClient)(nil).GetFunction), varargs...)
}

// GetGraphqlApi mocks base method.
func (m *MockAppsyncClient) GetGraphqlApi(arg0 context.Context, arg1 *appsync.GetGraphqlApiInput, arg2 ...func(*appsync.Options)) (*appsync.GetGraphqlApiOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetGraphqlApi")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGraphqlApi", varargs...)
	ret0, _ := ret[0].(*appsync.GetGraphqlApiOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGraphqlApi indicates an expected call of GetGraphqlApi.
func (mr *MockAppsyncClientMockRecorder) GetGraphqlApi(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGraphqlApi", reflect.TypeOf((*MockAppsyncClient)(nil).GetGraphqlApi), varargs...)
}

// GetGraphqlApiEnvironmentVariables mocks base method.
func (m *MockAppsyncClient) GetGraphqlApiEnvironmentVariables(arg0 context.Context, arg1 *appsync.GetGraphqlApiEnvironmentVariablesInput, arg2 ...func(*appsync.Options)) (*appsync.GetGraphqlApiEnvironmentVariablesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetGraphqlApiEnvironmentVariables")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGraphqlApiEnvironmentVariables", varargs...)
	ret0, _ := ret[0].(*appsync.GetGraphqlApiEnvironmentVariablesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGraphqlApiEnvironmentVariables indicates an expected call of GetGraphqlApiEnvironmentVariables.
func (mr *MockAppsyncClientMockRecorder) GetGraphqlApiEnvironmentVariables(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGraphqlApiEnvironmentVariables", reflect.TypeOf((*MockAppsyncClient)(nil).GetGraphqlApiEnvironmentVariables), varargs...)
}

// GetIntrospectionSchema mocks base method.
func (m *MockAppsyncClient) GetIntrospectionSchema(arg0 context.Context, arg1 *appsync.GetIntrospectionSchemaInput, arg2 ...func(*appsync.Options)) (*appsync.GetIntrospectionSchemaOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetIntrospectionSchema")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIntrospectionSchema", varargs...)
	ret0, _ := ret[0].(*appsync.GetIntrospectionSchemaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntrospectionSchema indicates an expected call of GetIntrospectionSchema.
func (mr *MockAppsyncClientMockRecorder) GetIntrospectionSchema(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntrospectionSchema", reflect.TypeOf((*MockAppsyncClient)(nil).GetIntrospectionSchema), varargs...)
}

// GetResolver mocks base method.
func (m *MockAppsyncClient) GetResolver(arg0 context.Context, arg1 *appsync.GetResolverInput, arg2 ...func(*appsync.Options)) (*appsync.GetResolverOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetResolver")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetResolver", varargs...)
	ret0, _ := ret[0].(*appsync.GetResolverOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResolver indicates an expected call of GetResolver.
func (mr *MockAppsyncClientMockRecorder) GetResolver(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResolver", reflect.TypeOf((*MockAppsyncClient)(nil).GetResolver), varargs...)
}

// GetSchemaCreationStatus mocks base method.
func (m *MockAppsyncClient) GetSchemaCreationStatus(arg0 context.Context, arg1 *appsync.GetSchemaCreationStatusInput, arg2 ...func(*appsync.Options)) (*appsync.GetSchemaCreationStatusOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSchemaCreationStatus")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSchemaCreationStatus", varargs...)
	ret0, _ := ret[0].(*appsync.GetSchemaCreationStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchemaCreationStatus indicates an expected call of GetSchemaCreationStatus.
func (mr *MockAppsyncClientMockRecorder) GetSchemaCreationStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchemaCreationStatus", reflect.TypeOf((*MockAppsyncClient)(nil).GetSchemaCreationStatus), varargs...)
}

// GetSourceApiAssociation mocks base method.
func (m *MockAppsyncClient) GetSourceApiAssociation(arg0 context.Context, arg1 *appsync.GetSourceApiAssociationInput, arg2 ...func(*appsync.Options)) (*appsync.GetSourceApiAssociationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSourceApiAssociation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSourceApiAssociation", varargs...)
	ret0, _ := ret[0].(*appsync.GetSourceApiAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSourceApiAssociation indicates an expected call of GetSourceApiAssociation.
func (mr *MockAppsyncClientMockRecorder) GetSourceApiAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSourceApiAssociation", reflect.TypeOf((*MockAppsyncClient)(nil).GetSourceApiAssociation), varargs...)
}

// GetType mocks base method.
func (m *MockAppsyncClient) GetType(arg0 context.Context, arg1 *appsync.GetTypeInput, arg2 ...func(*appsync.Options)) (*appsync.GetTypeOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetType")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetType", varargs...)
	ret0, _ := ret[0].(*appsync.GetTypeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetType indicates an expected call of GetType.
func (mr *MockAppsyncClientMockRecorder) GetType(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetType", reflect.TypeOf((*MockAppsyncClient)(nil).GetType), varargs...)
}

// ListApiKeys mocks base method.
func (m *MockAppsyncClient) ListApiKeys(arg0 context.Context, arg1 *appsync.ListApiKeysInput, arg2 ...func(*appsync.Options)) (*appsync.ListApiKeysOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListApiKeys")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApiKeys", varargs...)
	ret0, _ := ret[0].(*appsync.ListApiKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApiKeys indicates an expected call of ListApiKeys.
func (mr *MockAppsyncClientMockRecorder) ListApiKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApiKeys", reflect.TypeOf((*MockAppsyncClient)(nil).ListApiKeys), varargs...)
}

// ListDataSources mocks base method.
func (m *MockAppsyncClient) ListDataSources(arg0 context.Context, arg1 *appsync.ListDataSourcesInput, arg2 ...func(*appsync.Options)) (*appsync.ListDataSourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDataSources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDataSources", varargs...)
	ret0, _ := ret[0].(*appsync.ListDataSourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDataSources indicates an expected call of ListDataSources.
func (mr *MockAppsyncClientMockRecorder) ListDataSources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDataSources", reflect.TypeOf((*MockAppsyncClient)(nil).ListDataSources), varargs...)
}

// ListDomainNames mocks base method.
func (m *MockAppsyncClient) ListDomainNames(arg0 context.Context, arg1 *appsync.ListDomainNamesInput, arg2 ...func(*appsync.Options)) (*appsync.ListDomainNamesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListDomainNames")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDomainNames", varargs...)
	ret0, _ := ret[0].(*appsync.ListDomainNamesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDomainNames indicates an expected call of ListDomainNames.
func (mr *MockAppsyncClientMockRecorder) ListDomainNames(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDomainNames", reflect.TypeOf((*MockAppsyncClient)(nil).ListDomainNames), varargs...)
}

// ListFunctions mocks base method.
func (m *MockAppsyncClient) ListFunctions(arg0 context.Context, arg1 *appsync.ListFunctionsInput, arg2 ...func(*appsync.Options)) (*appsync.ListFunctionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListFunctions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFunctions", varargs...)
	ret0, _ := ret[0].(*appsync.ListFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFunctions indicates an expected call of ListFunctions.
func (mr *MockAppsyncClientMockRecorder) ListFunctions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFunctions", reflect.TypeOf((*MockAppsyncClient)(nil).ListFunctions), varargs...)
}

// ListGraphqlApis mocks base method.
func (m *MockAppsyncClient) ListGraphqlApis(arg0 context.Context, arg1 *appsync.ListGraphqlApisInput, arg2 ...func(*appsync.Options)) (*appsync.ListGraphqlApisOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListGraphqlApis")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGraphqlApis", varargs...)
	ret0, _ := ret[0].(*appsync.ListGraphqlApisOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGraphqlApis indicates an expected call of ListGraphqlApis.
func (mr *MockAppsyncClientMockRecorder) ListGraphqlApis(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGraphqlApis", reflect.TypeOf((*MockAppsyncClient)(nil).ListGraphqlApis), varargs...)
}

// ListResolvers mocks base method.
func (m *MockAppsyncClient) ListResolvers(arg0 context.Context, arg1 *appsync.ListResolversInput, arg2 ...func(*appsync.Options)) (*appsync.ListResolversOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListResolvers")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResolvers", varargs...)
	ret0, _ := ret[0].(*appsync.ListResolversOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResolvers indicates an expected call of ListResolvers.
func (mr *MockAppsyncClientMockRecorder) ListResolvers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResolvers", reflect.TypeOf((*MockAppsyncClient)(nil).ListResolvers), varargs...)
}

// ListResolversByFunction mocks base method.
func (m *MockAppsyncClient) ListResolversByFunction(arg0 context.Context, arg1 *appsync.ListResolversByFunctionInput, arg2 ...func(*appsync.Options)) (*appsync.ListResolversByFunctionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListResolversByFunction")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResolversByFunction", varargs...)
	ret0, _ := ret[0].(*appsync.ListResolversByFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResolversByFunction indicates an expected call of ListResolversByFunction.
func (mr *MockAppsyncClientMockRecorder) ListResolversByFunction(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResolversByFunction", reflect.TypeOf((*MockAppsyncClient)(nil).ListResolversByFunction), varargs...)
}

// ListSourceApiAssociations mocks base method.
func (m *MockAppsyncClient) ListSourceApiAssociations(arg0 context.Context, arg1 *appsync.ListSourceApiAssociationsInput, arg2 ...func(*appsync.Options)) (*appsync.ListSourceApiAssociationsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListSourceApiAssociations")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSourceApiAssociations", varargs...)
	ret0, _ := ret[0].(*appsync.ListSourceApiAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSourceApiAssociations indicates an expected call of ListSourceApiAssociations.
func (mr *MockAppsyncClientMockRecorder) ListSourceApiAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSourceApiAssociations", reflect.TypeOf((*MockAppsyncClient)(nil).ListSourceApiAssociations), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAppsyncClient) ListTagsForResource(arg0 context.Context, arg1 *appsync.ListTagsForResourceInput, arg2 ...func(*appsync.Options)) (*appsync.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
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
	ret0, _ := ret[0].(*appsync.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAppsyncClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppsyncClient)(nil).ListTagsForResource), varargs...)
}

// ListTypes mocks base method.
func (m *MockAppsyncClient) ListTypes(arg0 context.Context, arg1 *appsync.ListTypesInput, arg2 ...func(*appsync.Options)) (*appsync.ListTypesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTypes")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTypes", varargs...)
	ret0, _ := ret[0].(*appsync.ListTypesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTypes indicates an expected call of ListTypes.
func (mr *MockAppsyncClientMockRecorder) ListTypes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTypes", reflect.TypeOf((*MockAppsyncClient)(nil).ListTypes), varargs...)
}

// ListTypesByAssociation mocks base method.
func (m *MockAppsyncClient) ListTypesByAssociation(arg0 context.Context, arg1 *appsync.ListTypesByAssociationInput, arg2 ...func(*appsync.Options)) (*appsync.ListTypesByAssociationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &appsync.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTypesByAssociation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTypesByAssociation", varargs...)
	ret0, _ := ret[0].(*appsync.ListTypesByAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTypesByAssociation indicates an expected call of ListTypesByAssociation.
func (mr *MockAppsyncClientMockRecorder) ListTypesByAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTypesByAssociation", reflect.TypeOf((*MockAppsyncClient)(nil).ListTypesByAssociation), varargs...)
}
