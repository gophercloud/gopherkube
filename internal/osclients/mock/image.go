/*
Copyright 2024 The ORC Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by MockGen. DO NOT EDIT.
// Source: ../image.go
//
// Generated by this command:
//
//	mockgen -package mock -destination=image.go -source=../image.go github.com/k-orc/openstack-resource-controller/internal/osclients/mock ImageClient
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	io "io"
	iter "iter"
	reflect "reflect"

	imageimport "github.com/gophercloud/gophercloud/v2/openstack/image/v2/imageimport"
	images "github.com/gophercloud/gophercloud/v2/openstack/image/v2/images"
	gomock "go.uber.org/mock/gomock"
)

// MockImageClient is a mock of ImageClient interface.
type MockImageClient struct {
	ctrl     *gomock.Controller
	recorder *MockImageClientMockRecorder
}

// MockImageClientMockRecorder is the mock recorder for MockImageClient.
type MockImageClientMockRecorder struct {
	mock *MockImageClient
}

// NewMockImageClient creates a new mock instance.
func NewMockImageClient(ctrl *gomock.Controller) *MockImageClient {
	mock := &MockImageClient{ctrl: ctrl}
	mock.recorder = &MockImageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageClient) EXPECT() *MockImageClientMockRecorder {
	return m.recorder
}

// CreateImage mocks base method.
func (m *MockImageClient) CreateImage(ctx context.Context, createOpts images.CreateOptsBuilder) (*images.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImage", ctx, createOpts)
	ret0, _ := ret[0].(*images.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateImage indicates an expected call of CreateImage.
func (mr *MockImageClientMockRecorder) CreateImage(ctx, createOpts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImage", reflect.TypeOf((*MockImageClient)(nil).CreateImage), ctx, createOpts)
}

// CreateImport mocks base method.
func (m *MockImageClient) CreateImport(ctx context.Context, id string, createOpts imageimport.CreateOptsBuilder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateImport", ctx, id, createOpts)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateImport indicates an expected call of CreateImport.
func (mr *MockImageClientMockRecorder) CreateImport(ctx, id, createOpts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateImport", reflect.TypeOf((*MockImageClient)(nil).CreateImport), ctx, id, createOpts)
}

// DeleteImage mocks base method.
func (m *MockImageClient) DeleteImage(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage.
func (mr *MockImageClientMockRecorder) DeleteImage(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockImageClient)(nil).DeleteImage), ctx, id)
}

// GetImage mocks base method.
func (m *MockImageClient) GetImage(ctx context.Context, id string) (*images.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", ctx, id)
	ret0, _ := ret[0].(*images.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImage indicates an expected call of GetImage.
func (mr *MockImageClientMockRecorder) GetImage(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockImageClient)(nil).GetImage), ctx, id)
}

// GetImportInfo mocks base method.
func (m *MockImageClient) GetImportInfo(ctx context.Context) (*imageimport.ImportInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImportInfo", ctx)
	ret0, _ := ret[0].(*imageimport.ImportInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImportInfo indicates an expected call of GetImportInfo.
func (mr *MockImageClientMockRecorder) GetImportInfo(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImportInfo", reflect.TypeOf((*MockImageClient)(nil).GetImportInfo), ctx)
}

// ListImages mocks base method.
func (m *MockImageClient) ListImages(ctx context.Context, listOpts images.ListOptsBuilder) iter.Seq2[*images.Image, error] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", ctx, listOpts)
	ret0, _ := ret[0].(iter.Seq2[*images.Image, error])
	return ret0
}

// ListImages indicates an expected call of ListImages.
func (mr *MockImageClientMockRecorder) ListImages(ctx, listOpts any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockImageClient)(nil).ListImages), ctx, listOpts)
}

// UploadData mocks base method.
func (m *MockImageClient) UploadData(ctx context.Context, id string, data io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadData", ctx, id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadData indicates an expected call of UploadData.
func (mr *MockImageClientMockRecorder) UploadData(ctx, id, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadData", reflect.TypeOf((*MockImageClient)(nil).UploadData), ctx, id, data)
}
