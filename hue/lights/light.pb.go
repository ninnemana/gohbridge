// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lights/light.proto

/*
Package light is a generated protocol buffer package.

It is generated from these files:
	lights/light.proto

It has these top-level messages:
	GetParams
	ListParams
	NewParams
	SearchParams
	SetOperation
	Scan
	SearchResult
	State
	SoftwareUpdate
	Streaming
	Capabilities
	Light
	LightState
	SetStateParams
*/
package light

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

type GetParams struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	ID   int32  `protobuf:"varint,3,opt,name=ID" json:"ID,omitempty"`
}

func (m *GetParams) Reset()                    { *m = GetParams{} }
func (m *GetParams) String() string            { return proto.CompactTextString(m) }
func (*GetParams) ProtoMessage()               {}
func (*GetParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetParams) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *GetParams) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GetParams) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type ListParams struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *ListParams) Reset()                    { *m = ListParams{} }
func (m *ListParams) String() string            { return proto.CompactTextString(m) }
func (*ListParams) ProtoMessage()               {}
func (*ListParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ListParams) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ListParams) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type NewParams struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *NewParams) Reset()                    { *m = NewParams{} }
func (m *NewParams) String() string            { return proto.CompactTextString(m) }
func (*NewParams) ProtoMessage()               {}
func (*NewParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NewParams) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *NewParams) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type SearchParams struct {
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *SearchParams) Reset()                    { *m = SearchParams{} }
func (m *SearchParams) String() string            { return proto.CompactTextString(m) }
func (*SearchParams) ProtoMessage()               {}
func (*SearchParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchParams) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SearchParams) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type SetOperation struct {
	On        bool               `protobuf:"varint,1,opt,name=on" json:"on,omitempty"`
	Bri       int32              `protobuf:"varint,2,opt,name=bri" json:"bri,omitempty"`
	Hue       int32              `protobuf:"varint,3,opt,name=hue" json:"hue,omitempty"`
	Sat       int32              `protobuf:"varint,4,opt,name=sat" json:"sat,omitempty"`
	Xy        map[string]float32 `protobuf:"bytes,5,rep,name=xy" json:"xy,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed32,2,opt,name=value"`
	Ct        int32              `protobuf:"varint,6,opt,name=ct" json:"ct,omitempty"`
	Alert     string             `protobuf:"bytes,7,opt,name=alert" json:"alert,omitempty"`
	Effect    string             `protobuf:"bytes,8,opt,name=effect" json:"effect,omitempty"`
	Colormode string             `protobuf:"bytes,9,opt,name=colormode" json:"colormode,omitempty"`
	Reachable bool               `protobuf:"varint,10,opt,name=reachable" json:"reachable,omitempty"`
}

func (m *SetOperation) Reset()                    { *m = SetOperation{} }
func (m *SetOperation) String() string            { return proto.CompactTextString(m) }
func (*SetOperation) ProtoMessage()               {}
func (*SetOperation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SetOperation) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *SetOperation) GetBri() int32 {
	if m != nil {
		return m.Bri
	}
	return 0
}

func (m *SetOperation) GetHue() int32 {
	if m != nil {
		return m.Hue
	}
	return 0
}

func (m *SetOperation) GetSat() int32 {
	if m != nil {
		return m.Sat
	}
	return 0
}

func (m *SetOperation) GetXy() map[string]float32 {
	if m != nil {
		return m.Xy
	}
	return nil
}

func (m *SetOperation) GetCt() int32 {
	if m != nil {
		return m.Ct
	}
	return 0
}

func (m *SetOperation) GetAlert() string {
	if m != nil {
		return m.Alert
	}
	return ""
}

func (m *SetOperation) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *SetOperation) GetColormode() string {
	if m != nil {
		return m.Colormode
	}
	return ""
}

func (m *SetOperation) GetReachable() bool {
	if m != nil {
		return m.Reachable
	}
	return false
}

type Scan struct {
	Lastscan string `protobuf:"bytes,1,opt,name=lastscan" json:"lastscan,omitempty"`
}

