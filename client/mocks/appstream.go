// Code generated by MockGen. DO NOT EDIT.
// Source: appstream.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	appstream "github.com/aws/aws-sdk-go-v2/service/appstream"
	gomock "github.com/golang/mock/gomock"
)

// MockAppstreamClient is a mock of AppstreamClient interface.
type MockAppstreamClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppstreamClientMockRecorder
}

// MockAppstreamClientMockRecorder is the mock recorder for MockAppstreamClient.
type MockAppstreamClientMockRecorder struct {
	mock *MockAppstreamClient
}

// NewMockAppstreamClient creates a new mock instance.
func NewMockAppstreamClient(ctrl *gomock.Controller) *MockAppstreamClient {
	mock := &MockAppstreamClient{ctrl: ctrl}
	mock.recorder = &MockAppstreamClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppstreamClient) EXPECT() *MockAppstreamClientMockRecorder {
	return m.recorder
}

// DescribeAppBlockBuilderAppBlockAssociations mocks base method.
func (m *MockAppstreamClient) DescribeAppBlockBuilderAppBlockAssociations(arg0 context.Context, arg1 *appstream.DescribeAppBlockBuilderAppBlockAssociationsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppBlockBuilderAppBlockAssociations", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeAppBlockBuilderAppBlockAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppBlockBuilderAppBlockAssociations indicates an expected call of DescribeAppBlockBuilderAppBlockAssociations.
func (mr *MockAppstreamClientMockRecorder) DescribeAppBlockBuilderAppBlockAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppBlockBuilderAppBlockAssociations", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeAppBlockBuilderAppBlockAssociations), varargs...)
}

// DescribeAppBlockBuilders mocks base method.
func (m *MockAppstreamClient) DescribeAppBlockBuilders(arg0 context.Context, arg1 *appstream.DescribeAppBlockBuildersInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeAppBlockBuildersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppBlockBuilders", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeAppBlockBuildersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppBlockBuilders indicates an expected call of DescribeAppBlockBuilders.
func (mr *MockAppstreamClientMockRecorder) DescribeAppBlockBuilders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppBlockBuilders", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeAppBlockBuilders), varargs...)
}

// DescribeAppBlocks mocks base method.
func (m *MockAppstreamClient) DescribeAppBlocks(arg0 context.Context, arg1 *appstream.DescribeAppBlocksInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeAppBlocksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeAppBlocks", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeAppBlocksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAppBlocks indicates an expected call of DescribeAppBlocks.
func (mr *MockAppstreamClientMockRecorder) DescribeAppBlocks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAppBlocks", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeAppBlocks), varargs...)
}

// DescribeApplicationFleetAssociations mocks base method.
func (m *MockAppstreamClient) DescribeApplicationFleetAssociations(arg0 context.Context, arg1 *appstream.DescribeApplicationFleetAssociationsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeApplicationFleetAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApplicationFleetAssociations", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeApplicationFleetAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeApplicationFleetAssociations indicates an expected call of DescribeApplicationFleetAssociations.
func (mr *MockAppstreamClientMockRecorder) DescribeApplicationFleetAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplicationFleetAssociations", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeApplicationFleetAssociations), varargs...)
}

// DescribeApplications mocks base method.
func (m *MockAppstreamClient) DescribeApplications(arg0 context.Context, arg1 *appstream.DescribeApplicationsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeApplications", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeApplications indicates an expected call of DescribeApplications.
func (mr *MockAppstreamClientMockRecorder) DescribeApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeApplications", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeApplications), varargs...)
}

