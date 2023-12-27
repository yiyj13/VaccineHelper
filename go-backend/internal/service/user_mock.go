// Code generated by MockGen. DO NOT EDIT.
// Source: user.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"
	model "v-helper/internal/model"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// AddFollowingArticle mocks base method.
func (m *MockIUserService) AddFollowingArticle(userID, articleID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFollowingArticle", userID, articleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFollowingArticle indicates an expected call of AddFollowingArticle.
func (mr *MockIUserServiceMockRecorder) AddFollowingArticle(userID, articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFollowingArticle", reflect.TypeOf((*MockIUserService)(nil).AddFollowingArticle), userID, articleID)
}

// AddFollowingVaccine mocks base method.
func (m *MockIUserService) AddFollowingVaccine(userID, vaccineID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFollowingVaccine", userID, vaccineID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFollowingVaccine indicates an expected call of AddFollowingVaccine.
func (mr *MockIUserServiceMockRecorder) AddFollowingVaccine(userID, vaccineID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFollowingVaccine", reflect.TypeOf((*MockIUserService)(nil).AddFollowingVaccine), userID, vaccineID)
}

// CreateUser mocks base method.
func (m *MockIUserService) CreateUser(user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserServiceMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserService)(nil).CreateUser), user)
}

// DeleteUserByID mocks base method.
func (m *MockIUserService) DeleteUserByID(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserByID", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserByID indicates an expected call of DeleteUserByID.
func (mr *MockIUserServiceMockRecorder) DeleteUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserByID", reflect.TypeOf((*MockIUserService)(nil).DeleteUserByID), id)
}

// GetAllUsers mocks base method.
func (m *MockIUserService) GetAllUsers() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserServiceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserService)(nil).GetAllUsers))
}

// GetPublicUserByID mocks base method.
func (m *MockIUserService) GetPublicUserByID(id uint) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicUserByID", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicUserByID indicates an expected call of GetPublicUserByID.
func (mr *MockIUserServiceMockRecorder) GetPublicUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicUserByID", reflect.TypeOf((*MockIUserService)(nil).GetPublicUserByID), id)
}

// GetUserByID mocks base method.
func (m *MockIUserService) GetUserByID(id uint) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockIUserServiceMockRecorder) GetUserByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockIUserService)(nil).GetUserByID), id)
}

// GetUserByOpenID mocks base method.
func (m *MockIUserService) GetUserByOpenID(openID string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByOpenID", openID)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByOpenID indicates an expected call of GetUserByOpenID.
func (mr *MockIUserServiceMockRecorder) GetUserByOpenID(openID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByOpenID", reflect.TypeOf((*MockIUserService)(nil).GetUserByOpenID), openID)
}

// GetUserWithFollowings mocks base method.
func (m *MockIUserService) GetUserWithFollowings(id uint) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWithFollowings", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithFollowings indicates an expected call of GetUserWithFollowings.
func (mr *MockIUserServiceMockRecorder) GetUserWithFollowings(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithFollowings", reflect.TypeOf((*MockIUserService)(nil).GetUserWithFollowings), id)
}

// RemoveFollowingArticle mocks base method.
func (m *MockIUserService) RemoveFollowingArticle(userID, articleID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFollowingArticle", userID, articleID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFollowingArticle indicates an expected call of RemoveFollowingArticle.
func (mr *MockIUserServiceMockRecorder) RemoveFollowingArticle(userID, articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFollowingArticle", reflect.TypeOf((*MockIUserService)(nil).RemoveFollowingArticle), userID, articleID)
}

// RemoveFollowingVaccine mocks base method.
func (m *MockIUserService) RemoveFollowingVaccine(userID, vaccineID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFollowingVaccine", userID, vaccineID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFollowingVaccine indicates an expected call of RemoveFollowingVaccine.
func (mr *MockIUserServiceMockRecorder) RemoveFollowingVaccine(userID, vaccineID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFollowingVaccine", reflect.TypeOf((*MockIUserService)(nil).RemoveFollowingVaccine), userID, vaccineID)
}

// UpdateUserByID mocks base method.
func (m *MockIUserService) UpdateUserByID(user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByID", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByID indicates an expected call of UpdateUserByID.
func (mr *MockIUserServiceMockRecorder) UpdateUserByID(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByID", reflect.TypeOf((*MockIUserService)(nil).UpdateUserByID), user)
}