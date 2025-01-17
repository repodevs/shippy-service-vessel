// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package vessel

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

	context "context"

	api "github.com/micro/go-micro/v2/api"

	client "github.com/micro/go-micro/v2/client"

	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for VesselService service

func NewVesselServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VesselService service

type VesselService interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselService struct {
	c    client.Client
	name string
}

func NewVesselService(name string, c client.Client) VesselService {
	return &vesselService{
		c:    c,
		name: name,
	}
}

func (c *vesselService) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) error {
	type vesselService interface {
		FindAvailable(ctx context.Context, in *Specification, out *Response) error
	}
	type VesselService struct {
		vesselService
	}
	h := &vesselServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VesselService{h}, opts...))
}

type vesselServiceHandler struct {
	VesselServiceHandler
}

func (h *vesselServiceHandler) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}
