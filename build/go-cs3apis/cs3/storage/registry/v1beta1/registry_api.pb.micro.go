// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: cs3/storage/registry/v1beta1/registry_api.proto

package registryv1beta1

import (
	fmt "fmt"
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

// Api Endpoints for RegistryAPI service

func NewRegistryAPIEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RegistryAPI service

type RegistryAPIService interface {
	// Returns the storage provider that is reponsible for the given
	// resource reference.
	// MUST return CODE_NOT_FOUND if the reference does not exist.
	GetStorageProvider(ctx context.Context, in *GetStorageProviderRequest, opts ...client.CallOption) (*GetStorageProviderResponse, error)
	// Returns a list of the available storage providers known by this registry.
	ListStorageProviders(ctx context.Context, in *ListStorageProvidersRequest, opts ...client.CallOption) (*ListStorageProvidersResponse, error)
	// Gets the user home storage provider.
	GetHome(ctx context.Context, in *GetHomeRequest, opts ...client.CallOption) (*GetHomeResponse, error)
}

type registryAPIService struct {
	c    client.Client
	name string
}

func NewRegistryAPIService(name string, c client.Client) RegistryAPIService {
	return &registryAPIService{
		c:    c,
		name: name,
	}
}

func (c *registryAPIService) GetStorageProvider(ctx context.Context, in *GetStorageProviderRequest, opts ...client.CallOption) (*GetStorageProviderResponse, error) {
	req := c.c.NewRequest(c.name, "RegistryAPI.GetStorageProvider", in)
	out := new(GetStorageProviderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryAPIService) ListStorageProviders(ctx context.Context, in *ListStorageProvidersRequest, opts ...client.CallOption) (*ListStorageProvidersResponse, error) {
	req := c.c.NewRequest(c.name, "RegistryAPI.ListStorageProviders", in)
	out := new(ListStorageProvidersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registryAPIService) GetHome(ctx context.Context, in *GetHomeRequest, opts ...client.CallOption) (*GetHomeResponse, error) {
	req := c.c.NewRequest(c.name, "RegistryAPI.GetHome", in)
	out := new(GetHomeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RegistryAPI service

type RegistryAPIHandler interface {
	// Returns the storage provider that is reponsible for the given
	// resource reference.
	// MUST return CODE_NOT_FOUND if the reference does not exist.
	GetStorageProvider(context.Context, *GetStorageProviderRequest, *GetStorageProviderResponse) error
	// Returns a list of the available storage providers known by this registry.
	ListStorageProviders(context.Context, *ListStorageProvidersRequest, *ListStorageProvidersResponse) error
	// Gets the user home storage provider.
	GetHome(context.Context, *GetHomeRequest, *GetHomeResponse) error
}

func RegisterRegistryAPIHandler(s server.Server, hdlr RegistryAPIHandler, opts ...server.HandlerOption) error {
	type registryAPI interface {
		GetStorageProvider(ctx context.Context, in *GetStorageProviderRequest, out *GetStorageProviderResponse) error
		ListStorageProviders(ctx context.Context, in *ListStorageProvidersRequest, out *ListStorageProvidersResponse) error
		GetHome(ctx context.Context, in *GetHomeRequest, out *GetHomeResponse) error
	}
	type RegistryAPI struct {
		registryAPI
	}
	h := &registryAPIHandler{hdlr}
	return s.Handle(s.NewHandler(&RegistryAPI{h}, opts...))
}

type registryAPIHandler struct {
	RegistryAPIHandler
}

func (h *registryAPIHandler) GetStorageProvider(ctx context.Context, in *GetStorageProviderRequest, out *GetStorageProviderResponse) error {
	return h.RegistryAPIHandler.GetStorageProvider(ctx, in, out)
}

func (h *registryAPIHandler) ListStorageProviders(ctx context.Context, in *ListStorageProvidersRequest, out *ListStorageProvidersResponse) error {
	return h.RegistryAPIHandler.ListStorageProviders(ctx, in, out)
}

func (h *registryAPIHandler) GetHome(ctx context.Context, in *GetHomeRequest, out *GetHomeResponse) error {
	return h.RegistryAPIHandler.GetHome(ctx, in, out)
}
