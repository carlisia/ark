// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ObjectStore.proto

package generated

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type PutObjectRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Body                 []byte   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PutObjectRequest) Reset()         { *m = PutObjectRequest{} }
func (m *PutObjectRequest) String() string { return proto.CompactTextString(m) }
func (*PutObjectRequest) ProtoMessage()    {}
func (*PutObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{0}
}

func (m *PutObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PutObjectRequest.Unmarshal(m, b)
}
func (m *PutObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PutObjectRequest.Marshal(b, m, deterministic)
}
func (m *PutObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PutObjectRequest.Merge(m, src)
}
func (m *PutObjectRequest) XXX_Size() int {
	return xxx_messageInfo_PutObjectRequest.Size(m)
}
func (m *PutObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PutObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PutObjectRequest proto.InternalMessageInfo

func (m *PutObjectRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *PutObjectRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *PutObjectRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutObjectRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type ObjectExistsRequest struct {
	Bucket               string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectExistsRequest) Reset()         { *m = ObjectExistsRequest{} }
func (m *ObjectExistsRequest) String() string { return proto.CompactTextString(m) }
func (*ObjectExistsRequest) ProtoMessage()    {}
func (*ObjectExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{1}
}

func (m *ObjectExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectExistsRequest.Unmarshal(m, b)
}
func (m *ObjectExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectExistsRequest.Marshal(b, m, deterministic)
}
func (m *ObjectExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectExistsRequest.Merge(m, src)
}
func (m *ObjectExistsRequest) XXX_Size() int {
	return xxx_messageInfo_ObjectExistsRequest.Size(m)
}
func (m *ObjectExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectExistsRequest proto.InternalMessageInfo

func (m *ObjectExistsRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ObjectExistsRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type ObjectExistsResponse struct {
	Exists               bool     `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectExistsResponse) Reset()         { *m = ObjectExistsResponse{} }
func (m *ObjectExistsResponse) String() string { return proto.CompactTextString(m) }
func (*ObjectExistsResponse) ProtoMessage()    {}
func (*ObjectExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{2}
}

func (m *ObjectExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectExistsResponse.Unmarshal(m, b)
}
func (m *ObjectExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectExistsResponse.Marshal(b, m, deterministic)
}
func (m *ObjectExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectExistsResponse.Merge(m, src)
}
func (m *ObjectExistsResponse) XXX_Size() int {
	return xxx_messageInfo_ObjectExistsResponse.Size(m)
}
func (m *ObjectExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectExistsResponse proto.InternalMessageInfo

func (m *ObjectExistsResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type GetObjectRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectRequest) Reset()         { *m = GetObjectRequest{} }
func (m *GetObjectRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectRequest) ProtoMessage()    {}
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{3}
}

func (m *GetObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectRequest.Unmarshal(m, b)
}
func (m *GetObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectRequest.Merge(m, src)
}
func (m *GetObjectRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectRequest.Size(m)
}
func (m *GetObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectRequest proto.InternalMessageInfo

func (m *GetObjectRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *GetObjectRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *GetObjectRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Bytes struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{4}
}

func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (m *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(m, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ListCommonPrefixesRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Delimiter            string   `protobuf:"bytes,3,opt,name=delimiter,proto3" json:"delimiter,omitempty"`
	Prefix               string   `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCommonPrefixesRequest) Reset()         { *m = ListCommonPrefixesRequest{} }
func (m *ListCommonPrefixesRequest) String() string { return proto.CompactTextString(m) }
func (*ListCommonPrefixesRequest) ProtoMessage()    {}
func (*ListCommonPrefixesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{5}
}

func (m *ListCommonPrefixesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCommonPrefixesRequest.Unmarshal(m, b)
}
func (m *ListCommonPrefixesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCommonPrefixesRequest.Marshal(b, m, deterministic)
}
func (m *ListCommonPrefixesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCommonPrefixesRequest.Merge(m, src)
}
func (m *ListCommonPrefixesRequest) XXX_Size() int {
	return xxx_messageInfo_ListCommonPrefixesRequest.Size(m)
}
func (m *ListCommonPrefixesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCommonPrefixesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCommonPrefixesRequest proto.InternalMessageInfo

func (m *ListCommonPrefixesRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *ListCommonPrefixesRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ListCommonPrefixesRequest) GetDelimiter() string {
	if m != nil {
		return m.Delimiter
	}
	return ""
}

func (m *ListCommonPrefixesRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type ListCommonPrefixesResponse struct {
	Prefixes             []string `protobuf:"bytes,1,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCommonPrefixesResponse) Reset()         { *m = ListCommonPrefixesResponse{} }
func (m *ListCommonPrefixesResponse) String() string { return proto.CompactTextString(m) }
func (*ListCommonPrefixesResponse) ProtoMessage()    {}
func (*ListCommonPrefixesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{6}
}

func (m *ListCommonPrefixesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCommonPrefixesResponse.Unmarshal(m, b)
}
func (m *ListCommonPrefixesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCommonPrefixesResponse.Marshal(b, m, deterministic)
}
func (m *ListCommonPrefixesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCommonPrefixesResponse.Merge(m, src)
}
func (m *ListCommonPrefixesResponse) XXX_Size() int {
	return xxx_messageInfo_ListCommonPrefixesResponse.Size(m)
}
func (m *ListCommonPrefixesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCommonPrefixesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCommonPrefixesResponse proto.InternalMessageInfo

func (m *ListCommonPrefixesResponse) GetPrefixes() []string {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

type ListObjectsRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListObjectsRequest) Reset()         { *m = ListObjectsRequest{} }
func (m *ListObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ListObjectsRequest) ProtoMessage()    {}
func (*ListObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{7}
}

func (m *ListObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListObjectsRequest.Unmarshal(m, b)
}
func (m *ListObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListObjectsRequest.Marshal(b, m, deterministic)
}
func (m *ListObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListObjectsRequest.Merge(m, src)
}
func (m *ListObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ListObjectsRequest.Size(m)
}
func (m *ListObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListObjectsRequest proto.InternalMessageInfo

func (m *ListObjectsRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *ListObjectsRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ListObjectsRequest) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type ListObjectsResponse struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListObjectsResponse) Reset()         { *m = ListObjectsResponse{} }
func (m *ListObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*ListObjectsResponse) ProtoMessage()    {}
func (*ListObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{8}
}

func (m *ListObjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListObjectsResponse.Unmarshal(m, b)
}
func (m *ListObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListObjectsResponse.Marshal(b, m, deterministic)
}
func (m *ListObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListObjectsResponse.Merge(m, src)
}
func (m *ListObjectsResponse) XXX_Size() int {
	return xxx_messageInfo_ListObjectsResponse.Size(m)
}
func (m *ListObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListObjectsResponse proto.InternalMessageInfo

func (m *ListObjectsResponse) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

type DeleteObjectRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteObjectRequest) Reset()         { *m = DeleteObjectRequest{} }
func (m *DeleteObjectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteObjectRequest) ProtoMessage()    {}
func (*DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{9}
}

func (m *DeleteObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteObjectRequest.Unmarshal(m, b)
}
func (m *DeleteObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteObjectRequest.Marshal(b, m, deterministic)
}
func (m *DeleteObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteObjectRequest.Merge(m, src)
}
func (m *DeleteObjectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteObjectRequest.Size(m)
}
func (m *DeleteObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteObjectRequest proto.InternalMessageInfo

func (m *DeleteObjectRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *DeleteObjectRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *DeleteObjectRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type CreateSignedURLRequest struct {
	Plugin               string   `protobuf:"bytes,1,opt,name=plugin,proto3" json:"plugin,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Ttl                  int64    `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSignedURLRequest) Reset()         { *m = CreateSignedURLRequest{} }
func (m *CreateSignedURLRequest) String() string { return proto.CompactTextString(m) }
func (*CreateSignedURLRequest) ProtoMessage()    {}
func (*CreateSignedURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{10}
}

func (m *CreateSignedURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSignedURLRequest.Unmarshal(m, b)
}
func (m *CreateSignedURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSignedURLRequest.Marshal(b, m, deterministic)
}
func (m *CreateSignedURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSignedURLRequest.Merge(m, src)
}
func (m *CreateSignedURLRequest) XXX_Size() int {
	return xxx_messageInfo_CreateSignedURLRequest.Size(m)
}
func (m *CreateSignedURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSignedURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSignedURLRequest proto.InternalMessageInfo

func (m *CreateSignedURLRequest) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *CreateSignedURLRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *CreateSignedURLRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *CreateSignedURLRequest) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type CreateSignedURLResponse struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateSignedURLResponse) Reset()         { *m = CreateSignedURLResponse{} }
func (m *CreateSignedURLResponse) String() string { return proto.CompactTextString(m) }
func (*CreateSignedURLResponse) ProtoMessage()    {}
func (*CreateSignedURLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31b12e27a2e940f4, []int{11}
}

func (m *CreateSignedURLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateSignedURLResponse.Unmarshal(m, b)
}
func (m *CreateSignedURLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateSignedURLResponse.Marshal(b, m, deterministic)
}
func (m *CreateSignedURLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateSignedURLResponse.Merge(m, src)
}
func (m *CreateSignedURLResponse) XXX_Size() int {
	return xxx_messageInfo_CreateSignedURLResponse.Size(m)
}
func (m *CreateSignedURLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateSignedURLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateSignedURLResponse proto.InternalMessageInfo

func (m *CreateSignedURLResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*PutObjectRequest)(nil), "generated.PutObjectRequest")
	proto.RegisterType((*ObjectExistsRequest)(nil), "generated.ObjectExistsRequest")
	proto.RegisterType((*ObjectExistsResponse)(nil), "generated.ObjectExistsResponse")
	proto.RegisterType((*GetObjectRequest)(nil), "generated.GetObjectRequest")
	proto.RegisterType((*Bytes)(nil), "generated.Bytes")
	proto.RegisterType((*ListCommonPrefixesRequest)(nil), "generated.ListCommonPrefixesRequest")
	proto.RegisterType((*ListCommonPrefixesResponse)(nil), "generated.ListCommonPrefixesResponse")
	proto.RegisterType((*ListObjectsRequest)(nil), "generated.ListObjectsRequest")
	proto.RegisterType((*ListObjectsResponse)(nil), "generated.ListObjectsResponse")
	proto.RegisterType((*DeleteObjectRequest)(nil), "generated.DeleteObjectRequest")
	proto.RegisterType((*CreateSignedURLRequest)(nil), "generated.CreateSignedURLRequest")
	proto.RegisterType((*CreateSignedURLResponse)(nil), "generated.CreateSignedURLResponse")
}

func init() { proto.RegisterFile("ObjectStore.proto", fileDescriptor_31b12e27a2e940f4) }

var fileDescriptor_31b12e27a2e940f4 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xd1, 0x6b, 0xd4, 0x4e,
	0x10, 0x26, 0x4d, 0x7e, 0xa5, 0x99, 0x06, 0x7e, 0x71, 0x4f, 0xce, 0x98, 0x6a, 0x3d, 0x17, 0x85,
	0x13, 0xe1, 0x28, 0xfa, 0xe2, 0x43, 0x41, 0xb1, 0x16, 0x11, 0x0e, 0x5a, 0x72, 0x8a, 0x3e, 0xf8,
	0x92, 0x6b, 0xc6, 0x6b, 0xbc, 0x5c, 0x12, 0x93, 0x09, 0x34, 0x8f, 0xfe, 0x99, 0xfe, 0x37, 0xb2,
	0x9b, 0x35, 0xdd, 0x5c, 0x73, 0x0a, 0xe5, 0xde, 0x66, 0x66, 0x67, 0xbf, 0xf9, 0x76, 0xf2, 0x7d,
	0x81, 0x3b, 0x67, 0xf3, 0xef, 0x78, 0x41, 0x33, 0xca, 0x0a, 0x9c, 0xe4, 0x45, 0x46, 0x19, 0xb3,
	0x17, 0x98, 0x62, 0x11, 0x12, 0x46, 0xbe, 0x33, 0xbb, 0x0c, 0x0b, 0x8c, 0x9a, 0x03, 0x7e, 0x09,
	0xee, 0x79, 0x45, 0xcd, 0x85, 0x00, 0x7f, 0x54, 0x58, 0x12, 0x1b, 0xc2, 0x6e, 0x9e, 0x54, 0x8b,
	0x38, 0xf5, 0x8c, 0x91, 0x31, 0xb6, 0x03, 0x95, 0x89, 0xfa, 0xbc, 0xba, 0x58, 0x22, 0x79, 0x3b,
	0x4d, 0xbd, 0xc9, 0x98, 0x0b, 0xe6, 0x12, 0x6b, 0xcf, 0x94, 0x45, 0x11, 0x32, 0x06, 0xd6, 0x3c,
	0x8b, 0x6a, 0xcf, 0x1a, 0x19, 0x63, 0x27, 0x90, 0x31, 0x7f, 0x0d, 0x83, 0x66, 0xcc, 0xe9, 0x55,
	0x5c, 0x52, 0xa9, 0x0d, 0x53, 0xa0, 0x46, 0x1f, 0xe8, 0x4e, 0x0b, 0xca, 0x27, 0x70, 0xb7, 0x0b,
	0x50, 0xe6, 0x59, 0x5a, 0xa2, 0x40, 0x40, 0x59, 0x91, 0x08, 0x7b, 0x81, 0xca, 0xf8, 0x47, 0x70,
	0xdf, 0xe3, 0xb6, 0x9f, 0xc6, 0x0f, 0xe0, 0xbf, 0xb7, 0x35, 0x61, 0x29, 0xde, 0x18, 0x85, 0x14,
	0x4a, 0x20, 0x27, 0x90, 0x31, 0xff, 0x69, 0xc0, 0xfd, 0x69, 0x5c, 0xd2, 0x49, 0xb6, 0x5a, 0x65,
	0xe9, 0x79, 0x81, 0xdf, 0xe2, 0x2b, 0x2c, 0x6f, 0x3b, 0xfc, 0x01, 0xd8, 0x11, 0x26, 0xf1, 0x2a,
	0x26, 0x2c, 0x14, 0x85, 0xeb, 0x82, 0x44, 0x93, 0x03, 0xe4, 0x96, 0x05, 0x9a, 0xcc, 0xf8, 0x2b,
	0xf0, 0xfb, 0x28, 0xa8, 0x65, 0xf9, 0xb0, 0x97, 0xab, 0x9a, 0x67, 0x8c, 0xcc, 0xb1, 0x1d, 0xb4,
	0x39, 0xff, 0x0a, 0x4c, 0xdc, 0x6c, 0x36, 0x76, 0x6b, 0xd6, 0xd7, 0xbc, 0xcc, 0x0e, 0xaf, 0x67,
	0x30, 0xe8, 0xa0, 0x2b, 0x42, 0x0c, 0xac, 0x25, 0xd6, 0x7f, 0xc8, 0xc8, 0x98, 0x7f, 0x86, 0xc1,
	0x3b, 0x4c, 0x90, 0x70, 0xdb, 0x1f, 0x2f, 0x81, 0xe1, 0x49, 0x81, 0x21, 0xe1, 0x2c, 0x5e, 0xa4,
	0x18, 0x7d, 0x0a, 0xa6, 0xdb, 0xd3, 0xbc, 0x0b, 0x26, 0x51, 0x22, 0x3f, 0x86, 0x19, 0x88, 0x90,
	0x3f, 0x87, 0x7b, 0x37, 0xa6, 0xa9, 0x57, 0xbb, 0x60, 0x56, 0x45, 0xa2, 0x66, 0x89, 0xf0, 0xc5,
	0x2f, 0x0b, 0xf6, 0x35, 0xdf, 0xb2, 0x23, 0xb0, 0x3e, 0xa4, 0x31, 0xb1, 0xe1, 0xa4, 0xb5, 0xee,
	0x44, 0x14, 0x14, 0x61, 0xdf, 0xd5, 0xea, 0xa7, 0xab, 0x9c, 0x6a, 0x76, 0x0c, 0x76, 0x6b, 0x65,
	0x76, 0xa0, 0x1d, 0xaf, 0x1b, 0xfc, 0xe6, 0xdd, 0xb1, 0xc1, 0xce, 0xc0, 0xd1, 0xdd, 0xc5, 0x0e,
	0xb5, 0x9e, 0x1e, 0xdf, 0xfa, 0x8f, 0x36, 0x9e, 0xab, 0x27, 0x1e, 0x83, 0xdd, 0xda, 0xaf, 0x43,
	0x67, 0xdd, 0x94, 0x1d, 0x3a, 0xd2, 0x5b, 0x47, 0x06, 0x0b, 0x1b, 0x2d, 0x76, 0x55, 0xcc, 0x9e,
	0x68, 0x9d, 0x1b, 0x7d, 0xe6, 0x3f, 0xfd, 0x47, 0x97, 0x22, 0x38, 0x85, 0x7d, 0x4d, 0x90, 0xec,
	0xe1, 0xda, 0xad, 0xae, 0x0d, 0xfc, 0xc3, 0x4d, 0xc7, 0x0a, 0xed, 0x0d, 0x38, 0xba, 0x66, 0x3b,
	0xfb, 0xeb, 0x11, 0x73, 0xcf, 0xf7, 0xfb, 0x02, 0xff, 0xaf, 0xc9, 0x85, 0x3d, 0xd6, 0x9a, 0xfa,
	0x85, 0xeb, 0xf3, 0xbf, 0xb5, 0x34, 0xdc, 0xe6, 0xbb, 0xf2, 0x5f, 0xff, 0xf2, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xdf, 0x14, 0xf6, 0x79, 0x19, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

func errUnimplemented(methodName string) error {
	return status.Errorf(codes.Unimplemented, "method %s not implemented", methodName)
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ObjectStoreClient is the client API for ObjectStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ObjectStoreClient interface {
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error)
	PutObject(ctx context.Context, opts ...grpc.CallOption) (ObjectStore_PutObjectClient, error)
	ObjectExists(ctx context.Context, in *ObjectExistsRequest, opts ...grpc.CallOption) (*ObjectExistsResponse, error)
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (ObjectStore_GetObjectClient, error)
	ListCommonPrefixes(ctx context.Context, in *ListCommonPrefixesRequest, opts ...grpc.CallOption) (*ListCommonPrefixesResponse, error)
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error)
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*Empty, error)
	CreateSignedURL(ctx context.Context, in *CreateSignedURLRequest, opts ...grpc.CallOption) (*CreateSignedURLResponse, error)
}

type objectStoreClient struct {
	cc *grpc.ClientConn
}

func NewObjectStoreClient(cc *grpc.ClientConn) ObjectStoreClient {
	return &objectStoreClient{cc}
}

func (c *objectStoreClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) PutObject(ctx context.Context, opts ...grpc.CallOption) (ObjectStore_PutObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ObjectStore_serviceDesc.Streams[0], "/generated.ObjectStore/PutObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectStorePutObjectClient{stream}
	return x, nil
}

type ObjectStore_PutObjectClient interface {
	Send(*PutObjectRequest) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type objectStorePutObjectClient struct {
	grpc.ClientStream
}

func (x *objectStorePutObjectClient) Send(m *PutObjectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *objectStorePutObjectClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectStoreClient) ObjectExists(ctx context.Context, in *ObjectExistsRequest, opts ...grpc.CallOption) (*ObjectExistsResponse, error) {
	out := new(ObjectExistsResponse)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/ObjectExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (ObjectStore_GetObjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ObjectStore_serviceDesc.Streams[1], "/generated.ObjectStore/GetObject", opts...)
	if err != nil {
		return nil, err
	}
	x := &objectStoreGetObjectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ObjectStore_GetObjectClient interface {
	Recv() (*Bytes, error)
	grpc.ClientStream
}

type objectStoreGetObjectClient struct {
	grpc.ClientStream
}

func (x *objectStoreGetObjectClient) Recv() (*Bytes, error) {
	m := new(Bytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *objectStoreClient) ListCommonPrefixes(ctx context.Context, in *ListCommonPrefixesRequest, opts ...grpc.CallOption) (*ListCommonPrefixesResponse, error) {
	out := new(ListCommonPrefixesResponse)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/ListCommonPrefixes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error) {
	out := new(ListObjectsResponse)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/ListObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/DeleteObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) CreateSignedURL(ctx context.Context, in *CreateSignedURLRequest, opts ...grpc.CallOption) (*CreateSignedURLResponse, error) {
	out := new(CreateSignedURLResponse)
	err := c.cc.Invoke(ctx, "/generated.ObjectStore/CreateSignedURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectStoreServer is the server API for ObjectStore service.
type ObjectStoreServer interface {
	Init(context.Context, *InitRequest) (*Empty, error)
	PutObject(ObjectStore_PutObjectServer) error
	ObjectExists(context.Context, *ObjectExistsRequest) (*ObjectExistsResponse, error)
	GetObject(*GetObjectRequest, ObjectStore_GetObjectServer) error
	ListCommonPrefixes(context.Context, *ListCommonPrefixesRequest) (*ListCommonPrefixesResponse, error)
	ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error)
	DeleteObject(context.Context, *DeleteObjectRequest) (*Empty, error)
	CreateSignedURL(context.Context, *CreateSignedURLRequest) (*CreateSignedURLResponse, error)
}

// UnimplementedObjectStoreServer can be embedded to have forward compatible implementations.
type UnimplementedObjectStoreServer struct {
}

func (*UnimplementedObjectStoreServer) Init(ctx context.Context, req *InitRequest) (*Empty, error) {
	return nil, errUnimplemented("Init")
}
func (*UnimplementedObjectStoreServer) PutObject(srv ObjectStore_PutObjectServer) error {
	return errUnimplemented("PutObject")
}
func (*UnimplementedObjectStoreServer) ObjectExists(ctx context.Context, req *ObjectExistsRequest) (*ObjectExistsResponse, error) {
	return nil, errUnimplemented("ObjectExists")
}
func (*UnimplementedObjectStoreServer) GetObject(req *GetObjectRequest, srv ObjectStore_GetObjectServer) error {
	return errUnimplemented("GetObject")
}
func (*UnimplementedObjectStoreServer) ListCommonPrefixes(ctx context.Context, req *ListCommonPrefixesRequest) (*ListCommonPrefixesResponse, error) {
	return nil, errUnimplemented("ListCommonPrefixes")
}
func (*UnimplementedObjectStoreServer) ListObjects(ctx context.Context, req *ListObjectsRequest) (*ListObjectsResponse, error) {
	return nil, errUnimplemented("ListObjects")
}
func (*UnimplementedObjectStoreServer) DeleteObject(ctx context.Context, req *DeleteObjectRequest) (*Empty, error) {
	return nil, errUnimplemented("DeleteObject")
}
func (*UnimplementedObjectStoreServer) CreateSignedURL(ctx context.Context, req *CreateSignedURLRequest) (*CreateSignedURLResponse, error) {
	return nil, errUnimplemented("CreateSignedURL")
}

func RegisterObjectStoreServer(s *grpc.Server, srv ObjectStoreServer) {
	s.RegisterService(&_ObjectStore_serviceDesc, srv)
}

func _ObjectStore_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_PutObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ObjectStoreServer).PutObject(&objectStorePutObjectServer{stream})
}

type ObjectStore_PutObjectServer interface {
	SendAndClose(*Empty) error
	Recv() (*PutObjectRequest, error)
	grpc.ServerStream
}

type objectStorePutObjectServer struct {
	grpc.ServerStream
}

func (x *objectStorePutObjectServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *objectStorePutObjectServer) Recv() (*PutObjectRequest, error) {
	m := new(PutObjectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ObjectStore_ObjectExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).ObjectExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/ObjectExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).ObjectExists(ctx, req.(*ObjectExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_GetObject_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetObjectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ObjectStoreServer).GetObject(m, &objectStoreGetObjectServer{stream})
}

type ObjectStore_GetObjectServer interface {
	Send(*Bytes) error
	grpc.ServerStream
}

type objectStoreGetObjectServer struct {
	grpc.ServerStream
}

func (x *objectStoreGetObjectServer) Send(m *Bytes) error {
	return x.ServerStream.SendMsg(m)
}

func _ObjectStore_ListCommonPrefixes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommonPrefixesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).ListCommonPrefixes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/ListCommonPrefixes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).ListCommonPrefixes(ctx, req.(*ListCommonPrefixesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/ListObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).ListObjects(ctx, req.(*ListObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/DeleteObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_CreateSignedURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSignedURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).CreateSignedURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generated.ObjectStore/CreateSignedURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).CreateSignedURL(ctx, req.(*CreateSignedURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObjectStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "generated.ObjectStore",
	HandlerType: (*ObjectStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _ObjectStore_Init_Handler,
		},
		{
			MethodName: "ObjectExists",
			Handler:    _ObjectStore_ObjectExists_Handler,
		},
		{
			MethodName: "ListCommonPrefixes",
			Handler:    _ObjectStore_ListCommonPrefixes_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _ObjectStore_ListObjects_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _ObjectStore_DeleteObject_Handler,
		},
		{
			MethodName: "CreateSignedURL",
			Handler:    _ObjectStore_CreateSignedURL_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutObject",
			Handler:       _ObjectStore_PutObject_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetObject",
			Handler:       _ObjectStore_GetObject_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ObjectStore.proto",
}
