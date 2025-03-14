// Code generated by protoc-gen-goext. DO NOT EDIT.

package redis

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

func (m *UpdateUserRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateUserRequest) SetPasswords(v []string) {
	m.Passwords = v
}

func (m *UpdateUserRequest) SetPermissions(v *Permissions) {
	m.Permissions = v
}

func (m *UpdateUserRequest) SetEnabled(v bool) {
	m.Enabled = v
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
