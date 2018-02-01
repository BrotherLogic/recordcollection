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
	ReleaseMetadata_UNKNOWN       ReleaseMetadata_Category = 0
	ReleaseMetadata_PURCHASED     ReleaseMetadata_Category = 1
	ReleaseMetadata_UNLISTENED    ReleaseMetadata_Category = 2
	ReleaseMetadata_STAGED        ReleaseMetadata_Category = 3
	ReleaseMetadata_PRE_FRESHMAN  ReleaseMetadata_Category = 4
	ReleaseMetadata_FRESHMAN      ReleaseMetadata_Category = 5
	ReleaseMetadata_PRE_PROFESSOR ReleaseMetadata_Category = 6
	ReleaseMetadata_PROFESSOR     ReleaseMetadata_Category = 7
)

var ReleaseMetadata_Category_name = map[int32]string{
	0: "UNKNOWN",
	1: "PURCHASED",
	2: "UNLISTENED",
	3: "STAGED",
	4: "PRE_FRESHMAN",
	5: "FRESHMAN",
	6: "PRE_PROFESSOR",
	7: "PROFESSOR",
}
var ReleaseMetadata_Category_value = map[string]int32{
	"UNKNOWN":       0,
	"PURCHASED":     1,
	"UNLISTENED":    2,
	"STAGED":        3,
	"PRE_FRESHMAN":  4,
	"FRESHMAN":      5,
	"PRE_PROFESSOR": 6,
	"PROFESSOR":     7,
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
	// The time we were last Synced
	LastSyncTime int64 `protobuf:"varint,13,opt,name=last_sync_time,json=lastSyncTime" json:"last_sync_time,omitempty"`
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

func (m *ReleaseMetadata) GetLastSyncTime() int64 {
	if m != nil {
		return m.LastSyncTime
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
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xd1, 0x6e, 0xe3, 0x44,
	0x14, 0x5d, 0x27, 0xb5, 0xeb, 0xdc, 0xa4, 0xc5, 0xbd, 0x0b, 0x8b, 0x95, 0xa5, 0x22, 0xf5, 0x2e,
	0x22, 0x42, 0x4b, 0x0a, 0xe5, 0x19, 0x89, 0x28, 0x4d, 0xdb, 0xd5, 0xb2, 0x69, 0x35, 0x6e, 0x41,
	0x88, 0x07, 0xe3, 0x7a, 0xa6, 0x89, 0x85, 0xe3, 0x09, 0x9e, 0xe9, 0xa2, 0xbe, 0xf3, 0x29, 0xfc,
	0x03, 0xbf, 0x87, 0x66, 0xc6, 0x76, 0xb2, 0x4d, 0xb7, 0x41, 0x3c, 0xd5, 0xf7, 0xcc, 0x39, 0xe7,
	0x9e, 0x99, 0xb9, 0xd3, 0xc0, 0xb3, 0x82, 0x25, 0xbc, 0xa0, 0x09, 0xcf, 0x32, 0x96, 0xc8, 0x94,
	0xe7, 0x83, 0x45, 0xc1, 0x25, 0x47, 0xef, 0x3e, 0xde, 0xfd, 0x76, 0x9a, 0xca, 0xd9, 0xed, 0xf5,
	0x20, 0xe1, 0xf3, 0xc3, 0xeb, 0x82, 0xcb, 0x19, 0x2b, 0x32, 0x3e, 0x4d, 0x93, 0xc3, 0x29, 0xa7,
	0xa9, 0x48, 0xf8, 0x54, 0x2c, 0xbf, 0x8c, 0x49, 0xb0, 0x00, 0x8f, 0x68, 0x9b, 0x51, 0x6d, 0x83,
	0x47, 0xb0, 0x6d, 0xac, 0x85, 0x6f, 0xf5, 0x9a, 0xfd, 0xf6, 0x91, 0x3f, 0x58, 0x8b, 0x60, 0x44,
	0xa4, 0x22, 0x62, 0x1f, 0xec, 0x3f, 0xe3, 0x5c, 0x0a, 0xbf, 0xa1, 0x15, 0x38, 0x58, 0x36, 0x22,
	0x2c, 0x63, 0xb1, 0x60, 0xc4, 0x10, 0x82, 0x5b, 0x70, 0x8c, 0x18, 0x5f, 0xa9, 0x3e, 0x7a, 0xcd,
	0xb7, 0x7a, 0xd6, 0x07, 0x54, 0x15, 0x05, 0xbf, 0x07, 0x77, 0xce, 0x64, 0x4c, 0x63, 0x19, 0xfb,
	0x0d, 0x4d, 0x3f, 0x78, 0x28, 0x96, 0x26, 0xbf, 0x2d, 0x89, 0xa4, 0x96, 0x04, 0xfb, 0x60, 0x5f,
	0xf2, 0xdf, 0x59, 0x8e, 0x1f, 0x83, 0x2d, 0xd5, 0x87, 0xee, 0xd9, 0x22, 0xa6, 0x08, 0xfe, 0xd9,
	0x82, 0x8f, 0xee, 0x89, 0x71, 0x1f, 0x80, 0xc6, 0x92, 0x45, 0x31, 0xa5, 0x8c, 0x6a, 0x7a, 0x93,
	0xb4, 0x14, 0x32, 0x54, 0x00, 0x1e, 0x40, 0x47, 0x2f, 0xdf, 0x2e, 0xd4, 0x1f, 0xaa, 0x43, 0x35,
	0x49, 0x5b, 0x15, 0x57, 0x06, 0xc2, 0xe7, 0xd0, 0xba, 0x49, 0x33, 0x16, 0x2d, 0x62, 0x39, 0xf3,
	0x9b, 0xba, 0x9f, 0xab, 0x80, 0x8b, 0x58, 0xce, 0x10, 0x61, 0x2b, 0xe1, 0x42, 0xfa, 0x5b, 0x3d,
	0xab, 0x6f, 0x13, 0xfd, 0x8d, 0xcf, 0xc0, 0xd1, 0xd7, 0x26, 0x7c, 0xbb, 0x67, 0xf5, 0x5d, 0x52,
	0x56, 0xb8, 0x0b, 0x8d, 0x94, 0xfa, 0x8e, 0x66, 0x36, 0x52, 0xaa, 0xa2, 0x65, 0xb1, 0x90, 0x51,
	0x12, 0x27, 0x33, 0xe6, 0x6f, 0x9b, 0x68, 0x0a, 0x19, 0x29, 0x00, 0x4f, 0xc0, 0x4d, 0x62, 0xc9,
	0xa6, 0xbc, 0xb8, 0xf3, 0xdd, 0x9e, 0xd5, 0xdf, 0x3d, 0xfa, 0x6a, 0xe3, 0x59, 0x0d, 0x46, 0xa5,
	0x82, 0xd4, 0x5a, 0xfc, 0x1c, 0xda, 0x53, 0x1e, 0x67, 0xd1, 0x0d, 0xcf, 0x28, 0x2b, 0xfc, 0x96,
	0xee, 0x0f, 0x0a, 0x3a, 0xd1, 0x88, 0x3a, 0x4c, 0x9a, 0x16, 0xf2, 0xce, 0x07, 0x1d, 0xd7, 0x14,
	0x4a, 0x36, 0xe7, 0xef, 0x58, 0x25, 0x6b, 0x1b, 0x99, 0x82, 0x4a, 0xd9, 0x3e, 0x80, 0x60, 0x32,
	0x2a, 0x62, 0x99, 0xe6, 0x53, 0xbf, 0xa3, 0xd7, 0x5b, 0x82, 0x49, 0xa2, 0x01, 0x7c, 0x09, 0xbb,
	0x7a, 0x77, 0xe2, 0x2e, 0x4f, 0x22, 0x99, 0xce, 0x99, 0xbf, 0xa3, 0x77, 0xd8, 0x51, 0x68, 0x78,
	0x97, 0x27, 0x97, 0xe9, 0x9c, 0x05, 0x7f, 0x59, 0xe0, 0x56, 0x99, 0xb1, 0x0d, 0xdb, 0x57, 0x93,
	0x37, 0x93, 0xf3, 0x9f, 0x27, 0xde, 0x13, 0xdc, 0x81, 0xd6, 0xc5, 0x15, 0x19, 0x9d, 0x0d, 0xc3,
	0xf1, 0xb1, 0x67, 0xe1, 0x2e, 0xc0, 0xd5, 0xe4, 0xc7, 0xd7, 0xe1, 0xe5, 0x78, 0x32, 0x3e, 0xf6,
	0x1a, 0x08, 0xe0, 0x84, 0x97, 0xc3, 0xd3, 0xf1, 0xb1, 0xd7, 0x44, 0x0f, 0x3a, 0x17, 0x64, 0x1c,
	0x9d, 0x90, 0x71, 0x78, 0xf6, 0x76, 0x38, 0xf1, 0xb6, 0xb0, 0x03, 0x6e, 0x5d, 0xd9, 0xb8, 0x07,
	0x3b, 0x6a, 0xfd, 0x82, 0x9c, 0x9f, 0x8c, 0xc3, 0xf0, 0x9c, 0x78, 0x8e, 0x76, 0xaf, 0xcb, 0xed,
	0xe0, 0x57, 0xd8, 0x3b, 0x65, 0xd2, 0x8c, 0xb4, 0x20, 0xec, 0x8f, 0x5b, 0x26, 0x24, 0x7e, 0x03,
	0xce, 0x4d, 0x9a, 0x49, 0x56, 0x94, 0x93, 0xfd, 0xe1, 0x17, 0x54, 0xf2, 0xd4, 0x49, 0xde, 0xf0,
	0x22, 0x61, 0x7a, 0x8c, 0x5c, 0x62, 0x8a, 0xe0, 0x0c, 0x70, 0xd5, 0x5c, 0x2c, 0x78, 0x2e, 0xd8,
	0xff, 0x79, 0xa0, 0xc1, 0x29, 0x3c, 0x35, 0x53, 0x59, 0x2e, 0x2c, 0x83, 0x9a, 0xf9, 0xdd, 0x1c,
	0xd4, 0xf0, 0x82, 0x37, 0xf0, 0xc9, 0xaa, 0xd1, 0x7b, 0xa9, 0xaa, 0xa7, 0xb0, 0xc9, 0xab, 0x22,
	0x06, 0x23, 0xf0, 0x86, 0x94, 0xbe, 0x1f, 0xe9, 0x10, 0x1c, 0xc9, 0xd5, 0xa3, 0xdb, 0x68, 0x63,
	0x4b, 0x3e, 0xa4, 0xca, 0x64, 0x6f, 0xc5, 0xa4, 0x4c, 0x33, 0x00, 0xdb, 0xbc, 0xdb, 0xc6, 0x26,
	0x13, 0x4d, 0x0b, 0x3e, 0x83, 0x6e, 0x7d, 0xd2, 0xcb, 0xff, 0x85, 0x65, 0xa6, 0xe0, 0x07, 0x78,
	0xfe, 0xe0, 0x6a, 0xd9, 0xec, 0x00, 0x3a, 0x69, 0x2e, 0x64, 0x9c, 0x27, 0x2c, 0x4a, 0xcb, 0x5b,
	0xb1, 0x49, 0xbb, 0xc2, 0x5e, 0x53, 0x71, 0xf4, 0x77, 0x13, 0x3e, 0xbd, 0xaf, 0x0f, 0x59, 0xf1,
	0x2e, 0x4d, 0x18, 0xfe, 0x02, 0xb0, 0xbc, 0x65, 0x7c, 0xb1, 0x1e, 0x75, 0x6d, 0xc0, 0xba, 0x2f,
	0x1f, 0x27, 0x99, 0x5c, 0xc1, 0x13, 0xfc, 0x0d, 0x3a, 0xab, 0xb7, 0x85, 0x5f, 0xac, 0xeb, 0x1e,
	0x18, 0x8b, 0xee, 0x97, 0x8f, 0xd3, 0x56, 0x3b, 0xfc, 0x04, 0xad, 0xfa, 0xf4, 0x31, 0x58, 0xd7,
	0xdd, 0xbf, 0xdf, 0xee, 0x8b, 0x47, 0x39, 0xb5, 0xaf, 0x84, 0xa7, 0x0f, 0x1c, 0x39, 0xbe, 0x7a,
	0x64, 0xe3, 0x6b, 0xf7, 0xd6, 0xfd, 0xfa, 0x3f, 0xb2, 0xab, 0xae, 0xd7, 0x8e, 0xfe, 0x59, 0xfc,
	0xee, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xcc, 0x90, 0x2b, 0x75, 0x07, 0x00, 0x00,
}
