// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cs3/sharing/collaboration/v1beta1/collaboration_api.proto

package collaborationv1beta1

import (
	fmt "fmt"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/identity/user/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/rpc/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/storage/provider/v1beta1"
	_ "github.com/cs3org/go-cs3apis/build/go-cs3apis/cs3/types/v1beta1"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CollaborationAPI service

func NewCollaborationAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CollaborationAPI service

type CollaborationAPIService interface {
	// Creates a new share.
	// MUST return CODE_NOT_FOUND if the resource reference does not exist.
	// MUST return CODE_ALREADY_EXISTS if the share already exists for the 4-tuple consisting of
	// (owner, shared_resource, grantee).
	// New shares MUST be created in the state SHARE_STATE_PENDING.
	CreateShare(ctx context.Context, in *CreateShareRequest, opts ...client.CallOption) (*CreateShareResponse, error)
	// Removes a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	RemoveShare(ctx context.Context, in *RemoveShareRequest, opts ...client.CallOption) (*RemoveShareResponse, error)
	// Gets share information for a single share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	GetShare(ctx context.Context, in *GetShareRequest, opts ...client.CallOption) (*GetShareResponse, error)
	// List the shares the authenticated principal has created,
	// both as owner and creator. If a filter is specified, only
	// shares satisfying the filter MUST be returned.
	ListShares(ctx context.Context, in *ListSharesRequest, opts ...client.CallOption) (*ListSharesResponse, error)
	// Updates a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateShare(ctx context.Context, in *UpdateShareRequest, opts ...client.CallOption) (*UpdateShareResponse, error)
	// List all shares the authenticated principal has received.
	ListReceivedShares(ctx context.Context, in *ListReceivedSharesRequest, opts ...client.CallOption) (*ListReceivedSharesResponse, error)
	// Update the received share to change the share state or the display name.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateReceivedShare(ctx context.Context, in *UpdateReceivedShareRequest, opts ...client.CallOption) (*UpdateReceivedShareResponse, error)
	// Get the information for the given received share reference.
	// MUST return CODE_NOT_FOUND if the received share reference does not exist.
	GetReceivedShare(ctx context.Context, in *GetReceivedShareRequest, opts ...client.CallOption) (*GetReceivedShareResponse, error)
}

type collaborationAPIService struct {
	c    client.Client
	name string
}

func NewCollaborationAPIService(name string, c client.Client) CollaborationAPIService {
	return &collaborationAPIService{
		c:    c,
		name: name,
	}
}

func (c *collaborationAPIService) CreateShare(ctx context.Context, in *CreateShareRequest, opts ...client.CallOption) (*CreateShareResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.CreateShare", in)
	out := new(CreateShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIService) RemoveShare(ctx context.Context, in *RemoveShareRequest, opts ...client.CallOption) (*RemoveShareResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.RemoveShare", in)
	out := new(RemoveShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIService) GetShare(ctx context.Context, in *GetShareRequest, opts ...client.CallOption) (*GetShareResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.GetShare", in)
	out := new(GetShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIService) ListShares(ctx context.Context, in *ListSharesRequest, opts ...client.CallOption) (*ListSharesResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.ListShares", in)
	out := new(ListSharesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIService) UpdateShare(ctx context.Context, in *UpdateShareRequest, opts ...client.CallOption) (*UpdateShareResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.UpdateShare", in)
	out := new(UpdateShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIService) ListReceivedShares(ctx context.Context, in *ListReceivedSharesRequest, opts ...client.CallOption) (*ListReceivedSharesResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.ListReceivedShares", in)
	out := new(ListReceivedSharesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIService) UpdateReceivedShare(ctx context.Context, in *UpdateReceivedShareRequest, opts ...client.CallOption) (*UpdateReceivedShareResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.UpdateReceivedShare", in)
	out := new(UpdateReceivedShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationAPIService) GetReceivedShare(ctx context.Context, in *GetReceivedShareRequest, opts ...client.CallOption) (*GetReceivedShareResponse, error) {
	req := c.c.NewRequest(c.name, "CollaborationAPI.GetReceivedShare", in)
	out := new(GetReceivedShareResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CollaborationAPI service

type CollaborationAPIHandler interface {
	// Creates a new share.
	// MUST return CODE_NOT_FOUND if the resource reference does not exist.
	// MUST return CODE_ALREADY_EXISTS if the share already exists for the 4-tuple consisting of
	// (owner, shared_resource, grantee).
	// New shares MUST be created in the state SHARE_STATE_PENDING.
	CreateShare(context.Context, *CreateShareRequest, *CreateShareResponse) error
	// Removes a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	RemoveShare(context.Context, *RemoveShareRequest, *RemoveShareResponse) error
	// Gets share information for a single share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	GetShare(context.Context, *GetShareRequest, *GetShareResponse) error
	// List the shares the authenticated principal has created,
	// both as owner and creator. If a filter is specified, only
	// shares satisfying the filter MUST be returned.
	ListShares(context.Context, *ListSharesRequest, *ListSharesResponse) error
	// Updates a share.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateShare(context.Context, *UpdateShareRequest, *UpdateShareResponse) error
	// List all shares the authenticated principal has received.
	ListReceivedShares(context.Context, *ListReceivedSharesRequest, *ListReceivedSharesResponse) error
	// Update the received share to change the share state or the display name.
	// MUST return CODE_NOT_FOUND if the share reference does not exist.
	UpdateReceivedShare(context.Context, *UpdateReceivedShareRequest, *UpdateReceivedShareResponse) error
	// Get the information for the given received share reference.
	// MUST return CODE_NOT_FOUND if the received share reference does not exist.
	GetReceivedShare(context.Context, *GetReceivedShareRequest, *GetReceivedShareResponse) error
}

func RegisterCollaborationAPIHandler(s server.Server, hdlr CollaborationAPIHandler, opts ...server.HandlerOption) error {
	type collaborationAPI interface {
		CreateShare(ctx context.Context, in *CreateShareRequest, out *CreateShareResponse) error
		RemoveShare(ctx context.Context, in *RemoveShareRequest, out *RemoveShareResponse) error
		GetShare(ctx context.Context, in *GetShareRequest, out *GetShareResponse) error
		ListShares(ctx context.Context, in *ListSharesRequest, out *ListSharesResponse) error
		UpdateShare(ctx context.Context, in *UpdateShareRequest, out *UpdateShareResponse) error
		ListReceivedShares(ctx context.Context, in *ListReceivedSharesRequest, out *ListReceivedSharesResponse) error
		UpdateReceivedShare(ctx context.Context, in *UpdateReceivedShareRequest, out *UpdateReceivedShareResponse) error
		GetReceivedShare(ctx context.Context, in *GetReceivedShareRequest, out *GetReceivedShareResponse) error
	}
	type CollaborationAPI struct {
		collaborationAPI
	}
	h := &collaborationAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&CollaborationAPI{h}, opts...))
}

type collaborationAPIHandler struct {
	CollaborationAPIHandler
}

func (h *collaborationAPIHandler) CreateShare(ctx context.Context, in *CreateShareRequest, out *CreateShareResponse) error {
	return h.CollaborationAPIHandler.CreateShare(ctx, in, out)
}

func (h *collaborationAPIHandler) RemoveShare(ctx context.Context, in *RemoveShareRequest, out *RemoveShareResponse) error {
	return h.CollaborationAPIHandler.RemoveShare(ctx, in, out)
}

func (h *collaborationAPIHandler) GetShare(ctx context.Context, in *GetShareRequest, out *GetShareResponse) error {
	return h.CollaborationAPIHandler.GetShare(ctx, in, out)
}

func (h *collaborationAPIHandler) ListShares(ctx context.Context, in *ListSharesRequest, out *ListSharesResponse) error {
	return h.CollaborationAPIHandler.ListShares(ctx, in, out)
}

func (h *collaborationAPIHandler) UpdateShare(ctx context.Context, in *UpdateShareRequest, out *UpdateShareResponse) error {
	return h.CollaborationAPIHandler.UpdateShare(ctx, in, out)
}

func (h *collaborationAPIHandler) ListReceivedShares(ctx context.Context, in *ListReceivedSharesRequest, out *ListReceivedSharesResponse) error {
	return h.CollaborationAPIHandler.ListReceivedShares(ctx, in, out)
}

func (h *collaborationAPIHandler) UpdateReceivedShare(ctx context.Context, in *UpdateReceivedShareRequest, out *UpdateReceivedShareResponse) error {
	return h.CollaborationAPIHandler.UpdateReceivedShare(ctx, in, out)
}

func (h *collaborationAPIHandler) GetReceivedShare(ctx context.Context, in *GetReceivedShareRequest, out *GetReceivedShareResponse) error {
	return h.CollaborationAPIHandler.GetReceivedShare(ctx, in, out)
}
