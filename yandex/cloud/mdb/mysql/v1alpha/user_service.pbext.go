// Code generated by protoc-gen-goext. DO NOT EDIT.

package mysql

import (
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

func (m *GetUserRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetUserRequest) SetUserName(v string) {
	m.UserName = v
}

func (m *ListUsersRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListUsersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListUsersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListUsersResponse) SetUsers(v []*User) {
	m.Users = v
}

func (m *ListUsersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateUserRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateUserRequest) SetUserSpec(v *UserSpec) {
	m.UserSpec = v
}

func (m *CreateUserMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateUserMetadata) SetUserName(v string) {
	m.UserName = v
}

func (m *UpdateUserRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateUserRequest) SetUserName(v string) {
	m.UserName = v
}

func (m *UpdateUserRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateUserRequest) SetPassword(v string) {
	m.Password = v
}

func (m *UpdateUserRequest) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *UpdateUserMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateUserMetadata) SetUserName(v string) {
	m.UserName = v
}

func (m *DeleteUserRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteUserRequest) SetUserName(v string) {
	m.UserName = v
}

func (m *DeleteUserMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteUserMetadata) SetUserName(v string) {
	m.UserName = v
}

func (m *GrantUserPermissionRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GrantUserPermissionRequest) SetUserName(v string) {
	m.UserName = v
}

func (m *GrantUserPermissionRequest) SetPermission(v *Permission) {
	m.Permission = v
}

func (m *GrantUserPermissionMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GrantUserPermissionMetadata) SetUserName(v string) {
	m.UserName = v
}

func (m *RevokeUserPermissionRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RevokeUserPermissionRequest) SetUserName(v string) {
	m.UserName = v
}

func (m *RevokeUserPermissionRequest) SetDatabaseName(v string) {
	m.DatabaseName = v
}

func (m *RevokeUserPermissionMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RevokeUserPermissionMetadata) SetUserName(v string) {
	m.UserName = v
}