func (m *Scan) Reset()                    { *m = Scan{} }
func (m *Scan) String() string            { return proto.CompactTextString(m) }
func (*Scan) ProtoMessage()               {}
func (*Scan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Scan) GetLastscan() string {
	if m != nil {
		return m.Lastscan
	}
	return ""
}

type SearchResult struct {
	Success map[string]string `protobuf:"bytes,1,rep,name=success" json:"success,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SearchResult) Reset()                    { *m = SearchResult{} }
func (m *SearchResult) String() string            { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()               {}
func (*SearchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SearchResult) GetSuccess() map[string]string {
	if m != nil {
		return m.Success
	}
	return nil
}

type State struct {
	On             bool      `protobuf:"varint,1,opt,name=on" json:"on,omitempty"`
	Bri            float64   `protobuf:"fixed64,2,opt,name=bri" json:"bri,omitempty"`
	Alert          string    `protobuf:"bytes,3,opt,name=alert" json:"alert,omitempty"`
	Mode           string    `protobuf:"bytes,4,opt,name=mode" json:"mode,omitempty"`
	Reachable      bool      `protobuf:"varint,5,opt,name=reachable" json:"reachable,omitempty"`
	Hue            float64   `protobuf:"fixed64,6,opt,name=hue" json:"hue,omitempty"`
	Sat            float64   `protobuf:"fixed64,7,opt,name=sat" json:"sat,omitempty"`
	Xy             []float64 `protobuf:"fixed64,8,rep,packed,name=xy" json:"xy,omitempty"`
	Ct             float64   `protobuf:"fixed64,9,opt,name=ct" json:"ct,omitempty"`
	Effect         string    `protobuf:"bytes,10,opt,name=effect" json:"effect,omitempty"`
	Transitiontime float64   `protobuf:"fixed64,11,opt,name=transitiontime" json:"transitiontime,omitempty"`
	BriInc         float64   `protobuf:"fixed64,12,opt,name=bri_inc,json=briInc" json:"bri_inc,omitempty"`
	SatInc         float64   `protobuf:"fixed64,13,opt,name=sat_inc,json=satInc" json:"sat_inc,omitempty"`
	HueInc         float64   `protobuf:"fixed64,14,opt,name=hue_inc,json=hueInc" json:"hue_inc,omitempty"`
	CtInc          float64   `protobuf:"fixed64,15,opt,name=ct_inc,json=ctInc" json:"ct_inc,omitempty"`
	XyInc          []float64 `protobuf:"fixed64,16,rep,packed,name=xy_inc,json=xyInc" json:"xy_inc,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *State) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *State) GetBri() float64 {
	if m != nil {
		return m.Bri
	}
	return 0
}

func (m *State) GetAlert() string {
	if m != nil {
		return m.Alert
	}
	return ""
}

func (m *State) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *State) GetReachable() bool {
	if m != nil {
		return m.Reachable
	}
	return false
}

func (m *State) GetHue() float64 {
	if m != nil {
		return m.Hue
	}
	return 0
}

func (m *State) GetSat() float64 {
	if m != nil {
		return m.Sat
	}
	return 0
}

func (m *State) GetXy() []float64 {
	if m != nil {
		return m.Xy
	}
	return nil
}

func (m *State) GetCt() float64 {
	if m != nil {
		return m.Ct
	}
	return 0
}

func (m *State) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *State) GetTransitiontime() float64 {
	if m != nil {
		return m.Transitiontime
	}
	return 0
}

func (m *State) GetBriInc() float64 {
	if m != nil {
		return m.BriInc
	}
	return 0
}

func (m *State) GetSatInc() float64 {
	if m != nil {
		return m.SatInc
	}
	return 0
}

func (m *State) GetHueInc() float64 {
	if m != nil {
		return m.HueInc
	}
	return 0
}

func (m *State) GetCtInc() float64 {
	if m != nil {
		return m.CtInc
	}
	return 0
}

func (m *State) GetXyInc() []float64 {
	if m != nil {
		return m.XyInc
	}
	return nil
}

type SoftwareUpdate struct {
	State       string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	Lastinstall string `protobuf:"bytes,2,opt,name=lastinstall" json:"lastinstall,omitempty"`
}

