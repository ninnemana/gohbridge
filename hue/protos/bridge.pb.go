// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/bridge.proto

/*
Package hue is a generated protocol buffer package.

It is generated from these files:
	protos/bridge.proto

It has these top-level messages:
	BridgeNetwork
	Bridge
	DiscoverParams
	BridgeState
	Light
	ListLightParams
	GetLightParams
	Group
	Config
*/
package hue

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BridgeNetwork struct {
	InternalIP string `protobuf:"bytes,1,opt,name=InternalIP" json:"InternalIP,omitempty"`
}

func (m *BridgeNetwork) Reset()                    { *m = BridgeNetwork{} }
func (m *BridgeNetwork) String() string            { return proto.CompactTextString(m) }
func (*BridgeNetwork) ProtoMessage()               {}
func (*BridgeNetwork) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BridgeNetwork) GetInternalIP() string {
	if m != nil {
		return m.InternalIP
	}
	return ""
}

type Bridge struct {
	Network *BridgeNetwork `protobuf:"bytes,1,opt,name=network" json:"network,omitempty"`
	Id      string         `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	User    string         `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
}

func (m *Bridge) Reset()                    { *m = Bridge{} }
func (m *Bridge) String() string            { return proto.CompactTextString(m) }
func (*Bridge) ProtoMessage()               {}
func (*Bridge) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Bridge) GetNetwork() *BridgeNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *Bridge) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Bridge) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type DiscoverParams struct {
}

func (m *DiscoverParams) Reset()                    { *m = DiscoverParams{} }
func (m *DiscoverParams) String() string            { return proto.CompactTextString(m) }
func (*DiscoverParams) ProtoMessage()               {}
func (*DiscoverParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type BridgeState struct {
	Lights map[string]*Light `protobuf:"bytes,1,rep,name=lights" json:"lights,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Groups map[string]*Group `protobuf:"bytes,2,rep,name=groups" json:"groups,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Cfg    *Config           `protobuf:"bytes,3,opt,name=cfg" json:"cfg,omitempty"`
}

func (m *BridgeState) Reset()                    { *m = BridgeState{} }
func (m *BridgeState) String() string            { return proto.CompactTextString(m) }
func (*BridgeState) ProtoMessage()               {}
func (*BridgeState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BridgeState) GetLights() map[string]*Light {
	if m != nil {
		return m.Lights
	}
	return nil
}

func (m *BridgeState) GetGroups() map[string]*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *BridgeState) GetCfg() *Config {
	if m != nil {
		return m.Cfg
	}
	return nil
}

type Light struct {
}

func (m *Light) Reset()                    { *m = Light{} }
func (m *Light) String() string            { return proto.CompactTextString(m) }
func (*Light) ProtoMessage()               {}
func (*Light) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListLightParams struct {
}

