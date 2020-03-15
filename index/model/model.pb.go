// Code generated by protoc-gen-go. DO NOT EDIT.
// source: index/model/model.proto

package model

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type Entry_Type int32

const (
	Entry_TREE   Entry_Type = 0
	Entry_OBJECT Entry_Type = 1
)

var Entry_Type_name = map[int32]string{
	0: "TREE",
	1: "OBJECT",
}

var Entry_Type_value = map[string]int32{
	"TREE":   0,
	"OBJECT": 1,
}

func (x Entry_Type) String() string {
	return proto.EnumName(Entry_Type_name, int32(x))
}

func (Entry_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{4, 0}
}

// Data model
type Repo struct {
	RepoId               string   `protobuf:"bytes,1,opt,name=repo_id,json=repoId,proto3" json:"repo_id,omitempty"`
	BucketName           string   `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	CreationDate         int64    `protobuf:"varint,3,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	DefaultBranch        string   `protobuf:"bytes,4,opt,name=default_branch,json=defaultBranch,proto3" json:"default_branch,omitempty"`
	PartialCommitRatio   float32  `protobuf:"fixed32,5,opt,name=partial_commit_ratio,json=partialCommitRatio,proto3" json:"partial_commit_ratio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repo) Reset()         { *m = Repo{} }
func (m *Repo) String() string { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()    {}
func (*Repo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{0}
}

func (m *Repo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repo.Unmarshal(m, b)
}
func (m *Repo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repo.Marshal(b, m, deterministic)
}
func (m *Repo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repo.Merge(m, src)
}
func (m *Repo) XXX_Size() int {
	return xxx_messageInfo_Repo.Size(m)
}
func (m *Repo) XXX_DiscardUnknown() {
	xxx_messageInfo_Repo.DiscardUnknown(m)
}

var xxx_messageInfo_Repo proto.InternalMessageInfo

func (m *Repo) GetRepoId() string {
	if m != nil {
		return m.RepoId
	}
	return ""
}

func (m *Repo) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *Repo) GetCreationDate() int64 {
	if m != nil {
		return m.CreationDate
	}
	return 0
}

func (m *Repo) GetDefaultBranch() string {
	if m != nil {
		return m.DefaultBranch
	}
	return ""
}

func (m *Repo) GetPartialCommitRatio() float32 {
	if m != nil {
		return m.PartialCommitRatio
	}
	return 0
}

type Block struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Block) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Object struct {
	Blocks               []*Block          `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Checksum             string            `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Size                 int64             `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{2}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Object) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *Object) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Object) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Root struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Root) Reset()         { *m = Root{} }
func (m *Root) String() string { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()    {}
func (*Root) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{3}
}

func (m *Root) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Root.Unmarshal(m, b)
}
func (m *Root) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Root.Marshal(b, m, deterministic)
}
func (m *Root) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Root.Merge(m, src)
}
func (m *Root) XXX_Size() int {
	return xxx_messageInfo_Root.Size(m)
}
func (m *Root) XXX_DiscardUnknown() {
	xxx_messageInfo_Root.DiscardUnknown(m)
}

var xxx_messageInfo_Root proto.InternalMessageInfo

func (m *Root) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Root) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Entry struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address              string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Type                 Entry_Type `protobuf:"varint,3,opt,name=type,proto3,enum=Entry_Type" json:"type,omitempty"`
	Timestamp            int64      `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Size                 int64      `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Checksum             string     `protobuf:"bytes,6,opt,name=checksum,proto3" json:"checksum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{4}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Entry) GetType() Entry_Type {
	if m != nil {
		return m.Type
	}
	return Entry_TREE
}

func (m *Entry) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Entry) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Entry) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

type WorkspaceEntry struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Entry                *Entry   `protobuf:"bytes,2,opt,name=entry,proto3" json:"entry,omitempty"`
	Tombstone            bool     `protobuf:"varint,3,opt,name=tombstone,proto3" json:"tombstone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkspaceEntry) Reset()         { *m = WorkspaceEntry{} }
func (m *WorkspaceEntry) String() string { return proto.CompactTextString(m) }
func (*WorkspaceEntry) ProtoMessage()    {}
func (*WorkspaceEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{5}
}

