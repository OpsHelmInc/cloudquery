// Code generated by MockGen. DO NOT EDIT.
// Source: s3.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

// MockS3Client is a mock of S3Client interface.
type MockS3Client struct {
	ctrl     *gomock.Controller
	recorder *MockS3ClientMockRecorder
}

// MockS3ClientMockRecorder is the mock recorder for MockS3Client.
type MockS3ClientMockRecorder struct {
	mock *MockS3Client
}

// NewMockS3Client creates a new mock instance.
func NewMockS3Client(ctrl *gomock.Controller) *MockS3Client {
	mock := &MockS3Client{ctrl: ctrl}
	mock.recorder = &MockS3ClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3Client) EXPECT() *MockS3ClientMockRecorder {
	return m.recorder
}

// GetBucketAccelerateConfiguration mocks base method.
func (m *MockS3Client) GetBucketAccelerateConfiguration(arg0 context.Context, arg1 *s3.GetBucketAccelerateConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAccelerateConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAccelerateConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAccelerateConfiguration indicates an expected call of GetBucketAccelerateConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketAccelerateConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAccelerateConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketAccelerateConfiguration), varargs...)
}

// GetBucketAcl mocks base method.
func (m *MockS3Client) GetBucketAcl(arg0 context.Context, arg1 *s3.GetBucketAclInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAclOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAcl", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAclOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAcl indicates an expected call of GetBucketAcl.
func (mr *MockS3ClientMockRecorder) GetBucketAcl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAcl", reflect.TypeOf((*MockS3Client)(nil).GetBucketAcl), varargs...)
}

// GetBucketAnalyticsConfiguration mocks base method.
func (m *MockS3Client) GetBucketAnalyticsConfiguration(arg0 context.Context, arg1 *s3.GetBucketAnalyticsConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAnalyticsConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketAnalyticsConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAnalyticsConfiguration indicates an expected call of GetBucketAnalyticsConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketAnalyticsConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAnalyticsConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketAnalyticsConfiguration), varargs...)
}

// GetBucketCors mocks base method.
func (m *MockS3Client) GetBucketCors(arg0 context.Context, arg1 *s3.GetBucketCorsInput, arg2 ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketCors", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketCorsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketCors indicates an expected call of GetBucketCors.
func (mr *MockS3ClientMockRecorder) GetBucketCors(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketCors", reflect.TypeOf((*MockS3Client)(nil).GetBucketCors), varargs...)
}

// GetBucketEncryption mocks base method.
func (m *MockS3Client) GetBucketEncryption(arg0 context.Context, arg1 *s3.GetBucketEncryptionInput, arg2 ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketEncryption", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketEncryption indicates an expected call of GetBucketEncryption.
func (mr *MockS3ClientMockRecorder) GetBucketEncryption(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketEncryption", reflect.TypeOf((*MockS3Client)(nil).GetBucketEncryption), varargs...)
}

// GetBucketIntelligentTieringConfiguration mocks base method.
func (m *MockS3Client) GetBucketIntelligentTieringConfiguration(arg0 context.Context, arg1 *s3.GetBucketIntelligentTieringConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketIntelligentTieringConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketIntelligentTieringConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketIntelligentTieringConfiguration indicates an expected call of GetBucketIntelligentTieringConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketIntelligentTieringConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketIntelligentTieringConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketIntelligentTieringConfiguration), varargs...)
}

// GetBucketInventoryConfiguration mocks base method.
func (m *MockS3Client) GetBucketInventoryConfiguration(arg0 context.Context, arg1 *s3.GetBucketInventoryConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketInventoryConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketInventoryConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketInventoryConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketInventoryConfiguration indicates an expected call of GetBucketInventoryConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketInventoryConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketInventoryConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketInventoryConfiguration), varargs...)
}

// GetBucketLifecycleConfiguration mocks base method.
func (m *MockS3Client) GetBucketLifecycleConfiguration(arg0 context.Context, arg1 *s3.GetBucketLifecycleConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLifecycleConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLifecycleConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLifecycleConfiguration indicates an expected call of GetBucketLifecycleConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketLifecycleConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLifecycleConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketLifecycleConfiguration), varargs...)
}