func (m *SoftwareUpdate) Reset()                    { *m = SoftwareUpdate{} }
func (m *SoftwareUpdate) String() string            { return proto.CompactTextString(m) }
func (*SoftwareUpdate) ProtoMessage()               {}
func (*SoftwareUpdate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SoftwareUpdate) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *SoftwareUpdate) GetLastinstall() string {
	if m != nil {
		return m.Lastinstall
	}
	return ""
}

type Streaming struct {
	Renderer bool `protobuf:"varint,1,opt,name=renderer" json:"renderer,omitempty"`
	Proxy    bool `protobuf:"varint,2,opt,name=proxy" json:"proxy,omitempty"`
}

func (m *Streaming) Reset()                    { *m = Streaming{} }
func (m *Streaming) String() string            { return proto.CompactTextString(m) }
func (*Streaming) ProtoMessage()               {}
func (*Streaming) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Streaming) GetRenderer() bool {
	if m != nil {
		return m.Renderer
	}
	return false
}

func (m *Streaming) GetProxy() bool {
	if m != nil {
		return m.Proxy
	}
	return false
}

type Capabilities struct {
	Streaming *Streaming `protobuf:"bytes,1,opt,name=streaming" json:"streaming,omitempty"`
}

func (m *Capabilities) Reset()                    { *m = Capabilities{} }
func (m *Capabilities) String() string            { return proto.CompactTextString(m) }
func (*Capabilities) ProtoMessage()               {}
func (*Capabilities) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Capabilities) GetStreaming() *Streaming {
	if m != nil {
		return m.Streaming
	}
	return nil
}