func (m *ListLightParams) Reset()                    { *m = ListLightParams{} }
func (m *ListLightParams) String() string            { return proto.CompactTextString(m) }
func (*ListLightParams) ProtoMessage()               {}
func (*ListLightParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetLightParams struct {
}

func (m *GetLightParams) Reset()                    { *m = GetLightParams{} }
func (m *GetLightParams) String() string            { return proto.CompactTextString(m) }
func (*GetLightParams) ProtoMessage()               {}
func (*GetLightParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type Group struct {
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Config struct {
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*BridgeNetwork)(nil), "hue.BridgeNetwork")
	proto.RegisterType((*Bridge)(nil), "hue.Bridge")
	proto.RegisterType((*DiscoverParams)(nil), "hue.DiscoverParams")
	proto.RegisterType((*BridgeState)(nil), "hue.BridgeState")
	proto.RegisterType((*Light)(nil), "hue.Light")
	proto.RegisterType((*ListLightParams)(nil), "hue.ListLightParams")
	proto.RegisterType((*GetLightParams)(nil), "hue.GetLightParams")
	proto.RegisterType((*Group)(nil), "hue.Group")
	proto.RegisterType((*Config)(nil), "hue.Config")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hue service

type HueClient interface {
	Discover(ctx context.Context, in *DiscoverParams, opts ...grpc.CallOption) (Hue_DiscoverClient, error)
	GetBridgeState(ctx context.Context, in *Bridge, opts ...grpc.CallOption) (*BridgeState, error)
	ListLights(ctx context.Context, in *ListLightParams, opts ...grpc.CallOption) (Hue_ListLightsClient, error)
	GetLight(ctx context.Context, in *GetLightParams, opts ...grpc.CallOption) (*Light, error)
	SetLight(ctx context.Context, opts ...grpc.CallOption) (Hue_SetLightClient, error)
}

type hueClient struct {
	cc *grpc.ClientConn
}

func NewHueClient(cc *grpc.ClientConn) HueClient {
	return &hueClient{cc}
}

func (c *hueClient) Discover(ctx context.Context, in *DiscoverParams, opts ...grpc.CallOption) (Hue_DiscoverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hue_serviceDesc.Streams[0], c.cc, "/hue.Hue/Discover", opts...)
	if err != nil {
		return nil, err
	}
	x := &hueDiscoverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hue_DiscoverClient interface {
	Recv() (*Bridge, error)
	grpc.ClientStream
}

type hueDiscoverClient struct {
	grpc.ClientStream
}

func (x *hueDiscoverClient) Recv() (*Bridge, error) {
	m := new(Bridge)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hueClient) GetBridgeState(ctx context.Context, in *Bridge, opts ...grpc.CallOption) (*BridgeState, error) {
	out := new(BridgeState)
	err := grpc.Invoke(ctx, "/hue.Hue/GetBridgeState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hueClient) ListLights(ctx context.Context, in *ListLightParams, opts ...grpc.CallOption) (Hue_ListLightsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hue_serviceDesc.Streams[1], c.cc, "/hue.Hue/ListLights", opts...)
	if err != nil {
		return nil, err
	}
	x := &hueListLightsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Hue_ListLightsClient interface {
	Recv() (*Light, error)
	grpc.ClientStream
}

type hueListLightsClient struct {
	grpc.ClientStream
}

func (x *hueListLightsClient) Recv() (*Light, error) {
	m := new(Light)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *hueClient) GetLight(ctx context.Context, in *GetLightParams, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := grpc.Invoke(ctx, "/hue.Hue/GetLight", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hueClient) SetLight(ctx context.Context, opts ...grpc.CallOption) (Hue_SetLightClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Hue_serviceDesc.Streams[2], c.cc, "/hue.Hue/SetLight", opts...)
	if err != nil {
		return nil, err
	}
	x := &hueSetLightClient{stream}
	return x, nil
}

type Hue_SetLightClient interface {
	Send(*Light) error
	Recv() (*Light, error)
	grpc.ClientStream
}

type hueSetLightClient struct {
	grpc.ClientStream
}

func (x *hueSetLightClient) Send(m *Light) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hueSetLightClient) Recv() (*Light, error) {
	m := new(Light)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Hue service

type HueServer interface {
	Discover(*DiscoverParams, Hue_DiscoverServer) error
	GetBridgeState(context.Context, *Bridge) (*BridgeState, error)
	ListLights(*ListLightParams, Hue_ListLightsServer) error
	GetLight(context.Context, *GetLightParams) (*Light, error)
	SetLight(Hue_SetLightServer) error
}

func RegisterHueServer(s *grpc.Server, srv HueServer) {
	s.RegisterService(&_Hue_serviceDesc, srv)
}

func _Hue_Discover_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DiscoverParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HueServer).Discover(m, &hueDiscoverServer{stream})
}

type Hue_DiscoverServer interface {
	Send(*Bridge) error
	grpc.ServerStream
}

type hueDiscoverServer struct {
	grpc.ServerStream
}

func (x *hueDiscoverServer) Send(m *Bridge) error {
	return x.ServerStream.SendMsg(m)
}

func _Hue_GetBridgeState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bridge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HueServer).GetBridgeState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Hue/GetBridgeState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HueServer).GetBridgeState(ctx, req.(*Bridge))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hue_ListLights_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListLightParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HueServer).ListLights(m, &hueListLightsServer{stream})
}

type Hue_ListLightsServer interface {
	Send(*Light) error
	grpc.ServerStream
}

type hueListLightsServer struct {
	grpc.ServerStream
}

func (x *hueListLightsServer) Send(m *Light) error {
	return x.ServerStream.SendMsg(m)
}

func _Hue_GetLight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLightParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HueServer).GetLight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hue.Hue/GetLight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HueServer).GetLight(ctx, req.(*GetLightParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hue_SetLight_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HueServer).SetLight(&hueSetLightServer{stream})
}

type Hue_SetLightServer interface {
	Send(*Light) error
	Recv() (*Light, error)
	grpc.ServerStream
}

type hueSetLightServer struct {
	grpc.ServerStream
}

