// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab/gitlab.go

// Package mock_gitlab is a generated GoMock package.
package mock_gitlab

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gitlab0 "github.com/xanzy/go-gitlab"
	oauth2 "golang.org/x/oauth2"

	gitlab "github.com/annkuzn/mattermost-plugin-gitlab/server/gitlab"
)

// MockGitlab is a mock of Gitlab interface
type MockGitlab struct {
	ctrl     *gomock.Controller
	recorder *MockGitlabMockRecorder
}

// MockGitlabMockRecorder is the mock recorder for MockGitlab
type MockGitlabMockRecorder struct {
	mock *MockGitlab
}

// NewMockGitlab creates a new mock instance
func NewMockGitlab(ctrl *gomock.Controller) *MockGitlab {
	mock := &MockGitlab{ctrl: ctrl}
	mock.recorder = &MockGitlabMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitlab) EXPECT() *MockGitlabMockRecorder {
	return m.recorder
}

// GetCurrentUser mocks base method
func (m *MockGitlab) GetCurrentUser(userID string, token oauth2.Token) (*gitlab.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentUser", userID, token)
	ret0, _ := ret[0].(*gitlab.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentUser indicates an expected call of GetCurrentUser
func (mr *MockGitlabMockRecorder) GetCurrentUser(userID, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentUser", reflect.TypeOf((*MockGitlab)(nil).GetCurrentUser), userID, token)
}

// GetUserDetails mocks base method
func (m *MockGitlab) GetUserDetails(user *gitlab.UserInfo) (*gitlab0.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetails", user)
	ret0, _ := ret[0].(*gitlab0.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetails indicates an expected call of GetUserDetails
func (mr *MockGitlabMockRecorder) GetUserDetails(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetails", reflect.TypeOf((*MockGitlab)(nil).GetUserDetails), user)
}

// GetProject mocks base method
func (m *MockGitlab) GetProject(user *gitlab.UserInfo, owner, repo string) (*gitlab0.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", user, owner, repo)
	ret0, _ := ret[0].(*gitlab0.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject
func (mr *MockGitlabMockRecorder) GetProject(user, owner, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockGitlab)(nil).GetProject), user, owner, repo)
}

// GetReviews mocks base method
func (m *MockGitlab) GetReviews(user *gitlab.UserInfo) ([]*gitlab0.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReviews", user)
	ret0, _ := ret[0].([]*gitlab0.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReviews indicates an expected call of GetReviews
func (mr *MockGitlabMockRecorder) GetReviews(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReviews", reflect.TypeOf((*MockGitlab)(nil).GetReviews), user)
}

// GetYourPrs mocks base method
func (m *MockGitlab) GetYourPrs(user *gitlab.UserInfo) ([]*gitlab0.MergeRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetYourPrs", user)
	ret0, _ := ret[0].([]*gitlab0.MergeRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetYourPrs indicates an expected call of GetYourPrs
func (mr *MockGitlabMockRecorder) GetYourPrs(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetYourPrs", reflect.TypeOf((*MockGitlab)(nil).GetYourPrs), user)
}

// GetYourAssignments mocks base method
func (m *MockGitlab) GetYourAssignments(user *gitlab.UserInfo) ([]*gitlab0.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetYourAssignments", user)
	ret0, _ := ret[0].([]*gitlab0.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetYourAssignments indicates an expected call of GetYourAssignments
func (mr *MockGitlabMockRecorder) GetYourAssignments(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetYourAssignments", reflect.TypeOf((*MockGitlab)(nil).GetYourAssignments), user)
}

// GetUnreads mocks base method
func (m *MockGitlab) GetUnreads(user *gitlab.UserInfo) ([]*gitlab0.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnreads", user)
	ret0, _ := ret[0].([]*gitlab0.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnreads indicates an expected call of GetUnreads
func (mr *MockGitlabMockRecorder) GetUnreads(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnreads", reflect.TypeOf((*MockGitlab)(nil).GetUnreads), user)
}

// GetProjectHooks mocks base method
func (m *MockGitlab) GetProjectHooks(user *gitlab.UserInfo, owner, repo string) ([]*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectHooks", user, owner, repo)
	ret0, _ := ret[0].([]*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectHooks indicates an expected call of GetProjectHooks
func (mr *MockGitlabMockRecorder) GetProjectHooks(user, owner, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectHooks", reflect.TypeOf((*MockGitlab)(nil).GetProjectHooks), user, owner, repo)
}

// GetGroupHooks mocks base method
func (m *MockGitlab) GetGroupHooks(user *gitlab.UserInfo, owner string) ([]*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupHooks", user, owner)
	ret0, _ := ret[0].([]*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupHooks indicates an expected call of GetGroupHooks
func (mr *MockGitlabMockRecorder) GetGroupHooks(user, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupHooks", reflect.TypeOf((*MockGitlab)(nil).GetGroupHooks), user, owner)
}

// NewProjectHook mocks base method
func (m *MockGitlab) NewProjectHook(user *gitlab.UserInfo, projectID interface{}, projectHookOptions *gitlab.AddWebhookOptions) (*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProjectHook", user, projectID, projectHookOptions)
	ret0, _ := ret[0].(*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewProjectHook indicates an expected call of NewProjectHook
func (mr *MockGitlabMockRecorder) NewProjectHook(user, projectID, projectHookOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProjectHook", reflect.TypeOf((*MockGitlab)(nil).NewProjectHook), user, projectID, projectHookOptions)
}

// NewGroupHook mocks base method
func (m *MockGitlab) NewGroupHook(user *gitlab.UserInfo, groupName string, groupHookOptions *gitlab.AddWebhookOptions) (*gitlab.WebhookInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGroupHook", user, groupName, groupHookOptions)
	ret0, _ := ret[0].(*gitlab.WebhookInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGroupHook indicates an expected call of NewGroupHook
func (mr *MockGitlabMockRecorder) NewGroupHook(user, groupName, groupHookOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGroupHook", reflect.TypeOf((*MockGitlab)(nil).NewGroupHook), user, groupName, groupHookOptions)
}

// ResolveNamespaceAndProject mocks base method
func (m *MockGitlab) ResolveNamespaceAndProject(userInfo *gitlab.UserInfo, fullPath string, allowPrivate bool) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveNamespaceAndProject", userInfo, fullPath, allowPrivate)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ResolveNamespaceAndProject indicates an expected call of ResolveNamespaceAndProject
func (mr *MockGitlabMockRecorder) ResolveNamespaceAndProject(userInfo, fullPath, allowPrivate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveNamespaceAndProject", reflect.TypeOf((*MockGitlab)(nil).ResolveNamespaceAndProject), userInfo, fullPath, allowPrivate)
}
