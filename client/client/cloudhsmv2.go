// Code generated by MockGen. DO NOT EDIT.
// Source: cloudhsmv2.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	cloudhsmv2 "github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	gomock "github.com/golang/mock/gomock"
)

// MockCloudhsmv2Client is a mock of Cloudhsmv2Client interface.
type MockCloudhsmv2Client struct {
	ctrl     *gomock.Controller
	recorder *MockCloudhsmv2ClientMockRecorder
}

// MockCloudhsmv2ClientMockRecorder is the mock recorder for MockCloudhsmv2Client.
type MockCloudhsmv2ClientMockRecorder struct {
	mock *MockCloudhsmv2Client
}

// NewMockCloudhsmv2Client creates a new mock instance.
func NewMockCloudhsmv2Client(ctrl *gomock.Controller) *MockCloudhsmv2Client {
	mock := &MockCloudhsmv2Client{ctrl: ctrl}
	mock.recorder = &MockCloudhsmv2ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloudhsmv2Client) EXPECT() *MockCloudhsmv2ClientMockRecorder {
	return m.recorder
}

// DescribeBackups mocks base method.
func (m *MockCloudhsmv2Client) DescribeBackups(arg0 context.Context, arg1 *cloudhsmv2.DescribeBackupsInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeBackupsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudhsmv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeBackups")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeBackups", varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.DescribeBackupsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeBackups indicates an expected call of DescribeBackups.
func (mr *MockCloudhsmv2ClientMockRecorder) DescribeBackups(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeBackups", reflect.TypeOf((*MockCloudhsmv2Client)(nil).DescribeBackups), varargs...)
}

// DescribeClusters mocks base method.
func (m *MockCloudhsmv2Client) DescribeClusters(arg0 context.Context, arg1 *cloudhsmv2.DescribeClustersInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.DescribeClustersOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudhsmv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to DescribeClusters")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeClusters", varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.DescribeClustersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeClusters indicates an expected call of DescribeClusters.
func (mr *MockCloudhsmv2ClientMockRecorder) DescribeClusters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeClusters", reflect.TypeOf((*MockCloudhsmv2Client)(nil).DescribeClusters), varargs...)
}

// ListTags mocks base method.
func (m *MockCloudhsmv2Client) ListTags(arg0 context.Context, arg1 *cloudhsmv2.ListTagsInput, arg2 ...func(*cloudhsmv2.Options)) (*cloudhsmv2.ListTagsOutput, error) {

	// Assertion inserted by client/mockgen/main.go
	o := &cloudhsmv2.Options{}
	for _, f := range arg2 {
		f(o)
	}
	if o.Region == "" {
		m.ctrl.T.Errorf("Region not set in call to ListTags")
	}

	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTags", varargs...)
	ret0, _ := ret[0].(*cloudhsmv2.ListTagsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTags indicates an expected call of ListTags.
func (mr *MockCloudhsmv2ClientMockRecorder) ListTags(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTags", reflect.TypeOf((*MockCloudhsmv2Client)(nil).ListTags), varargs...)
}
