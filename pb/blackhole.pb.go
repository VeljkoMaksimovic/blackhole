// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blackhole.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	blackhole.proto

It has these top-level messages:
	Strategy
	KV
	Selector
	Data
	Payload
	PutReq
	GetReq
	Resp
*/
package service

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

// Request helper messages
type Strategy struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Kind string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
}

func (m *Strategy) Reset()                    { *m = Strategy{} }
func (m *Strategy) String() string            { return proto.CompactTextString(m) }
func (*Strategy) ProtoMessage()               {}
func (*Strategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Strategy) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Strategy) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type KV struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KV) Reset()                    { *m = KV{} }
func (m *KV) String() string            { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()               {}
func (*KV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KV) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Selector struct {
	CmpKind string `protobuf:"bytes,1,opt,name=cmp_kind,json=cmpKind" json:"cmp_kind,omitempty"`
	Labels  []*KV  `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
}

func (m *Selector) Reset()                    { *m = Selector{} }
func (m *Selector) String() string            { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()               {}
func (*Selector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Selector) GetCmpKind() string {
	if m != nil {
		return m.CmpKind
	}
	return ""
}

func (m *Selector) GetLabels() []*KV {
	if m != nil {
		return m.Labels
	}
	return nil
}

type Data struct {
	Kind  string   `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Value []string `protobuf:"bytes,2,rep,name=value" json:"value,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Data) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Data) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type Payload struct {
	Strategy *Strategy `protobuf:"bytes,1,opt,name=strategy" json:"strategy,omitempty"`
	Selector *Selector `protobuf:"bytes,2,opt,name=selector" json:"selector,omitempty"`
	Data     *Data     `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Payload) GetStrategy() *Strategy {
	if m != nil {
		return m.Strategy
	}
	return nil
}

func (m *Payload) GetSelector() *Selector {
	if m != nil {
		return m.Selector
	}
	return nil
}

func (m *Payload) GetData() *Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// Reques messages
type PutReq struct {
	Id        string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Kind      string   `protobuf:"bytes,3,opt,name=kind" json:"kind,omitempty"`
	Timestamp int64    `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Namespace string   `protobuf:"bytes,5,opt,name=namespace" json:"namespace,omitempty"`
	RegionId  string   `protobuf:"bytes,6,opt,name=region_id,json=regionId" json:"region_id,omitempty"`
	ClusterId string   `protobuf:"bytes,7,opt,name=cluster_id,json=clusterId" json:"cluster_id,omitempty"`
	UserId    string   `protobuf:"bytes,8,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Payload   *Payload `protobuf:"bytes,9,opt,name=payload" json:"payload,omitempty"`
}

func (m *PutReq) Reset()                    { *m = PutReq{} }
func (m *PutReq) String() string            { return proto.CompactTextString(m) }
func (*PutReq) ProtoMessage()               {}
func (*PutReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *PutReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PutReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PutReq) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *PutReq) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PutReq) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *PutReq) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *PutReq) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *PutReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PutReq) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type GetReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Kind string `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetReq) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

// Response message
type Resp struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Resp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Strategy)(nil), "service.Strategy")
	proto.RegisterType((*KV)(nil), "service.KV")
	proto.RegisterType((*Selector)(nil), "service.Selector")
	proto.RegisterType((*Data)(nil), "service.Data")
	proto.RegisterType((*Payload)(nil), "service.Payload")
	proto.RegisterType((*PutReq)(nil), "service.PutReq")
	proto.RegisterType((*GetReq)(nil), "service.GetReq")
	proto.RegisterType((*Resp)(nil), "service.Resp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlackHoleService service

type BlackHoleServiceClient interface {
	Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*Resp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Resp, error)
}

type blackHoleServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlackHoleServiceClient(cc *grpc.ClientConn) BlackHoleServiceClient {
	return &blackHoleServiceClient{cc}
}

func (c *blackHoleServiceClient) Put(ctx context.Context, in *PutReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/service.BlackHoleService/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blackHoleServiceClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/service.BlackHoleService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BlackHoleService service

type BlackHoleServiceServer interface {
	Put(context.Context, *PutReq) (*Resp, error)
	Get(context.Context, *GetReq) (*Resp, error)
}

func RegisterBlackHoleServiceServer(s *grpc.Server, srv BlackHoleServiceServer) {
	s.RegisterService(&_BlackHoleService_serviceDesc, srv)
}

func _BlackHoleService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackHoleServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BlackHoleService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackHoleServiceServer).Put(ctx, req.(*PutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlackHoleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlackHoleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.BlackHoleService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlackHoleServiceServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlackHoleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.BlackHoleService",
	HandlerType: (*BlackHoleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _BlackHoleService_Put_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BlackHoleService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blackhole.proto",
}

func init() { proto.RegisterFile("blackhole.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 450 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x35, 0x1f, 0x9b, 0xa4, 0xb7, 0xe8, 0xd6, 0x41, 0x70, 0xfc, 0x82, 0x1a, 0x1f, 0x2c, 0xa2,
	0x65, 0xa9, 0xff, 0x40, 0x84, 0xb5, 0xf6, 0xa5, 0x4c, 0x61, 0x5f, 0x97, 0x69, 0xe6, 0x52, 0x43,
	0x27, 0x4d, 0xcc, 0x4c, 0x16, 0xfa, 0x17, 0xfc, 0xcf, 0x82, 0xcc, 0x47, 0x92, 0x8a, 0xbe, 0xdd,
	0x9c, 0x73, 0xee, 0xdc, 0x73, 0x6e, 0x6f, 0xe1, 0x7a, 0x2f, 0x79, 0x71, 0xfc, 0x51, 0x4b, 0x5c,
	0x36, 0x6d, 0xad, 0x6b, 0x92, 0x2a, 0x6c, 0x1f, 0xca, 0x02, 0xf3, 0x15, 0x64, 0x3b, 0xdd, 0x72,
	0x8d, 0x87, 0x33, 0x21, 0x10, 0xeb, 0x73, 0x83, 0x34, 0x98, 0x07, 0x8b, 0x09, 0xb3, 0xb5, 0xc1,
	0x8e, 0xe5, 0x49, 0xd0, 0xd0, 0x61, 0xa6, 0xce, 0x3f, 0x42, 0xb8, 0xb9, 0x23, 0x33, 0x88, 0x8e,
	0x78, 0xf6, 0x62, 0x53, 0x92, 0x67, 0x70, 0xf5, 0xc0, 0x65, 0x87, 0x5e, 0xec, 0x3e, 0xf2, 0xef,
	0x90, 0xed, 0x50, 0x62, 0xa1, 0xeb, 0x96, 0xbc, 0x80, 0xac, 0xa8, 0x9a, 0x7b, 0xfb, 0xa2, 0x6b,
	0x4c, 0x8b, 0xaa, 0xd9, 0x94, 0x27, 0x41, 0xde, 0x41, 0x22, 0xf9, 0x1e, 0xa5, 0xa2, 0xe1, 0x3c,
	0x5a, 0x4c, 0x57, 0xd3, 0xa5, 0xb7, 0xb8, 0xdc, 0xdc, 0x31, 0x4f, 0xe5, 0x37, 0x10, 0x7f, 0xe5,
	0x9a, 0x0f, 0xae, 0x82, 0xd1, 0xd5, 0xe5, 0xf4, 0x68, 0x9c, 0xfe, 0x2b, 0x80, 0x74, 0xcb, 0xcf,
	0xb2, 0xe6, 0x82, 0x7c, 0x82, 0x4c, 0xf9, 0xac, 0xb6, 0x73, 0xba, 0x7a, 0x3a, 0x0c, 0xe9, 0x97,
	0xc0, 0x06, 0x89, 0x95, 0x7b, 0xe3, 0x36, 0xd1, 0x5f, 0x72, 0x4f, 0xb0, 0x41, 0x42, 0xde, 0x42,
	0x2c, 0xb8, 0xe6, 0x34, 0xb2, 0xd2, 0xc7, 0x83, 0xd4, 0x18, 0x66, 0x96, 0xca, 0x7f, 0x07, 0x90,
	0x6c, 0x3b, 0xcd, 0xf0, 0x27, 0x79, 0x02, 0x61, 0xd9, 0xfb, 0x0f, 0x4b, 0x61, 0x12, 0x9d, 0x78,
	0xd5, 0xaf, 0xce, 0xd6, 0x43, 0xca, 0xe8, 0x22, 0xe5, 0x6b, 0x98, 0xe8, 0xb2, 0x42, 0xa5, 0x79,
	0xd5, 0xd0, 0x78, 0x1e, 0x2c, 0x22, 0x36, 0x02, 0x86, 0x35, 0x9d, 0xaa, 0xe1, 0x05, 0xd2, 0x2b,
	0xdb, 0x36, 0x02, 0xe4, 0x15, 0x4c, 0x5a, 0x3c, 0x94, 0xf5, 0xe9, 0xbe, 0x14, 0x34, 0xb1, 0x6c,
	0xe6, 0x80, 0xb5, 0x20, 0x6f, 0x00, 0x0a, 0xd9, 0x29, 0x8d, 0xad, 0x61, 0x53, 0xd7, 0xeb, 0x91,
	0xb5, 0x20, 0xcf, 0x21, 0xed, 0x94, 0xe3, 0x32, 0xcb, 0x25, 0xe6, 0x73, 0x2d, 0xc8, 0x07, 0x48,
	0x1b, 0xb7, 0x5f, 0x3a, 0xb1, 0xc9, 0x67, 0x43, 0x72, 0xbf, 0x77, 0xd6, 0x0b, 0xf2, 0x1b, 0x48,
	0x6e, 0xd1, 0xc6, 0xef, 0xe3, 0x06, 0xff, 0x89, 0x7b, 0x79, 0x6a, 0x14, 0x62, 0x86, 0xaa, 0x31,
	0xc7, 0x56, 0xa9, 0x43, 0x7f, 0x6c, 0x95, 0x3a, 0xac, 0x04, 0xcc, 0xbe, 0x98, 0xa3, 0xfe, 0x56,
	0x4b, 0xdc, 0xb9, 0x81, 0xe4, 0x3d, 0x44, 0xdb, 0x4e, 0x93, 0xeb, 0xd1, 0x81, 0x5d, 0xf6, 0xcb,
	0xf1, 0xc7, 0x30, 0x8f, 0xe5, 0x8f, 0x8c, 0xf0, 0x16, 0x2f, 0x85, 0xce, 0xd6, 0x3f, 0xc2, 0x7d,
	0x62, 0xff, 0x2e, 0x9f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x8c, 0x76, 0x09, 0x41, 0x03,
	0x00, 0x00,
}
