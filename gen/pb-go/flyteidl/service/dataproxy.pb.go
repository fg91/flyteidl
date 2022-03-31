// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/dataproxy.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type CreateUploadLocationResponse struct {
	// SignedUrl specifies the url to use to upload content to (e.g. https://my-bucket.s3.amazonaws.com/randomstring/suffix.tar?X-...)
	SignedUrl string `protobuf:"bytes,1,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
	// NativeUrl specifies the url in the format of the configured storage provider (e.g. s3://my-bucket/randomstring/suffix.tar)
	NativeUrl string `protobuf:"bytes,2,opt,name=native_url,json=nativeUrl,proto3" json:"native_url,omitempty"`
	// ExpiresAt defines when will the signed URL expires.
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateUploadLocationResponse) Reset()         { *m = CreateUploadLocationResponse{} }
func (m *CreateUploadLocationResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUploadLocationResponse) ProtoMessage()    {}
func (*CreateUploadLocationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bffb71366d75dab0, []int{0}
}

func (m *CreateUploadLocationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUploadLocationResponse.Unmarshal(m, b)
}
func (m *CreateUploadLocationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUploadLocationResponse.Marshal(b, m, deterministic)
}
func (m *CreateUploadLocationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUploadLocationResponse.Merge(m, src)
}
func (m *CreateUploadLocationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUploadLocationResponse.Size(m)
}
func (m *CreateUploadLocationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUploadLocationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUploadLocationResponse proto.InternalMessageInfo

func (m *CreateUploadLocationResponse) GetSignedUrl() string {
	if m != nil {
		return m.SignedUrl
	}
	return ""
}

func (m *CreateUploadLocationResponse) GetNativeUrl() string {
	if m != nil {
		return m.NativeUrl
	}
	return ""
}

func (m *CreateUploadLocationResponse) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

// CreateUploadLocationRequest specified request for the CreateUploadLocation API.
type CreateUploadLocationRequest struct {
	// Project to create the upload location for
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Domain to create the upload location for.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Suffix specifies a desired suffix for the generated location. E.g. `/file.py` or `pre/fix/file.zip`.
	// +optional. By default, the service will generate a random file name.
	Suffix string `protobuf:"bytes,3,opt,name=suffix,proto3" json:"suffix,omitempty"`
	// ExpiresIn defines a requested expiration duration for the generated url. The request will be rejected if this
	// exceeds the platform allowed max.
	// +optional. The default value comes from a global config.
	ExpiresIn            *duration.Duration `protobuf:"bytes,4,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateUploadLocationRequest) Reset()         { *m = CreateUploadLocationRequest{} }
func (m *CreateUploadLocationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUploadLocationRequest) ProtoMessage()    {}
func (*CreateUploadLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bffb71366d75dab0, []int{1}
}

func (m *CreateUploadLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUploadLocationRequest.Unmarshal(m, b)
}
func (m *CreateUploadLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUploadLocationRequest.Marshal(b, m, deterministic)
}
func (m *CreateUploadLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUploadLocationRequest.Merge(m, src)
}
func (m *CreateUploadLocationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUploadLocationRequest.Size(m)
}
func (m *CreateUploadLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUploadLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUploadLocationRequest proto.InternalMessageInfo

func (m *CreateUploadLocationRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *CreateUploadLocationRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *CreateUploadLocationRequest) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *CreateUploadLocationRequest) GetExpiresIn() *duration.Duration {
	if m != nil {
		return m.ExpiresIn
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateUploadLocationResponse)(nil), "flyteidl.service.CreateUploadLocationResponse")
	proto.RegisterType((*CreateUploadLocationRequest)(nil), "flyteidl.service.CreateUploadLocationRequest")
}

func init() { proto.RegisterFile("flyteidl/service/dataproxy.proto", fileDescriptor_bffb71366d75dab0) }

