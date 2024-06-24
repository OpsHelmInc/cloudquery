// Code generated by MockGen. DO NOT EDIT.
// Source: securityhub.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	securityhub "github.com/aws/aws-sdk-go-v2/service/securityhub"
	gomock "github.com/golang/mock/gomock"
)

// MockSecurityhubClient is a mock of SecurityhubClient interface.
type MockSecurityhubClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityhubClientMockRecorder
}

// MockSecurityhubClientMockRecorder is the mock recorder for MockSecurityhubClient.
type MockSecurityhubClientMockRecorder struct {
	mock *MockSecurityhubClient
}

// NewMockSecurityhubClient creates a new mock instance.
func NewMockSecurityhubClient(ctrl *gomock.Controller) *MockSecurityhubClient {
	mock := &MockSecurityhubClient{ctrl: ctrl}
	mock.recorder = &MockSecurityhubClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecurityhubClient) EXPECT() *MockSecurityhubClientMockRecorder {
	return m.recorder
}

// BatchGetAutomationRules mocks base method.
func (m *MockSecurityhubClient) BatchGetAutomationRules(arg0 context.Context, arg1 *securityhub.BatchGetAutomationRulesInput, arg2 ...func(*securityhub.Options)) (*securityhub.BatchGetAutomationRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetAutomationRules", varargs...)
	ret0, _ := ret[0].(*securityhub.BatchGetAutomationRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetAutomationRules indicates an expected call of BatchGetAutomationRules.
func (mr *MockSecurityhubClientMockRecorder) BatchGetAutomationRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetAutomationRules", reflect.TypeOf((*MockSecurityhubClient)(nil).BatchGetAutomationRules), varargs...)
}

// BatchGetConfigurationPolicyAssociations mocks base method.
func (m *MockSecurityhubClient) BatchGetConfigurationPolicyAssociations(arg0 context.Context, arg1 *securityhub.BatchGetConfigurationPolicyAssociationsInput, arg2 ...func(*securityhub.Options)) (*securityhub.BatchGetConfigurationPolicyAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetConfigurationPolicyAssociations", varargs...)
	ret0, _ := ret[0].(*securityhub.BatchGetConfigurationPolicyAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetConfigurationPolicyAssociations indicates an expected call of BatchGetConfigurationPolicyAssociations.
func (mr *MockSecurityhubClientMockRecorder) BatchGetConfigurationPolicyAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetConfigurationPolicyAssociations", reflect.TypeOf((*MockSecurityhubClient)(nil).BatchGetConfigurationPolicyAssociations), varargs...)
}

// BatchGetSecurityControls mocks base method.
func (m *MockSecurityhubClient) BatchGetSecurityControls(arg0 context.Context, arg1 *securityhub.BatchGetSecurityControlsInput, arg2 ...func(*securityhub.Options)) (*securityhub.BatchGetSecurityControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetSecurityControls", varargs...)
	ret0, _ := ret[0].(*securityhub.BatchGetSecurityControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetSecurityControls indicates an expected call of BatchGetSecurityControls.
func (mr *MockSecurityhubClientMockRecorder) BatchGetSecurityControls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetSecurityControls", reflect.TypeOf((*MockSecurityhubClient)(nil).BatchGetSecurityControls), varargs...)
}

// BatchGetStandardsControlAssociations mocks base method.
func (m *MockSecurityhubClient) BatchGetStandardsControlAssociations(arg0 context.Context, arg1 *securityhub.BatchGetStandardsControlAssociationsInput, arg2 ...func(*securityhub.Options)) (*securityhub.BatchGetStandardsControlAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetStandardsControlAssociations", varargs...)
	ret0, _ := ret[0].(*securityhub.BatchGetStandardsControlAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetStandardsControlAssociations indicates an expected call of BatchGetStandardsControlAssociations.
func (mr *MockSecurityhubClientMockRecorder) BatchGetStandardsControlAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetStandardsControlAssociations", reflect.TypeOf((*MockSecurityhubClient)(nil).BatchGetStandardsControlAssociations), varargs...)
}

// DescribeActionTargets mocks base method.
func (m *MockSecurityhubClient) DescribeActionTargets(arg0 context.Context, arg1 *securityhub.DescribeActionTargetsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeActionTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeActionTargets", varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeActionTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeActionTargets indicates an expected call of DescribeActionTargets.
func (mr *MockSecurityhubClientMockRecorder) DescribeActionTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeActionTargets", reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeActionTargets), varargs...)
}

// DescribeHub mocks base method.
func (m *MockSecurityhubClient) DescribeHub(arg0 context.Context, arg1 *securityhub.DescribeHubInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeHubOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHub", varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeHubOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHub indicates an expected call of DescribeHub.
func (mr *MockSecurityhubClientMockRecorder) DescribeHub(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHub", reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeHub), varargs...)
}

// DescribeOrganizationConfiguration mocks base method.
func (m *MockSecurityhubClient) DescribeOrganizationConfiguration(arg0 context.Context, arg1 *securityhub.DescribeOrganizationConfigurationInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeOrganizationConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeOrganizationConfiguration", varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeOrganizationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeOrganizationConfiguration indicates an expected call of DescribeOrganizationConfiguration.
func (mr *MockSecurityhubClientMockRecorder) DescribeOrganizationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeOrganizationConfiguration", reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeOrganizationConfiguration), varargs...)
}

// DescribeProducts mocks base method.
func (m *MockSecurityhubClient) DescribeProducts(arg0 context.Context, arg1 *securityhub.DescribeProductsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeProductsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeProducts", varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeProductsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeProducts indicates an expected call of DescribeProducts.
func (mr *MockSecurityhubClientMockRecorder) DescribeProducts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeProducts", reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeProducts), varargs...)
}

// DescribeStandards mocks base method.
func (m *MockSecurityhubClient) DescribeStandards(arg0 context.Context, arg1 *securityhub.DescribeStandardsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeStandardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStandards", varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeStandardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStandards indicates an expected call of DescribeStandards.
func (mr *MockSecurityhubClientMockRecorder) DescribeStandards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStandards", reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeStandards), varargs...)
}

// DescribeStandardsControls mocks base method.
func (m *MockSecurityhubClient) DescribeStandardsControls(arg0 context.Context, arg1 *securityhub.DescribeStandardsControlsInput, arg2 ...func(*securityhub.Options)) (*securityhub.DescribeStandardsControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStandardsControls", varargs...)
	ret0, _ := ret[0].(*securityhub.DescribeStandardsControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStandardsControls indicates an expected call of DescribeStandardsControls.
func (mr *MockSecurityhubClientMockRecorder) DescribeStandardsControls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStandardsControls", reflect.TypeOf((*MockSecurityhubClient)(nil).DescribeStandardsControls), varargs...)
}

// GetAdministratorAccount mocks base method.
func (m *MockSecurityhubClient) GetAdministratorAccount(arg0 context.Context, arg1 *securityhub.GetAdministratorAccountInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetAdministratorAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAdministratorAccount", varargs...)
	ret0, _ := ret[0].(*securityhub.GetAdministratorAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdministratorAccount indicates an expected call of GetAdministratorAccount.
func (mr *MockSecurityhubClientMockRecorder) GetAdministratorAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdministratorAccount", reflect.TypeOf((*MockSecurityhubClient)(nil).GetAdministratorAccount), varargs...)
}

// GetConfigurationPolicy mocks base method.
func (m *MockSecurityhubClient) GetConfigurationPolicy(arg0 context.Context, arg1 *securityhub.GetConfigurationPolicyInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetConfigurationPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigurationPolicy", varargs...)
	ret0, _ := ret[0].(*securityhub.GetConfigurationPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigurationPolicy indicates an expected call of GetConfigurationPolicy.
func (mr *MockSecurityhubClientMockRecorder) GetConfigurationPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigurationPolicy", reflect.TypeOf((*MockSecurityhubClient)(nil).GetConfigurationPolicy), varargs...)
}

// GetConfigurationPolicyAssociation mocks base method.
func (m *MockSecurityhubClient) GetConfigurationPolicyAssociation(arg0 context.Context, arg1 *securityhub.GetConfigurationPolicyAssociationInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetConfigurationPolicyAssociationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetConfigurationPolicyAssociation", varargs...)
	ret0, _ := ret[0].(*securityhub.GetConfigurationPolicyAssociationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigurationPolicyAssociation indicates an expected call of GetConfigurationPolicyAssociation.
func (mr *MockSecurityhubClientMockRecorder) GetConfigurationPolicyAssociation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigurationPolicyAssociation", reflect.TypeOf((*MockSecurityhubClient)(nil).GetConfigurationPolicyAssociation), varargs...)
}

// GetEnabledStandards mocks base method.
func (m *MockSecurityhubClient) GetEnabledStandards(arg0 context.Context, arg1 *securityhub.GetEnabledStandardsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetEnabledStandardsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnabledStandards", varargs...)
	ret0, _ := ret[0].(*securityhub.GetEnabledStandardsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnabledStandards indicates an expected call of GetEnabledStandards.
func (mr *MockSecurityhubClientMockRecorder) GetEnabledStandards(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnabledStandards", reflect.TypeOf((*MockSecurityhubClient)(nil).GetEnabledStandards), varargs...)
}

// GetFindingAggregator mocks base method.
func (m *MockSecurityhubClient) GetFindingAggregator(arg0 context.Context, arg1 *securityhub.GetFindingAggregatorInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetFindingAggregatorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFindingAggregator", varargs...)
	ret0, _ := ret[0].(*securityhub.GetFindingAggregatorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFindingAggregator indicates an expected call of GetFindingAggregator.
func (mr *MockSecurityhubClientMockRecorder) GetFindingAggregator(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFindingAggregator", reflect.TypeOf((*MockSecurityhubClient)(nil).GetFindingAggregator), varargs...)
}

// GetFindingHistory mocks base method.
func (m *MockSecurityhubClient) GetFindingHistory(arg0 context.Context, arg1 *securityhub.GetFindingHistoryInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetFindingHistoryOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFindingHistory", varargs...)
	ret0, _ := ret[0].(*securityhub.GetFindingHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFindingHistory indicates an expected call of GetFindingHistory.
func (mr *MockSecurityhubClientMockRecorder) GetFindingHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFindingHistory", reflect.TypeOf((*MockSecurityhubClient)(nil).GetFindingHistory), varargs...)
}

// GetFindings mocks base method.
func (m *MockSecurityhubClient) GetFindings(arg0 context.Context, arg1 *securityhub.GetFindingsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetFindingsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFindings", varargs...)
	ret0, _ := ret[0].(*securityhub.GetFindingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFindings indicates an expected call of GetFindings.
func (mr *MockSecurityhubClientMockRecorder) GetFindings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFindings", reflect.TypeOf((*MockSecurityhubClient)(nil).GetFindings), varargs...)
}

// GetInsightResults mocks base method.
func (m *MockSecurityhubClient) GetInsightResults(arg0 context.Context, arg1 *securityhub.GetInsightResultsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetInsightResultsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsightResults", varargs...)
	ret0, _ := ret[0].(*securityhub.GetInsightResultsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightResults indicates an expected call of GetInsightResults.
func (mr *MockSecurityhubClientMockRecorder) GetInsightResults(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightResults", reflect.TypeOf((*MockSecurityhubClient)(nil).GetInsightResults), varargs...)
}

// GetInsights mocks base method.
func (m *MockSecurityhubClient) GetInsights(arg0 context.Context, arg1 *securityhub.GetInsightsInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetInsightsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsights", varargs...)
	ret0, _ := ret[0].(*securityhub.GetInsightsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsights indicates an expected call of GetInsights.
func (mr *MockSecurityhubClientMockRecorder) GetInsights(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsights", reflect.TypeOf((*MockSecurityhubClient)(nil).GetInsights), varargs...)
}

// GetInvitationsCount mocks base method.
func (m *MockSecurityhubClient) GetInvitationsCount(arg0 context.Context, arg1 *securityhub.GetInvitationsCountInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetInvitationsCountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInvitationsCount", varargs...)
	ret0, _ := ret[0].(*securityhub.GetInvitationsCountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInvitationsCount indicates an expected call of GetInvitationsCount.
func (mr *MockSecurityhubClientMockRecorder) GetInvitationsCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInvitationsCount", reflect.TypeOf((*MockSecurityhubClient)(nil).GetInvitationsCount), varargs...)
}

// GetMasterAccount mocks base method.
func (m *MockSecurityhubClient) GetMasterAccount(arg0 context.Context, arg1 *securityhub.GetMasterAccountInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetMasterAccountOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMasterAccount", varargs...)
	ret0, _ := ret[0].(*securityhub.GetMasterAccountOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMasterAccount indicates an expected call of GetMasterAccount.
func (mr *MockSecurityhubClientMockRecorder) GetMasterAccount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterAccount", reflect.TypeOf((*MockSecurityhubClient)(nil).GetMasterAccount), varargs...)
}

// GetMembers mocks base method.
func (m *MockSecurityhubClient) GetMembers(arg0 context.Context, arg1 *securityhub.GetMembersInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMembers", varargs...)
	ret0, _ := ret[0].(*securityhub.GetMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMembers indicates an expected call of GetMembers.
func (mr *MockSecurityhubClientMockRecorder) GetMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembers", reflect.TypeOf((*MockSecurityhubClient)(nil).GetMembers), varargs...)
}

// GetSecurityControlDefinition mocks base method.
func (m *MockSecurityhubClient) GetSecurityControlDefinition(arg0 context.Context, arg1 *securityhub.GetSecurityControlDefinitionInput, arg2 ...func(*securityhub.Options)) (*securityhub.GetSecurityControlDefinitionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecurityControlDefinition", varargs...)
	ret0, _ := ret[0].(*securityhub.GetSecurityControlDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecurityControlDefinition indicates an expected call of GetSecurityControlDefinition.
func (mr *MockSecurityhubClientMockRecorder) GetSecurityControlDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecurityControlDefinition", reflect.TypeOf((*MockSecurityhubClient)(nil).GetSecurityControlDefinition), varargs...)
}

// ListAutomationRules mocks base method.
func (m *MockSecurityhubClient) ListAutomationRules(arg0 context.Context, arg1 *securityhub.ListAutomationRulesInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListAutomationRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAutomationRules", varargs...)
	ret0, _ := ret[0].(*securityhub.ListAutomationRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAutomationRules indicates an expected call of ListAutomationRules.
func (mr *MockSecurityhubClientMockRecorder) ListAutomationRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAutomationRules", reflect.TypeOf((*MockSecurityhubClient)(nil).ListAutomationRules), varargs...)
}

// ListConfigurationPolicies mocks base method.
func (m *MockSecurityhubClient) ListConfigurationPolicies(arg0 context.Context, arg1 *securityhub.ListConfigurationPoliciesInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListConfigurationPoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConfigurationPolicies", varargs...)
	ret0, _ := ret[0].(*securityhub.ListConfigurationPoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurationPolicies indicates an expected call of ListConfigurationPolicies.
func (mr *MockSecurityhubClientMockRecorder) ListConfigurationPolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurationPolicies", reflect.TypeOf((*MockSecurityhubClient)(nil).ListConfigurationPolicies), varargs...)
}

// ListConfigurationPolicyAssociations mocks base method.
func (m *MockSecurityhubClient) ListConfigurationPolicyAssociations(arg0 context.Context, arg1 *securityhub.ListConfigurationPolicyAssociationsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListConfigurationPolicyAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListConfigurationPolicyAssociations", varargs...)
	ret0, _ := ret[0].(*securityhub.ListConfigurationPolicyAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListConfigurationPolicyAssociations indicates an expected call of ListConfigurationPolicyAssociations.
func (mr *MockSecurityhubClientMockRecorder) ListConfigurationPolicyAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConfigurationPolicyAssociations", reflect.TypeOf((*MockSecurityhubClient)(nil).ListConfigurationPolicyAssociations), varargs...)
}

// ListEnabledProductsForImport mocks base method.
func (m *MockSecurityhubClient) ListEnabledProductsForImport(arg0 context.Context, arg1 *securityhub.ListEnabledProductsForImportInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListEnabledProductsForImportOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnabledProductsForImport", varargs...)
	ret0, _ := ret[0].(*securityhub.ListEnabledProductsForImportOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnabledProductsForImport indicates an expected call of ListEnabledProductsForImport.
func (mr *MockSecurityhubClientMockRecorder) ListEnabledProductsForImport(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnabledProductsForImport", reflect.TypeOf((*MockSecurityhubClient)(nil).ListEnabledProductsForImport), varargs...)
}

// ListFindingAggregators mocks base method.
func (m *MockSecurityhubClient) ListFindingAggregators(arg0 context.Context, arg1 *securityhub.ListFindingAggregatorsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListFindingAggregatorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListFindingAggregators", varargs...)
	ret0, _ := ret[0].(*securityhub.ListFindingAggregatorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFindingAggregators indicates an expected call of ListFindingAggregators.
func (mr *MockSecurityhubClientMockRecorder) ListFindingAggregators(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFindingAggregators", reflect.TypeOf((*MockSecurityhubClient)(nil).ListFindingAggregators), varargs...)
}

// ListInvitations mocks base method.
func (m *MockSecurityhubClient) ListInvitations(arg0 context.Context, arg1 *securityhub.ListInvitationsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListInvitationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListInvitations", varargs...)
	ret0, _ := ret[0].(*securityhub.ListInvitationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInvitations indicates an expected call of ListInvitations.
func (mr *MockSecurityhubClientMockRecorder) ListInvitations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInvitations", reflect.TypeOf((*MockSecurityhubClient)(nil).ListInvitations), varargs...)
}

// ListMembers mocks base method.
func (m *MockSecurityhubClient) ListMembers(arg0 context.Context, arg1 *securityhub.ListMembersInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListMembersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMembers", varargs...)
	ret0, _ := ret[0].(*securityhub.ListMembersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockSecurityhubClientMockRecorder) ListMembers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockSecurityhubClient)(nil).ListMembers), varargs...)
}

// ListOrganizationAdminAccounts mocks base method.
func (m *MockSecurityhubClient) ListOrganizationAdminAccounts(arg0 context.Context, arg1 *securityhub.ListOrganizationAdminAccountsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListOrganizationAdminAccountsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizationAdminAccounts", varargs...)
	ret0, _ := ret[0].(*securityhub.ListOrganizationAdminAccountsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizationAdminAccounts indicates an expected call of ListOrganizationAdminAccounts.
func (mr *MockSecurityhubClientMockRecorder) ListOrganizationAdminAccounts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizationAdminAccounts", reflect.TypeOf((*MockSecurityhubClient)(nil).ListOrganizationAdminAccounts), varargs...)
}

// ListSecurityControlDefinitions mocks base method.
func (m *MockSecurityhubClient) ListSecurityControlDefinitions(arg0 context.Context, arg1 *securityhub.ListSecurityControlDefinitionsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListSecurityControlDefinitionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecurityControlDefinitions", varargs...)
	ret0, _ := ret[0].(*securityhub.ListSecurityControlDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecurityControlDefinitions indicates an expected call of ListSecurityControlDefinitions.
func (mr *MockSecurityhubClientMockRecorder) ListSecurityControlDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecurityControlDefinitions", reflect.TypeOf((*MockSecurityhubClient)(nil).ListSecurityControlDefinitions), varargs...)
}

// ListStandardsControlAssociations mocks base method.
func (m *MockSecurityhubClient) ListStandardsControlAssociations(arg0 context.Context, arg1 *securityhub.ListStandardsControlAssociationsInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListStandardsControlAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStandardsControlAssociations", varargs...)
	ret0, _ := ret[0].(*securityhub.ListStandardsControlAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStandardsControlAssociations indicates an expected call of ListStandardsControlAssociations.
func (mr *MockSecurityhubClientMockRecorder) ListStandardsControlAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStandardsControlAssociations", reflect.TypeOf((*MockSecurityhubClient)(nil).ListStandardsControlAssociations), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockSecurityhubClient) ListTagsForResource(arg0 context.Context, arg1 *securityhub.ListTagsForResourceInput, arg2 ...func(*securityhub.Options)) (*securityhub.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*securityhub.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockSecurityhubClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockSecurityhubClient)(nil).ListTagsForResource), varargs...)
}
