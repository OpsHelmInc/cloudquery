// Code generated by MockGen. DO NOT EDIT.
// Source: appconfig.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	appconfig "github.com/aws/aws-sdk-go-v2/service/appconfig"
	gomock "github.com/golang/mock/gomock"
)

// MockAppconfigClient is a mock of AppconfigClient interface.
type MockAppconfigClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppconfigClientMockRecorder
}

// MockAppconfigClientMockRecorder is the mock recorder for MockAppconfigClient.
type MockAppconfigClientMockRecorder struct {
	mock *MockAppconfigClient
}

// NewMockAppconfigClient creates a new mock instance.
func NewMockAppconfigClient(ctrl *gomock.Controller) *MockAppconfigClient {
	mock := &MockAppconfigClient{ctrl: ctrl}
	mock.recorder = &MockAppconfigClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppconfigClient) EXPECT() *MockAppconfigClientMockRecorder {
	return m.recorder
}

// GetApplication mocks base method.
func (m *MockAppconfigClient) GetApplication(arg0 context.Context, arg1 *appconfig.GetApplicationInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetApplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApplication", varargs...)
	ret0, _ := ret[0].(*appconfig.GetApplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockAppconfigClientMockRecorder) GetApplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockAppconfigClient)(nil).GetApplication), varargs...)
}

// GetConfiguration mocks base method.
func (m *MockAppconfigClient) GetConfiguration(arg0 context.Context, arg1 *appconfig.GetConfigurationInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfiguration", varargs...)
	ret0, _ := ret[0].(*appconfig.GetConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfiguration indicates an expected call of GetConfiguration.
func (mr *MockAppconfigClientMockRecorder) GetConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfiguration", reflect.TypeOf((*MockAppconfigClient)(nil).GetConfiguration), varargs...)
}

// GetConfigurationProfile mocks base method.
func (m *MockAppconfigClient) GetConfigurationProfile(arg0 context.Context, arg1 *appconfig.GetConfigurationProfileInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetConfigurationProfileOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigurationProfile", varargs...)
	ret0, _ := ret[0].(*appconfig.GetConfigurationProfileOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigurationProfile indicates an expected call of GetConfigurationProfile.
func (mr *MockAppconfigClientMockRecorder) GetConfigurationProfile(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigurationProfile", reflect.TypeOf((*MockAppconfigClient)(nil).GetConfigurationProfile), varargs...)
}

// GetDeployment mocks base method.
func (m *MockAppconfigClient) GetDeployment(arg0 context.Context, arg1 *appconfig.GetDeploymentInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetDeploymentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeployment", varargs...)
	ret0, _ := ret[0].(*appconfig.GetDeploymentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockAppconfigClientMockRecorder) GetDeployment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockAppconfigClient)(nil).GetDeployment), varargs...)
}

// GetDeploymentStrategy mocks base method.
func (m *MockAppconfigClient) GetDeploymentStrategy(arg0 context.Context, arg1 *appconfig.GetDeploymentStrategyInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetDeploymentStrategyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeploymentStrategy", varargs...)
	ret0, _ := ret[0].(*appconfig.GetDeploymentStrategyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentStrategy indicates an expected call of GetDeploymentStrategy.
func (mr *MockAppconfigClientMockRecorder) GetDeploymentStrategy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentStrategy", reflect.TypeOf((*MockAppconfigClient)(nil).GetDeploymentStrategy), varargs...)
}

// GetEnvironment mocks base method.
func (m *MockAppconfigClient) GetEnvironment(arg0 context.Context, arg1 *appconfig.GetEnvironmentInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetEnvironmentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnvironment", varargs...)
	ret0, _ := ret[0].(*appconfig.GetEnvironmentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment.
func (mr *MockAppconfigClientMockRecorder) GetEnvironment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockAppconfigClient)(nil).GetEnvironment), varargs...)
}

