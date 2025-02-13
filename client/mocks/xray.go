// Code generated by MockGen. DO NOT EDIT.
// Source: xray.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	xray "github.com/aws/aws-sdk-go-v2/service/xray"
	gomock "github.com/golang/mock/gomock"
)

// MockXrayClient is a mock of XrayClient interface.
type MockXrayClient struct {
	ctrl     *gomock.Controller
	recorder *MockXrayClientMockRecorder
}

// MockXrayClientMockRecorder is the mock recorder for MockXrayClient.
type MockXrayClientMockRecorder struct {
	mock *MockXrayClient
}

// NewMockXrayClient creates a new mock instance.
func NewMockXrayClient(ctrl *gomock.Controller) *MockXrayClient {
	mock := &MockXrayClient{ctrl: ctrl}
	mock.recorder = &MockXrayClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockXrayClient) EXPECT() *MockXrayClientMockRecorder {
	return m.recorder
}

// BatchGetTraces mocks base method.
func (m *MockXrayClient) BatchGetTraces(arg0 context.Context, arg1 *xray.BatchGetTracesInput, arg2 ...func(*xray.Options)) (*xray.BatchGetTracesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "BatchGetTraces", varargs...)
	ret0, _ := ret[0].(*xray.BatchGetTracesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchGetTraces indicates an expected call of BatchGetTraces.
func (mr *MockXrayClientMockRecorder) BatchGetTraces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchGetTraces", reflect.TypeOf((*MockXrayClient)(nil).BatchGetTraces), varargs...)
}

