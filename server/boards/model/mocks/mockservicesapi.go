// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-server/server/v8/boards/model (interfaces: ServicesAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mux "github.com/gorilla/mux"
	mlog "github.com/mattermost/mattermost-server/server/v8/platform/shared/mlog"
	model "github.com/mattermost/mattermost-server/server/v8/public/model"
)

// MockServicesAPI is a mock of ServicesAPI interface.
type MockServicesAPI struct {
	ctrl     *gomock.Controller
	recorder *MockServicesAPIMockRecorder
}

// MockServicesAPIMockRecorder is the mock recorder for MockServicesAPI.
type MockServicesAPIMockRecorder struct {
	mock *MockServicesAPI
}

// NewMockServicesAPI creates a new mock instance.
func NewMockServicesAPI(ctrl *gomock.Controller) *MockServicesAPI {
	mock := &MockServicesAPI{ctrl: ctrl}
	mock.recorder = &MockServicesAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServicesAPI) EXPECT() *MockServicesAPIMockRecorder {
	return m.recorder
}

// CreateMember mocks base method.
func (m *MockServicesAPI) CreateMember(arg0, arg1 string) (*model.TeamMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", arg0, arg1)
	ret0, _ := ret[0].(*model.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMember indicates an expected call of CreateMember.
func (mr *MockServicesAPIMockRecorder) CreateMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockServicesAPI)(nil).CreateMember), arg0, arg1)
}

// CreatePost mocks base method.
func (m *MockServicesAPI) CreatePost(arg0 *model.Post) (*model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", arg0)
	ret0, _ := ret[0].(*model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockServicesAPIMockRecorder) CreatePost(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockServicesAPI)(nil).CreatePost), arg0)
}

// DeletePreferencesForUser mocks base method.
func (m *MockServicesAPI) DeletePreferencesForUser(arg0 string, arg1 model.Preferences) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePreferencesForUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePreferencesForUser indicates an expected call of DeletePreferencesForUser.
func (mr *MockServicesAPIMockRecorder) DeletePreferencesForUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePreferencesForUser", reflect.TypeOf((*MockServicesAPI)(nil).DeletePreferencesForUser), arg0, arg1)
}

// EnsureBot mocks base method.
func (m *MockServicesAPI) EnsureBot(arg0 *model.Bot) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureBot", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureBot indicates an expected call of EnsureBot.
func (mr *MockServicesAPIMockRecorder) EnsureBot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureBot", reflect.TypeOf((*MockServicesAPI)(nil).EnsureBot), arg0)
}

// GetChannelByID mocks base method.
func (m *MockServicesAPI) GetChannelByID(arg0 string) (*model.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelByID", arg0)
	ret0, _ := ret[0].(*model.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannelByID indicates an expected call of GetChannelByID.
func (mr *MockServicesAPIMockRecorder) GetChannelByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelByID", reflect.TypeOf((*MockServicesAPI)(nil).GetChannelByID), arg0)
}

// GetChannelMember mocks base method.
func (m *MockServicesAPI) GetChannelMember(arg0, arg1 string) (*model.ChannelMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelMember", arg0, arg1)
	ret0, _ := ret[0].(*model.ChannelMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannelMember indicates an expected call of GetChannelMember.
func (mr *MockServicesAPIMockRecorder) GetChannelMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelMember", reflect.TypeOf((*MockServicesAPI)(nil).GetChannelMember), arg0, arg1)
}

// GetChannelsForTeamForUser mocks base method.
func (m *MockServicesAPI) GetChannelsForTeamForUser(arg0, arg1 string, arg2 bool) (model.ChannelList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelsForTeamForUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(model.ChannelList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannelsForTeamForUser indicates an expected call of GetChannelsForTeamForUser.
func (mr *MockServicesAPIMockRecorder) GetChannelsForTeamForUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelsForTeamForUser", reflect.TypeOf((*MockServicesAPI)(nil).GetChannelsForTeamForUser), arg0, arg1, arg2)
}

// GetCloudLimits mocks base method.
func (m *MockServicesAPI) GetCloudLimits() (*model.ProductLimits, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCloudLimits")
	ret0, _ := ret[0].(*model.ProductLimits)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCloudLimits indicates an expected call of GetCloudLimits.
func (mr *MockServicesAPIMockRecorder) GetCloudLimits() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCloudLimits", reflect.TypeOf((*MockServicesAPI)(nil).GetCloudLimits))
}

// GetConfig mocks base method.
func (m *MockServicesAPI) GetConfig() *model.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*model.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockServicesAPIMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockServicesAPI)(nil).GetConfig))
}

