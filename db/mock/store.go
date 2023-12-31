// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/patrickchap/clipsapi/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/patrickchap/clipsapi/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(arg0 context.Context, arg1 db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), arg0, arg1)
}

// CreateLike mocks base method.
func (m *MockStore) CreateLike(arg0 context.Context, arg1 db.CreateLikeParams) (db.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLike", arg0, arg1)
	ret0, _ := ret[0].(db.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLike indicates an expected call of CreateLike.
func (mr *MockStoreMockRecorder) CreateLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLike", reflect.TypeOf((*MockStore)(nil).CreateLike), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// CreateVideo mocks base method.
func (m *MockStore) CreateVideo(arg0 context.Context, arg1 db.CreateVideoParams) (db.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVideo", arg0, arg1)
	ret0, _ := ret[0].(db.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVideo indicates an expected call of CreateVideo.
func (mr *MockStoreMockRecorder) CreateVideo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVideo", reflect.TypeOf((*MockStore)(nil).CreateVideo), arg0, arg1)
}

// DeleteCommentsByVideo mocks base method.
func (m *MockStore) DeleteCommentsByVideo(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCommentsByVideo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCommentsByVideo indicates an expected call of DeleteCommentsByVideo.
func (mr *MockStoreMockRecorder) DeleteCommentsByVideo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCommentsByVideo", reflect.TypeOf((*MockStore)(nil).DeleteCommentsByVideo), arg0, arg1)
}

// DeleteLikesByVideoAndUser mocks base method.
func (m *MockStore) DeleteLikesByVideoAndUser(arg0 context.Context, arg1 db.DeleteLikesByVideoAndUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLikesByVideoAndUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLikesByVideoAndUser indicates an expected call of DeleteLikesByVideoAndUser.
func (mr *MockStoreMockRecorder) DeleteLikesByVideoAndUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLikesByVideoAndUser", reflect.TypeOf((*MockStore)(nil).DeleteLikesByVideoAndUser), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// DeleteVideo mocks base method.
func (m *MockStore) DeleteVideo(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVideo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVideo indicates an expected call of DeleteVideo.
func (mr *MockStoreMockRecorder) DeleteVideo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVideo", reflect.TypeOf((*MockStore)(nil).DeleteVideo), arg0, arg1)
}

// GetCommentsByVideo mocks base method.
func (m *MockStore) GetCommentsByVideo(arg0 context.Context, arg1 db.GetCommentsByVideoParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsByVideo", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByVideo indicates an expected call of GetCommentsByVideo.
func (mr *MockStoreMockRecorder) GetCommentsByVideo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByVideo", reflect.TypeOf((*MockStore)(nil).GetCommentsByVideo), arg0, arg1)
}

// GetLikesByVideo mocks base method.
func (m *MockStore) GetLikesByVideo(arg0 context.Context, arg1 db.GetLikesByVideoParams) ([]db.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikesByVideo", arg0, arg1)
	ret0, _ := ret[0].([]db.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikesByVideo indicates an expected call of GetLikesByVideo.
func (mr *MockStoreMockRecorder) GetLikesByVideo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikesByVideo", reflect.TypeOf((*MockStore)(nil).GetLikesByVideo), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// GetUserVideoWithLikes mocks base method.
func (m *MockStore) GetUserVideoWithLikes(arg0 context.Context, arg1 db.GetUserVideoWithLikesParams) ([]db.GetUserVideoWithLikesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserVideoWithLikes", arg0, arg1)
	ret0, _ := ret[0].([]db.GetUserVideoWithLikesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserVideoWithLikes indicates an expected call of GetUserVideoWithLikes.
func (mr *MockStoreMockRecorder) GetUserVideoWithLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserVideoWithLikes", reflect.TypeOf((*MockStore)(nil).GetUserVideoWithLikes), arg0, arg1)
}

// GetVideo mocks base method.
func (m *MockStore) GetVideo(arg0 context.Context, arg1 int64) (db.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideo", arg0, arg1)
	ret0, _ := ret[0].(db.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideo indicates an expected call of GetVideo.
func (mr *MockStoreMockRecorder) GetVideo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideo", reflect.TypeOf((*MockStore)(nil).GetVideo), arg0, arg1)
}

// GetVideoWithLikes mocks base method.
func (m *MockStore) GetVideoWithLikes(arg0 context.Context, arg1 int64) (db.GetVideoWithLikesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideoWithLikes", arg0, arg1)
	ret0, _ := ret[0].(db.GetVideoWithLikesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideoWithLikes indicates an expected call of GetVideoWithLikes.
func (mr *MockStoreMockRecorder) GetVideoWithLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideoWithLikes", reflect.TypeOf((*MockStore)(nil).GetVideoWithLikes), arg0, arg1)
}

// GetVideoWithLikesWithSearch mocks base method.
func (m *MockStore) GetVideoWithLikesWithSearch(arg0 context.Context, arg1 int64) (db.GetVideoWithLikesWithSearchRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVideoWithLikesWithSearch", arg0, arg1)
	ret0, _ := ret[0].(db.GetVideoWithLikesWithSearchRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVideoWithLikesWithSearch indicates an expected call of GetVideoWithLikesWithSearch.
func (mr *MockStoreMockRecorder) GetVideoWithLikesWithSearch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVideoWithLikesWithSearch", reflect.TypeOf((*MockStore)(nil).GetVideoWithLikesWithSearch), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// ListVideos mocks base method.
func (m *MockStore) ListVideos(arg0 context.Context, arg1 db.ListVideosParams) ([]db.Video, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVideos", arg0, arg1)
	ret0, _ := ret[0].([]db.Video)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVideos indicates an expected call of ListVideos.
func (mr *MockStoreMockRecorder) ListVideos(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVideos", reflect.TypeOf((*MockStore)(nil).ListVideos), arg0, arg1)
}

// ListVideosWithLikesAndSearch mocks base method.
func (m *MockStore) ListVideosWithLikesAndSearch(arg0 context.Context, arg1 db.ListVideosWithLikesAndSearchParams) ([]db.ListVideosWithLikesAndSearchRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVideosWithLikesAndSearch", arg0, arg1)
	ret0, _ := ret[0].([]db.ListVideosWithLikesAndSearchRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVideosWithLikesAndSearch indicates an expected call of ListVideosWithLikesAndSearch.
func (mr *MockStoreMockRecorder) ListVideosWithLikesAndSearch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVideosWithLikesAndSearch", reflect.TypeOf((*MockStore)(nil).ListVideosWithLikesAndSearch), arg0, arg1)
}

// UpdateVideo mocks base method.
func (m *MockStore) UpdateVideo(arg0 context.Context, arg1 db.UpdateVideoParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVideo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVideo indicates an expected call of UpdateVideo.
func (mr *MockStoreMockRecorder) UpdateVideo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVideo", reflect.TypeOf((*MockStore)(nil).UpdateVideo), arg0, arg1)
}