// DescribeDirectoryConfigs mocks base method.
func (m *MockAppstreamClient) DescribeDirectoryConfigs(arg0 context.Context, arg1 *appstream.DescribeDirectoryConfigsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeDirectoryConfigsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDirectoryConfigs", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeDirectoryConfigsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDirectoryConfigs indicates an expected call of DescribeDirectoryConfigs.
func (mr *MockAppstreamClientMockRecorder) DescribeDirectoryConfigs(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDirectoryConfigs", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeDirectoryConfigs), varargs...)
}

// DescribeEntitlements mocks base method.
func (m *MockAppstreamClient) DescribeEntitlements(arg0 context.Context, arg1 *appstream.DescribeEntitlementsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeEntitlementsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeEntitlements", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeEntitlementsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntitlements indicates an expected call of DescribeEntitlements.
func (mr *MockAppstreamClientMockRecorder) DescribeEntitlements(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntitlements", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeEntitlements), varargs...)
}

// DescribeFleets mocks base method.
func (m *MockAppstreamClient) DescribeFleets(arg0 context.Context, arg1 *appstream.DescribeFleetsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeFleets", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeFleets indicates an expected call of DescribeFleets.
func (mr *MockAppstreamClientMockRecorder) DescribeFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeFleets", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeFleets), varargs...)
}

// DescribeImageBuilders mocks base method.
func (m *MockAppstreamClient) DescribeImageBuilders(arg0 context.Context, arg1 *appstream.DescribeImageBuildersInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeImageBuildersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImageBuilders", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeImageBuildersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImageBuilders indicates an expected call of DescribeImageBuilders.
func (mr *MockAppstreamClientMockRecorder) DescribeImageBuilders(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImageBuilders", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeImageBuilders), varargs...)
}

// DescribeImagePermissions mocks base method.
func (m *MockAppstreamClient) DescribeImagePermissions(arg0 context.Context, arg1 *appstream.DescribeImagePermissionsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeImagePermissionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImagePermissions", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeImagePermissionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImagePermissions indicates an expected call of DescribeImagePermissions.
func (mr *MockAppstreamClientMockRecorder) DescribeImagePermissions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImagePermissions", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeImagePermissions), varargs...)
}

// DescribeImages mocks base method.
func (m *MockAppstreamClient) DescribeImages(arg0 context.Context, arg1 *appstream.DescribeImagesInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeImagesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeImages", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeImagesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeImages indicates an expected call of DescribeImages.
func (mr *MockAppstreamClientMockRecorder) DescribeImages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeImages", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeImages), varargs...)
}

// DescribeSessions mocks base method.
func (m *MockAppstreamClient) DescribeSessions(arg0 context.Context, arg1 *appstream.DescribeSessionsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeSessionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeSessions", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeSessionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSessions indicates an expected call of DescribeSessions.
func (mr *MockAppstreamClientMockRecorder) DescribeSessions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSessions", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeSessions), varargs...)
}

// DescribeStacks mocks base method.
func (m *MockAppstreamClient) DescribeStacks(arg0 context.Context, arg1 *appstream.DescribeStacksInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStacks", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStacks indicates an expected call of DescribeStacks.
func (mr *MockAppstreamClientMockRecorder) DescribeStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStacks", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeStacks), varargs...)
}

// DescribeUsageReportSubscriptions mocks base method.
func (m *MockAppstreamClient) DescribeUsageReportSubscriptions(arg0 context.Context, arg1 *appstream.DescribeUsageReportSubscriptionsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUsageReportSubscriptions", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeUsageReportSubscriptionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUsageReportSubscriptions indicates an expected call of DescribeUsageReportSubscriptions.
func (mr *MockAppstreamClientMockRecorder) DescribeUsageReportSubscriptions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUsageReportSubscriptions", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeUsageReportSubscriptions), varargs...)
}

// DescribeUserStackAssociations mocks base method.
func (m *MockAppstreamClient) DescribeUserStackAssociations(arg0 context.Context, arg1 *appstream.DescribeUserStackAssociationsInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeUserStackAssociationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUserStackAssociations", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeUserStackAssociationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUserStackAssociations indicates an expected call of DescribeUserStackAssociations.
func (mr *MockAppstreamClientMockRecorder) DescribeUserStackAssociations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUserStackAssociations", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeUserStackAssociations), varargs...)
}

// DescribeUsers mocks base method.
func (m *MockAppstreamClient) DescribeUsers(arg0 context.Context, arg1 *appstream.DescribeUsersInput, arg2 ...func(*appstream.Options)) (*appstream.DescribeUsersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeUsers", varargs...)
	ret0, _ := ret[0].(*appstream.DescribeUsersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeUsers indicates an expected call of DescribeUsers.
func (mr *MockAppstreamClientMockRecorder) DescribeUsers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeUsers", reflect.TypeOf((*MockAppstreamClient)(nil).DescribeUsers), varargs...)
}

// ListAssociatedFleets mocks base method.
func (m *MockAppstreamClient) ListAssociatedFleets(arg0 context.Context, arg1 *appstream.ListAssociatedFleetsInput, arg2 ...func(*appstream.Options)) (*appstream.ListAssociatedFleetsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssociatedFleets", varargs...)
	ret0, _ := ret[0].(*appstream.ListAssociatedFleetsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssociatedFleets indicates an expected call of ListAssociatedFleets.
func (mr *MockAppstreamClientMockRecorder) ListAssociatedFleets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssociatedFleets", reflect.TypeOf((*MockAppstreamClient)(nil).ListAssociatedFleets), varargs...)
}

// ListAssociatedStacks mocks base method.
func (m *MockAppstreamClient) ListAssociatedStacks(arg0 context.Context, arg1 *appstream.ListAssociatedStacksInput, arg2 ...func(*appstream.Options)) (*appstream.ListAssociatedStacksOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListAssociatedStacks", varargs...)
	ret0, _ := ret[0].(*appstream.ListAssociatedStacksOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAssociatedStacks indicates an expected call of ListAssociatedStacks.
func (mr *MockAppstreamClientMockRecorder) ListAssociatedStacks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAssociatedStacks", reflect.TypeOf((*MockAppstreamClient)(nil).ListAssociatedStacks), varargs...)
}

// ListEntitledApplications mocks base method.
func (m *MockAppstreamClient) ListEntitledApplications(arg0 context.Context, arg1 *appstream.ListEntitledApplicationsInput, arg2 ...func(*appstream.Options)) (*appstream.ListEntitledApplicationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEntitledApplications", varargs...)
	ret0, _ := ret[0].(*appstream.ListEntitledApplicationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntitledApplications indicates an expected call of ListEntitledApplications.
func (mr *MockAppstreamClientMockRecorder) ListEntitledApplications(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntitledApplications", reflect.TypeOf((*MockAppstreamClient)(nil).ListEntitledApplications), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAppstreamClient) ListTagsForResource(arg0 context.Context, arg1 *appstream.ListTagsForResourceInput, arg2 ...func(*appstream.Options)) (*appstream.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*appstream.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAppstreamClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppstreamClient)(nil).ListTagsForResource), varargs...)
}