// GetEncryptionConfig mocks base method.
func (m *MockXrayClient) GetEncryptionConfig(arg0 context.Context, arg1 *xray.GetEncryptionConfigInput, arg2 ...func(*xray.Options)) (*xray.GetEncryptionConfigOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEncryptionConfig", varargs...)
	ret0, _ := ret[0].(*xray.GetEncryptionConfigOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEncryptionConfig indicates an expected call of GetEncryptionConfig.
func (mr *MockXrayClientMockRecorder) GetEncryptionConfig(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEncryptionConfig", reflect.TypeOf((*MockXrayClient)(nil).GetEncryptionConfig), varargs...)
}

// GetGroup mocks base method.
func (m *MockXrayClient) GetGroup(arg0 context.Context, arg1 *xray.GetGroupInput, arg2 ...func(*xray.Options)) (*xray.GetGroupOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroup", varargs...)
	ret0, _ := ret[0].(*xray.GetGroupOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroup indicates an expected call of GetGroup.
func (mr *MockXrayClientMockRecorder) GetGroup(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroup", reflect.TypeOf((*MockXrayClient)(nil).GetGroup), varargs...)
}

// GetGroups mocks base method.
func (m *MockXrayClient) GetGroups(arg0 context.Context, arg1 *xray.GetGroupsInput, arg2 ...func(*xray.Options)) (*xray.GetGroupsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGroups", varargs...)
	ret0, _ := ret[0].(*xray.GetGroupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroups indicates an expected call of GetGroups.
func (mr *MockXrayClientMockRecorder) GetGroups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroups", reflect.TypeOf((*MockXrayClient)(nil).GetGroups), varargs...)
}

// GetIndexingRules mocks base method.
func (m *MockXrayClient) GetIndexingRules(arg0 context.Context, arg1 *xray.GetIndexingRulesInput, arg2 ...func(*xray.Options)) (*xray.GetIndexingRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIndexingRules", varargs...)
	ret0, _ := ret[0].(*xray.GetIndexingRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIndexingRules indicates an expected call of GetIndexingRules.
func (mr *MockXrayClientMockRecorder) GetIndexingRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIndexingRules", reflect.TypeOf((*MockXrayClient)(nil).GetIndexingRules), varargs...)
}

// GetInsight mocks base method.
func (m *MockXrayClient) GetInsight(arg0 context.Context, arg1 *xray.GetInsightInput, arg2 ...func(*xray.Options)) (*xray.GetInsightOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsight", varargs...)
	ret0, _ := ret[0].(*xray.GetInsightOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsight indicates an expected call of GetInsight.
func (mr *MockXrayClientMockRecorder) GetInsight(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsight", reflect.TypeOf((*MockXrayClient)(nil).GetInsight), varargs...)
}

// GetInsightEvents mocks base method.
func (m *MockXrayClient) GetInsightEvents(arg0 context.Context, arg1 *xray.GetInsightEventsInput, arg2 ...func(*xray.Options)) (*xray.GetInsightEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsightEvents", varargs...)
	ret0, _ := ret[0].(*xray.GetInsightEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightEvents indicates an expected call of GetInsightEvents.
func (mr *MockXrayClientMockRecorder) GetInsightEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightEvents", reflect.TypeOf((*MockXrayClient)(nil).GetInsightEvents), varargs...)
}

// GetInsightImpactGraph mocks base method.
func (m *MockXrayClient) GetInsightImpactGraph(arg0 context.Context, arg1 *xray.GetInsightImpactGraphInput, arg2 ...func(*xray.Options)) (*xray.GetInsightImpactGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsightImpactGraph", varargs...)
	ret0, _ := ret[0].(*xray.GetInsightImpactGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightImpactGraph indicates an expected call of GetInsightImpactGraph.
func (mr *MockXrayClientMockRecorder) GetInsightImpactGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightImpactGraph", reflect.TypeOf((*MockXrayClient)(nil).GetInsightImpactGraph), varargs...)
}

// GetInsightSummaries mocks base method.
func (m *MockXrayClient) GetInsightSummaries(arg0 context.Context, arg1 *xray.GetInsightSummariesInput, arg2 ...func(*xray.Options)) (*xray.GetInsightSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInsightSummaries", varargs...)
	ret0, _ := ret[0].(*xray.GetInsightSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInsightSummaries indicates an expected call of GetInsightSummaries.
func (mr *MockXrayClientMockRecorder) GetInsightSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInsightSummaries", reflect.TypeOf((*MockXrayClient)(nil).GetInsightSummaries), varargs...)
}

// GetRetrievedTracesGraph mocks base method.
func (m *MockXrayClient) GetRetrievedTracesGraph(arg0 context.Context, arg1 *xray.GetRetrievedTracesGraphInput, arg2 ...func(*xray.Options)) (*xray.GetRetrievedTracesGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRetrievedTracesGraph", varargs...)
	ret0, _ := ret[0].(*xray.GetRetrievedTracesGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRetrievedTracesGraph indicates an expected call of GetRetrievedTracesGraph.
func (mr *MockXrayClientMockRecorder) GetRetrievedTracesGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRetrievedTracesGraph", reflect.TypeOf((*MockXrayClient)(nil).GetRetrievedTracesGraph), varargs...)
}

// GetSamplingRules mocks base method.
func (m *MockXrayClient) GetSamplingRules(arg0 context.Context, arg1 *xray.GetSamplingRulesInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingRulesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSamplingRules", varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingRulesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamplingRules indicates an expected call of GetSamplingRules.
func (mr *MockXrayClientMockRecorder) GetSamplingRules(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamplingRules", reflect.TypeOf((*MockXrayClient)(nil).GetSamplingRules), varargs...)
}

// GetSamplingStatisticSummaries mocks base method.
func (m *MockXrayClient) GetSamplingStatisticSummaries(arg0 context.Context, arg1 *xray.GetSamplingStatisticSummariesInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingStatisticSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSamplingStatisticSummaries", varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingStatisticSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamplingStatisticSummaries indicates an expected call of GetSamplingStatisticSummaries.
func (mr *MockXrayClientMockRecorder) GetSamplingStatisticSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamplingStatisticSummaries", reflect.TypeOf((*MockXrayClient)(nil).GetSamplingStatisticSummaries), varargs...)
}

// GetSamplingTargets mocks base method.
func (m *MockXrayClient) GetSamplingTargets(arg0 context.Context, arg1 *xray.GetSamplingTargetsInput, arg2 ...func(*xray.Options)) (*xray.GetSamplingTargetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSamplingTargets", varargs...)
	ret0, _ := ret[0].(*xray.GetSamplingTargetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamplingTargets indicates an expected call of GetSamplingTargets.
func (mr *MockXrayClientMockRecorder) GetSamplingTargets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamplingTargets", reflect.TypeOf((*MockXrayClient)(nil).GetSamplingTargets), varargs...)
}

// GetServiceGraph mocks base method.
func (m *MockXrayClient) GetServiceGraph(arg0 context.Context, arg1 *xray.GetServiceGraphInput, arg2 ...func(*xray.Options)) (*xray.GetServiceGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetServiceGraph", varargs...)
	ret0, _ := ret[0].(*xray.GetServiceGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceGraph indicates an expected call of GetServiceGraph.
func (mr *MockXrayClientMockRecorder) GetServiceGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceGraph", reflect.TypeOf((*MockXrayClient)(nil).GetServiceGraph), varargs...)
}

// GetTimeSeriesServiceStatistics mocks base method.
func (m *MockXrayClient) GetTimeSeriesServiceStatistics(arg0 context.Context, arg1 *xray.GetTimeSeriesServiceStatisticsInput, arg2 ...func(*xray.Options)) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTimeSeriesServiceStatistics", varargs...)
	ret0, _ := ret[0].(*xray.GetTimeSeriesServiceStatisticsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimeSeriesServiceStatistics indicates an expected call of GetTimeSeriesServiceStatistics.
func (mr *MockXrayClientMockRecorder) GetTimeSeriesServiceStatistics(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeSeriesServiceStatistics", reflect.TypeOf((*MockXrayClient)(nil).GetTimeSeriesServiceStatistics), varargs...)
}

// GetTraceGraph mocks base method.
func (m *MockXrayClient) GetTraceGraph(arg0 context.Context, arg1 *xray.GetTraceGraphInput, arg2 ...func(*xray.Options)) (*xray.GetTraceGraphOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTraceGraph", varargs...)
	ret0, _ := ret[0].(*xray.GetTraceGraphOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTraceGraph indicates an expected call of GetTraceGraph.
func (mr *MockXrayClientMockRecorder) GetTraceGraph(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraceGraph", reflect.TypeOf((*MockXrayClient)(nil).GetTraceGraph), varargs...)
}

// GetTraceSegmentDestination mocks base method.
func (m *MockXrayClient) GetTraceSegmentDestination(arg0 context.Context, arg1 *xray.GetTraceSegmentDestinationInput, arg2 ...func(*xray.Options)) (*xray.GetTraceSegmentDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTraceSegmentDestination", varargs...)
	ret0, _ := ret[0].(*xray.GetTraceSegmentDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTraceSegmentDestination indicates an expected call of GetTraceSegmentDestination.
func (mr *MockXrayClientMockRecorder) GetTraceSegmentDestination(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraceSegmentDestination", reflect.TypeOf((*MockXrayClient)(nil).GetTraceSegmentDestination), varargs...)
}

// GetTraceSummaries mocks base method.
func (m *MockXrayClient) GetTraceSummaries(arg0 context.Context, arg1 *xray.GetTraceSummariesInput, arg2 ...func(*xray.Options)) (*xray.GetTraceSummariesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTraceSummaries", varargs...)
	ret0, _ := ret[0].(*xray.GetTraceSummariesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTraceSummaries indicates an expected call of GetTraceSummaries.
func (mr *MockXrayClientMockRecorder) GetTraceSummaries(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTraceSummaries", reflect.TypeOf((*MockXrayClient)(nil).GetTraceSummaries), varargs...)
}

// ListResourcePolicies mocks base method.
func (m *MockXrayClient) ListResourcePolicies(arg0 context.Context, arg1 *xray.ListResourcePoliciesInput, arg2 ...func(*xray.Options)) (*xray.ListResourcePoliciesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListResourcePolicies", varargs...)
	ret0, _ := ret[0].(*xray.ListResourcePoliciesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListResourcePolicies indicates an expected call of ListResourcePolicies.
func (mr *MockXrayClientMockRecorder) ListResourcePolicies(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcePolicies", reflect.TypeOf((*MockXrayClient)(nil).ListResourcePolicies), varargs...)
}

// ListRetrievedTraces mocks base method.
func (m *MockXrayClient) ListRetrievedTraces(arg0 context.Context, arg1 *xray.ListRetrievedTracesInput, arg2 ...func(*xray.Options)) (*xray.ListRetrievedTracesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRetrievedTraces", varargs...)
	ret0, _ := ret[0].(*xray.ListRetrievedTracesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRetrievedTraces indicates an expected call of ListRetrievedTraces.
func (mr *MockXrayClientMockRecorder) ListRetrievedTraces(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRetrievedTraces", reflect.TypeOf((*MockXrayClient)(nil).ListRetrievedTraces), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockXrayClient) ListTagsForResource(arg0 context.Context, arg1 *xray.ListTagsForResourceInput, arg2 ...func(*xray.Options)) (*xray.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*xray.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockXrayClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockXrayClient)(nil).ListTagsForResource), varargs...)
}