// GetBucketLocation mocks base method.
func (m *MockS3Client) GetBucketLocation(arg0 context.Context, arg1 *s3.GetBucketLocationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLocation", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLocationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLocation indicates an expected call of GetBucketLocation.
func (mr *MockS3ClientMockRecorder) GetBucketLocation(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLocation", reflect.TypeOf((*MockS3Client)(nil).GetBucketLocation), varargs...)
}

// GetBucketLogging mocks base method.
func (m *MockS3Client) GetBucketLogging(arg0 context.Context, arg1 *s3.GetBucketLoggingInput, arg2 ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketLogging", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketLoggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketLogging indicates an expected call of GetBucketLogging.
func (mr *MockS3ClientMockRecorder) GetBucketLogging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketLogging", reflect.TypeOf((*MockS3Client)(nil).GetBucketLogging), varargs...)
}

// GetBucketMetadataTableConfiguration mocks base method.
func (m *MockS3Client) GetBucketMetadataTableConfiguration(arg0 context.Context, arg1 *s3.GetBucketMetadataTableConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketMetadataTableConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketMetadataTableConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketMetadataTableConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketMetadataTableConfiguration indicates an expected call of GetBucketMetadataTableConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketMetadataTableConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketMetadataTableConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketMetadataTableConfiguration), varargs...)
}

// GetBucketMetricsConfiguration mocks base method.
func (m *MockS3Client) GetBucketMetricsConfiguration(arg0 context.Context, arg1 *s3.GetBucketMetricsConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketMetricsConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketMetricsConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketMetricsConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketMetricsConfiguration indicates an expected call of GetBucketMetricsConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketMetricsConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketMetricsConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketMetricsConfiguration), varargs...)
}

// GetBucketNotificationConfiguration mocks base method.
func (m *MockS3Client) GetBucketNotificationConfiguration(arg0 context.Context, arg1 *s3.GetBucketNotificationConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketNotificationConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketNotificationConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketNotificationConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketNotificationConfiguration indicates an expected call of GetBucketNotificationConfiguration.
func (mr *MockS3ClientMockRecorder) GetBucketNotificationConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketNotificationConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetBucketNotificationConfiguration), varargs...)
}

