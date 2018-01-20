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
	GetRecordCollectionRequest
	GetRecordCollectionResponse
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

type ReleaseMetadata_Category int32

const (
	ReleaseMetadata_UNKNOWN      ReleaseMetadata_Category = 0
	ReleaseMetadata_PURCHASED    ReleaseMetadata_Category = 1
	ReleaseMetadata_UNLISTENED   ReleaseMetadata_Category = 2
	ReleaseMetadata_STAGED       ReleaseMetadata_Category = 3
	ReleaseMetadata_PRE_FRESHMAN ReleaseMetadata_Category = 4
	ReleaseMetadata_FRESHMAN     ReleaseMetadata_Category = 5
)

var ReleaseMetadata_Category_name = map[int32]string{
	0: "UNKNOWN",
	1: "PURCHASED",
	2: "UNLISTENED",
	3: "STAGED",
	4: "PRE_FRESHMAN",
	5: "FRESHMAN",
}
var ReleaseMetadata_Category_value = map[string]int32{
	"UNKNOWN":      0,
	"PURCHASED":    1,
	"UNLISTENED":   2,
	"STAGED":       3,
	"PRE_FRESHMAN": 4,
	"FRESHMAN":     5,
}

func (x ReleaseMetadata_Category) String() string {
	return proto.EnumName(ReleaseMetadata_Category_name, int32(x))
}
func (ReleaseMetadata_Category) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

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
	// The cost of the record in cents
	Cost int32 `protobuf:"varint,4,opt,name=cost" json:"cost,omitempty"`
	// If we have other copies of this
	Others bool `protobuf:"varint,5,opt,name=others" json:"others,omitempty"`
	// The id of the release this relates to
	Id int32 `protobuf:"varint,6,opt,name=id" json:"id,omitempty"`
	// The data we last recached this release
	LastCache int64                    `protobuf:"varint,7,opt,name=last_cache,json=lastCache" json:"last_cache,omitempty"`
	Category  ReleaseMetadata_Category `protobuf:"varint,8,opt,name=category,enum=recordcollection.ReleaseMetadata_Category" json:"category,omitempty"`
	// The folder this record should be placed in
	GoalFolder int32 `protobuf:"varint,9,opt,name=goal_folder,json=goalFolder" json:"goal_folder,omitempty"`
	// If a record is dirty, it needs to send updates to discogs
	Dirty bool `protobuf:"varint,10,opt,name=dirty" json:"dirty,omitempty"`
	// The folder this record should move to
	MoveFolder int32 `protobuf:"varint,11,opt,name=move_folder,json=moveFolder" json:"move_folder,omitempty"`
	// This is the rating that should be set on the record
	SetRating int32 `protobuf:"varint,12,opt,name=set_rating,json=setRating" json:"set_rating,omitempty"`
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

func (m *ReleaseMetadata) GetCategory() ReleaseMetadata_Category {
	if m != nil {
		return m.Category
	}
	return ReleaseMetadata_UNKNOWN
}

func (m *ReleaseMetadata) GetGoalFolder() int32 {
	if m != nil {
		return m.GoalFolder
	}
	return 0
}

func (m *ReleaseMetadata) GetDirty() bool {
	if m != nil {
		return m.Dirty
	}
	return false
}

func (m *ReleaseMetadata) GetMoveFolder() int32 {
	if m != nil {
		return m.MoveFolder
	}
	return 0
}

func (m *ReleaseMetadata) GetSetRating() int32 {
	if m != nil {
		return m.SetRating
	}
	return 0
}

