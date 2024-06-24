// Code generated by MockGen. DO NOT EDIT.
// Source: appmesh.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	appmesh "github.com/aws/aws-sdk-go-v2/service/appmesh"
	gomock "github.com/golang/mock/gomock"
)

// MockAppmeshClient is a mock of AppmeshClient interface.
type MockAppmeshClient struct {
	ctrl     *gomock.Controller
	recorder *MockAppmeshClientMockRecorder
}

// MockAppmeshClientMockRecorder is the mock recorder for MockAppmeshClient.
type MockAppmeshClientMockRecorder struct {
	mock *MockAppmeshClient
}

// NewMockAppmeshClient creates a new mock instance.
func NewMockAppmeshClient(ctrl *gomock.Controller) *MockAppmeshClient {
	mock := &MockAppmeshClient{ctrl: ctrl}
	mock.recorder = &MockAppmeshClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppmeshClient) EXPECT() *MockAppmeshClientMockRecorder {
	return m.recorder
}

// DescribeGatewayRoute mocks base method.
func (m *MockAppmeshClient) DescribeGatewayRoute(arg0 context.Context, arg1 *appmesh.DescribeGatewayRouteInput, arg2 ...func(*appmesh.Options)) (*appmesh.DescribeGatewayRouteOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeGatewayRoute", varargs...)
	ret0, _ := ret[0].(*appmesh.DescribeGatewayRouteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeGatewayRoute indicates an expected call of DescribeGatewayRoute.
func (mr *MockAppmeshClientMockRecorder) DescribeGatewayRoute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeGatewayRoute", reflect.TypeOf((*MockAppmeshClient)(nil).DescribeGatewayRoute), varargs...)
}

// DescribeMesh mocks base method.
func (m *MockAppmeshClient) DescribeMesh(arg0 context.Context, arg1 *appmesh.DescribeMeshInput, arg2 ...func(*appmesh.Options)) (*appmesh.DescribeMeshOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMesh", varargs...)
	ret0, _ := ret[0].(*appmesh.DescribeMeshOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMesh indicates an expected call of DescribeMesh.
func (mr *MockAppmeshClientMockRecorder) DescribeMesh(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMesh", reflect.TypeOf((*MockAppmeshClient)(nil).DescribeMesh), varargs...)
}

// DescribeRoute mocks base method.
func (m *MockAppmeshClient) DescribeRoute(arg0 context.Context, arg1 *appmesh.DescribeRouteInput, arg2 ...func(*appmesh.Options)) (*appmesh.DescribeRouteOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeRoute", varargs...)
	ret0, _ := ret[0].(*appmesh.DescribeRouteOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRoute indicates an expected call of DescribeRoute.
func (mr *MockAppmeshClientMockRecorder) DescribeRoute(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRoute", reflect.TypeOf((*MockAppmeshClient)(nil).DescribeRoute), varargs...)
}

// DescribeVirtualGateway mocks base method.
func (m *MockAppmeshClient) DescribeVirtualGateway(arg0 context.Context, arg1 *appmesh.DescribeVirtualGatewayInput, arg2 ...func(*appmesh.Options)) (*appmesh.DescribeVirtualGatewayOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualGateway", varargs...)
	ret0, _ := ret[0].(*appmesh.DescribeVirtualGatewayOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualGateway indicates an expected call of DescribeVirtualGateway.
func (mr *MockAppmeshClientMockRecorder) DescribeVirtualGateway(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualGateway", reflect.TypeOf((*MockAppmeshClient)(nil).DescribeVirtualGateway), varargs...)
}

// DescribeVirtualNode mocks base method.
func (m *MockAppmeshClient) DescribeVirtualNode(arg0 context.Context, arg1 *appmesh.DescribeVirtualNodeInput, arg2 ...func(*appmesh.Options)) (*appmesh.DescribeVirtualNodeOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualNode", varargs...)
	ret0, _ := ret[0].(*appmesh.DescribeVirtualNodeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualNode indicates an expected call of DescribeVirtualNode.
func (mr *MockAppmeshClientMockRecorder) DescribeVirtualNode(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualNode", reflect.TypeOf((*MockAppmeshClient)(nil).DescribeVirtualNode), varargs...)
}

// DescribeVirtualRouter mocks base method.
func (m *MockAppmeshClient) DescribeVirtualRouter(arg0 context.Context, arg1 *appmesh.DescribeVirtualRouterInput, arg2 ...func(*appmesh.Options)) (*appmesh.DescribeVirtualRouterOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualRouter", varargs...)
	ret0, _ := ret[0].(*appmesh.DescribeVirtualRouterOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualRouter indicates an expected call of DescribeVirtualRouter.
func (mr *MockAppmeshClientMockRecorder) DescribeVirtualRouter(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualRouter", reflect.TypeOf((*MockAppmeshClient)(nil).DescribeVirtualRouter), varargs...)
}

// DescribeVirtualService mocks base method.
func (m *MockAppmeshClient) DescribeVirtualService(arg0 context.Context, arg1 *appmesh.DescribeVirtualServiceInput, arg2 ...func(*appmesh.Options)) (*appmesh.DescribeVirtualServiceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeVirtualService", varargs...)
	ret0, _ := ret[0].(*appmesh.DescribeVirtualServiceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeVirtualService indicates an expected call of DescribeVirtualService.
func (mr *MockAppmeshClientMockRecorder) DescribeVirtualService(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeVirtualService", reflect.TypeOf((*MockAppmeshClient)(nil).DescribeVirtualService), varargs...)
}

// ListGatewayRoutes mocks base method.
func (m *MockAppmeshClient) ListGatewayRoutes(arg0 context.Context, arg1 *appmesh.ListGatewayRoutesInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListGatewayRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListGatewayRoutes", varargs...)
	ret0, _ := ret[0].(*appmesh.ListGatewayRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGatewayRoutes indicates an expected call of ListGatewayRoutes.
func (mr *MockAppmeshClientMockRecorder) ListGatewayRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGatewayRoutes", reflect.TypeOf((*MockAppmeshClient)(nil).ListGatewayRoutes), varargs...)
}

// ListMeshes mocks base method.
func (m *MockAppmeshClient) ListMeshes(arg0 context.Context, arg1 *appmesh.ListMeshesInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListMeshesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListMeshes", varargs...)
	ret0, _ := ret[0].(*appmesh.ListMeshesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMeshes indicates an expected call of ListMeshes.
func (mr *MockAppmeshClientMockRecorder) ListMeshes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMeshes", reflect.TypeOf((*MockAppmeshClient)(nil).ListMeshes), varargs...)
}

// ListRoutes mocks base method.
func (m *MockAppmeshClient) ListRoutes(arg0 context.Context, arg1 *appmesh.ListRoutesInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListRoutesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListRoutes", varargs...)
	ret0, _ := ret[0].(*appmesh.ListRoutesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRoutes indicates an expected call of ListRoutes.
func (mr *MockAppmeshClientMockRecorder) ListRoutes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRoutes", reflect.TypeOf((*MockAppmeshClient)(nil).ListRoutes), varargs...)
}

// ListTagsForResource mocks base method.
func (m *MockAppmeshClient) ListTagsForResource(arg0 context.Context, arg1 *appmesh.ListTagsForResourceInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListTagsForResourceOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForResource", varargs...)
	ret0, _ := ret[0].(*appmesh.ListTagsForResourceOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForResource indicates an expected call of ListTagsForResource.
func (mr *MockAppmeshClientMockRecorder) ListTagsForResource(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForResource", reflect.TypeOf((*MockAppmeshClient)(nil).ListTagsForResource), varargs...)
}

// ListVirtualGateways mocks base method.
func (m *MockAppmeshClient) ListVirtualGateways(arg0 context.Context, arg1 *appmesh.ListVirtualGatewaysInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListVirtualGatewaysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualGateways", varargs...)
	ret0, _ := ret[0].(*appmesh.ListVirtualGatewaysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualGateways indicates an expected call of ListVirtualGateways.
func (mr *MockAppmeshClientMockRecorder) ListVirtualGateways(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualGateways", reflect.TypeOf((*MockAppmeshClient)(nil).ListVirtualGateways), varargs...)
}

// ListVirtualNodes mocks base method.
func (m *MockAppmeshClient) ListVirtualNodes(arg0 context.Context, arg1 *appmesh.ListVirtualNodesInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListVirtualNodesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualNodes", varargs...)
	ret0, _ := ret[0].(*appmesh.ListVirtualNodesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualNodes indicates an expected call of ListVirtualNodes.
func (mr *MockAppmeshClientMockRecorder) ListVirtualNodes(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualNodes", reflect.TypeOf((*MockAppmeshClient)(nil).ListVirtualNodes), varargs...)
}

// ListVirtualRouters mocks base method.
func (m *MockAppmeshClient) ListVirtualRouters(arg0 context.Context, arg1 *appmesh.ListVirtualRoutersInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListVirtualRoutersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualRouters", varargs...)
	ret0, _ := ret[0].(*appmesh.ListVirtualRoutersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualRouters indicates an expected call of ListVirtualRouters.
func (mr *MockAppmeshClientMockRecorder) ListVirtualRouters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualRouters", reflect.TypeOf((*MockAppmeshClient)(nil).ListVirtualRouters), varargs...)
}

// ListVirtualServices mocks base method.
func (m *MockAppmeshClient) ListVirtualServices(arg0 context.Context, arg1 *appmesh.ListVirtualServicesInput, arg2 ...func(*appmesh.Options)) (*appmesh.ListVirtualServicesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualServices", varargs...)
	ret0, _ := ret[0].(*appmesh.ListVirtualServicesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualServices indicates an expected call of ListVirtualServices.
func (mr *MockAppmeshClientMockRecorder) ListVirtualServices(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualServices", reflect.TypeOf((*MockAppmeshClient)(nil).ListVirtualServices), varargs...)
}