func (m *WorkspaceEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkspaceEntry.Unmarshal(m, b)
}
func (m *WorkspaceEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkspaceEntry.Marshal(b, m, deterministic)
}
func (m *WorkspaceEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkspaceEntry.Merge(m, src)
}
func (m *WorkspaceEntry) XXX_Size() int {
	return xxx_messageInfo_WorkspaceEntry.Size(m)
}
func (m *WorkspaceEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkspaceEntry.DiscardUnknown(m)
}

var xxx_messageInfo_WorkspaceEntry proto.InternalMessageInfo

func (m *WorkspaceEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WorkspaceEntry) GetEntry() *Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *WorkspaceEntry) GetTombstone() bool {
	if m != nil {
		return m.Tombstone
	}
	return false
}

type Commit struct {
	Address              string            `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Tree                 string            `protobuf:"bytes,2,opt,name=tree,proto3" json:"tree,omitempty"`
	Parents              []string          `protobuf:"bytes,3,rep,name=parents,proto3" json:"parents,omitempty"`
	Committer            string            `protobuf:"bytes,4,opt,name=committer,proto3" json:"committer,omitempty"`
	Message              string            `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            int64             `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{6}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Commit) GetTree() string {
	if m != nil {
		return m.Tree
	}
	return ""
}

func (m *Commit) GetParents() []string {
	if m != nil {
		return m.Parents
	}
	return nil
}

func (m *Commit) GetCommitter() string {
	if m != nil {
		return m.Committer
	}
	return ""
}

func (m *Commit) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Commit) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Commit) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Branch struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Commit               string   `protobuf:"bytes,2,opt,name=commit,proto3" json:"commit,omitempty"`
	CommitRoot           string   `protobuf:"bytes,3,opt,name=commit_root,json=commitRoot,proto3" json:"commit_root,omitempty"`
	WorkspaceRoot        string   `protobuf:"bytes,4,opt,name=workspace_root,json=workspaceRoot,proto3" json:"workspace_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Branch) Reset()         { *m = Branch{} }
func (m *Branch) String() string { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()    {}
func (*Branch) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{7}
}

func (m *Branch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Branch.Unmarshal(m, b)
}
func (m *Branch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Branch.Marshal(b, m, deterministic)
}
func (m *Branch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Branch.Merge(m, src)
}
func (m *Branch) XXX_Size() int {
	return xxx_messageInfo_Branch.Size(m)
}
func (m *Branch) XXX_DiscardUnknown() {
	xxx_messageInfo_Branch.DiscardUnknown(m)
}

var xxx_messageInfo_Branch proto.InternalMessageInfo

func (m *Branch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Branch) GetCommit() string {
	if m != nil {
		return m.Commit
	}
	return ""
}

func (m *Branch) GetCommitRoot() string {
	if m != nil {
		return m.CommitRoot
	}
	return ""
}

func (m *Branch) GetWorkspaceRoot() string {
	if m != nil {
		return m.WorkspaceRoot
	}
	return ""
}

type MultipartUpload struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultipartUpload) Reset()         { *m = MultipartUpload{} }
func (m *MultipartUpload) String() string { return proto.CompactTextString(m) }
func (*MultipartUpload) ProtoMessage()    {}
func (*MultipartUpload) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{8}
}

func (m *MultipartUpload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipartUpload.Unmarshal(m, b)
}
func (m *MultipartUpload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipartUpload.Marshal(b, m, deterministic)
}
func (m *MultipartUpload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipartUpload.Merge(m, src)
}
func (m *MultipartUpload) XXX_Size() int {
	return xxx_messageInfo_MultipartUpload.Size(m)
}
func (m *MultipartUpload) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipartUpload.DiscardUnknown(m)
}

var xxx_messageInfo_MultipartUpload proto.InternalMessageInfo