type Light struct {
	State            *State          `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	Swupdate         *SoftwareUpdate `protobuf:"bytes,2,opt,name=swupdate" json:"swupdate,omitempty"`
	Type             string          `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	Name             string          `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Modelid          string          `protobuf:"bytes,5,opt,name=modelid" json:"modelid,omitempty"`
	Manufacturername string          `protobuf:"bytes,6,opt,name=manufacturername" json:"manufacturername,omitempty"`
	Capabilities     *Capabilities   `protobuf:"bytes,7,opt,name=capabilities" json:"capabilities,omitempty"`
	Uniqueid         string          `protobuf:"bytes,8,opt,name=uniqueid" json:"uniqueid,omitempty"`
	Swversion        string          `protobuf:"bytes,9,opt,name=swversion" json:"swversion,omitempty"`
	Swconfigid       string          `protobuf:"bytes,10,opt,name=swconfigid" json:"swconfigid,omitempty"`
	Productid        string          `protobuf:"bytes,11,opt,name=productid" json:"productid,omitempty"`
	ID               int32           `protobuf:"varint,12,opt,name=ID" json:"ID,omitempty"`
}

func (m *Light) Reset()                    { *m = Light{} }
func (m *Light) String() string            { return proto.CompactTextString(m) }
func (*Light) ProtoMessage()               {}
func (*Light) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Light) GetState() *State {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *Light) GetSwupdate() *SoftwareUpdate {
	if m != nil {
		return m.Swupdate
	}
	return nil
}

func (m *Light) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Light) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Light) GetModelid() string {
	if m != nil {
		return m.Modelid
	}
	return ""
}

func (m *Light) GetManufacturername() string {
	if m != nil {
		return m.Manufacturername
	}
	return ""
}

func (m *Light) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func (m *Light) GetUniqueid() string {
	if m != nil {
		return m.Uniqueid
	}
	return ""
}

func (m *Light) GetSwversion() string {
	if m != nil {
		return m.Swversion
	}
	return ""
}

func (m *Light) GetSwconfigid() string {
	if m != nil {
		return m.Swconfigid
	}
	return ""
}

func (m *Light) GetProductid() string {
	if m != nil {
		return m.Productid
	}
	return ""
}

func (m *Light) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type LightState struct {
	On             bool      `protobuf:"varint,1,opt,name=on" json:"on,omitempty"`
	Bri            float64   `protobuf:"fixed64,2,opt,name=bri" json:"bri,omitempty"`
	Alert          string    `protobuf:"bytes,3,opt,name=alert" json:"alert,omitempty"`
	Hue            float64   `protobuf:"fixed64,4,opt,name=hue" json:"hue,omitempty"`
	Sat            float64   `protobuf:"fixed64,5,opt,name=sat" json:"sat,omitempty"`
	Xy             []float64 `protobuf:"fixed64,6,rep,packed,name=xy" json:"xy,omitempty"`
	Ct             float64   `protobuf:"fixed64,7,opt,name=ct" json:"ct,omitempty"`
	Effect         string    `protobuf:"bytes,8,opt,name=effect" json:"effect,omitempty"`
	Transitiontime float64   `protobuf:"fixed64,9,opt,name=transitiontime" json:"transitiontime,omitempty"`
	BriInc         float64   `protobuf:"fixed64,10,opt,name=bri_inc,json=briInc" json:"bri_inc,omitempty"`
	SatInc         float64   `protobuf:"fixed64,11,opt,name=sat_inc,json=satInc" json:"sat_inc,omitempty"`
	HueInc         float64   `protobuf:"fixed64,12,opt,name=hue_inc,json=hueInc" json:"hue_inc,omitempty"`
	CtInc          float64   `protobuf:"fixed64,13,opt,name=ct_inc,json=ctInc" json:"ct_inc,omitempty"`
	XyInc          []float64 `protobuf:"fixed64,14,rep,packed,name=xy_inc,json=xyInc" json:"xy_inc,omitempty"`
}

func (m *LightState) Reset()                    { *m = LightState{} }
func (m *LightState) String() string            { return proto.CompactTextString(m) }
func (*LightState) ProtoMessage()               {}
func (*LightState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *LightState) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *LightState) GetBri() float64 {
	if m != nil {
		return m.Bri
	}
	return 0
}

func (m *LightState) GetAlert() string {
	if m != nil {
		return m.Alert
	}
	return ""
}

func (m *LightState) GetHue() float64 {
	if m != nil {
		return m.Hue
	}
	return 0
}

func (m *LightState) GetSat() float64 {
	if m != nil {
		return m.Sat
	}
	return 0
}

func (m *LightState) GetXy() []float64 {
	if m != nil {
		return m.Xy
	}
	return nil
}

func (m *LightState) GetCt() float64 {
	if m != nil {
		return m.Ct
	}
	return 0
}

func (m *LightState) GetEffect() string {
	if m != nil {
		return m.Effect
	}
	return ""
}

func (m *LightState) GetTransitiontime() float64 {
	if m != nil {
		return m.Transitiontime
	}
	return 0
}

func (m *LightState) GetBriInc() float64 {
	if m != nil {
		return m.BriInc
	}
	return 0
}

func (m *LightState) GetSatInc() float64 {
	if m != nil {
		return m.SatInc
	}
	return 0
}

func (m *LightState) GetHueInc() float64 {
	if m != nil {
		return m.HueInc
	}
	return 0
}

func (m *LightState) GetCtInc() float64 {
	if m != nil {
		return m.CtInc
	}
	return 0
}

func (m *LightState) GetXyInc() []float64 {
	if m != nil {
		return m.XyInc
	}
	return nil
}

type SetStateParams struct {
	Update *LightState `protobuf:"bytes,1,opt,name=update" json:"update,omitempty"`
	Host   string      `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	User   string      `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	ID     int32       `protobuf:"varint,4,opt,name=ID" json:"ID,omitempty"`
}

func (m *SetStateParams) Reset()                    { *m = SetStateParams{} }
func (m *SetStateParams) String() string            { return proto.CompactTextString(m) }
func (*SetStateParams) ProtoMessage()               {}
func (*SetStateParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SetStateParams) GetUpdate() *LightState {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *SetStateParams) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *SetStateParams) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *SetStateParams) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterType((*GetParams)(nil), "light.GetParams")
	proto.RegisterType((*ListParams)(nil), "light.ListParams")
	proto.RegisterType((*NewParams)(nil), "light.NewParams")
	proto.RegisterType((*SearchParams)(nil), "light.SearchParams")
	proto.RegisterType((*SetOperation)(nil), "light.SetOperation")
	proto.RegisterType((*Scan)(nil), "light.Scan")
	proto.RegisterType((*SearchResult)(nil), "light.SearchResult")
	proto.RegisterType((*State)(nil), "light.State")
	proto.RegisterType((*SoftwareUpdate)(nil), "light.SoftwareUpdate")
	proto.RegisterType((*Streaming)(nil), "light.Streaming")
	proto.RegisterType((*Capabilities)(nil), "light.Capabilities")
	proto.RegisterType((*Light)(nil), "light.Light")
	proto.RegisterType((*LightState)(nil), "light.LightState")
	proto.RegisterType((*SetStateParams)(nil), "light.SetStateParams")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Service service

type ServiceClient interface {
	All(ctx context.Context, in *ListParams, opts ...grpc.CallOption) (Service_AllClient, error)
	New(ctx context.Context, in *NewParams, opts ...grpc.CallOption) (*Scan, error)
	Search(ctx context.Context, in *SearchParams, opts ...grpc.CallOption) (*SearchResult, error)
	Get(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Light, error)
	SetState(ctx context.Context, in *SetStateParams, opts ...grpc.CallOption) (*Light, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) All(ctx context.Context, in *ListParams, opts ...grpc.CallOption) (Service_AllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Service_serviceDesc.Streams[0], c.cc, "/light.Service/All", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_AllClient interface {
	Recv() (*Light, error)
	grpc.ClientStream
}

type serviceAllClient struct {
	grpc.ClientStream
}

func (x *serviceAllClient) Recv() (*Light, error) {
	m := new(Light)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) New(ctx context.Context, in *NewParams, opts ...grpc.CallOption) (*Scan, error) {
	out := new(Scan)
	err := grpc.Invoke(ctx, "/light.Service/New", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Search(ctx context.Context, in *SearchParams, opts ...grpc.CallOption) (*SearchResult, error) {
	out := new(SearchResult)
	err := grpc.Invoke(ctx, "/light.Service/Search", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Get(ctx context.Context, in *GetParams, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := grpc.Invoke(ctx, "/light.Service/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SetState(ctx context.Context, in *SetStateParams, opts ...grpc.CallOption) (*Light, error) {
	out := new(Light)
	err := grpc.Invoke(ctx, "/light.Service/SetState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceServer interface {
	All(*ListParams, Service_AllServer) error
	New(context.Context, *NewParams) (*Scan, error)
	Search(context.Context, *SearchParams) (*SearchResult, error)
	Get(context.Context, *GetParams) (*Light, error)
	SetState(context.Context, *SetStateParams) (*Light, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_All_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).All(m, &serviceAllServer{stream})
}

type Service_AllServer interface {
	Send(*Light) error
	grpc.ServerStream
}

type serviceAllServer struct {
	grpc.ServerStream
}

func (x *serviceAllServer) Send(m *Light) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.Service/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).New(ctx, req.(*NewParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.Service/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Search(ctx, req.(*SearchParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.Service/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Get(ctx, req.(*GetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_SetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetStateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/light.Service/SetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SetState(ctx, req.(*SetStateParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "light.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Service_New_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Service_Search_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Service_Get_Handler,
		},
		{
			MethodName: "SetState",
			Handler:    _Service_SetState_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "All",
			Handler:       _Service_All_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lights/light.proto",
}

func init() { proto.RegisterFile("lights/light.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 972 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdb, 0x6e, 0x23, 0x45,
	0x10, 0xcd, 0xcc, 0x78, 0xc6, 0x9e, 0xb2, 0xd7, 0x84, 0xce, 0x06, 0x46, 0x06, 0x21, 0x6b, 0x1e,
	0x96, 0xb0, 0x48, 0x06, 0xbc, 0xcb, 0x45, 0x91, 0x40, 0x42, 0xbb, 0x68, 0x89, 0xb4, 0x5a, 0xd0,
	0x58, 0x48, 0xbc, 0xa1, 0x76, 0xbb, 0x1d, 0xb7, 0x18, 0xf7, 0x98, 0xee, 0x9e, 0xd8, 0xfe, 0x01,
	0xde, 0xf8, 0x19, 0xfe, 0x80, 0x1f, 0x82, 0x5f, 0x40, 0x7d, 0x99, 0x4b, 0x9c, 0x84, 0x28, 0xe2,
	0xc9, 0x5d, 0xa7, 0xaa, 0xfa, 0x72, 0xea, 0x94, 0x6b, 0x00, 0xe5, 0xec, 0x72, 0xa5, 0xe4, 0x27,
	0xe6, 0x67, 0xb2, 0x11, 0x85, 0x2a, 0x50, 0x68, 0x8c, 0xf4, 0x05, 0xc4, 0xaf, 0xa8, 0xfa, 0x11,
	0x0b, 0xbc, 0x96, 0x08, 0x41, 0xa7, 0x94, 0x54, 0x24, 0xde, 0xd8, 0x3b, 0x8b, 0x33, 0xb3, 0xd6,
	0xd8, 0xaa, 0x90, 0x2a, 0xf1, 0x2d, 0xa6, 0xd7, 0x68, 0x08, 0xfe, 0xc5, 0xcb, 0x24, 0x18, 0x7b,
	0x67, 0x61, 0xe6, 0x5f, 0xbc, 0x4c, 0x9f, 0x03, 0xbc, 0x66, 0xf2, 0x81, 0xbb, 0xa4, 0xcf, 0x20,
	0x7e, 0x43, 0xb7, 0x0f, 0x4c, 0xfa, 0x02, 0x06, 0x33, 0x8a, 0x05, 0x59, 0x3d, 0x30, 0xef, 0x4f,
	0x5f, 0x27, 0xaa, 0x1f, 0x36, 0x54, 0x60, 0xc5, 0x0a, 0xae, 0xdf, 0x50, 0x70, 0x93, 0xd6, 0xcb,
	0xfc, 0x82, 0xa3, 0x63, 0x08, 0xe6, 0x82, 0x99, 0x9c, 0x30, 0xd3, 0x4b, 0x8d, 0xac, 0x4a, 0xea,
	0x9e, 0xa9, 0x97, 0x1a, 0x91, 0x58, 0x25, 0x1d, 0x8b, 0x48, 0xac, 0xd0, 0xc7, 0xe0, 0xef, 0xf6,
	0x49, 0x38, 0x0e, 0xce, 0xfa, 0xd3, 0xf7, 0x26, 0x96, 0xdf, 0xf6, 0x31, 0x93, 0x9f, 0xf7, 0xdf,
	0x71, 0x25, 0xf6, 0x99, 0xbf, 0xdb, 0xeb, 0x23, 0x89, 0x4a, 0x22, 0x4b, 0x1b, 0x51, 0xe8, 0x31,
	0x84, 0x38, 0xa7, 0x42, 0x25, 0x5d, 0x73, 0x51, 0x6b, 0xa0, 0x77, 0x20, 0xa2, 0xcb, 0x25, 0x25,
	0x2a, 0xe9, 0x19, 0xd8, 0x59, 0xe8, 0x7d, 0x88, 0x49, 0x91, 0x17, 0x62, 0x5d, 0x2c, 0x68, 0x12,
	0x1b, 0x57, 0x03, 0x68, 0xaf, 0xa0, 0x98, 0xac, 0xf0, 0x3c, 0xa7, 0x09, 0x98, 0x57, 0x35, 0xc0,
	0xe8, 0x73, 0xe8, 0xba, 0x8b, 0xe8, 0x37, 0xfc, 0x4a, 0xf7, 0x8e, 0x2f, 0xbd, 0xd4, 0xd7, 0xb8,
	0xc2, 0x79, 0x49, 0xcd, 0xdb, 0xfd, 0xcc, 0x1a, 0xe7, 0xfe, 0x57, 0x5e, 0x9a, 0x42, 0x67, 0x46,
	0x30, 0x47, 0x23, 0xe8, 0xe5, 0x58, 0x2a, 0x49, 0x30, 0x77, 0x89, 0xb5, 0x9d, 0xfe, 0xee, 0x55,
	0x15, 0xc9, 0xa8, 0x2c, 0x73, 0x85, 0xce, 0xa1, 0x2b, 0x4b, 0x42, 0xa8, 0x94, 0x89, 0x67, 0x78,
	0x19, 0xd7, 0xbc, 0x34, 0x51, 0x93, 0x99, 0x0d, 0xb1, 0xe4, 0x54, 0x09, 0xa3, 0x73, 0x18, 0xb4,
	0x1d, 0xf7, 0x5d, 0x36, 0x6e, 0x5f, 0xf6, 0x6f, 0x1f, 0xc2, 0x99, 0xc2, 0x8a, 0xfe, 0x57, 0x69,
	0x3d, 0x5b, 0xda, 0x9a, 0xf9, 0xa0, 0xcd, 0x3c, 0x82, 0x8e, 0x21, 0xb7, 0x63, 0x75, 0x73, 0x93,
	0xd7, 0xf0, 0x80, 0xd7, 0x4a, 0x22, 0x91, 0xdd, 0xb9, 0x25, 0x91, 0xae, 0x45, 0xb4, 0x44, 0x86,
	0x46, 0x22, 0xbd, 0x71, 0x70, 0xe6, 0xb5, 0x54, 0x10, 0x9b, 0x00, 0xad, 0x82, 0xa6, 0xde, 0x70,
	0xad, 0xde, 0x4f, 0x60, 0xa8, 0x04, 0xe6, 0x92, 0x69, 0x1d, 0x29, 0xb6, 0xa6, 0x49, 0xdf, 0xe4,
	0x1c, 0xa0, 0xe8, 0x5d, 0xe8, 0xce, 0x05, 0xfb, 0x85, 0x71, 0x92, 0x0c, 0x4c, 0x40, 0x34, 0x17,
	0xec, 0x82, 0x13, 0xed, 0x90, 0x58, 0x19, 0xc7, 0x23, 0xeb, 0x90, 0x58, 0x39, 0xc7, 0xaa, 0xa4,
	0xc6, 0x31, 0xb4, 0x8e, 0x55, 0x49, 0xb5, 0xe3, 0x14, 0x22, 0x62, 0x13, 0xde, 0x32, 0x78, 0x48,
	0x94, 0x83, 0x77, 0x7b, 0x03, 0x1f, 0x9b, 0x57, 0x84, 0xbb, 0xfd, 0x05, 0x27, 0xe9, 0xf7, 0x30,
	0x9c, 0x15, 0x4b, 0xb5, 0xc5, 0x82, 0xfe, 0xb4, 0x59, 0x68, 0xe2, 0x1f, 0x43, 0x28, 0x75, 0x05,
	0x5c, 0xc1, 0xac, 0x81, 0xc6, 0xd0, 0xd7, 0x6a, 0x61, 0x5c, 0x2a, 0x9c, 0xe7, 0xae, 0x70, 0x6d,
	0x28, 0xfd, 0x1a, 0xe2, 0x99, 0x12, 0x14, 0xaf, 0x19, 0xbf, 0xd4, 0x62, 0x13, 0x94, 0x2f, 0xa8,
	0x70, 0x5d, 0xdd, 0xcb, 0x6a, 0x5b, 0x1f, 0xb0, 0x11, 0xc5, 0x6e, 0x6f, 0x36, 0xe9, 0x65, 0xd6,
	0x48, 0xbf, 0x81, 0xc1, 0x0b, 0xbc, 0xc1, 0x73, 0x96, 0x33, 0xc5, 0xa8, 0x44, 0x13, 0x88, 0x65,
	0xb5, 0x9d, 0xd9, 0xa2, 0x3f, 0x3d, 0xae, 0x34, 0x58, 0xe1, 0x59, 0x13, 0x92, 0xfe, 0x11, 0x40,
	0xf8, 0x5a, 0xbb, 0x51, 0xda, 0x7e, 0x40, 0x7f, 0x3a, 0xa8, 0xb3, 0xb0, 0xa2, 0xd5, 0x73, 0x3e,
	0x83, 0x9e, 0xdc, 0x96, 0xe6, 0xc1, 0xe6, 0x1a, 0xfd, 0xe9, 0x69, 0x15, 0x76, 0x8d, 0x8d, 0xac,
	0x0e, 0xd3, 0xc2, 0x52, 0xfb, 0x0d, 0x75, 0x6a, 0x33, 0x6b, 0x8d, 0x71, 0xbc, 0xae, 0xc5, 0xa6,
	0xd7, 0x28, 0x81, 0xae, 0x16, 0x5d, 0xce, 0x16, 0x46, 0x6a, 0x71, 0x56, 0x99, 0xe8, 0x29, 0x1c,
	0xaf, 0x31, 0x2f, 0x97, 0x98, 0xa8, 0x52, 0x50, 0x61, 0x32, 0x23, 0x13, 0x72, 0x03, 0x47, 0x5f,
	0xc2, 0x80, 0xb4, 0xe8, 0x30, 0x5a, 0xec, 0x4f, 0x4f, 0xdc, 0x25, 0xdb, 0x4c, 0x65, 0xd7, 0x02,
	0x35, 0xf3, 0x25, 0x67, 0xbf, 0x95, 0x94, 0x2d, 0xdc, 0x7f, 0x4f, 0x6d, 0xeb, 0x3e, 0x90, 0xdb,
	0x2b, 0x2a, 0x24, 0x2b, 0x78, 0xf5, 0xef, 0x53, 0x03, 0xe8, 0x03, 0x00, 0xb9, 0x25, 0x05, 0x5f,
	0xb2, 0x4b, 0xb6, 0x70, 0x3a, 0x6e, 0x21, 0x3a, 0x7b, 0x23, 0x8a, 0x45, 0x49, 0x14, 0x5b, 0x18,
	0x19, 0xc7, 0x59, 0x03, 0xb8, 0x71, 0x32, 0xa8, 0xc7, 0xc9, 0x5f, 0xbe, 0x9e, 0x27, 0x97, 0x2b,
	0xf5, 0xff, 0xda, 0xd9, 0x35, 0x67, 0xe7, 0x46, 0x73, 0x86, 0x87, 0xcd, 0x19, 0x1d, 0x34, 0x67,
	0xf7, 0x96, 0xe6, 0xec, 0xdd, 0xd3, 0x9c, 0xf1, 0x7d, 0xcd, 0x09, 0x77, 0x35, 0x67, 0xff, 0xae,
	0xe6, 0x1c, 0xdc, 0xd1, 0x9c, 0x8f, 0x6e, 0x6f, 0xce, 0x61, 0xbb, 0x39, 0x25, 0x0c, 0x67, 0xd4,
	0x12, 0xe8, 0x26, 0xe5, 0x47, 0x10, 0x39, 0xd5, 0x5a, 0x71, 0xbf, 0xed, 0x04, 0xd1, 0x30, 0x9d,
	0x45, 0x8d, 0x5e, 0x6f, 0xcc, 0xfc, 0x6a, 0xd0, 0x06, 0xad, 0x41, 0x6b, 0x0b, 0xd7, 0xa9, 0x0a,
	0x37, 0xfd, 0xc7, 0x83, 0xee, 0x8c, 0x8a, 0x2b, 0x46, 0x28, 0x7a, 0x0a, 0xc1, 0xb7, 0x79, 0x8e,
	0x9a, 0x53, 0xaa, 0xef, 0x83, 0xd1, 0xa0, 0x7d, 0x70, 0x7a, 0xf4, 0xa9, 0x87, 0x9e, 0x40, 0xf0,
	0x86, 0x6e, 0x51, 0xd5, 0xa4, 0xf5, 0x57, 0xc1, 0xa8, 0x5f, 0x75, 0x96, 0x9e, 0x34, 0x47, 0xe8,
	0x39, 0x44, 0x76, 0x88, 0xa0, 0x93, 0x6b, 0x33, 0xc5, 0x45, 0x9f, 0xdc, 0x32, 0x68, 0xd2, 0x23,
	0xf4, 0x21, 0x04, 0xaf, 0xa8, 0xaa, 0x77, 0xaf, 0x3f, 0x77, 0x0e, 0x2f, 0xa2, 0x3b, 0xbb, 0xe2,
	0x0c, 0x9d, 0x36, 0xc3, 0xbc, 0x45, 0xe2, 0x61, 0xca, 0x3c, 0x32, 0x1f, 0x53, 0xcf, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x99, 0x3f, 0xc0, 0x41, 0x62, 0x09, 0x00, 0x00,
}