// GetDiagnosticID mocks base method.
func (m *MockServicesAPI) GetDiagnosticID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiagnosticID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDiagnosticID indicates an expected call of GetDiagnosticID.
func (mr *MockServicesAPIMockRecorder) GetDiagnosticID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiagnosticID", reflect.TypeOf((*MockServicesAPI)(nil).GetDiagnosticID))
}

// GetDirectChannel mocks base method.
func (m *MockServicesAPI) GetDirectChannel(arg0, arg1 string) (*model.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirectChannel", arg0, arg1)
	ret0, _ := ret[0].(*model.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDirectChannel indicates an expected call of GetDirectChannel.
func (mr *MockServicesAPIMockRecorder) GetDirectChannel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirectChannel", reflect.TypeOf((*MockServicesAPI)(nil).GetDirectChannel), arg0, arg1)
}

// GetDirectChannelOrCreate mocks base method.
func (m *MockServicesAPI) GetDirectChannelOrCreate(arg0, arg1 string) (*model.Channel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirectChannelOrCreate", arg0, arg1)
	ret0, _ := ret[0].(*model.Channel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDirectChannelOrCreate indicates an expected call of GetDirectChannelOrCreate.
func (mr *MockServicesAPIMockRecorder) GetDirectChannelOrCreate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirectChannelOrCreate", reflect.TypeOf((*MockServicesAPI)(nil).GetDirectChannelOrCreate), arg0, arg1)
}

// GetFileInfo mocks base method.
func (m *MockServicesAPI) GetFileInfo(arg0 string) (*model.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileInfo", arg0)
	ret0, _ := ret[0].(*model.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileInfo indicates an expected call of GetFileInfo.
func (mr *MockServicesAPIMockRecorder) GetFileInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileInfo", reflect.TypeOf((*MockServicesAPI)(nil).GetFileInfo), arg0)
}

// GetLicense mocks base method.
func (m *MockServicesAPI) GetLicense() *model.License {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLicense")
	ret0, _ := ret[0].(*model.License)
	return ret0
}

// GetLicense indicates an expected call of GetLicense.
func (mr *MockServicesAPIMockRecorder) GetLicense() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLicense", reflect.TypeOf((*MockServicesAPI)(nil).GetLicense))
}

// GetLogger mocks base method.
func (m *MockServicesAPI) GetLogger() mlog.LoggerIFace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(mlog.LoggerIFace)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockServicesAPIMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockServicesAPI)(nil).GetLogger))
}

// GetMasterDB mocks base method.
func (m *MockServicesAPI) GetMasterDB() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMasterDB")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMasterDB indicates an expected call of GetMasterDB.
func (mr *MockServicesAPIMockRecorder) GetMasterDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMasterDB", reflect.TypeOf((*MockServicesAPI)(nil).GetMasterDB))
}

// GetPreferencesForUser mocks base method.
func (m *MockServicesAPI) GetPreferencesForUser(arg0 string) (model.Preferences, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreferencesForUser", arg0)
	ret0, _ := ret[0].(model.Preferences)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreferencesForUser indicates an expected call of GetPreferencesForUser.
func (mr *MockServicesAPIMockRecorder) GetPreferencesForUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreferencesForUser", reflect.TypeOf((*MockServicesAPI)(nil).GetPreferencesForUser), arg0)
}

// GetTeamMember mocks base method.
func (m *MockServicesAPI) GetTeamMember(arg0, arg1 string) (*model.TeamMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTeamMember", arg0, arg1)
	ret0, _ := ret[0].(*model.TeamMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTeamMember indicates an expected call of GetTeamMember.
func (mr *MockServicesAPIMockRecorder) GetTeamMember(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTeamMember", reflect.TypeOf((*MockServicesAPI)(nil).GetTeamMember), arg0, arg1)
}

// GetUserByEmail mocks base method.
func (m *MockServicesAPI) GetUserByEmail(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockServicesAPIMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockServicesAPI)(nil).GetUserByEmail), arg0)
}

// GetUserByID mocks base method.
func (m *MockServicesAPI) GetUserByID(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockServicesAPIMockRecorder) GetUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockServicesAPI)(nil).GetUserByID), arg0)
}

// GetUserByUsername mocks base method.
func (m *MockServicesAPI) GetUserByUsername(arg0 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockServicesAPIMockRecorder) GetUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockServicesAPI)(nil).GetUserByUsername), arg0)
}