// A request to get some records
type GetRecordsRequest struct {
	Filter *Record `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	// Forces a refresh on this record
	Force bool `protobuf:"varint,2,opt,name=force" json:"force,omitempty"`
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

func (m *GetRecordsRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
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

type GetRecordCollectionRequest struct {
}

func (m *GetRecordCollectionRequest) Reset()                    { *m = GetRecordCollectionRequest{} }
func (m *GetRecordCollectionRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRecordCollectionRequest) ProtoMessage()               {}
func (*GetRecordCollectionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type GetRecordCollectionResponse struct {
	InstanceIds []int32 `protobuf:"varint,1,rep,packed,name=instance_ids,json=instanceIds" json:"instance_ids,omitempty"`
}

func (m *GetRecordCollectionResponse) Reset()                    { *m = GetRecordCollectionResponse{} }
func (m *GetRecordCollectionResponse) String() string            { return proto.CompactTextString(m) }
func (*GetRecordCollectionResponse) ProtoMessage()               {}
func (*GetRecordCollectionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *GetRecordCollectionResponse) GetInstanceIds() []int32 {
	if m != nil {
		return m.InstanceIds
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
	proto.RegisterType((*GetRecordCollectionRequest)(nil), "recordcollection.GetRecordCollectionRequest")
	proto.RegisterType((*GetRecordCollectionResponse)(nil), "recordcollection.GetRecordCollectionResponse")
	proto.RegisterEnum("recordcollection.ReleaseMetadata_Category", ReleaseMetadata_Category_name, ReleaseMetadata_Category_value)
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
	GetRecordCollection(ctx context.Context, in *GetRecordCollectionRequest, opts ...grpc.CallOption) (*GetRecordCollectionResponse, error)
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

func (c *recordCollectionServiceClient) GetRecordCollection(ctx context.Context, in *GetRecordCollectionRequest, opts ...grpc.CallOption) (*GetRecordCollectionResponse, error) {
	out := new(GetRecordCollectionResponse)
	err := grpc.Invoke(ctx, "/recordcollection.RecordCollectionService/GetRecordCollection", in, out, c.cc, opts...)
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
	GetRecordCollection(context.Context, *GetRecordCollectionRequest) (*GetRecordCollectionResponse, error)
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

func _RecordCollectionService_GetRecordCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecordCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordCollectionServiceServer).GetRecordCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/recordcollection.RecordCollectionService/GetRecordCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordCollectionServiceServer).GetRecordCollection(ctx, req.(*GetRecordCollectionRequest))
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
		{
			MethodName: "GetRecordCollection",
			Handler:    _RecordCollectionService_GetRecordCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recordcollection.proto",
}

func init() { proto.RegisterFile("recordcollection.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x6e, 0xda, 0x48,
	0x14, 0x0d, 0x10, 0x13, 0x73, 0x61, 0xb3, 0xce, 0x64, 0x37, 0x6b, 0x91, 0x8d, 0x96, 0x38, 0xbb,
	0x5a, 0xb4, 0xca, 0x92, 0x96, 0xfe, 0xae, 0x54, 0x44, 0xc8, 0x87, 0xd2, 0xd0, 0x68, 0x08, 0xad,
	0xaa, 0xfe, 0xa0, 0x83, 0x67, 0x00, 0xab, 0x0e, 0x43, 0x3d, 0x43, 0xaa, 0xbc, 0x4f, 0x9f, 0xaf,
	0xcf, 0x50, 0xcd, 0x8c, 0x0d, 0x04, 0x08, 0x54, 0xfd, 0x85, 0xe7, 0xcc, 0x39, 0xe7, 0x1e, 0xee,
	0xbd, 0x06, 0xd8, 0x8b, 0x98, 0xcf, 0x23, 0xea, 0xf3, 0x30, 0x64, 0xbe, 0x0c, 0xf8, 0xb0, 0x32,
	0x8a, 0xb8, 0xe4, 0xc8, 0x99, 0xc7, 0x8b, 0xcf, 0xfb, 0x81, 0x1c, 0x8c, 0xbb, 0x15, 0x9f, 0xdf,
	0x9d, 0x74, 0x23, 0x2e, 0x07, 0x2c, 0x0a, 0x79, 0x3f, 0xf0, 0x4f, 0xfa, 0x9c, 0x06, 0xc2, 0xe7,
	0x7d, 0x31, 0x7d, 0x32, 0x26, 0xde, 0x08, 0x1c, 0xac, 0x6d, 0xea, 0x13, 0x1b, 0x54, 0x85, 0x2d,
	0x63, 0x2d, 0xdc, 0x54, 0x29, 0x53, 0xce, 0x57, 0xdd, 0xca, 0x42, 0x04, 0x23, 0xc2, 0x09, 0x11,
	0x95, 0xc1, 0xfa, 0x42, 0x86, 0x52, 0xb8, 0x69, 0xad, 0x40, 0x95, 0x69, 0x21, 0xcc, 0x42, 0x46,
	0x04, 0xc3, 0x86, 0xe0, 0x8d, 0x21, 0x6b, 0xc4, 0xe8, 0x58, 0xd5, 0xd1, 0x77, 0x6e, 0xaa, 0x94,
	0x7a, 0x42, 0x95, 0x50, 0xd0, 0x4b, 0xb0, 0xef, 0x98, 0x24, 0x94, 0x48, 0xe2, 0xa6, 0x35, 0xfd,
	0x70, 0x59, 0x2c, 0x4d, 0xbe, 0x8e, 0x89, 0x78, 0x22, 0xf1, 0x0e, 0xc0, 0xba, 0xe5, 0x9f, 0xd8,
	0x10, 0xfd, 0x06, 0x96, 0x54, 0x0f, 0xba, 0x66, 0x0e, 0x9b, 0x83, 0xf7, 0x2d, 0x03, 0xbf, 0xce,
	0x89, 0xd1, 0x01, 0x00, 0x25, 0x92, 0x75, 0x08, 0xa5, 0x8c, 0x6a, 0x7a, 0x06, 0xe7, 0x14, 0x52,
	0x53, 0x00, 0x3a, 0x84, 0x82, 0xbe, 0x1e, 0x8f, 0xd4, 0x07, 0xd5, 0xa1, 0x32, 0x38, 0xaf, 0x0e,
	0x6d, 0x03, 0xa1, 0x7d, 0xc8, 0xf5, 0x82, 0x90, 0x75, 0x46, 0x44, 0x0e, 0xdc, 0x8c, 0xae, 0x67,
	0x2b, 0xe0, 0x86, 0xc8, 0x01, 0x42, 0xb0, 0xe9, 0x73, 0x21, 0xdd, 0xcd, 0x52, 0xaa, 0x6c, 0x61,
	0xfd, 0x8c, 0xf6, 0x20, 0xab, 0xc7, 0x26, 0x5c, 0xab, 0x94, 0x2a, 0xdb, 0x38, 0x3e, 0xa1, 0x6d,
	0x48, 0x07, 0xd4, 0xcd, 0x6a, 0x66, 0x3a, 0xa0, 0x2a, 0x5a, 0x48, 0x84, 0xec, 0xf8, 0xc4, 0x1f,
	0x30, 0x77, 0xcb, 0x44, 0x53, 0x48, 0x5d, 0x01, 0xe8, 0x0c, 0x6c, 0x9f, 0x48, 0xd6, 0xe7, 0xd1,
	0x83, 0x6b, 0x97, 0x52, 0xe5, 0xed, 0xea, 0x7f, 0x6b, 0x7b, 0x55, 0xa9, 0xc7, 0x0a, 0x3c, 0xd1,
	0xa2, 0xbf, 0x20, 0xdf, 0xe7, 0x24, 0xec, 0xf4, 0x78, 0x48, 0x59, 0xe4, 0xe6, 0x74, 0x7d, 0x50,
	0xd0, 0x99, 0x46, 0x54, 0x33, 0x69, 0x10, 0xc9, 0x07, 0x17, 0x74, 0x5c, 0x73, 0x50, 0xb2, 0x3b,
	0x7e, 0xcf, 0x12, 0x59, 0xde, 0xc8, 0x14, 0x14, 0xcb, 0x0e, 0x00, 0x04, 0x93, 0x9d, 0x88, 0xc8,
	0x60, 0xd8, 0x77, 0x0b, 0xfa, 0x3e, 0x27, 0x98, 0xc4, 0x1a, 0xf0, 0xba, 0x60, 0x27, 0x61, 0x50,
	0x1e, 0xb6, 0xda, 0xcd, 0xab, 0xe6, 0x9b, 0x77, 0x4d, 0x67, 0x03, 0xfd, 0x02, 0xb9, 0x9b, 0x36,
	0xae, 0x5f, 0xd4, 0x5a, 0x8d, 0x53, 0x27, 0x85, 0xb6, 0x01, 0xda, 0xcd, 0xd7, 0x97, 0xad, 0xdb,
	0x46, 0xb3, 0x71, 0xea, 0xa4, 0x11, 0x40, 0xb6, 0x75, 0x5b, 0x3b, 0x6f, 0x9c, 0x3a, 0x19, 0xe4,
	0x40, 0xe1, 0x06, 0x37, 0x3a, 0x67, 0xb8, 0xd1, 0xba, 0xb8, 0xae, 0x35, 0x9d, 0x4d, 0x54, 0x00,
	0x7b, 0x72, 0xb2, 0xbc, 0x0f, 0xb0, 0x73, 0xce, 0xa4, 0xd9, 0x44, 0x81, 0xd9, 0xe7, 0x31, 0x13,
	0x12, 0x3d, 0x83, 0x6c, 0x2f, 0x08, 0x25, 0x8b, 0xe2, 0x85, 0x7c, 0x7a, 0xf1, 0x63, 0x9e, 0x6a,
	0x40, 0x8f, 0x47, 0x3e, 0xd3, 0xd3, 0xb7, 0xb1, 0x39, 0x78, 0x17, 0x80, 0x66, 0xcd, 0xc5, 0x88,
	0x0f, 0x05, 0xfb, 0x99, 0xf7, 0xca, 0x3b, 0x87, 0x5d, 0xb3, 0x4c, 0xf1, 0xc5, 0x34, 0xa8, 0x59,
	0xbb, 0xf5, 0x41, 0x0d, 0xcf, 0xbb, 0x82, 0xdf, 0x67, 0x8d, 0x1e, 0xa5, 0x4a, 0x36, 0x78, 0x9d,
	0x57, 0x42, 0xf4, 0xea, 0xe0, 0xd4, 0x28, 0x7d, 0x1c, 0xe9, 0x04, 0xb2, 0x92, 0xab, 0x77, 0x65,
	0xad, 0x8d, 0x25, 0x79, 0x8d, 0x2a, 0x93, 0x9d, 0x19, 0x93, 0x38, 0x4d, 0x05, 0x2c, 0xf3, 0xba,
	0xa5, 0xd7, 0x99, 0x68, 0x9a, 0xf7, 0x27, 0x14, 0x27, 0x9d, 0x9e, 0xfe, 0x84, 0xc5, 0x99, 0xbc,
	0x57, 0xb0, 0xbf, 0xf4, 0x36, 0x2e, 0x76, 0x08, 0x85, 0x60, 0x28, 0x24, 0x19, 0xfa, 0xac, 0x13,
	0xc4, 0x53, 0xb1, 0x70, 0x3e, 0xc1, 0x2e, 0xa9, 0xa8, 0x7e, 0xcd, 0xc0, 0x1f, 0xf3, 0xfa, 0x16,
	0x8b, 0xee, 0x03, 0x9f, 0xa1, 0xf7, 0x00, 0xd3, 0x29, 0xa3, 0xa3, 0xc5, 0xa8, 0x0b, 0x0b, 0x56,
	0xfc, 0x7b, 0x35, 0xc9, 0xe4, 0xf2, 0x36, 0xd0, 0x47, 0x28, 0xcc, 0x4e, 0x0b, 0xfd, 0xb3, 0xa8,
	0x5b, 0xb2, 0x16, 0xc5, 0x7f, 0x57, 0xd3, 0x66, 0x2b, 0xbc, 0x85, 0xdc, 0xa4, 0xfb, 0xc8, 0x5b,
	0xd4, 0xcd, 0xcf, 0xb7, 0x78, 0xb4, 0x92, 0x33, 0xf1, 0x95, 0xb0, 0xbb, 0xa4, 0xe5, 0xe8, 0x78,
	0xc5, 0x17, 0x5f, 0x98, 0x5b, 0xf1, 0xff, 0x1f, 0x64, 0x27, 0x55, 0xbb, 0x59, 0xfd, 0x6f, 0xf6,
	0xe2, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x9e, 0xa1, 0x06, 0x2c, 0x07, 0x00, 0x00,
}