// GetBucketOwnershipControls mocks base method.
func (m *MockS3Client) GetBucketOwnershipControls(arg0 context.Context, arg1 *s3.GetBucketOwnershipControlsInput, arg2 ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketOwnershipControls", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketOwnershipControlsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketOwnershipControls indicates an expected call of GetBucketOwnershipControls.
func (mr *MockS3ClientMockRecorder) GetBucketOwnershipControls(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketOwnershipControls", reflect.TypeOf((*MockS3Client)(nil).GetBucketOwnershipControls), varargs...)
}

// GetBucketPolicy mocks base method.
func (m *MockS3Client) GetBucketPolicy(arg0 context.Context, arg1 *s3.GetBucketPolicyInput, arg2 ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketPolicy", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketPolicy indicates an expected call of GetBucketPolicy.
func (mr *MockS3ClientMockRecorder) GetBucketPolicy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketPolicy", reflect.TypeOf((*MockS3Client)(nil).GetBucketPolicy), varargs...)
}

// GetBucketPolicyStatus mocks base method.
func (m *MockS3Client) GetBucketPolicyStatus(arg0 context.Context, arg1 *s3.GetBucketPolicyStatusInput, arg2 ...func(*s3.Options)) (*s3.GetBucketPolicyStatusOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketPolicyStatus", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketPolicyStatusOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketPolicyStatus indicates an expected call of GetBucketPolicyStatus.
func (mr *MockS3ClientMockRecorder) GetBucketPolicyStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketPolicyStatus", reflect.TypeOf((*MockS3Client)(nil).GetBucketPolicyStatus), varargs...)
}

// GetBucketReplication mocks base method.
func (m *MockS3Client) GetBucketReplication(arg0 context.Context, arg1 *s3.GetBucketReplicationInput, arg2 ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketReplication", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketReplicationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketReplication indicates an expected call of GetBucketReplication.
func (mr *MockS3ClientMockRecorder) GetBucketReplication(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketReplication", reflect.TypeOf((*MockS3Client)(nil).GetBucketReplication), varargs...)
}

// GetBucketRequestPayment mocks base method.
func (m *MockS3Client) GetBucketRequestPayment(arg0 context.Context, arg1 *s3.GetBucketRequestPaymentInput, arg2 ...func(*s3.Options)) (*s3.GetBucketRequestPaymentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketRequestPayment", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketRequestPaymentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketRequestPayment indicates an expected call of GetBucketRequestPayment.
func (mr *MockS3ClientMockRecorder) GetBucketRequestPayment(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketRequestPayment", reflect.TypeOf((*MockS3Client)(nil).GetBucketRequestPayment), varargs...)
}

// GetBucketTagging mocks base method.
func (m *MockS3Client) GetBucketTagging(arg0 context.Context, arg1 *s3.GetBucketTaggingInput, arg2 ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketTagging", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketTagging indicates an expected call of GetBucketTagging.
func (mr *MockS3ClientMockRecorder) GetBucketTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketTagging", reflect.TypeOf((*MockS3Client)(nil).GetBucketTagging), varargs...)
}

// GetBucketVersioning mocks base method.
func (m *MockS3Client) GetBucketVersioning(arg0 context.Context, arg1 *s3.GetBucketVersioningInput, arg2 ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketVersioning", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketVersioningOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketVersioning indicates an expected call of GetBucketVersioning.
func (mr *MockS3ClientMockRecorder) GetBucketVersioning(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketVersioning", reflect.TypeOf((*MockS3Client)(nil).GetBucketVersioning), varargs...)
}

// GetBucketWebsite mocks base method.
func (m *MockS3Client) GetBucketWebsite(arg0 context.Context, arg1 *s3.GetBucketWebsiteInput, arg2 ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketWebsite", varargs...)
	ret0, _ := ret[0].(*s3.GetBucketWebsiteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketWebsite indicates an expected call of GetBucketWebsite.
func (mr *MockS3ClientMockRecorder) GetBucketWebsite(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketWebsite", reflect.TypeOf((*MockS3Client)(nil).GetBucketWebsite), varargs...)
}

// GetObject mocks base method.
func (m *MockS3Client) GetObject(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3ClientMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3Client)(nil).GetObject), varargs...)
}

// GetObjectAcl mocks base method.
func (m *MockS3Client) GetObjectAcl(arg0 context.Context, arg1 *s3.GetObjectAclInput, arg2 ...func(*s3.Options)) (*s3.GetObjectAclOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectAcl", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectAclOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectAcl indicates an expected call of GetObjectAcl.
func (mr *MockS3ClientMockRecorder) GetObjectAcl(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectAcl", reflect.TypeOf((*MockS3Client)(nil).GetObjectAcl), varargs...)
}

// GetObjectAttributes mocks base method.
func (m *MockS3Client) GetObjectAttributes(arg0 context.Context, arg1 *s3.GetObjectAttributesInput, arg2 ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectAttributes", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectAttributesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectAttributes indicates an expected call of GetObjectAttributes.
func (mr *MockS3ClientMockRecorder) GetObjectAttributes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectAttributes", reflect.TypeOf((*MockS3Client)(nil).GetObjectAttributes), varargs...)
}

// GetObjectLegalHold mocks base method.
func (m *MockS3Client) GetObjectLegalHold(arg0 context.Context, arg1 *s3.GetObjectLegalHoldInput, arg2 ...func(*s3.Options)) (*s3.GetObjectLegalHoldOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectLegalHold", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectLegalHoldOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectLegalHold indicates an expected call of GetObjectLegalHold.
func (mr *MockS3ClientMockRecorder) GetObjectLegalHold(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectLegalHold", reflect.TypeOf((*MockS3Client)(nil).GetObjectLegalHold), varargs...)
}

// GetObjectLockConfiguration mocks base method.
func (m *MockS3Client) GetObjectLockConfiguration(arg0 context.Context, arg1 *s3.GetObjectLockConfigurationInput, arg2 ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectLockConfiguration", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectLockConfigurationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectLockConfiguration indicates an expected call of GetObjectLockConfiguration.
func (mr *MockS3ClientMockRecorder) GetObjectLockConfiguration(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectLockConfiguration", reflect.TypeOf((*MockS3Client)(nil).GetObjectLockConfiguration), varargs...)
}

// GetObjectRetention mocks base method.
func (m *MockS3Client) GetObjectRetention(arg0 context.Context, arg1 *s3.GetObjectRetentionInput, arg2 ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectRetention", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectRetentionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectRetention indicates an expected call of GetObjectRetention.
func (mr *MockS3ClientMockRecorder) GetObjectRetention(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectRetention", reflect.TypeOf((*MockS3Client)(nil).GetObjectRetention), varargs...)
}

// GetObjectTagging mocks base method.
func (m *MockS3Client) GetObjectTagging(arg0 context.Context, arg1 *s3.GetObjectTaggingInput, arg2 ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectTagging", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectTaggingOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectTagging indicates an expected call of GetObjectTagging.
func (mr *MockS3ClientMockRecorder) GetObjectTagging(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectTagging", reflect.TypeOf((*MockS3Client)(nil).GetObjectTagging), varargs...)
}

// GetObjectTorrent mocks base method.
func (m *MockS3Client) GetObjectTorrent(arg0 context.Context, arg1 *s3.GetObjectTorrentInput, arg2 ...func(*s3.Options)) (*s3.GetObjectTorrentOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObjectTorrent", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectTorrentOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObjectTorrent indicates an expected call of GetObjectTorrent.
func (mr *MockS3ClientMockRecorder) GetObjectTorrent(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectTorrent", reflect.TypeOf((*MockS3Client)(nil).GetObjectTorrent), varargs...)
}

// GetPublicAccessBlock mocks base method.
func (m *MockS3Client) GetPublicAccessBlock(arg0 context.Context, arg1 *s3.GetPublicAccessBlockInput, arg2 ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicAccessBlock", varargs...)
	ret0, _ := ret[0].(*s3.GetPublicAccessBlockOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicAccessBlock indicates an expected call of GetPublicAccessBlock.
func (mr *MockS3ClientMockRecorder) GetPublicAccessBlock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicAccessBlock", reflect.TypeOf((*MockS3Client)(nil).GetPublicAccessBlock), varargs...)
}

// ListBucketAnalyticsConfigurations mocks base method.
func (m *MockS3Client) ListBucketAnalyticsConfigurations(arg0 context.Context, arg1 *s3.ListBucketAnalyticsConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketAnalyticsConfigurations", varargs...)
	ret0, _ := ret[0].(*s3.ListBucketAnalyticsConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketAnalyticsConfigurations indicates an expected call of ListBucketAnalyticsConfigurations.
func (mr *MockS3ClientMockRecorder) ListBucketAnalyticsConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketAnalyticsConfigurations", reflect.TypeOf((*MockS3Client)(nil).ListBucketAnalyticsConfigurations), varargs...)
}

// ListBucketIntelligentTieringConfigurations mocks base method.
func (m *MockS3Client) ListBucketIntelligentTieringConfigurations(arg0 context.Context, arg1 *s3.ListBucketIntelligentTieringConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketIntelligentTieringConfigurations", varargs...)
	ret0, _ := ret[0].(*s3.ListBucketIntelligentTieringConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketIntelligentTieringConfigurations indicates an expected call of ListBucketIntelligentTieringConfigurations.
func (mr *MockS3ClientMockRecorder) ListBucketIntelligentTieringConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketIntelligentTieringConfigurations", reflect.TypeOf((*MockS3Client)(nil).ListBucketIntelligentTieringConfigurations), varargs...)
}

// ListBucketInventoryConfigurations mocks base method.
func (m *MockS3Client) ListBucketInventoryConfigurations(arg0 context.Context, arg1 *s3.ListBucketInventoryConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketInventoryConfigurations", varargs...)
	ret0, _ := ret[0].(*s3.ListBucketInventoryConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketInventoryConfigurations indicates an expected call of ListBucketInventoryConfigurations.
func (mr *MockS3ClientMockRecorder) ListBucketInventoryConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketInventoryConfigurations", reflect.TypeOf((*MockS3Client)(nil).ListBucketInventoryConfigurations), varargs...)
}

// ListBucketMetricsConfigurations mocks base method.
func (m *MockS3Client) ListBucketMetricsConfigurations(arg0 context.Context, arg1 *s3.ListBucketMetricsConfigurationsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBucketMetricsConfigurations", varargs...)
	ret0, _ := ret[0].(*s3.ListBucketMetricsConfigurationsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBucketMetricsConfigurations indicates an expected call of ListBucketMetricsConfigurations.
func (mr *MockS3ClientMockRecorder) ListBucketMetricsConfigurations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBucketMetricsConfigurations", reflect.TypeOf((*MockS3Client)(nil).ListBucketMetricsConfigurations), varargs...)
}

// ListBuckets mocks base method.
func (m *MockS3Client) ListBuckets(arg0 context.Context, arg1 *s3.ListBucketsInput, arg2 ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListBuckets", varargs...)
	ret0, _ := ret[0].(*s3.ListBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuckets indicates an expected call of ListBuckets.
func (mr *MockS3ClientMockRecorder) ListBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuckets", reflect.TypeOf((*MockS3Client)(nil).ListBuckets), varargs...)
}

// ListDirectoryBuckets mocks base method.
func (m *MockS3Client) ListDirectoryBuckets(arg0 context.Context, arg1 *s3.ListDirectoryBucketsInput, arg2 ...func(*s3.Options)) (*s3.ListDirectoryBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDirectoryBuckets", varargs...)
	ret0, _ := ret[0].(*s3.ListDirectoryBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDirectoryBuckets indicates an expected call of ListDirectoryBuckets.
func (mr *MockS3ClientMockRecorder) ListDirectoryBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDirectoryBuckets", reflect.TypeOf((*MockS3Client)(nil).ListDirectoryBuckets), varargs...)
}

// ListMultipartUploads mocks base method.
func (m *MockS3Client) ListMultipartUploads(arg0 context.Context, arg1 *s3.ListMultipartUploadsInput, arg2 ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMultipartUploads", varargs...)
	ret0, _ := ret[0].(*s3.ListMultipartUploadsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMultipartUploads indicates an expected call of ListMultipartUploads.
func (mr *MockS3ClientMockRecorder) ListMultipartUploads(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMultipartUploads", reflect.TypeOf((*MockS3Client)(nil).ListMultipartUploads), varargs...)
}

// ListObjectVersions mocks base method.
func (m *MockS3Client) ListObjectVersions(arg0 context.Context, arg1 *s3.ListObjectVersionsInput, arg2 ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjectVersions", varargs...)
	ret0, _ := ret[0].(*s3.ListObjectVersionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectVersions indicates an expected call of ListObjectVersions.
func (mr *MockS3ClientMockRecorder) ListObjectVersions(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectVersions", reflect.TypeOf((*MockS3Client)(nil).ListObjectVersions), varargs...)
}

// ListObjects mocks base method.
func (m *MockS3Client) ListObjects(arg0 context.Context, arg1 *s3.ListObjectsInput, arg2 ...func(*s3.Options)) (*s3.ListObjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjects", varargs...)
	ret0, _ := ret[0].(*s3.ListObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects.
func (mr *MockS3ClientMockRecorder) ListObjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockS3Client)(nil).ListObjects), varargs...)
}

// ListObjectsV2 mocks base method.
func (m *MockS3Client) ListObjectsV2(arg0 context.Context, arg1 *s3.ListObjectsV2Input, arg2 ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListObjectsV2", varargs...)
	ret0, _ := ret[0].(*s3.ListObjectsV2Output)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjectsV2 indicates an expected call of ListObjectsV2.
func (mr *MockS3ClientMockRecorder) ListObjectsV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjectsV2", reflect.TypeOf((*MockS3Client)(nil).ListObjectsV2), varargs...)
}

// ListParts mocks base method.
func (m *MockS3Client) ListParts(arg0 context.Context, arg1 *s3.ListPartsInput, arg2 ...func(*s3.Options)) (*s3.ListPartsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListParts", varargs...)
	ret0, _ := ret[0].(*s3.ListPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListParts indicates an expected call of ListParts.
func (mr *MockS3ClientMockRecorder) ListParts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListParts", reflect.TypeOf((*MockS3Client)(nil).ListParts), varargs...)
}