func (x *hueSetLightServer) Send(m *Light) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hueSetLightServer) Recv() (*Light, error) {
	m := new(Light)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Hue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hue.Hue",
	HandlerType: (*HueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBridgeState",
			Handler:    _Hue_GetBridgeState_Handler,
		},
		{
			MethodName: "GetLight",
			Handler:    _Hue_GetLight_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Discover",
			Handler:       _Hue_Discover_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListLights",
			Handler:       _Hue_ListLights_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SetLight",
			Handler:       _Hue_SetLight_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protos/bridge.proto",
}

func init() { proto.RegisterFile("protos/bridge.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x8f, 0x93, 0x40,
	0x14, 0xdf, 0x01, 0x4b, 0xf1, 0x4d, 0xac, 0xf8, 0xd6, 0x03, 0x21, 0x6a, 0x9a, 0x39, 0x71, 0x50,
	0xb6, 0xa2, 0x07, 0xe3, 0x51, 0xdd, 0xac, 0x9b, 0x6c, 0xcc, 0x86, 0xde, 0xbc, 0xd1, 0x76, 0x4a,
	0x49, 0x2b, 0x34, 0x33, 0x43, 0x4d, 0xbf, 0x90, 0x5f, 0xd1, 0xab, 0xe1, 0x01, 0x66, 0x30, 0x31,
	0xd9, 0x1b, 0xf3, 0xde, 0xef, 0xcf, 0x7b, 0xbf, 0x17, 0xe0, 0xf2, 0xa8, 0x6a, 0x53, 0xeb, 0xab,
	0x95, 0x2a, 0x37, 0x85, 0x4c, 0xe8, 0x85, 0xee, 0xae, 0x91, 0xe2, 0x0a, 0x9e, 0x7c, 0xa2, 0xe2,
	0x37, 0x69, 0x7e, 0xd6, 0x6a, 0x8f, 0xaf, 0x00, 0x6e, 0x2b, 0x23, 0x55, 0x95, 0x1f, 0x6e, 0xef,
	0x43, 0x36, 0x67, 0xf1, 0xe3, 0xcc, 0xaa, 0x88, 0xef, 0xe0, 0x75, 0x04, 0x7c, 0x0d, 0xd3, 0xaa,
	0x23, 0x11, 0x8c, 0xa7, 0x98, 0xec, 0x1a, 0x99, 0x8c, 0xe4, 0xb2, 0x01, 0x82, 0x33, 0x70, 0xca,
	0x4d, 0xe8, 0x90, 0x9e, 0x53, 0x6e, 0x10, 0xe1, 0x51, 0xa3, 0xa5, 0x0a, 0x5d, 0xaa, 0xd0, 0xb7,
	0x08, 0x60, 0xf6, 0xa5, 0xd4, 0xeb, 0xfa, 0x24, 0xd5, 0x7d, 0xae, 0xf2, 0x1f, 0x5a, 0xfc, 0x72,
	0x80, 0x77, 0x82, 0x4b, 0x93, 0x1b, 0x89, 0xef, 0xc1, 0x3b, 0x94, 0xc5, 0xce, 0xe8, 0x90, 0xcd,
	0xdd, 0x98, 0xa7, 0x2f, 0x2c, 0x4b, 0x42, 0x24, 0x77, 0xd4, 0xbe, 0xae, 0x8c, 0x3a, 0x67, 0x3d,
	0xb6, 0x65, 0x15, 0xaa, 0x6e, 0x8e, 0x3a, 0x74, 0xfe, 0xc3, 0xba, 0xa1, 0x76, 0xcf, 0xea, 0xb0,
	0xf8, 0x12, 0xdc, 0xf5, 0xb6, 0xa0, 0x01, 0x79, 0xca, 0x89, 0xf2, 0xb9, 0xae, 0xb6, 0x65, 0x91,
	0xb5, 0xf5, 0xe8, 0x1a, 0xb8, 0xe5, 0x85, 0x01, 0xb8, 0x7b, 0x79, 0xee, 0x03, 0x6b, 0x3f, 0x71,
	0x0e, 0x93, 0x53, 0x7e, 0x68, 0x24, 0x2d, 0xcd, 0x53, 0x20, 0x05, 0xa2, 0x64, 0x5d, 0xe3, 0xa3,
	0xf3, 0x81, 0xb5, 0x32, 0x96, 0xf9, 0x43, 0x65, 0x88, 0x62, 0xc9, 0x88, 0x29, 0x4c, 0x48, 0x5a,
	0x3c, 0x83, 0xa7, 0x77, 0xa5, 0x36, 0xf4, 0xe8, 0x43, 0x0c, 0x60, 0x76, 0x23, 0x47, 0x95, 0x29,
	0x4c, 0x48, 0x41, 0xf8, 0xe0, 0x75, 0x3b, 0xa5, 0xbf, 0x19, 0xb8, 0x5f, 0x1b, 0x89, 0x0b, 0xf0,
	0x87, 0x1b, 0xe0, 0x25, 0x79, 0x8d, 0x4f, 0x12, 0x71, 0x2b, 0x3c, 0x71, 0xb1, 0x60, 0xf8, 0x96,
	0xe4, 0xed, 0x2b, 0xd9, 0x90, 0x28, 0xf8, 0x37, 0x6c, 0x71, 0x81, 0x29, 0xc0, 0xdf, 0x21, 0x35,
	0x3e, 0xef, 0x93, 0x19, 0x4d, 0x1d, 0x59, 0x79, 0x91, 0xcd, 0x1b, 0xf0, 0x87, 0x2d, 0xfa, 0xc1,
	0xc6, 0x4b, 0x8d, 0x09, 0x18, 0x83, 0xbf, 0x1c, 0xe0, 0x56, 0x67, 0x8c, 0x8a, 0xd9, 0x82, 0xad,
	0x3c, 0xfa, 0x1d, 0xde, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x08, 0xd9, 0xbd, 0x2d, 0x25, 0x03,
	0x00, 0x00,
}