func (m *MultipartUpload) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MultipartUpload) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MultipartUpload) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type MultipartUploadPart struct {
	Blocks               []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Checksum             string   `protobuf:"bytes,2,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Timestamp            int64    `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Size                 int64    `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultipartUploadPart) Reset()         { *m = MultipartUploadPart{} }
func (m *MultipartUploadPart) String() string { return proto.CompactTextString(m) }
func (*MultipartUploadPart) ProtoMessage()    {}
func (*MultipartUploadPart) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{9}
}

func (m *MultipartUploadPart) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipartUploadPart.Unmarshal(m, b)
}
func (m *MultipartUploadPart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipartUploadPart.Marshal(b, m, deterministic)
}
func (m *MultipartUploadPart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipartUploadPart.Merge(m, src)
}
func (m *MultipartUploadPart) XXX_Size() int {
	return xxx_messageInfo_MultipartUploadPart.Size(m)
}
func (m *MultipartUploadPart) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipartUploadPart.DiscardUnknown(m)
}

var xxx_messageInfo_MultipartUploadPart proto.InternalMessageInfo

func (m *MultipartUploadPart) GetBlocks() []*Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *MultipartUploadPart) GetChecksum() string {
	if m != nil {
		return m.Checksum
	}
	return ""
}

func (m *MultipartUploadPart) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MultipartUploadPart) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type MultipartUploadPartRequest struct {
	PartNumber           int32    `protobuf:"varint,1,opt,name=partNumber,proto3" json:"partNumber,omitempty"`
	Etag                 string   `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultipartUploadPartRequest) Reset()         { *m = MultipartUploadPartRequest{} }
func (m *MultipartUploadPartRequest) String() string { return proto.CompactTextString(m) }
func (*MultipartUploadPartRequest) ProtoMessage()    {}
func (*MultipartUploadPartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d1d9db5dfb08471, []int{10}
}

func (m *MultipartUploadPartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultipartUploadPartRequest.Unmarshal(m, b)
}
func (m *MultipartUploadPartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultipartUploadPartRequest.Marshal(b, m, deterministic)
}
func (m *MultipartUploadPartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultipartUploadPartRequest.Merge(m, src)
}
func (m *MultipartUploadPartRequest) XXX_Size() int {
	return xxx_messageInfo_MultipartUploadPartRequest.Size(m)
}
func (m *MultipartUploadPartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MultipartUploadPartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MultipartUploadPartRequest proto.InternalMessageInfo

func (m *MultipartUploadPartRequest) GetPartNumber() int32 {
	if m != nil {
		return m.PartNumber
	}
	return 0
}

func (m *MultipartUploadPartRequest) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

func init() {
	proto.RegisterEnum("Entry_Type", Entry_Type_name, Entry_Type_value)
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*Object)(nil), "Object")
	proto.RegisterMapType((map[string]string)(nil), "Object.MetadataEntry")
	proto.RegisterType((*Root)(nil), "Root")
	proto.RegisterType((*Entry)(nil), "Entry")
	proto.RegisterType((*WorkspaceEntry)(nil), "WorkspaceEntry")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterMapType((map[string]string)(nil), "Commit.MetadataEntry")
	proto.RegisterType((*Branch)(nil), "Ref")
	proto.RegisterType((*MultipartUpload)(nil), "MultipartUpload")
	proto.RegisterType((*MultipartUploadPart)(nil), "MultipartUploadPart")
	proto.RegisterType((*MultipartUploadPartRequest)(nil), "MultipartUploadPartRequest")
}

func init() {
	proto.RegisterFile("index/model/model.proto", fileDescriptor_9d1d9db5dfb08471)
}

