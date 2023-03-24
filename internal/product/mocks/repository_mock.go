// Code generated by MockGen. DO NOT EDIT.
// Source: internal/product/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "mceasy/service-demo/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetProducts mocks base method.
func (m *MockRepository) GetProducts(ctx context.Context) ([]*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", ctx)
	ret0, _ := ret[0].([]*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockRepositoryMockRecorder) GetProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockRepository)(nil).GetProducts), ctx)
}

// StoreProduct mocks base method.
func (m *MockRepository) StoreProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreProduct", ctx, product)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreProduct indicates an expected call of StoreProduct.
func (mr *MockRepositoryMockRecorder) StoreProduct(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreProduct", reflect.TypeOf((*MockRepository)(nil).StoreProduct), ctx, product)
}
