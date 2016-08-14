// Code generated by protoc-gen-go.
// source: gosuv.proto
// DO NOT EDIT!

/*
Package gosuvpb is a generated protocol buffer package.

It is generated from these files:
	gosuv.proto

It has these top-level messages:
	NopRequest
	Response
	Request
	TailRequest
	ProgramInfo
	ProgramStatus
	StatusResponse
	LogLine
*/
package gosuvpb

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
const _ = proto.ProtoPackageIsVersion1

type NopRequest struct {
}

func (m *NopRequest) Reset()                    { *m = NopRequest{} }
func (m *NopRequest) String() string            { return proto.CompactTextString(m) }
func (*NopRequest) ProtoMessage()               {}
func (*NopRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Response struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type TailRequest struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	Follow bool   `protobuf:"varint,3,opt,name=follow" json:"follow,omitempty"`
}

func (m *TailRequest) Reset()                    { *m = TailRequest{} }
func (m *TailRequest) String() string            { return proto.CompactTextString(m) }
func (*TailRequest) ProtoMessage()               {}
func (*TailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ProgramInfo struct {
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Directory string   `protobuf:"bytes,3,opt,name=directory" json:"directory,omitempty"`
	Command   string   `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
	Environ   []string `protobuf:"bytes,4,rep,name=environ" json:"environ,omitempty"`
}

func (m *ProgramInfo) Reset()                    { *m = ProgramInfo{} }
func (m *ProgramInfo) String() string            { return proto.CompactTextString(m) }
func (*ProgramInfo) ProtoMessage()               {}
func (*ProgramInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ProgramStatus struct {
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Extra  string `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
}

func (m *ProgramStatus) Reset()                    { *m = ProgramStatus{} }
func (m *ProgramStatus) String() string            { return proto.CompactTextString(m) }
func (*ProgramStatus) ProtoMessage()               {}
func (*ProgramStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type StatusResponse struct {
	Programs []*ProgramStatus `protobuf:"bytes,1,rep,name=programs" json:"programs,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StatusResponse) GetPrograms() []*ProgramStatus {
	if m != nil {
		return m.Programs
	}
	return nil
}

type LogLine struct {
	Line string `protobuf:"bytes,1,opt,name=line" json:"line,omitempty"`
}

func (m *LogLine) Reset()                    { *m = LogLine{} }
func (m *LogLine) String() string            { return proto.CompactTextString(m) }
func (*LogLine) ProtoMessage()               {}
func (*LogLine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*NopRequest)(nil), "gosuvpb.NopRequest")
	proto.RegisterType((*Response)(nil), "gosuvpb.Response")
	proto.RegisterType((*Request)(nil), "gosuvpb.Request")
	proto.RegisterType((*TailRequest)(nil), "gosuvpb.TailRequest")
	proto.RegisterType((*ProgramInfo)(nil), "gosuvpb.ProgramInfo")
	proto.RegisterType((*ProgramStatus)(nil), "gosuvpb.ProgramStatus")
	proto.RegisterType((*StatusResponse)(nil), "gosuvpb.StatusResponse")
	proto.RegisterType((*LogLine)(nil), "gosuvpb.LogLine")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for GoSuv service

type GoSuvClient interface {
	Shutdown(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error)
	Version(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error)
	Status(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	Create(ctx context.Context, in *ProgramInfo, opts ...grpc.CallOption) (*Response, error)
}

type goSuvClient struct {
	cc *grpc.ClientConn
}

func NewGoSuvClient(cc *grpc.ClientConn) GoSuvClient {
	return &goSuvClient{cc}
}

func (c *goSuvClient) Shutdown(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Shutdown", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Version(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Status(ctx context.Context, in *NopRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSuvClient) Create(ctx context.Context, in *ProgramInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.GoSuv/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoSuv service

type GoSuvServer interface {
	Shutdown(context.Context, *NopRequest) (*Response, error)
	Version(context.Context, *NopRequest) (*Response, error)
	Status(context.Context, *NopRequest) (*StatusResponse, error)
	Create(context.Context, *ProgramInfo) (*Response, error)
}

func RegisterGoSuvServer(s *grpc.Server, srv GoSuvServer) {
	s.RegisterService(&_GoSuv_serviceDesc, srv)
}

func _GoSuv_Shutdown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Shutdown(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Version(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(NopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Status(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _GoSuv_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ProgramInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(GoSuvServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _GoSuv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosuvpb.GoSuv",
	HandlerType: (*GoSuvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Shutdown",
			Handler:    _GoSuv_Shutdown_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _GoSuv_Version_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _GoSuv_Status_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GoSuv_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for Program service

type ProgramClient interface {
	Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Remove(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Tail(ctx context.Context, in *TailRequest, opts ...grpc.CallOption) (Program_TailClient, error)
}

type programClient struct {
	cc *grpc.ClientConn
}

func NewProgramClient(cc *grpc.ClientConn) ProgramClient {
	return &programClient{cc}
}

func (c *programClient) Start(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.Program/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programClient) Stop(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.Program/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programClient) Remove(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/gosuvpb.Program/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programClient) Tail(ctx context.Context, in *TailRequest, opts ...grpc.CallOption) (Program_TailClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Program_serviceDesc.Streams[0], c.cc, "/gosuvpb.Program/Tail", opts...)
	if err != nil {
		return nil, err
	}
	x := &programTailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Program_TailClient interface {
	Recv() (*LogLine, error)
	grpc.ClientStream
}

type programTailClient struct {
	grpc.ClientStream
}

func (x *programTailClient) Recv() (*LogLine, error) {
	m := new(LogLine)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Program service

type ProgramServer interface {
	Start(context.Context, *Request) (*Response, error)
	Stop(context.Context, *Request) (*Response, error)
	Remove(context.Context, *Request) (*Response, error)
	Tail(*TailRequest, Program_TailServer) error
}

func RegisterProgramServer(s *grpc.Server, srv ProgramServer) {
	s.RegisterService(&_Program_serviceDesc, srv)
}

func _Program_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProgramServer).Start(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Program_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProgramServer).Stop(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Program_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ProgramServer).Remove(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Program_Tail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TailRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProgramServer).Tail(m, &programTailServer{stream})
}

type Program_TailServer interface {
	Send(*LogLine) error
	grpc.ServerStream
}

type programTailServer struct {
	grpc.ServerStream
}

func (x *programTailServer) Send(m *LogLine) error {
	return x.ServerStream.SendMsg(m)
}

var _Program_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gosuvpb.Program",
	HandlerType: (*ProgramServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Program_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Program_Stop_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Program_Remove_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tail",
			Handler:       _Program_Tail_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0x89, 0x3f, 0xc7, 0x6d, 0xa1, 0x0b, 0xa2, 0x56, 0x4f, 0xd5, 0x9e, 0xc2, 0x01, 0x83,
	0x5c, 0x0e, 0x08, 0xb8, 0x71, 0x40, 0x48, 0x15, 0xaa, 0x62, 0xc4, 0x7d, 0x13, 0x4f, 0x8d, 0x25,
	0x7b, 0xc7, 0xec, 0xae, 0x5d, 0xf8, 0x91, 0xfc, 0x08, 0xfe, 0x09, 0xfe, 0xc2, 0x4d, 0x4b, 0x84,
	0xc2, 0x29, 0x9a, 0xc9, 0x7b, 0x6f, 0xde, 0xbc, 0x59, 0x43, 0x98, 0x93, 0x6e, 0xda, 0xb8, 0x56,
	0x64, 0x88, 0x79, 0x43, 0x51, 0xaf, 0xf9, 0x21, 0xc0, 0x27, 0xaa, 0x57, 0xf8, 0xad, 0x41, 0x6d,
	0xf8, 0x33, 0xf0, 0x57, 0xa8, 0x6b, 0x92, 0x1a, 0xd9, 0x21, 0xd8, 0x1b, 0xca, 0x30, 0xb2, 0xce,
	0xad, 0xa5, 0xc3, 0x1e, 0x82, 0x57, 0xa1, 0xd6, 0x22, 0xc7, 0xe8, 0x41, 0xd7, 0x08, 0xf8, 0x29,
	0x78, 0x13, 0xab, 0x47, 0x4a, 0x51, 0x8d, 0xc8, 0x80, 0xbf, 0x85, 0xf0, 0xb3, 0x28, 0xca, 0x9d,
	0x7f, 0xb2, 0x63, 0x70, 0x65, 0x53, 0xad, 0x51, 0x0d, 0x2a, 0x4e, 0x5f, 0x5f, 0x53, 0x59, 0xd2,
	0x4d, 0xb4, 0xe8, 0x6a, 0x9f, 0x5f, 0x41, 0x78, 0xa5, 0x28, 0x57, 0xa2, 0xfa, 0x28, 0xaf, 0xe9,
	0x1e, 0xf9, 0x04, 0x82, 0xac, 0x50, 0xb8, 0x31, 0xa4, 0x7e, 0x0c, 0xf8, 0xa0, 0xb7, 0xb5, 0xa1,
	0xaa, 0x12, 0x32, 0x1b, 0x6d, 0xf5, 0x0d, 0x94, 0x6d, 0xa1, 0x48, 0x46, 0xf6, 0xf9, 0xa2, 0xb3,
	0xf3, 0x0e, 0x8e, 0x26, 0xc5, 0xd4, 0x08, 0xd3, 0xe8, 0xbf, 0x0d, 0xe9, 0xa1, 0x3f, 0xf1, 0x8f,
	0xc0, 0xc1, 0xef, 0x46, 0x89, 0x51, 0x9f, 0xbf, 0x81, 0xe3, 0x91, 0x36, 0xc7, 0xb2, 0x04, 0xbf,
	0x1e, 0xf5, 0x74, 0x27, 0xb1, 0x58, 0x86, 0xc9, 0xd3, 0x78, 0x0a, 0x33, 0xbe, 0x33, 0xa8, 0x4f,
	0xe8, 0x92, 0xf2, 0xcb, 0x42, 0x0e, 0x59, 0x96, 0xdd, 0xef, 0x38, 0x33, 0xf9, 0x65, 0x81, 0xf3,
	0x81, 0xd2, 0xa6, 0x65, 0xaf, 0xc0, 0x4f, 0xbf, 0x36, 0x26, 0xa3, 0x1b, 0xc9, 0x1e, 0xcf, 0x32,
	0xb7, 0x07, 0x39, 0x3b, 0x99, 0x9b, 0x7f, 0x0c, 0xf0, 0x03, 0x76, 0x01, 0xde, 0x17, 0x54, 0xba,
	0xa0, 0xff, 0x21, 0xbd, 0x06, 0x77, 0x0a, 0x60, 0x27, 0xe7, 0x74, 0x6e, 0xde, 0xdd, 0x77, 0x18,
	0xe7, 0xbe, 0x57, 0x28, 0x0c, 0xb2, 0x27, 0xf7, 0x37, 0xed, 0x8f, 0xb4, 0x73, 0x5c, 0xf2, 0xd3,
	0x02, 0x6f, 0x02, 0xb1, 0x18, 0x9c, 0x4e, 0x54, 0x19, 0xf6, 0x68, 0x0b, 0xf9, 0x0f, 0xab, 0xcf,
	0xc1, 0x4e, 0x0d, 0xd5, 0xfb, 0xc2, 0x5f, 0x80, 0xbb, 0xc2, 0x8a, 0x5a, 0xdc, 0x97, 0x90, 0x80,
	0xdd, 0xbf, 0xd0, 0xad, 0x75, 0xb6, 0x1e, 0xec, 0xd9, 0xad, 0xc8, 0x74, 0x3d, 0x7e, 0xf0, 0xd2,
	0x5a, 0xbb, 0xc3, 0x77, 0x73, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x35, 0x2e, 0xee, 0x46,
	0x03, 0x00, 0x00,
}