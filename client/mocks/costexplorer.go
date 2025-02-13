// Code generated by MockGen. DO NOT EDIT.
// Source: costexplorer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	costexplorer "github.com/aws/aws-sdk-go-v2/service/costexplorer"
	gomock "github.com/golang/mock/gomock"
)

// MockCostexplorerClient is a mock of CostexplorerClient interface.
type MockCostexplorerClient struct {
	ctrl     *gomock.Controller
	recorder *MockCostexplorerClientMockRecorder
}

// MockCostexplorerClientMockRecorder is the mock recorder for MockCostexplorerClient.
type MockCostexplorerClientMockRecorder struct {
	mock *MockCostexplorerClient
}

// NewMockCostexplorerClient creates a new mock instance.
func NewMockCostexplorerClient(ctrl *gomock.Controller) *MockCostexplorerClient {
	mock := &MockCostexplorerClient{ctrl: ctrl}
	mock.recorder = &MockCostexplorerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCostexplorerClient) EXPECT() *MockCostexplorerClientMockRecorder {
	return m.recorder
}

// DescribeCostCategoryDefinition mocks base method.
func (m *MockCostexplorerClient) DescribeCostCategoryDefinition(arg0 context.Context, arg1 *costexplorer.DescribeCostCategoryDefinitionInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeCostCategoryDefinition")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeCostCategoryDefinition", varargs...)
	ret0, _ := ret[0].(*costexplorer.DescribeCostCategoryDefinitionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeCostCategoryDefinition indicates an expected call of DescribeCostCategoryDefinition.
func (mr *MockCostexplorerClientMockRecorder) DescribeCostCategoryDefinition(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeCostCategoryDefinition", reflect.TypeOf((*MockCostexplorerClient)(nil).DescribeCostCategoryDefinition), varargs...)
}

// GetAnomalies mocks base method.
func (m *MockCostexplorerClient) GetAnomalies(arg0 context.Context, arg1 *costexplorer.GetAnomaliesInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetAnomaliesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAnomalies")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAnomalies", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetAnomaliesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnomalies indicates an expected call of GetAnomalies.
func (mr *MockCostexplorerClientMockRecorder) GetAnomalies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnomalies", reflect.TypeOf((*MockCostexplorerClient)(nil).GetAnomalies), varargs...)
}

// GetAnomalyMonitors mocks base method.
func (m *MockCostexplorerClient) GetAnomalyMonitors(arg0 context.Context, arg1 *costexplorer.GetAnomalyMonitorsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetAnomalyMonitorsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAnomalyMonitors")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAnomalyMonitors", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetAnomalyMonitorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnomalyMonitors indicates an expected call of GetAnomalyMonitors.
func (mr *MockCostexplorerClientMockRecorder) GetAnomalyMonitors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnomalyMonitors", reflect.TypeOf((*MockCostexplorerClient)(nil).GetAnomalyMonitors), varargs...)
}

// GetAnomalySubscriptions mocks base method.
func (m *MockCostexplorerClient) GetAnomalySubscriptions(arg0 context.Context, arg1 *costexplorer.GetAnomalySubscriptionsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetAnomalySubscriptionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetAnomalySubscriptions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAnomalySubscriptions", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetAnomalySubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAnomalySubscriptions indicates an expected call of GetAnomalySubscriptions.
func (mr *MockCostexplorerClientMockRecorder) GetAnomalySubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAnomalySubscriptions", reflect.TypeOf((*MockCostexplorerClient)(nil).GetAnomalySubscriptions), varargs...)
}

// GetApproximateUsageRecords mocks base method.
func (m *MockCostexplorerClient) GetApproximateUsageRecords(arg0 context.Context, arg1 *costexplorer.GetApproximateUsageRecordsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetApproximateUsageRecordsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetApproximateUsageRecords")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApproximateUsageRecords", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetApproximateUsageRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApproximateUsageRecords indicates an expected call of GetApproximateUsageRecords.
func (mr *MockCostexplorerClientMockRecorder) GetApproximateUsageRecords(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApproximateUsageRecords", reflect.TypeOf((*MockCostexplorerClient)(nil).GetApproximateUsageRecords), varargs...)
}

// GetCostAndUsage mocks base method.
func (m *MockCostexplorerClient) GetCostAndUsage(arg0 context.Context, arg1 *costexplorer.GetCostAndUsageInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetCostAndUsageOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCostAndUsage")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCostAndUsage", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetCostAndUsageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostAndUsage indicates an expected call of GetCostAndUsage.
func (mr *MockCostexplorerClientMockRecorder) GetCostAndUsage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostAndUsage", reflect.TypeOf((*MockCostexplorerClient)(nil).GetCostAndUsage), varargs...)
}

// GetCostAndUsageWithResources mocks base method.
func (m *MockCostexplorerClient) GetCostAndUsageWithResources(arg0 context.Context, arg1 *costexplorer.GetCostAndUsageWithResourcesInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCostAndUsageWithResources")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCostAndUsageWithResources", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetCostAndUsageWithResourcesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostAndUsageWithResources indicates an expected call of GetCostAndUsageWithResources.
func (mr *MockCostexplorerClientMockRecorder) GetCostAndUsageWithResources(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostAndUsageWithResources", reflect.TypeOf((*MockCostexplorerClient)(nil).GetCostAndUsageWithResources), varargs...)
}

// GetCostCategories mocks base method.
func (m *MockCostexplorerClient) GetCostCategories(arg0 context.Context, arg1 *costexplorer.GetCostCategoriesInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetCostCategoriesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCostCategories")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCostCategories", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetCostCategoriesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostCategories indicates an expected call of GetCostCategories.
func (mr *MockCostexplorerClientMockRecorder) GetCostCategories(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostCategories", reflect.TypeOf((*MockCostexplorerClient)(nil).GetCostCategories), varargs...)
}

// GetCostForecast mocks base method.
func (m *MockCostexplorerClient) GetCostForecast(arg0 context.Context, arg1 *costexplorer.GetCostForecastInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetCostForecastOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetCostForecast")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCostForecast", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetCostForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCostForecast indicates an expected call of GetCostForecast.
func (mr *MockCostexplorerClientMockRecorder) GetCostForecast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCostForecast", reflect.TypeOf((*MockCostexplorerClient)(nil).GetCostForecast), varargs...)
}

// GetDimensionValues mocks base method.
func (m *MockCostexplorerClient) GetDimensionValues(arg0 context.Context, arg1 *costexplorer.GetDimensionValuesInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetDimensionValuesOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetDimensionValues")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDimensionValues", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetDimensionValuesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionValues indicates an expected call of GetDimensionValues.
func (mr *MockCostexplorerClientMockRecorder) GetDimensionValues(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionValues", reflect.TypeOf((*MockCostexplorerClient)(nil).GetDimensionValues), varargs...)
}

// GetReservationCoverage mocks base method.
func (m *MockCostexplorerClient) GetReservationCoverage(arg0 context.Context, arg1 *costexplorer.GetReservationCoverageInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetReservationCoverageOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetReservationCoverage")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReservationCoverage", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetReservationCoverageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationCoverage indicates an expected call of GetReservationCoverage.
func (mr *MockCostexplorerClientMockRecorder) GetReservationCoverage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationCoverage", reflect.TypeOf((*MockCostexplorerClient)(nil).GetReservationCoverage), varargs...)
}

// GetReservationPurchaseRecommendation mocks base method.
func (m *MockCostexplorerClient) GetReservationPurchaseRecommendation(arg0 context.Context, arg1 *costexplorer.GetReservationPurchaseRecommendationInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetReservationPurchaseRecommendation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReservationPurchaseRecommendation", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetReservationPurchaseRecommendationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationPurchaseRecommendation indicates an expected call of GetReservationPurchaseRecommendation.
func (mr *MockCostexplorerClientMockRecorder) GetReservationPurchaseRecommendation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationPurchaseRecommendation", reflect.TypeOf((*MockCostexplorerClient)(nil).GetReservationPurchaseRecommendation), varargs...)
}

// GetReservationUtilization mocks base method.
func (m *MockCostexplorerClient) GetReservationUtilization(arg0 context.Context, arg1 *costexplorer.GetReservationUtilizationInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetReservationUtilizationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetReservationUtilization")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReservationUtilization", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetReservationUtilizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReservationUtilization indicates an expected call of GetReservationUtilization.
func (mr *MockCostexplorerClientMockRecorder) GetReservationUtilization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReservationUtilization", reflect.TypeOf((*MockCostexplorerClient)(nil).GetReservationUtilization), varargs...)
}

// GetRightsizingRecommendation mocks base method.
func (m *MockCostexplorerClient) GetRightsizingRecommendation(arg0 context.Context, arg1 *costexplorer.GetRightsizingRecommendationInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetRightsizingRecommendationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetRightsizingRecommendation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRightsizingRecommendation", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetRightsizingRecommendationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRightsizingRecommendation indicates an expected call of GetRightsizingRecommendation.
func (mr *MockCostexplorerClientMockRecorder) GetRightsizingRecommendation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRightsizingRecommendation", reflect.TypeOf((*MockCostexplorerClient)(nil).GetRightsizingRecommendation), varargs...)
}

// GetSavingsPlanPurchaseRecommendationDetails mocks base method.
func (m *MockCostexplorerClient) GetSavingsPlanPurchaseRecommendationDetails(arg0 context.Context, arg1 *costexplorer.GetSavingsPlanPurchaseRecommendationDetailsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlanPurchaseRecommendationDetailsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSavingsPlanPurchaseRecommendationDetails")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSavingsPlanPurchaseRecommendationDetails", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetSavingsPlanPurchaseRecommendationDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavingsPlanPurchaseRecommendationDetails indicates an expected call of GetSavingsPlanPurchaseRecommendationDetails.
func (mr *MockCostexplorerClientMockRecorder) GetSavingsPlanPurchaseRecommendationDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavingsPlanPurchaseRecommendationDetails", reflect.TypeOf((*MockCostexplorerClient)(nil).GetSavingsPlanPurchaseRecommendationDetails), varargs...)
}

// GetSavingsPlansCoverage mocks base method.
func (m *MockCostexplorerClient) GetSavingsPlansCoverage(arg0 context.Context, arg1 *costexplorer.GetSavingsPlansCoverageInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansCoverageOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSavingsPlansCoverage")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSavingsPlansCoverage", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetSavingsPlansCoverageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavingsPlansCoverage indicates an expected call of GetSavingsPlansCoverage.
func (mr *MockCostexplorerClientMockRecorder) GetSavingsPlansCoverage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavingsPlansCoverage", reflect.TypeOf((*MockCostexplorerClient)(nil).GetSavingsPlansCoverage), varargs...)
}

// GetSavingsPlansPurchaseRecommendation mocks base method.
func (m *MockCostexplorerClient) GetSavingsPlansPurchaseRecommendation(arg0 context.Context, arg1 *costexplorer.GetSavingsPlansPurchaseRecommendationInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSavingsPlansPurchaseRecommendation")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSavingsPlansPurchaseRecommendation", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetSavingsPlansPurchaseRecommendationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavingsPlansPurchaseRecommendation indicates an expected call of GetSavingsPlansPurchaseRecommendation.
func (mr *MockCostexplorerClientMockRecorder) GetSavingsPlansPurchaseRecommendation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavingsPlansPurchaseRecommendation", reflect.TypeOf((*MockCostexplorerClient)(nil).GetSavingsPlansPurchaseRecommendation), varargs...)
}

// GetSavingsPlansUtilization mocks base method.
func (m *MockCostexplorerClient) GetSavingsPlansUtilization(arg0 context.Context, arg1 *costexplorer.GetSavingsPlansUtilizationInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSavingsPlansUtilization")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSavingsPlansUtilization", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetSavingsPlansUtilizationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavingsPlansUtilization indicates an expected call of GetSavingsPlansUtilization.
func (mr *MockCostexplorerClientMockRecorder) GetSavingsPlansUtilization(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavingsPlansUtilization", reflect.TypeOf((*MockCostexplorerClient)(nil).GetSavingsPlansUtilization), varargs...)
}

// GetSavingsPlansUtilizationDetails mocks base method.
func (m *MockCostexplorerClient) GetSavingsPlansUtilizationDetails(arg0 context.Context, arg1 *costexplorer.GetSavingsPlansUtilizationDetailsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetSavingsPlansUtilizationDetails")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSavingsPlansUtilizationDetails", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetSavingsPlansUtilizationDetailsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSavingsPlansUtilizationDetails indicates an expected call of GetSavingsPlansUtilizationDetails.
func (mr *MockCostexplorerClientMockRecorder) GetSavingsPlansUtilizationDetails(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSavingsPlansUtilizationDetails", reflect.TypeOf((*MockCostexplorerClient)(nil).GetSavingsPlansUtilizationDetails), varargs...)
}

// GetTags mocks base method.
func (m *MockCostexplorerClient) GetTags(arg0 context.Context, arg1 *costexplorer.GetTagsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetTagsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetTags")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTags", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockCostexplorerClientMockRecorder) GetTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockCostexplorerClient)(nil).GetTags), varargs...)
}

// GetUsageForecast mocks base method.
func (m *MockCostexplorerClient) GetUsageForecast(arg0 context.Context, arg1 *costexplorer.GetUsageForecastInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.GetUsageForecastOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to GetUsageForecast")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUsageForecast", varargs...)
	ret0, _ := ret[0].(*costexplorer.GetUsageForecastOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsageForecast indicates an expected call of GetUsageForecast.
func (mr *MockCostexplorerClientMockRecorder) GetUsageForecast(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsageForecast", reflect.TypeOf((*MockCostexplorerClient)(nil).GetUsageForecast), varargs...)
}

// ListCostAllocationTagBackfillHistory mocks base method.
func (m *MockCostexplorerClient) ListCostAllocationTagBackfillHistory(arg0 context.Context, arg1 *costexplorer.ListCostAllocationTagBackfillHistoryInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.ListCostAllocationTagBackfillHistoryOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListCostAllocationTagBackfillHistory")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCostAllocationTagBackfillHistory", varargs...)
	ret0, _ := ret[0].(*costexplorer.ListCostAllocationTagBackfillHistoryOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCostAllocationTagBackfillHistory indicates an expected call of ListCostAllocationTagBackfillHistory.
func (mr *MockCostexplorerClientMockRecorder) ListCostAllocationTagBackfillHistory(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCostAllocationTagBackfillHistory", reflect.TypeOf((*MockCostexplorerClient)(nil).ListCostAllocationTagBackfillHistory), varargs...)
}

// ListCostAllocationTags mocks base method.
func (m *MockCostexplorerClient) ListCostAllocationTags(arg0 context.Context, arg1 *costexplorer.ListCostAllocationTagsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.ListCostAllocationTagsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListCostAllocationTags")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCostAllocationTags", varargs...)
	ret0, _ := ret[0].(*costexplorer.ListCostAllocationTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCostAllocationTags indicates an expected call of ListCostAllocationTags.
func (mr *MockCostexplorerClientMockRecorder) ListCostAllocationTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCostAllocationTags", reflect.TypeOf((*MockCostexplorerClient)(nil).ListCostAllocationTags), varargs...)
}

// ListCostCategoryDefinitions mocks base method.
func (m *MockCostexplorerClient) ListCostCategoryDefinitions(arg0 context.Context, arg1 *costexplorer.ListCostCategoryDefinitionsInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListCostCategoryDefinitions")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListCostCategoryDefinitions", varargs...)
	ret0, _ := ret[0].(*costexplorer.ListCostCategoryDefinitionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCostCategoryDefinitions indicates an expected call of ListCostCategoryDefinitions.
func (mr *MockCostexplorerClientMockRecorder) ListCostCategoryDefinitions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCostCategoryDefinitions", reflect.TypeOf((*MockCostexplorerClient)(nil).ListCostCategoryDefinitions), varargs...)
}

// ListSavingsPlansPurchaseRecommendationGeneration mocks base method.
func (m *MockCostexplorerClient) ListSavingsPlansPurchaseRecommendationGeneration(arg0 context.Context, arg1 *costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListSavingsPlansPurchaseRecommendationGeneration")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSavingsPlansPurchaseRecommendationGeneration", varargs...)
	ret0, _ := ret[0].(*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSavingsPlansPurchaseRecommendationGeneration indicates an expected call of ListSavingsPlansPurchaseRecommendationGeneration.
func (mr *MockCostexplorerClientMockRecorder) ListSavingsPlansPurchaseRecommendationGeneration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSavingsPlansPurchaseRecommendationGeneration", reflect.TypeOf((*MockCostexplorerClient)(nil).ListSavingsPlansPurchaseRecommendationGeneration), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockCostexplorerClient) ListTagsForResource(arg0 context.Context, arg1 *costexplorer.ListTagsForResourceInput, arg2 ...func(*costexplorer.Options)) (*costexplorer.ListTagsForResourceOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &costexplorer.Options{}
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
	ret0, _ := ret[0].(*costexplorer.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockCostexplorerClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockCostexplorerClient)(nil).ListTagsForResource), varargs...)
}
