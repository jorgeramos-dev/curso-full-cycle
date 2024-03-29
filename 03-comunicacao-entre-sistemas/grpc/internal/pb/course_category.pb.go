// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course_category.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Blank struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blank) Reset()         { *m = Blank{} }
func (m *Blank) String() string { return proto.CompactTextString(m) }
func (*Blank) ProtoMessage()    {}
func (*Blank) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{0}
}

func (m *Blank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blank.Unmarshal(m, b)
}
func (m *Blank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blank.Marshal(b, m, deterministic)
}
func (m *Blank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blank.Merge(m, src)
}
func (m *Blank) XXX_Size() int {
	return xxx_messageInfo_Blank.Size(m)
}
func (m *Blank) XXX_DiscardUnknown() {
	xxx_messageInfo_Blank.DiscardUnknown(m)
}

var xxx_messageInfo_Blank proto.InternalMessageInfo

type Category struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{1}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateCategoryRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCategoryRequest) Reset()         { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()    {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{2}
}

func (m *CreateCategoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateCategoryRequest.Unmarshal(m, b)
}
func (m *CreateCategoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateCategoryRequest.Marshal(b, m, deterministic)
}
func (m *CreateCategoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCategoryRequest.Merge(m, src)
}
func (m *CreateCategoryRequest) XXX_Size() int {
	return xxx_messageInfo_CreateCategoryRequest.Size(m)
}
func (m *CreateCategoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCategoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCategoryRequest proto.InternalMessageInfo

func (m *CreateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCategoryRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CategoryList struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CategoryList) Reset()         { *m = CategoryList{} }
func (m *CategoryList) String() string { return proto.CompactTextString(m) }
func (*CategoryList) ProtoMessage()    {}
func (*CategoryList) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{3}
}

func (m *CategoryList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryList.Unmarshal(m, b)
}
func (m *CategoryList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryList.Marshal(b, m, deterministic)
}
func (m *CategoryList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryList.Merge(m, src)
}
func (m *CategoryList) XXX_Size() int {
	return xxx_messageInfo_CategoryList.Size(m)
}
func (m *CategoryList) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryList.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryList proto.InternalMessageInfo

func (m *CategoryList) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type CategoryGetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryGetRequest) Reset()         { *m = CategoryGetRequest{} }
func (m *CategoryGetRequest) String() string { return proto.CompactTextString(m) }
func (*CategoryGetRequest) ProtoMessage()    {}
func (*CategoryGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde79cb73866aac8, []int{4}
}

func (m *CategoryGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryGetRequest.Unmarshal(m, b)
}
func (m *CategoryGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryGetRequest.Marshal(b, m, deterministic)
}
func (m *CategoryGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryGetRequest.Merge(m, src)
}
func (m *CategoryGetRequest) XXX_Size() int {
	return xxx_messageInfo_CategoryGetRequest.Size(m)
}
func (m *CategoryGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryGetRequest proto.InternalMessageInfo

func (m *CategoryGetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Blank)(nil), "pb.Blank")
	proto.RegisterType((*Category)(nil), "pb.Category")
	proto.RegisterType((*CreateCategoryRequest)(nil), "pb.CreateCategoryRequest")
	proto.RegisterType((*CategoryList)(nil), "pb.CategoryList")
	proto.RegisterType((*CategoryGetRequest)(nil), "pb.CategoryGetRequest")
}

func init() {
	proto.RegisterFile("proto/course_category.proto", fileDescriptor_fde79cb73866aac8)
}

