// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recordcollection.proto

/*
Package recordcollection is a generated protocol buffer package.

It is generated from these files:
	recordcollection.proto

It has these top-level messages:
	RecordCollection
	Record
	Token
	ReleaseMetadata
	GetRecordsRequest
	GetRecordsResponse
	UpdateRecordRequest
	UpdateRecordsResponse
	AddRecordRequest
	AddRecordResponse
*/
package recordcollection

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import godiscogs "github.com/brotherlogic/godiscogs"

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

type RecordCollection struct {
	Records []*Record            `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
	Wants   []*godiscogs.Release `protobuf:"bytes,2,rep,name=wants" json:"wants,omitempty"`
}

func (m *RecordCollection) Reset()                    { *m = RecordCollection{} }
func (m *RecordCollection) String() string            { return proto.CompactTextString(m) }
func (*RecordCollection) ProtoMessage()               {}
func (*RecordCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RecordCollection) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *RecordCollection) GetWants() []*godiscogs.Release {
	if m != nil {
		return m.Wants
	}
	return nil
}

// This is a record that we own
type Record struct {
	Release  *godiscogs.Release `protobuf:"bytes,1,opt,name=release" json:"release,omitempty"`
	Metadata *ReleaseMetadata   `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Record) GetRelease() *godiscogs.Release {
	if m != nil {
		return m.Release
	}
	return nil
}

func (m *Record) GetMetadata() *ReleaseMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// This is the discogs token
type Token struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ReleaseMetadata struct {
	// The date the release was added
	DateAdded int64 `protobuf:"varint,1,opt,name=date_added,json=dateAdded" json:"date_added,omitempty"`
	// The date the release was last updated
	DateUpdated int64 `protobuf:"varint,2,opt,name=date_updated,json=dateUpdated" json:"date_updated,omitempty"`
	// The relative path to the release
	FilePath string `protobuf:"bytes,3,opt,name=file_path,json=filePath" json:"file_path,omitempty"`
	// The cost of the record in pence
	Cost int32 `protobuf:"varint,4,opt,name=cost" json:"cost,omitempty"`
	// If we have other copies of this
	Others bool `protobuf:"varint,5,opt,name=others" json:"others,omitempty"`
	// The id of the release this relates to
	Id int32 `protobuf:"varint,6,opt,name=id" json:"id,omitempty"`
	// The data we last recached this release
	LastCache int64 `protobuf:"varint,7,opt,name=last_cache,json=lastCache" json:"last_cache,omitempty"`
}

func (m *ReleaseMetadata) Reset()                    { *m = ReleaseMetadata{} }
func (m *ReleaseMetadata) String() string            { return proto.CompactTextString(m) }
func (*ReleaseMetadata) ProtoMessage()               {}
func (*ReleaseMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReleaseMetadata) GetDateAdded() int64 {
	if m != nil {
		return m.DateAdded
	}
	return 0
}

func (m *ReleaseMetadata) GetDateUpdated() int64 {
	if m != nil {
		return m.DateUpdated
	}
	return 0
}

func (m *ReleaseMetadata) GetFilePath() string {
	if m != nil {
		return m.FilePath
	}
	return ""
}

func (m *ReleaseMetadata) GetCost() int32 {
	if m != nil {
		return m.Cost
	}
	return 0
}

func (m *ReleaseMetadata) GetOthers() bool {
	if m != nil {
		return m.Others
	}
	return false
}

func (m *ReleaseMetadata) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReleaseMetadata) GetLastCache() int64 {
	if m != nil {
		return m.LastCache
	}
	return 0
}

// A request to get some records
type GetRecordsRequest struct {
	Filter *Record `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
}

func (m *GetRecordsRequest) Reset()                    { *m = GetRecordsRequest{} }
func (m *GetRecordsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRecordsRequest) ProtoMessage()               {}
func (*GetRecordsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetRecordsRequest) GetFilter() *Record {
	if m != nil {
		return m.Filter
	}
	return nil
}

// A list of records in response
type GetRecordsResponse struct {
	Records []*Record `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
}

