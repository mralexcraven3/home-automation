// Code generated by jrpc. DO NOT EDIT.

package deviceregistrydef

import (
	context "context"

	rpc "github.com/jakewright/home-automation/libraries/go/rpc"
)

// Do performs the request
func (m *GetDeviceRequest) Do(ctx context.Context) (*GetDeviceResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/device",
		Body:   m,
	}

	rsp := &GetDeviceResponse{}
	_, err := rpc.Do(ctx, req, rsp)
	return rsp, err
}

// Do performs the request
func (m *ListDevicesRequest) Do(ctx context.Context) (*ListDevicesResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/devices",
		Body:   m,
	}

	rsp := &ListDevicesResponse{}
	_, err := rpc.Do(ctx, req, rsp)
	return rsp, err
}

// Do performs the request
func (m *GetRoomRequest) Do(ctx context.Context) (*GetRoomResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/room",
		Body:   m,
	}

	rsp := &GetRoomResponse{}
	_, err := rpc.Do(ctx, req, rsp)
	return rsp, err
}

// Do performs the request
func (m *ListRoomsRequest) Do(ctx context.Context) (*ListRoomsResponse, error) {
	req := &rpc.Request{
		Method: "GET",
		URL:    "service.device-registry/rooms",
		Body:   m,
	}

	rsp := &ListRoomsResponse{}
	_, err := rpc.Do(ctx, req, rsp)
	return rsp, err
}