var fileDescriptor_9d1d9db5dfb08471 = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x3f, 0x71, 0x93, 0x29, 0x0d, 0xd5, 0x52, 0x68, 0x54, 0x55, 0xb4, 0x32, 0xaa, 0x94,
	0x53, 0x80, 0x22, 0xa4, 0x0a, 0x6e, 0x2d, 0x39, 0x80, 0xd4, 0x1f, 0x2d, 0x45, 0x1c, 0xc3, 0xda,
	0x1e, 0x5a, 0x13, 0xdb, 0x6b, 0xd6, 0x6b, 0x20, 0x88, 0x1b, 0xef, 0xc0, 0xe3, 0x70, 0xe0, 0xc4,
	0x63, 0xa1, 0xdd, 0xb5, 0x9d, 0xb8, 0x4a, 0x7b, 0xe9, 0xc5, 0x9a, 0xf9, 0x66, 0x77, 0x76, 0xe6,
	0x9b, 0xcf, 0x03, 0x9b, 0x71, 0x16, 0xe1, 0xf7, 0x27, 0x29, 0x8f, 0x30, 0x31, 0xdf, 0x51, 0x2e,
	0xb8, 0xe4, 0xfe, 0x1f, 0x0b, 0x5c, 0x8a, 0x39, 0x27, 0x9b, 0xb0, 0x22, 0x30, 0xe7, 0x93, 0x38,
	0x1a, 0x58, 0xbb, 0xd6, 0xb0, 0x47, 0x3d, 0xe5, 0xbe, 0x89, 0xc8, 0x0e, 0xac, 0x06, 0x65, 0x38,
	0x45, 0x39, 0xc9, 0x58, 0x8a, 0x03, 0x5b, 0x07, 0xc1, 0x40, 0x27, 0x2c, 0x45, 0xf2, 0x18, 0xd6,
	0x42, 0x81, 0x4c, 0xc6, 0x3c, 0x9b, 0x44, 0x4c, 0xe2, 0xc0, 0xd9, 0xb5, 0x86, 0x0e, 0xbd, 0x5b,
	0x83, 0xaf, 0x99, 0x44, 0xb2, 0x07, 0xfd, 0x08, 0x3f, 0xb1, 0x32, 0x91, 0x93, 0x40, 0xb0, 0x2c,
	0xbc, 0x1c, 0xb8, 0x3a, 0xd1, 0x5a, 0x85, 0x1e, 0x6a, 0x90, 0x3c, 0x85, 0x8d, 0x9c, 0x09, 0x19,
	0xb3, 0x64, 0x12, 0xf2, 0x34, 0x8d, 0xe5, 0x44, 0xa8, 0x1c, 0x83, 0xce, 0xae, 0x35, 0xb4, 0x29,
	0xa9, 0x62, 0x47, 0x3a, 0x44, 0x55, 0xc4, 0x7f, 0x01, 0x9d, 0xc3, 0x84, 0x87, 0x53, 0x32, 0x80,
	0x15, 0x16, 0x45, 0x02, 0x8b, 0xa2, 0x6a, 0xa0, 0x76, 0x09, 0x01, 0xb7, 0x88, 0x7f, 0x98, 0xd2,
	0x1d, 0xaa, 0x6d, 0xff, 0x9f, 0x05, 0xde, 0x69, 0xf0, 0x19, 0x43, 0x49, 0x1e, 0x81, 0x17, 0xa8,
	0x0c, 0xea, 0x9e, 0x33, 0x5c, 0xdd, 0xf7, 0x46, 0x3a, 0x21, 0xad, 0x50, 0xb2, 0x05, 0xdd, 0xf0,
	0x12, 0xc3, 0x69, 0x51, 0xa6, 0x55, 0xf7, 0x8d, 0x4f, 0x9e, 0x41, 0x37, 0x45, 0xc9, 0x22, 0x26,
	0xd9, 0xc0, 0xd1, 0xb7, 0x1f, 0x8c, 0x4c, 0xda, 0xd1, 0x71, 0x85, 0x8f, 0x33, 0x29, 0x66, 0xb4,
	0x39, 0xd6, 0x54, 0xe3, 0xce, 0xab, 0xd9, 0x7a, 0x05, 0x6b, 0xad, 0xe3, 0x64, 0x1d, 0x9c, 0x29,
	0xce, 0xaa, 0x46, 0x94, 0x49, 0x36, 0xa0, 0xf3, 0x95, 0x25, 0x65, 0x3d, 0x00, 0xe3, 0xbc, 0xb4,
	0x0f, 0x2c, 0xff, 0x00, 0x5c, 0xca, 0xb9, 0x24, 0xdb, 0xd0, 0x93, 0x71, 0x8a, 0x85, 0x64, 0x69,
	0xae, 0x6f, 0x3a, 0x74, 0x0e, 0x2c, 0x25, 0xe1, 0xaf, 0x05, 0x1d, 0xf3, 0x1e, 0x01, 0x57, 0x4f,
	0xd7, 0x3c, 0xa8, 0xed, 0x45, 0x42, 0xed, 0x36, 0xa1, 0x3b, 0xe0, 0xca, 0x59, 0x6e, 0x06, 0xdd,
	0xdf, 0x5f, 0x1d, 0xe9, 0x1c, 0xa3, 0xf3, 0x59, 0x8e, 0x54, 0x07, 0xda, 0xa5, 0xb8, 0xd7, 0x95,
	0xd2, 0x99, 0x97, 0xd2, 0x22, 0xd9, 0x6b, 0x93, 0xec, 0x6f, 0x83, 0xab, 0x72, 0x93, 0x2e, 0xb8,
	0xe7, 0x74, 0x3c, 0x5e, 0xbf, 0x43, 0x00, 0xbc, 0xd3, 0xc3, 0xb7, 0xe3, 0xa3, 0xf3, 0x75, 0xcb,
	0xff, 0x08, 0xfd, 0x0f, 0x5c, 0x4c, 0x8b, 0x9c, 0x85, 0xd8, 0x34, 0x93, 0x33, 0x79, 0x59, 0x37,
	0xa3, 0x6c, 0xb2, 0x0d, 0x1d, 0x54, 0x41, 0xdd, 0x8a, 0x9a, 0xb1, 0x19, 0x8b, 0x01, 0x75, 0xbd,
	0x3c, 0x0d, 0x0a, 0xc9, 0x33, 0xd3, 0x55, 0x97, 0xce, 0x01, 0xff, 0xb7, 0x0d, 0x9e, 0x91, 0xdc,
	0xcd, 0x22, 0x93, 0x02, 0xeb, 0xf1, 0x68, 0x5b, 0x9d, 0xce, 0x99, 0xc0, 0x4c, 0x16, 0x5a, 0x1c,
	0x3d, 0x5a, 0xbb, 0xea, 0x41, 0xa3, 0x6f, 0x89, 0xa2, 0xfa, 0x13, 0xe6, 0x80, 0xba, 0x97, 0x62,
	0x51, 0xb0, 0x0b, 0xc3, 0x51, 0x8f, 0xd6, 0x6e, 0x9b, 0x58, 0xef, 0x2a, 0xb1, 0x8b, 0x6a, 0x5c,
	0xa9, 0xd4, 0x68, 0x0a, 0xbf, 0x4e, 0x8d, 0xb7, 0x53, 0xde, 0x4f, 0xf0, 0xaa, 0xff, 0x76, 0x99,
	0x7e, 0x1e, 0x82, 0x67, 0x5a, 0xaa, 0x2e, 0x56, 0x9e, 0x5a, 0x28, 0xf5, 0xbf, 0xcd, 0xb9, 0xd4,
	0x74, 0xf7, 0x28, 0x18, 0x48, 0x0b, 0x79, 0x0f, 0xfa, 0xdf, 0xea, 0x89, 0x9a, 0x33, 0xd5, 0xae,
	0x68, 0x50, 0x75, 0xcc, 0x7f, 0x07, 0xf7, 0x8e, 0xcb, 0x44, 0xc6, 0x6a, 0x29, 0xbc, 0xcf, 0x13,
	0xce, 0x22, 0xd2, 0x07, 0xbb, 0xd9, 0x5f, 0x76, 0x1c, 0x35, 0x4a, 0xb0, 0x5b, 0x4a, 0x58, 0xa0,
	0xd0, 0xb9, 0x42, 0xa1, 0xff, 0xcb, 0x82, 0xfb, 0x57, 0xb2, 0x9e, 0x31, 0x71, 0xbb, 0x25, 0x71,
	0xe3, 0x8b, 0xcb, 0xf6, 0x81, 0x7f, 0x06, 0x5b, 0x4b, 0x8a, 0xa0, 0xf8, 0xa5, 0xc4, 0x42, 0xd5,
	0x02, 0x2a, 0x70, 0x52, 0xa6, 0x01, 0x0a, 0xdd, 0x6d, 0x87, 0x2e, 0x20, 0x2a, 0x23, 0x4a, 0x76,
	0x51, 0x77, 0xad, 0xec, 0xc0, 0xd3, 0xeb, 0xfe, 0xf9, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65,
	0x07, 0x3c, 0xdd, 0x09, 0x06, 0x00, 0x00,
}