var fileDescriptor_bffb71366d75dab0 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xb5, 0x05, 0x15, 0xc5, 0x5c, 0xaa, 0x15, 0x42, 0x21, 0x94, 0x62, 0x85, 0x4b, 0x85,
	0x88, 0x2d, 0xca, 0x01, 0xca, 0xad, 0xd0, 0x0b, 0x02, 0x24, 0x14, 0xe8, 0x85, 0x4b, 0x34, 0xd9,
	0x9d, 0x75, 0x0c, 0x1b, 0xdb, 0xd8, 0xb3, 0x69, 0x72, 0x45, 0xe2, 0x05, 0xe0, 0xc0, 0x99, 0x67,
	0xe2, 0x15, 0x78, 0x00, 0x1e, 0x01, 0xad, 0xd7, 0x69, 0xa5, 0x10, 0x21, 0x4e, 0xd1, 0xcc, 0xff,
	0x39, 0xfa, 0xd6, 0x33, 0x66, 0xbc, 0xaa, 0x57, 0x84, 0xba, 0xac, 0x65, 0x40, 0xbf, 0xd0, 0x05,
	0xca, 0x12, 0x08, 0x9c, 0xb7, 0xcb, 0x95, 0x70, 0xde, 0x92, 0xcd, 0xf7, 0xd6, 0x84, 0x48, 0xc4,
	0x60, 0x5f, 0x59, 0xab, 0x6a, 0x94, 0xe0, 0xb4, 0x04, 0x63, 0x2c, 0x01, 0x69, 0x6b, 0x42, 0xc7,
	0x0f, 0x1e, 0xc4, 0x9f, 0x62, 0xa4, 0xd0, 0x8c, 0xc2, 0x39, 0x28, 0x85, 0x5e, 0x5a, 0x17, 0x89,
	0x2d, 0xf4, 0x41, 0xfa, 0xaf, 0x58, 0x4d, 0x9b, 0x4a, 0x96, 0x8d, 0x8f, 0x40, 0xca, 0xef, 0x6e,
	0xe6, 0xa4, 0xe7, 0x18, 0x08, 0xe6, 0xae, 0x03, 0x86, 0xdf, 0x33, 0xb6, 0xff, 0xdc, 0x23, 0x10,
	0x9e, 0xb9, 0xda, 0x42, 0xf9, 0xca, 0x16, 0xf1, 0xfc, 0x18, 0x83, 0xb3, 0x26, 0x60, 0x7e, 0x87,
	0xb1, 0xa0, 0x95, 0xc1, 0x72, 0xd2, 0xf8, 0xba, 0x9f, 0xf1, 0xec, 0xb0, 0x37, 0xee, 0x75, 0x9d,
	0x33, 0x5f, 0xb7, 0xb1, 0x01, 0xd2, 0x0b, 0x8c, 0xf1, 0x4e, 0x17, 0x77, 0x9d, 0x36, 0x3e, 0x66,
	0x0c, 0x97, 0x4e, 0x7b, 0x0c, 0x13, 0xa0, 0xfe, 0x15, 0x9e, 0x1d, 0x5e, 0x3f, 0x1a, 0x88, 0x4e,
	0x4a, 0xac, 0xa5, 0xc4, 0xbb, 0xb5, 0xd4, 0xb8, 0x97, 0xe8, 0x13, 0x1a, 0xfe, 0xc8, 0xd8, 0xed,
	0xed, 0x66, 0x9f, 0x1a, 0x0c, 0x94, 0xf7, 0xd9, 0x35, 0xe7, 0xed, 0x07, 0x2c, 0x28, 0x59, 0xad,
	0xcb, 0xfc, 0x26, 0xdb, 0x2d, 0xed, 0x1c, 0xb4, 0x49, 0x3e, 0xa9, 0x6a, 0xfb, 0xa1, 0xa9, 0x2a,
	0xbd, 0x8c, 0x22, 0xbd, 0x71, 0xaa, 0xf2, 0x27, 0x97, 0x92, 0xda, 0xf4, 0xaf, 0x46, 0xc9, 0x5b,
	0x7f, 0x49, 0x9e, 0xa6, 0x9b, 0xbd, 0x70, 0x7c, 0x61, 0x8e, 0xbe, 0xec, 0xb0, 0xbd, 0x53, 0x20,
	0x78, 0xd3, 0x0e, 0xfc, 0x6d, 0x37, 0xdf, 0xfc, 0x77, 0xc6, 0x6e, 0x6c, 0x13, 0xcf, 0x47, 0x62,
	0x73, 0x17, 0xc4, 0x3f, 0x3e, 0x70, 0x20, 0xfe, 0x17, 0xef, 0x26, 0x35, 0x5c, 0x7d, 0x3d, 0x79,
	0x3d, 0x78, 0xd9, 0x21, 0x81, 0x03, 0x3f, 0xf7, 0x9a, 0x70, 0x64, 0x4d, 0xbd, 0xe2, 0x33, 0x22,
	0xc7, 0xeb, 0x74, 0x80, 0xd3, 0x0c, 0x88, 0xeb, 0xc0, 0xa1, 0x28, 0x30, 0x04, 0x3d, 0xad, 0x91,
	0x57, 0xd6, 0x73, 0x82, 0xf0, 0x31, 0x70, 0x20, 0xee, 0x1b, 0xd3, 0xae, 0x89, 0xf8, 0xfc, 0xf3,
	0xd7, 0xb7, 0x9d, 0x7b, 0xc3, 0x83, 0xb8, 0xa9, 0x8b, 0x87, 0x97, 0xab, 0x2d, 0xc1, 0x93, 0xae,
	0xa0, 0xa0, 0x49, 0xe3, 0xcd, 0xd3, 0xec, 0xfe, 0xb3, 0xe3, 0xf7, 0x8f, 0x95, 0xa6, 0x59, 0x33,
	0x15, 0x85, 0x9d, 0xcb, 0xa8, 0x6d, 0xbd, 0x92, 0x17, 0x8f, 0x43, 0xa1, 0x91, 0x6e, 0x3a, 0x52,
	0x56, 0x6e, 0xbe, 0x97, 0xe9, 0x6e, 0xbc, 0xe0, 0x47, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x28,
	0xb1, 0xe0, 0xaf, 0x4a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataProxyServiceClient is the client API for DataProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataProxyServiceClient interface {
	// CreateUploadLocation creates a signed url to upload artifacts to for a given project/domain.
	CreateUploadLocation(ctx context.Context, in *CreateUploadLocationRequest, opts ...grpc.CallOption) (*CreateUploadLocationResponse, error)
}

type dataProxyServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataProxyServiceClient(cc *grpc.ClientConn) DataProxyServiceClient {
	return &dataProxyServiceClient{cc}
}

func (c *dataProxyServiceClient) CreateUploadLocation(ctx context.Context, in *CreateUploadLocationRequest, opts ...grpc.CallOption) (*CreateUploadLocationResponse, error) {
	out := new(CreateUploadLocationResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.DataProxyService/CreateUploadLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataProxyServiceServer is the server API for DataProxyService service.
type DataProxyServiceServer interface {
	// CreateUploadLocation creates a signed url to upload artifacts to for a given project/domain.
	CreateUploadLocation(context.Context, *CreateUploadLocationRequest) (*CreateUploadLocationResponse, error)
}

// UnimplementedDataProxyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataProxyServiceServer struct {
}

func (*UnimplementedDataProxyServiceServer) CreateUploadLocation(ctx context.Context, req *CreateUploadLocationRequest) (*CreateUploadLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUploadLocation not implemented")
}

func RegisterDataProxyServiceServer(s *grpc.Server, srv DataProxyServiceServer) {
	s.RegisterService(&_DataProxyService_serviceDesc, srv)
}

func _DataProxyService_CreateUploadLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUploadLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataProxyServiceServer).CreateUploadLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.DataProxyService/CreateUploadLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataProxyServiceServer).CreateUploadLocation(ctx, req.(*CreateUploadLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataProxyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.DataProxyService",
	HandlerType: (*DataProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUploadLocation",
			Handler:    _DataProxyService_CreateUploadLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/dataproxy.proto",
}