// GetExtension mocks base method.
func (m *MockAppconfigClient) GetExtension(arg0 context.Context, arg1 *appconfig.GetExtensionInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetExtensionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExtension", varargs...)
	ret0, _ := ret[0].(*appconfig.GetExtensionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtension indicates an expected call of GetExtension.
func (mr *MockAppconfigClientMockRecorder) GetExtension(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtension", reflect.TypeOf((*MockAppconfigClient)(nil).GetExtension), varargs...)
}

// GetExtensionAssociation mocks base method.
func (m *MockAppconfigClient) GetExtensionAssociation(arg0 context.Context, arg1 *appconfig.GetExtensionAssociationInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetExtensionAssociationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetExtensionAssociation", varargs...)
	ret0, _ := ret[0].(*appconfig.GetExtensionAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExtensionAssociation indicates an expected call of GetExtensionAssociation.
func (mr *MockAppconfigClientMockRecorder) GetExtensionAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtensionAssociation", reflect.TypeOf((*MockAppconfigClient)(nil).GetExtensionAssociation), varargs...)
}

// GetHostedConfigurationVersion mocks base method.
func (m *MockAppconfigClient) GetHostedConfigurationVersion(arg0 context.Context, arg1 *appconfig.GetHostedConfigurationVersionInput, arg2 ...func(*appconfig.Options)) (*appconfig.GetHostedConfigurationVersionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetHostedConfigurationVersion", varargs...)
	ret0, _ := ret[0].(*appconfig.GetHostedConfigurationVersionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHostedConfigurationVersion indicates an expected call of GetHostedConfigurationVersion.
func (mr *MockAppconfigClientMockRecorder) GetHostedConfigurationVersion(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostedConfigurationVersion", reflect.TypeOf((*MockAppconfigClient)(nil).GetHostedConfigurationVersion), varargs...)
}

// ListApplications mocks base method.
func (m *MockAppconfigClient) ListApplications(arg0 context.Context, arg1 *appconfig.ListApplicationsInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApplications", varargs...)
	ret0, _ := ret[0].(*appconfig.ListApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockAppconfigClientMockRecorder) ListApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockAppconfigClient)(nil).ListApplications), varargs...)
}

// ListConfigurationProfiles mocks base method.
func (m *MockAppconfigClient) ListConfigurationProfiles(arg0 context.Context, arg1 *appconfig.ListConfigurationProfilesInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListConfigurationProfilesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConfigurationProfiles", varargs...)
	ret0, _ := ret[0].(*appconfig.ListConfigurationProfilesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurationProfiles indicates an expected call of ListConfigurationProfiles.
func (mr *MockAppconfigClientMockRecorder) ListConfigurationProfiles(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurationProfiles", reflect.TypeOf((*MockAppconfigClient)(nil).ListConfigurationProfiles), varargs...)
}

// ListDeploymentStrategies mocks base method.
func (m *MockAppconfigClient) ListDeploymentStrategies(arg0 context.Context, arg1 *appconfig.ListDeploymentStrategiesInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListDeploymentStrategiesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeploymentStrategies", varargs...)
	ret0, _ := ret[0].(*appconfig.ListDeploymentStrategiesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploymentStrategies indicates an expected call of ListDeploymentStrategies.
func (mr *MockAppconfigClientMockRecorder) ListDeploymentStrategies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploymentStrategies", reflect.TypeOf((*MockAppconfigClient)(nil).ListDeploymentStrategies), varargs...)
}

// ListDeployments mocks base method.
func (m *MockAppconfigClient) ListDeployments(arg0 context.Context, arg1 *appconfig.ListDeploymentsInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListDeploymentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeployments", varargs...)
	ret0, _ := ret[0].(*appconfig.ListDeploymentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployments indicates an expected call of ListDeployments.
func (mr *MockAppconfigClientMockRecorder) ListDeployments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockAppconfigClient)(nil).ListDeployments), varargs...)
}

// ListEnvironments mocks base method.
func (m *MockAppconfigClient) ListEnvironments(arg0 context.Context, arg1 *appconfig.ListEnvironmentsInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListEnvironmentsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvironments", varargs...)
	ret0, _ := ret[0].(*appconfig.ListEnvironmentsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments.
func (mr *MockAppconfigClientMockRecorder) ListEnvironments(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockAppconfigClient)(nil).ListEnvironments), varargs...)
}

// ListExtensionAssociations mocks base method.
func (m *MockAppconfigClient) ListExtensionAssociations(arg0 context.Context, arg1 *appconfig.ListExtensionAssociationsInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListExtensionAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListExtensionAssociations", varargs...)
	ret0, _ := ret[0].(*appconfig.ListExtensionAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExtensionAssociations indicates an expected call of ListExtensionAssociations.
func (mr *MockAppconfigClientMockRecorder) ListExtensionAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExtensionAssociations", reflect.TypeOf((*MockAppconfigClient)(nil).ListExtensionAssociations), varargs...)
}

// ListExtensions mocks base method.
func (m *MockAppconfigClient) ListExtensions(arg0 context.Context, arg1 *appconfig.ListExtensionsInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListExtensionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListExtensions", varargs...)
	ret0, _ := ret[0].(*appconfig.ListExtensionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListExtensions indicates an expected call of ListExtensions.
func (mr *MockAppconfigClientMockRecorder) ListExtensions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListExtensions", reflect.TypeOf((*MockAppconfigClient)(nil).ListExtensions), varargs...)
}

// ListHostedConfigurationVersions mocks base method.
func (m *MockAppconfigClient) ListHostedConfigurationVersions(arg0 context.Context, arg1 *appconfig.ListHostedConfigurationVersionsInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListHostedConfigurationVersions", varargs...)
	ret0, _ := ret[0].(*appconfig.ListHostedConfigurationVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedConfigurationVersions indicates an expected call of ListHostedConfigurationVersions.
func (mr *MockAppconfigClientMockRecorder) ListHostedConfigurationVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedConfigurationVersions", reflect.TypeOf((*MockAppconfigClient)(nil).ListHostedConfigurationVersions), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAppconfigClient) ListTagsForResource(arg0 context.Context, arg1 *appconfig.ListTagsForResourceInput, arg2 ...func(*appconfig.Options)) (*appconfig.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*appconfig.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAppconfigClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppconfigClient)(nil).ListTagsForResource), varargs...)
}
