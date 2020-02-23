// Code generated by jrpc. DO NOT EDIT.

package userdef

import (
	context "context"

	rpc "github.com/jakewright/home-automation/libraries/go/rpc"
)

// Do performs the request
func (m *GetUserRequest) Do(ctx context.Context) (*GetUserResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.user/user",
		Body:   m,
	}

	rsp := &GetUserResponse{}
	_, err := rpc.Do(ctx, req, rsp)
	return rsp, err
}

// Do performs the request
func (m *ListUsersRequest) Do(ctx context.Context) (*ListUsersResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.user/users",
		Body:   m,
	}

	rsp := &ListUsersResponse{}
	_, err := rpc.Do(ctx, req, rsp)
	return rsp, err
}
