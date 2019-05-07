// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/osconfig/v1alpha1/assignments.proto

package osconfig

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

// An Assignment maps configurations to instance guest environments.
type Assignment struct {
	// Identifying name for this Assignment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Assignment. Length of the description is limited
	// to 1024 characters.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. Time this Assignment was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Last time this Assignment was updated.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Represents cloud resource labels.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// List of OsConfigs to configure on the instances. These are relative
	// resource names of OsConfigs. For example 'organizations/1234/osConfigs/foo'
	// or 'projects/12345/osConfigs/foo'. If an OsConfig referenced here is
	// deleted it will be ignored when instances lookup their configs.
	OsConfigs []string `protobuf:"bytes,6,rep,name=os_configs,json=osConfigs,proto3" json:"os_configs,omitempty"`
	// Optional. A [CEL](https://github.com/google/cel-spec) expression used to
	// filter instances when determining which configs apply. If omitted, the
	// OsConfigs specified in this assignment will apply to all instances under
	// this resource.
	Expression           string   `protobuf:"bytes,7,opt,name=expression,proto3" json:"expression,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e4b5881c4fd4af0, []int{0}
}

func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (m *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(m, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Assignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Assignment) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Assignment) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Assignment) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Assignment) GetOsConfigs() []string {
	if m != nil {
		return m.OsConfigs
	}
	return nil
}

func (m *Assignment) GetExpression() string {
	if m != nil {
		return m.Expression
	}
	return ""
}

// A request message for creating an Assignment.
type CreateAssignmentRequest struct {
	// The resource name of the parent.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The Assignment to create.
	Assignment           *Assignment `protobuf:"bytes,2,opt,name=assignment,proto3" json:"assignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateAssignmentRequest) Reset()         { *m = CreateAssignmentRequest{} }
func (m *CreateAssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAssignmentRequest) ProtoMessage()    {}
func (*CreateAssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e4b5881c4fd4af0, []int{1}
}

func (m *CreateAssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAssignmentRequest.Unmarshal(m, b)
}
func (m *CreateAssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAssignmentRequest.Marshal(b, m, deterministic)
}
func (m *CreateAssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAssignmentRequest.Merge(m, src)
}
func (m *CreateAssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAssignmentRequest.Size(m)
}
func (m *CreateAssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAssignmentRequest proto.InternalMessageInfo

func (m *CreateAssignmentRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateAssignmentRequest) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

// A request message for retrieving a Assignment.
type GetAssignmentRequest struct {
	// The resource name of the Assignment.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAssignmentRequest) Reset()         { *m = GetAssignmentRequest{} }
func (m *GetAssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*GetAssignmentRequest) ProtoMessage()    {}
func (*GetAssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e4b5881c4fd4af0, []int{2}
}

func (m *GetAssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAssignmentRequest.Unmarshal(m, b)
}
func (m *GetAssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAssignmentRequest.Marshal(b, m, deterministic)
}
func (m *GetAssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAssignmentRequest.Merge(m, src)
}
func (m *GetAssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_GetAssignmentRequest.Size(m)
}
func (m *GetAssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAssignmentRequest proto.InternalMessageInfo

func (m *GetAssignmentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// A request message for listing Assignments.
type ListAssignmentsRequest struct {
	// The resource name of the parent.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of Assignments to return.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A pagination token returned from a previous call to ListAssignments
	// that indicates where this listing should continue from.
	// This field is optional.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssignmentsRequest) Reset()         { *m = ListAssignmentsRequest{} }
func (m *ListAssignmentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListAssignmentsRequest) ProtoMessage()    {}
func (*ListAssignmentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e4b5881c4fd4af0, []int{3}
}

func (m *ListAssignmentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssignmentsRequest.Unmarshal(m, b)
}
func (m *ListAssignmentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssignmentsRequest.Marshal(b, m, deterministic)
}
func (m *ListAssignmentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssignmentsRequest.Merge(m, src)
}
func (m *ListAssignmentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListAssignmentsRequest.Size(m)
}
func (m *ListAssignmentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssignmentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssignmentsRequest proto.InternalMessageInfo

func (m *ListAssignmentsRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListAssignmentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListAssignmentsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// A response message for listing Assignments.
type ListAssignmentsResponse struct {
	// The list of Assignments.
	Assignments []*Assignment `protobuf:"bytes,1,rep,name=assignments,proto3" json:"assignments,omitempty"`
	// A pagination token that can be used to get the next page
	// of Assignments.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAssignmentsResponse) Reset()         { *m = ListAssignmentsResponse{} }
func (m *ListAssignmentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListAssignmentsResponse) ProtoMessage()    {}
func (*ListAssignmentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e4b5881c4fd4af0, []int{4}
}

func (m *ListAssignmentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAssignmentsResponse.Unmarshal(m, b)
}
func (m *ListAssignmentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAssignmentsResponse.Marshal(b, m, deterministic)
}
func (m *ListAssignmentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAssignmentsResponse.Merge(m, src)
}
func (m *ListAssignmentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListAssignmentsResponse.Size(m)
}
func (m *ListAssignmentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAssignmentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAssignmentsResponse proto.InternalMessageInfo

func (m *ListAssignmentsResponse) GetAssignments() []*Assignment {
	if m != nil {
		return m.Assignments
	}
	return nil
}

func (m *ListAssignmentsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// A request message for updating a Assignment.
type UpdateAssignmentRequest struct {
	// The resource name of the Assignment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The updated Assignment.
	Assignment *Assignment `protobuf:"bytes,2,opt,name=assignment,proto3" json:"assignment,omitempty"`
	// Field mask that controls which fields of the Assignment should be
	// updated.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateAssignmentRequest) Reset()         { *m = UpdateAssignmentRequest{} }
func (m *UpdateAssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAssignmentRequest) ProtoMessage()    {}
func (*UpdateAssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e4b5881c4fd4af0, []int{5}
}

func (m *UpdateAssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAssignmentRequest.Unmarshal(m, b)
}
func (m *UpdateAssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAssignmentRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAssignmentRequest.Merge(m, src)
}
func (m *UpdateAssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAssignmentRequest.Size(m)
}
func (m *UpdateAssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAssignmentRequest proto.InternalMessageInfo

func (m *UpdateAssignmentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateAssignmentRequest) GetAssignment() *Assignment {
	if m != nil {
		return m.Assignment
	}
	return nil
}

func (m *UpdateAssignmentRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

// A request message for deleting a Assignment.
type DeleteAssignmentRequest struct {
	// The resource name of the Assignment.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAssignmentRequest) Reset()         { *m = DeleteAssignmentRequest{} }
func (m *DeleteAssignmentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAssignmentRequest) ProtoMessage()    {}
func (*DeleteAssignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e4b5881c4fd4af0, []int{6}
}

func (m *DeleteAssignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAssignmentRequest.Unmarshal(m, b)
}
func (m *DeleteAssignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAssignmentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAssignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAssignmentRequest.Merge(m, src)
}
func (m *DeleteAssignmentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAssignmentRequest.Size(m)
}
func (m *DeleteAssignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAssignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAssignmentRequest proto.InternalMessageInfo

func (m *DeleteAssignmentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Assignment)(nil), "google.cloud.osconfig.v1alpha1.Assignment")
	proto.RegisterMapType((map[string]string)(nil), "google.cloud.osconfig.v1alpha1.Assignment.LabelsEntry")
	proto.RegisterType((*CreateAssignmentRequest)(nil), "google.cloud.osconfig.v1alpha1.CreateAssignmentRequest")
	proto.RegisterType((*GetAssignmentRequest)(nil), "google.cloud.osconfig.v1alpha1.GetAssignmentRequest")
	proto.RegisterType((*ListAssignmentsRequest)(nil), "google.cloud.osconfig.v1alpha1.ListAssignmentsRequest")
	proto.RegisterType((*ListAssignmentsResponse)(nil), "google.cloud.osconfig.v1alpha1.ListAssignmentsResponse")
	proto.RegisterType((*UpdateAssignmentRequest)(nil), "google.cloud.osconfig.v1alpha1.UpdateAssignmentRequest")
	proto.RegisterType((*DeleteAssignmentRequest)(nil), "google.cloud.osconfig.v1alpha1.DeleteAssignmentRequest")
}

func init() {
	proto.RegisterFile("google/cloud/osconfig/v1alpha1/assignments.proto", fileDescriptor_7e4b5881c4fd4af0)
}

var fileDescriptor_7e4b5881c4fd4af0 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x51, 0x6b, 0xd4, 0x40,
	0x10, 0x26, 0xbd, 0xf6, 0xf4, 0x26, 0x88, 0xb2, 0x94, 0x5e, 0x3c, 0xb1, 0x86, 0x3c, 0xc8, 0x51,
	0x30, 0xb1, 0x15, 0x44, 0xed, 0x93, 0xad, 0x56, 0x90, 0x2a, 0x12, 0xeb, 0x8b, 0x2f, 0xc7, 0x5e,
	0x6e, 0x2e, 0x86, 0x4b, 0x76, 0x63, 0x66, 0x53, 0x7a, 0x05, 0x5f, 0x7d, 0xf5, 0xdf, 0xf8, 0xfb,
	0x64, 0x37, 0xc9, 0x5d, 0xb0, 0xd6, 0x3b, 0xc1, 0xb7, 0xdd, 0xd9, 0xef, 0x9b, 0xf9, 0xf6, 0xdb,
	0x99, 0x85, 0xc7, 0xb1, 0x94, 0x71, 0x8a, 0x41, 0x94, 0xca, 0x72, 0x12, 0x48, 0x8a, 0xa4, 0x98,
	0x26, 0x71, 0x70, 0xbe, 0xcf, 0xd3, 0xfc, 0x0b, 0xdf, 0x0f, 0x38, 0x51, 0x12, 0x8b, 0x0c, 0x85,
	0x22, 0x3f, 0x2f, 0xa4, 0x92, 0x6c, 0xb7, 0x62, 0xf8, 0x86, 0xe1, 0x37, 0x0c, 0xbf, 0x61, 0x0c,
	0xee, 0xd6, 0x19, 0x79, 0x9e, 0x04, 0x05, 0x92, 0x2c, 0x8b, 0x08, 0x2b, 0xea, 0xc0, 0xad, 0x8f,
	0xcc, 0x6e, 0x5c, 0x4e, 0x83, 0x69, 0x82, 0xe9, 0x64, 0x94, 0x71, 0x9a, 0xd5, 0x88, 0x07, 0xbf,
	0x23, 0x54, 0x92, 0x21, 0x29, 0x9e, 0xe5, 0x15, 0xc0, 0xfb, 0xde, 0x01, 0x78, 0xb9, 0xd0, 0xc4,
	0x18, 0x6c, 0x0a, 0x9e, 0xa1, 0x63, 0xb9, 0xd6, 0xb0, 0x17, 0x9a, 0x35, 0x73, 0xc1, 0x9e, 0x20,
	0x45, 0x45, 0x92, 0xab, 0x44, 0x0a, 0x67, 0xc3, 0x1c, 0xb5, 0x43, 0xec, 0x10, 0xec, 0xa8, 0x40,
	0xae, 0x70, 0xa4, 0xd3, 0x3b, 0x1d, 0xd7, 0x1a, 0xda, 0x07, 0x03, 0xbf, 0xbe, 0x58, 0x53, 0xdb,
	0x3f, 0x6b, 0x6a, 0x87, 0x50, 0xc1, 0x75, 0x40, 0x93, 0xcb, 0x7c, 0xb2, 0x20, 0x6f, 0xae, 0x26,
	0x57, 0x70, 0x43, 0x7e, 0x0f, 0xdd, 0x94, 0x8f, 0x31, 0x25, 0x67, 0xcb, 0xed, 0x0c, 0xed, 0x83,
	0xa7, 0xfe, 0xdf, 0xdd, 0xf4, 0x97, 0x77, 0xf5, 0x4f, 0x0d, 0xf1, 0xb5, 0x50, 0xc5, 0x3c, 0xac,
	0xb3, 0xb0, 0xfb, 0x00, 0x92, 0x46, 0x15, 0x89, 0x9c, 0xae, 0xdb, 0x19, 0xf6, 0xc2, 0x9e, 0xa4,
	0xe3, 0x2a, 0xc0, 0x76, 0x01, 0xf0, 0x22, 0x2f, 0x90, 0x48, 0x3b, 0x71, 0xc3, 0x38, 0xd1, 0x8a,
	0x0c, 0x9e, 0x83, 0xdd, 0xca, 0xca, 0xee, 0x40, 0x67, 0x86, 0xf3, 0xda, 0x4c, 0xbd, 0x64, 0xdb,
	0xb0, 0x75, 0xce, 0xd3, 0x12, 0x6b, 0x17, 0xab, 0xcd, 0x8b, 0x8d, 0x67, 0x96, 0xf7, 0x0d, 0xfa,
	0xc7, 0xc6, 0x94, 0xa5, 0xc2, 0x10, 0xbf, 0x96, 0x48, 0x8a, 0xed, 0x40, 0x37, 0xe7, 0x05, 0x0a,
	0x55, 0x67, 0xaa, 0x77, 0xec, 0x2d, 0xc0, 0xb2, 0x9d, 0x4c, 0x46, 0xfb, 0x60, 0x6f, 0x7d, 0x03,
	0xc2, 0x16, 0xdb, 0xdb, 0x83, 0xed, 0x37, 0xa8, 0xae, 0xd6, 0xfe, 0x43, 0x43, 0x78, 0x29, 0xec,
	0x9c, 0x26, 0xd4, 0x02, 0xd3, 0x2a, 0xa5, 0xf7, 0xa0, 0x97, 0xf3, 0x18, 0x47, 0x94, 0x5c, 0x56,
	0x57, 0xdf, 0x0a, 0x6f, 0xea, 0xc0, 0xc7, 0xe4, 0x12, 0xb5, 0xe7, 0xe6, 0x50, 0xc9, 0x19, 0x0a,
	0xd3, 0x3c, 0xbd, 0xd0, 0xc0, 0xcf, 0x74, 0xc0, 0xfb, 0x61, 0x41, 0xff, 0x4a, 0x39, 0xca, 0xa5,
	0x20, 0x64, 0xa7, 0x60, 0xb7, 0x06, 0xca, 0xb1, 0x4c, 0x0f, 0xfc, 0x8b, 0x05, 0x6d, 0x3a, 0x7b,
	0x08, 0xb7, 0x05, 0x5e, 0xa8, 0x51, 0x4b, 0x4d, 0xf5, 0x4c, 0xb7, 0x74, 0xf8, 0xc3, 0x42, 0xd1,
	0x4f, 0x0b, 0xfa, 0x9f, 0x4c, 0x0f, 0xae, 0xe5, 0xd7, 0xff, 0x7c, 0xa7, 0xd6, 0xb4, 0xe8, 0x29,
	0xbf, 0x76, 0xd4, 0x4e, 0xf4, 0x47, 0xf0, 0x8e, 0xd3, 0xac, 0x99, 0x16, 0xbd, 0xf6, 0x1e, 0x41,
	0xff, 0x15, 0xa6, 0xb8, 0xa6, 0xee, 0xa3, 0x39, 0x78, 0x91, 0xcc, 0x56, 0x08, 0x3d, 0xb2, 0x5b,
	0x0f, 0xf3, 0xf9, 0xa4, 0x06, 0xc7, 0x32, 0xe5, 0x22, 0xf6, 0x65, 0x11, 0x07, 0x31, 0x0a, 0x23,
	0x2b, 0xa8, 0x8e, 0x78, 0x9e, 0xd0, 0x75, 0xbf, 0xe3, 0x61, 0x13, 0x19, 0x77, 0x0d, 0xe5, 0xc9,
	0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x56, 0x54, 0xd2, 0x4f, 0x05, 0x00, 0x00,
}