var fileDescriptor_fde79cb73866aac8 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbb, 0xdb, 0xff, 0x5f, 0xed, 0xa4, 0x56, 0x19, 0x54, 0xa2, 0x5e, 0xea, 0xe2, 0x21,
	0x07, 0x69, 0xa5, 0xe2, 0x49, 0x4f, 0xed, 0xa1, 0x17, 0x05, 0x49, 0x6f, 0x5e, 0x64, 0x93, 0x0c,
	0xb2, 0x98, 0x26, 0x71, 0xb3, 0x15, 0xfc, 0x9c, 0x7e, 0x21, 0xc9, 0xd2, 0xc4, 0xa4, 0x46, 0xc1,
	0xdb, 0x32, 0xf3, 0xf8, 0xcd, 0x7b, 0x8f, 0x85, 0xd3, 0x4c, 0xa7, 0x26, 0x1d, 0x87, 0xe9, 0x4a,
	0xe7, 0xf4, 0x14, 0x4a, 0x43, 0xcf, 0xa9, 0x7e, 0x1f, 0xd9, 0x29, 0xf2, 0x2c, 0x10, 0xdb, 0xf0,
	0x7f, 0x1a, 0xcb, 0xe4, 0x45, 0x3c, 0xc0, 0xce, 0x6c, 0xbd, 0xc6, 0x01, 0x70, 0x15, 0xb9, 0x6c,
	0xc8, 0xbc, 0x9e, 0xcf, 0x55, 0x84, 0x08, 0xff, 0x12, 0xb9, 0x24, 0x97, 0xdb, 0x89, 0x7d, 0xe3,
	0x10, 0x9c, 0x88, 0xf2, 0x50, 0xab, 0xcc, 0xa8, 0x34, 0x71, 0xbb, 0x76, 0x55, 0x1f, 0x89, 0x7b,
	0x38, 0x9c, 0x69, 0x92, 0x86, 0x4a, 0xae, 0x4f, 0xaf, 0x2b, 0xca, 0x4d, 0x85, 0x63, 0x3f, 0xe3,
	0xf8, 0x77, 0xdc, 0x2d, 0xf4, 0x4b, 0xd0, 0x9d, 0xca, 0x0d, 0x5e, 0x00, 0xac, 0xf3, 0x28, 0xca,
	0x5d, 0x36, 0xec, 0x7a, 0xce, 0xa4, 0x3f, 0xca, 0x82, 0x51, 0x75, 0xae, 0xb6, 0x17, 0xe7, 0x80,
	0xe5, 0x7c, 0x4e, 0xa6, 0x74, 0xb2, 0x11, 0x74, 0xf2, 0xc1, 0x61, 0xaf, 0x94, 0x2d, 0x48, 0xbf,
	0xa9, 0x90, 0xf0, 0x06, 0x06, 0xcd, 0x18, 0x78, 0x6c, 0xaf, 0xb4, 0x45, 0x3b, 0x69, 0x18, 0x10,
	0x1d, 0x9c, 0xc3, 0x41, 0x53, 0xb8, 0x30, 0x9a, 0xe4, 0xf2, 0x37, 0xc4, 0x7e, 0x1d, 0x51, 0x24,
	0x15, 0x1d, 0x8f, 0xa1, 0x0f, 0x67, 0x6d, 0xa0, 0xa9, 0x8a, 0x94, 0xa6, 0xb0, 0x28, 0x48, 0xc6,
	0x7f, 0x30, 0xe6, 0xb1, 0x4b, 0x86, 0x63, 0x18, 0x14, 0xfc, 0x59, 0xd5, 0x12, 0xf6, 0x0a, 0x95,
	0xfd, 0x0f, 0x6d, 0x36, 0xf0, 0x1a, 0x9c, 0x39, 0x99, 0xaa, 0x87, 0xa3, 0xba, 0xe4, 0xab, 0xd5,
	0xcd, 0x5b, 0xd3, 0xdd, 0x47, 0x47, 0x25, 0x86, 0x74, 0x22, 0xe3, 0x71, 0x16, 0x04, 0x5b, 0xf6,
	0xf7, 0x5d, 0x7d, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe5, 0xb9, 0xdd, 0x47, 0x9c, 0x02, 0x00, 0x00,
}