func (m *GetRecordsResponse) Reset()                    { *m = GetRecordsResponse{} }
func (m *GetRecordsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRecordsResponse) ProtoMessage()               {}
func (*GetRecordsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetRecordsResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type UpdateRecordRequest struct {
	Update *Record `protobuf:"bytes,1,opt,name=update" json:"update,omitempty"`
}

func (m *UpdateRecordRequest) Reset()                    { *m = UpdateRecordRequest{} }
func (m *UpdateRecordRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRecordRequest) ProtoMessage()               {}
func (*UpdateRecordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateRecordRequest) GetUpdate() *Record {
	if m != nil {
		return m.Update
	}
	return nil
}

type UpdateRecordsResponse struct {
	Updated *Record `protobuf:"bytes,1,opt,name=updated" json:"updated,omitempty"`
}

func (m *UpdateRecordsResponse) Reset()                    { *m = UpdateRecordsResponse{} }
func (m *UpdateRecordsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateRecordsResponse) ProtoMessage()               {}
func (*UpdateRecordsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateRecordsResponse) GetUpdated() *Record {
	if m != nil {
		return m.Updated
	}
	return nil
}

type AddRecordRequest struct {
	ToAdd *Record `protobuf:"bytes,1,opt,name=to_add,json=toAdd" json:"to_add,omitempty"`
}

func (m *AddRecordRequest) Reset()                    { *m = AddRecordRequest{} }
func (m *AddRecordRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRecordRequest) ProtoMessage()               {}
func (*AddRecordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddRecordRequest) GetToAdd() *Record {
	if m != nil {
		return m.ToAdd
	}
	return nil
}

type AddRecordResponse struct {
	Added *Record `protobuf:"bytes,2,opt,name=added" json:"added,omitempty"`
}

func (m *AddRecordResponse) Reset()                    { *m = AddRecordResponse{} }
func (m *AddRecordResponse) String() string            { return proto.CompactTextString(m) }
func (*AddRecordResponse) ProtoMessage()               {}
func (*AddRecordResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AddRecordResponse) GetAdded() *Record {
	if m != nil {
		return m.Added
	}
	return nil
}

func init() {
	proto.RegisterType((*RecordCollection)(nil), "recordcollection.RecordCollection")
	proto.RegisterType((*Record)(nil), "recordcollection.Record")
	proto.RegisterType((*Token)(nil), "recordcollection.Token")
	proto.RegisterType((*ReleaseMetadata)(nil), "recordcollection.ReleaseMetadata")
	proto.RegisterType((*GetRecordsRequest)(nil), "recordcollection.GetRecordsRequest")
	proto.RegisterType((*GetRecordsResponse)(nil), "recordcollection.GetRecordsResponse")
	proto.RegisterType((*UpdateRecordRequest)(nil), "recordcollection.UpdateRecordRequest")
	proto.RegisterType((*UpdateRecordsResponse)(nil), "recordcollection.UpdateRecordsResponse")
	proto.RegisterType((*AddRecordRequest)(nil), "recordcollection.AddRecordRequest")
	proto.RegisterType((*AddRecordResponse)(nil), "recordcollection.AddRecordResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RecordCollectionService service

type RecordCollectionServiceClient interface {
	GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsResponse, error)
	UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordsResponse, error)
	AddRecord(ctx context.Context, in *AddRecordRequest, opts ...grpc.CallOption) (*AddRecordResponse, error)
}

type recordCollectionServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecordCollectionServiceClient(cc *grpc.ClientConn) RecordCollectionServiceClient {
	return &recordCollectionServiceClient{cc}
}

func (c *recordCollectionServiceClient) GetRecords(ctx context.Context, in *GetRecordsRequest, opts ...grpc.CallOption) (*GetRecordsResponse, error) {
	out := new(GetRecordsResponse)
	err := grpc.Invoke(ctx, "/recordcollection.RecordCollectionService/GetRecords", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordCollectionServiceClient) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordsResponse, error) {
	out := new(UpdateRecordsResponse)
	err := grpc.Invoke(ctx, "/recordcollection.RecordCollectionService/UpdateRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordCollectionServiceClient) AddRecord(ctx context.Context, in *AddRecordRequest, opts ...grpc.CallOption) (*AddRecordResponse, error) {
	out := new(AddRecordResponse)
	err := grpc.Invoke(ctx, "/recordcollection.RecordCollectionService/AddRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RecordCollectionService service

type RecordCollectionServiceServer interface {
	GetRecords(context.Context, *GetRecordsRequest) (*GetRecordsResponse, error)
	UpdateRecord(context.Context, *UpdateRecordRequest) (*UpdateRecordsResponse, error)
	AddRecord(context.Context, *AddRecordRequest) (*AddRecordResponse, error)
}

func RegisterRecordCollectionServiceServer(s *grpc.Server, srv RecordCollectionServiceServer) {
	s.RegisterService(&_RecordCollectionService_serviceDesc, srv)
}

func _RecordCollectionService_GetRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordCollectionServiceServer).GetRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recordcollection.RecordCollectionService/GetRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordCollectionServiceServer).GetRecords(ctx, req.(*GetRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordCollectionService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordCollectionServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recordcollection.RecordCollectionService/UpdateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordCollectionServiceServer).UpdateRecord(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecordCollectionService_AddRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordCollectionServiceServer).AddRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recordcollection.RecordCollectionService/AddRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordCollectionServiceServer).AddRecord(ctx, req.(*AddRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecordCollectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recordcollection.RecordCollectionService",
	HandlerType: (*RecordCollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRecords",
			Handler:    _RecordCollectionService_GetRecords_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _RecordCollectionService_UpdateRecord_Handler,
		},
		{
			MethodName: "AddRecord",
			Handler:    _RecordCollectionService_AddRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recordcollection.proto",
}

func init() { proto.RegisterFile("recordcollection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xdd, 0x6e, 0x13, 0x3f,
	0x10, 0xc5, 0xff, 0xbb, 0xe9, 0x6e, 0x93, 0x69, 0xf5, 0x27, 0x1d, 0xa0, 0x58, 0x41, 0x95, 0x52,
	0x17, 0x44, 0x2e, 0x50, 0x02, 0xe1, 0x9a, 0x8b, 0x28, 0x42, 0x45, 0x42, 0x48, 0xc8, 0x7c, 0x48,
	0x5c, 0x05, 0xc7, 0x76, 0x93, 0x15, 0xdb, 0x38, 0xac, 0x1d, 0x78, 0x2b, 0x5e, 0x84, 0x97, 0x42,
	0xb6, 0x37, 0x1f, 0xcd, 0xb6, 0x89, 0xc4, 0x55, 0xed, 0xf1, 0xef, 0x9c, 0x39, 0x9a, 0xd9, 0x06,
	0x4e, 0x0b, 0x25, 0x74, 0x21, 0x85, 0xce, 0x73, 0x25, 0x6c, 0xa6, 0x67, 0xdd, 0x79, 0xa1, 0xad,
	0xc6, 0xe6, 0x76, 0xbd, 0xf5, 0x72, 0x92, 0xd9, 0xe9, 0x62, 0xdc, 0x15, 0xfa, 0xba, 0x37, 0x2e,
	0xb4, 0x9d, 0xaa, 0x22, 0xd7, 0x93, 0x4c, 0xf4, 0x26, 0x5a, 0x66, 0x46, 0xe8, 0x89, 0x59, 0x9f,
	0x82, 0x09, 0x9d, 0x43, 0x93, 0x79, 0x9b, 0xe1, 0xca, 0x06, 0xfb, 0x70, 0x18, 0xac, 0x0d, 0x89,
	0xda, 0xb5, 0xce, 0x51, 0x9f, 0x74, 0x2b, 0x11, 0x82, 0x88, 0x2d, 0x41, 0xec, 0x40, 0xf2, 0x8b,
	0xcf, 0xac, 0x21, 0xb1, 0x57, 0x60, 0x77, 0xdd, 0x88, 0xa9, 0x5c, 0x71, 0xa3, 0x58, 0x00, 0xe8,
	0x02, 0xd2, 0x20, 0xc6, 0xe7, 0xae, 0x8f, 0x7f, 0x23, 0x51, 0x3b, 0xba, 0x43, 0xb5, 0x44, 0xf0,
	0x35, 0xd4, 0xaf, 0x95, 0xe5, 0x92, 0x5b, 0x4e, 0x62, 0x8f, 0x9f, 0xdf, 0x16, 0xcb, 0xc3, 0xef,
	0x4b, 0x90, 0xad, 0x24, 0xf4, 0x0c, 0x92, 0x4f, 0xfa, 0xbb, 0x9a, 0xe1, 0x03, 0x48, 0xac, 0x3b,
	0xf8, 0x9e, 0x0d, 0x16, 0x2e, 0xf4, 0x4f, 0x04, 0xf7, 0xb6, 0xc4, 0x78, 0x06, 0x20, 0xb9, 0x55,
	0x23, 0x2e, 0xa5, 0x92, 0x1e, 0xaf, 0xb1, 0x86, 0xab, 0x0c, 0x5c, 0x01, 0xcf, 0xe1, 0xd8, 0x3f,
	0x2f, 0xe6, 0xee, 0x8f, 0xf4, 0xa1, 0x6a, 0xec, 0xc8, 0x5d, 0x3e, 0x87, 0x12, 0x3e, 0x86, 0xc6,
	0x55, 0x96, 0xab, 0xd1, 0x9c, 0xdb, 0x29, 0xa9, 0xf9, 0x7e, 0x75, 0x57, 0xf8, 0xc0, 0xed, 0x14,
	0x11, 0x0e, 0x84, 0x36, 0x96, 0x1c, 0xb4, 0xa3, 0x4e, 0xc2, 0xfc, 0x19, 0x4f, 0x21, 0xf5, 0x6b,
	0x33, 0x24, 0x69, 0x47, 0x9d, 0x3a, 0x2b, 0x6f, 0xf8, 0x3f, 0xc4, 0x99, 0x24, 0xa9, 0x27, 0xe3,
	0x4c, 0xba, 0x68, 0x39, 0x37, 0x76, 0x24, 0xb8, 0x98, 0x2a, 0x72, 0x18, 0xa2, 0xb9, 0xca, 0xd0,
	0x15, 0xe8, 0x1b, 0x38, 0xb9, 0x54, 0x36, 0x8c, 0xd9, 0x30, 0xf5, 0x63, 0xa1, 0x8c, 0xc5, 0x17,
	0x90, 0x5e, 0x65, 0xb9, 0x55, 0x45, 0x39, 0xed, 0xbb, 0xb7, 0x5a, 0x72, 0xf4, 0x2d, 0xe0, 0xa6,
	0x8d, 0x99, 0xeb, 0x99, 0x51, 0xff, 0xf2, 0x79, 0xd0, 0x4b, 0xb8, 0x1f, 0x66, 0x52, 0x3e, 0xac,
	0x23, 0x85, 0xe9, 0xed, 0x8f, 0x14, 0x38, 0xfa, 0x0e, 0x1e, 0x6e, 0x1a, 0xdd, 0x48, 0xb5, 0x5c,
	0xc4, 0x3e, 0xaf, 0x25, 0x48, 0x87, 0xd0, 0x1c, 0x48, 0x79, 0x33, 0x52, 0x0f, 0x52, 0xab, 0xdd,
	0xca, 0xf7, 0xda, 0x24, 0x56, 0x0f, 0xa4, 0x33, 0x39, 0xd9, 0x30, 0x29, 0xd3, 0x74, 0x21, 0x09,
	0x5f, 0x4d, 0xbc, 0xcf, 0xc4, 0x63, 0xfd, 0xdf, 0x31, 0x3c, 0xda, 0xfe, 0x3f, 0xfc, 0xa8, 0x8a,
	0x9f, 0x99, 0x50, 0xf8, 0x15, 0x60, 0xbd, 0x05, 0xbc, 0xa8, 0x5a, 0x55, 0x56, 0xdd, 0x7a, 0xb2,
	0x1b, 0x0a, 0x21, 0xe9, 0x7f, 0xf8, 0x0d, 0x8e, 0x37, 0xa7, 0x89, 0x4f, 0xab, 0xba, 0x5b, 0xd6,
	0xd6, 0x7a, 0xb6, 0x1b, 0xdb, 0xec, 0xf0, 0x05, 0x1a, 0xab, 0xe9, 0x20, 0xad, 0xea, 0xb6, 0xe7,
	0xdf, 0xba, 0xd8, 0xc9, 0x2c, 0x7d, 0xc7, 0xa9, 0xff, 0xf9, 0x7a, 0xf5, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x01, 0xb9, 0x92, 0x46, 0x1d, 0x05, 0x00, 0x00,
}