// GetUsersFromProfiles mocks base method.
func (m *MockServicesAPI) GetUsersFromProfiles(arg0 *model.UserGetOptions) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersFromProfiles", arg0)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersFromProfiles indicates an expected call of GetUsersFromProfiles.
func (mr *MockServicesAPIMockRecorder) GetUsersFromProfiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersFromProfiles", reflect.TypeOf((*MockServicesAPI)(nil).GetUsersFromProfiles), arg0)
}

// HasPermissionTo mocks base method.
func (m *MockServicesAPI) HasPermissionTo(arg0 string, arg1 *model.Permission) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermissionTo", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPermissionTo indicates an expected call of HasPermissionTo.
func (mr *MockServicesAPIMockRecorder) HasPermissionTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermissionTo", reflect.TypeOf((*MockServicesAPI)(nil).HasPermissionTo), arg0, arg1)
}

// HasPermissionToChannel mocks base method.
func (m *MockServicesAPI) HasPermissionToChannel(arg0, arg1 string, arg2 *model.Permission) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermissionToChannel", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPermissionToChannel indicates an expected call of HasPermissionToChannel.
func (mr *MockServicesAPIMockRecorder) HasPermissionToChannel(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermissionToChannel", reflect.TypeOf((*MockServicesAPI)(nil).HasPermissionToChannel), arg0, arg1, arg2)
}

// HasPermissionToTeam mocks base method.
func (m *MockServicesAPI) HasPermissionToTeam(arg0, arg1 string, arg2 *model.Permission) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermissionToTeam", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasPermissionToTeam indicates an expected call of HasPermissionToTeam.
func (mr *MockServicesAPIMockRecorder) HasPermissionToTeam(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermissionToTeam", reflect.TypeOf((*MockServicesAPI)(nil).HasPermissionToTeam), arg0, arg1, arg2)
}

// KVSetWithOptions mocks base method.
func (m *MockServicesAPI) KVSetWithOptions(arg0 string, arg1 []byte, arg2 model.PluginKVSetOptions) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KVSetWithOptions", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KVSetWithOptions indicates an expected call of KVSetWithOptions.
func (mr *MockServicesAPIMockRecorder) KVSetWithOptions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KVSetWithOptions", reflect.TypeOf((*MockServicesAPI)(nil).KVSetWithOptions), arg0, arg1, arg2)
}

// PublishPluginClusterEvent mocks base method.
func (m *MockServicesAPI) PublishPluginClusterEvent(arg0 model.PluginClusterEvent, arg1 model.PluginClusterEventSendOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishPluginClusterEvent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishPluginClusterEvent indicates an expected call of PublishPluginClusterEvent.
func (mr *MockServicesAPIMockRecorder) PublishPluginClusterEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishPluginClusterEvent", reflect.TypeOf((*MockServicesAPI)(nil).PublishPluginClusterEvent), arg0, arg1)
}

// PublishWebSocketEvent mocks base method.
func (m *MockServicesAPI) PublishWebSocketEvent(arg0 string, arg1 map[string]interface{}, arg2 *model.WebsocketBroadcast) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PublishWebSocketEvent", arg0, arg1, arg2)
}

// PublishWebSocketEvent indicates an expected call of PublishWebSocketEvent.
func (mr *MockServicesAPIMockRecorder) PublishWebSocketEvent(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishWebSocketEvent", reflect.TypeOf((*MockServicesAPI)(nil).PublishWebSocketEvent), arg0, arg1, arg2)
}

// RegisterRouter mocks base method.
func (m *MockServicesAPI) RegisterRouter(arg0 *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRouter", arg0)
}

// RegisterRouter indicates an expected call of RegisterRouter.
func (mr *MockServicesAPIMockRecorder) RegisterRouter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRouter", reflect.TypeOf((*MockServicesAPI)(nil).RegisterRouter), arg0)
}

// UpdatePreferencesForUser mocks base method.
func (m *MockServicesAPI) UpdatePreferencesForUser(arg0 string, arg1 model.Preferences) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePreferencesForUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePreferencesForUser indicates an expected call of UpdatePreferencesForUser.
func (mr *MockServicesAPIMockRecorder) UpdatePreferencesForUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePreferencesForUser", reflect.TypeOf((*MockServicesAPI)(nil).UpdatePreferencesForUser), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockServicesAPI) UpdateUser(arg0 *model.User) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServicesAPIMockRecorder) UpdateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockServicesAPI)(nil).UpdateUser), arg0)
}
