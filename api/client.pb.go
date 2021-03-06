// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Whether these limits are cachable, and if they are, by what facet of the token.
type ResourceLimits_Cachability int32

const (
	// Do not cache these limits.
	ResourceLimits_NONE ResourceLimits_Cachability = 0
	// These limits can be cached and applied to this and any of the user's
	// other repls.
	ResourceLimits_USER ResourceLimits_Cachability = 1
	// These limits can be cached and applied only to this repl.
	ResourceLimits_REPL ResourceLimits_Cachability = 2
)

var ResourceLimits_Cachability_name = map[int32]string{
	0: "NONE",
	1: "USER",
	2: "REPL",
}

var ResourceLimits_Cachability_value = map[string]int32{
	"NONE": 0,
	"USER": 1,
	"REPL": 2,
}

func (x ResourceLimits_Cachability) String() string {
	return proto.EnumName(ResourceLimits_Cachability_name, int32(x))
}

func (ResourceLimits_Cachability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1, 0}
}

// Whether to persist filesystem, metadata, or both.
type ReplToken_Persistence int32

const (
	// This is the usual mode of operation: both filesystem and metadata will be
	// persisted.
	ReplToken_PERSISTENT ReplToken_Persistence = 0
	// The ephemeral flag indicates the repl being connected to will have a time
	// restriction on stored metadata.  This has the consequence that repl will
	// be unable to wakeup or serve static traffic once the metadata has timed
	// out. This option does NOT affect filesystem and other data persistence.
	//
	// For context, this value is used on the client when repls are created for:
	// - replrun
	// - guests
	// - anon users
	// - temp vnc repls
	// - users with non-verified emails
	ReplToken_EPHEMERAL ReplToken_Persistence = 1
	// This indicates that the repl being connected does not have the ability to
	// persist files or be woken up after the lifetime of this repl expires.
	//
	// For context, this value is used on the client when repls are created for:
	// - replrun
	// - guests
	// - language pages
	ReplToken_NONE ReplToken_Persistence = 2
)

var ReplToken_Persistence_name = map[int32]string{
	0: "PERSISTENT",
	1: "EPHEMERAL",
	2: "NONE",
}

var ReplToken_Persistence_value = map[string]int32{
	"PERSISTENT": 0,
	"EPHEMERAL":  1,
	"NONE":       2,
}

func (x ReplToken_Persistence) String() string {
	return proto.EnumName(ReplToken_Persistence_name, int32(x))
}

func (ReplToken_Persistence) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3, 0}
}

// allows the client to choose a wire format.
type ReplToken_WireFormat int32

const (
	// The default wire format: Protobuf-over-WebSocket.
	ReplToken_PROTOBUF ReplToken_WireFormat = 0
	// Legacy protocol.
	ReplToken_JSON ReplToken_WireFormat = 1 // Deprecated: Do not use.
)

var ReplToken_WireFormat_name = map[int32]string{
	0: "PROTOBUF",
	1: "JSON",
}

var ReplToken_WireFormat_value = map[string]int32{
	"PROTOBUF": 0,
	"JSON":     1,
}

func (x ReplToken_WireFormat) String() string {
	return proto.EnumName(ReplToken_WireFormat_name, int32(x))
}

func (ReplToken_WireFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3, 1}
}

// This message constitutes the repl metadata and define the repl we're
// connecting to. All fields are required unless otherwise stated.
type Repl struct {
	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Language string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Bucket   string `protobuf:"bytes,3,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Slug     string `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	User     string `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	// (Optional) The replID of a repl to be used as the source filesystem. All
	// writes will still go to the actual repl. This is intended to be a
	// replacement for guest repls, giving us cheap COW semantics so all
	// connections can have a real repl.
	//
	// One exception:
	//
	// It's important to note that data is not implicitly copied from src to
	// dest. Only what is explicitly written when talking to pid1 (either
	// gcsfiles or snapshots) will persist. This makes it slightly different
	// than just forking.
	//
	// It's unclear what the behaviour should be if:
	// - the dest and src repl both exist
	// - the dest and src are the same
	// - we have an src but no dest
	// consider these unsupported/undefined for now.
	SourceRepl           string   `protobuf:"bytes,6,opt,name=sourceRepl,proto3" json:"sourceRepl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repl) Reset()         { *m = Repl{} }
func (m *Repl) String() string { return proto.CompactTextString(m) }
func (*Repl) ProtoMessage()    {}
func (*Repl) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

func (m *Repl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repl.Unmarshal(m, b)
}
func (m *Repl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repl.Marshal(b, m, deterministic)
}
func (m *Repl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repl.Merge(m, src)
}
func (m *Repl) XXX_Size() int {
	return xxx_messageInfo_Repl.Size(m)
}
func (m *Repl) XXX_DiscardUnknown() {
	xxx_messageInfo_Repl.DiscardUnknown(m)
}

var xxx_messageInfo_Repl proto.InternalMessageInfo

func (m *Repl) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Repl) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Repl) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Repl) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Repl) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Repl) GetSourceRepl() string {
	if m != nil {
		return m.SourceRepl
	}
	return ""
}

// The resource limits that should be applied to the Repl's container.
type ResourceLimits struct {
	// Whether the repl has network access.
	Net bool `protobuf:"varint,1,opt,name=net,proto3" json:"net,omitempty"`
	// The amount of RAM in bytes that this repl will have.
	Memory int64 `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	// The number of cores that the container will be allowed to have.
	Threads float64 `protobuf:"fixed64,3,opt,name=threads,proto3" json:"threads,omitempty"`
	// The Docker container weight factor for the scheduler. Similar to the
	// `--cpu-shares` commandline flag.
	Shares float64 `protobuf:"fixed64,4,opt,name=shares,proto3" json:"shares,omitempty"`
	// The size of the disk in bytes.
	Disk                 int64                      `protobuf:"varint,5,opt,name=disk,proto3" json:"disk,omitempty"`
	Cache                ResourceLimits_Cachability `protobuf:"varint,6,opt,name=cache,proto3,enum=api.ResourceLimits_Cachability" json:"cache,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ResourceLimits) Reset()         { *m = ResourceLimits{} }
func (m *ResourceLimits) String() string { return proto.CompactTextString(m) }
func (*ResourceLimits) ProtoMessage()    {}
func (*ResourceLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1}
}

func (m *ResourceLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceLimits.Unmarshal(m, b)
}
func (m *ResourceLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceLimits.Marshal(b, m, deterministic)
}
func (m *ResourceLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceLimits.Merge(m, src)
}
func (m *ResourceLimits) XXX_Size() int {
	return xxx_messageInfo_ResourceLimits.Size(m)
}
func (m *ResourceLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceLimits.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceLimits proto.InternalMessageInfo

func (m *ResourceLimits) GetNet() bool {
	if m != nil {
		return m.Net
	}
	return false
}

func (m *ResourceLimits) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ResourceLimits) GetThreads() float64 {
	if m != nil {
		return m.Threads
	}
	return 0
}

func (m *ResourceLimits) GetShares() float64 {
	if m != nil {
		return m.Shares
	}
	return 0
}

func (m *ResourceLimits) GetDisk() int64 {
	if m != nil {
		return m.Disk
	}
	return 0
}

func (m *ResourceLimits) GetCache() ResourceLimits_Cachability {
	if m != nil {
		return m.Cache
	}
	return ResourceLimits_NONE
}

// Permissions allow tokens to perform certain actions.
type Permissions struct {
	// This token has permission to toggle the always on state of a container.
	// For a connection to send the AlwaysOn message, it must have this permission.
	ToggleAlwaysOn       bool     `protobuf:"varint,1,opt,name=toggleAlwaysOn,proto3" json:"toggleAlwaysOn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{2}
}

func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (m *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(m, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetToggleAlwaysOn() bool {
	if m != nil {
		return m.ToggleAlwaysOn
	}
	return false
}

// ReplToken is the expected client options during the handshake. This is encoded
// into the token that is used to connect using WebSocket.
type ReplToken struct {
	// Issue timestamp. Equivalent to JWT's "iat" (Issued At) claim.  Tokens with
	// no `iat` field will be treated as if they had been issed at the UNIX epoch
	// (1970-01-01T00:00:00Z).
	Iat *timestamp.Timestamp `protobuf:"bytes,1,opt,name=iat,proto3" json:"iat,omitempty"`
	// Expiration timestamp. Equivalent to JWT's "exp" (Expiration Time) Claim.
	// If unset, will default to one hour after `iat`.
	Exp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=exp,proto3" json:"exp,omitempty"`
	// An arbitrary string that helps prevent replay attacks by ensuring that all
	// tokens are distinct.
	Salt string `protobuf:"bytes,3,opt,name=salt,proto3" json:"salt,omitempty"`
	// The cluster that a repl is located in. This prevents replay attacks in
	// which a user is given a token for one cluster and then presents that same
	// token to a conman instance in another token, which could lead to a case
	// where multiple containers are associated with a repl.
	//
	// Conman therefore needs to validate that this parameter matches the
	// `-cluster` flag it was started with.
	Cluster string `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Whether to persist filesystem, metadata, or both.  When connecting to an
	// already running/existing repl, its settings will be updated to match this
	// mode.
	Persistence ReplToken_Persistence `protobuf:"varint,6,opt,name=persistence,proto3,enum=api.ReplToken_Persistence" json:"persistence,omitempty"`
	// One of the three ways to identify a repl in goval.
	//
	// Types that are valid to be assigned to Metadata:
	//	*ReplToken_Repl
	//	*ReplToken_Id
	//	*ReplToken_Classroom
	Metadata isReplToken_Metadata `protobuf_oneof:"metadata"`
	// The resource limits for the container.
	ResourceLimits *ResourceLimits      `protobuf:"bytes,10,opt,name=resourceLimits,proto3" json:"resourceLimits,omitempty"`
	Format         ReplToken_WireFormat `protobuf:"varint,12,opt,name=format,proto3,enum=api.ReplToken_WireFormat" json:"format,omitempty"`
	Presenced      *ReplToken_Presenced `protobuf:"bytes,13,opt,name=presenced,proto3" json:"presenced,omitempty"`
	// Flags are handy for passing arbitrary configs along. Mostly used so
	// the client can try out new features
	Flags                []string     `protobuf:"bytes,14,rep,name=flags,proto3" json:"flags,omitempty"`
	Permissions          *Permissions `protobuf:"bytes,15,opt,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplToken) Reset()         { *m = ReplToken{} }
func (m *ReplToken) String() string { return proto.CompactTextString(m) }
func (*ReplToken) ProtoMessage()    {}
func (*ReplToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3}
}

func (m *ReplToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplToken.Unmarshal(m, b)
}
func (m *ReplToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplToken.Marshal(b, m, deterministic)
}
func (m *ReplToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplToken.Merge(m, src)
}
func (m *ReplToken) XXX_Size() int {
	return xxx_messageInfo_ReplToken.Size(m)
}
func (m *ReplToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplToken.DiscardUnknown(m)
}

var xxx_messageInfo_ReplToken proto.InternalMessageInfo

func (m *ReplToken) GetIat() *timestamp.Timestamp {
	if m != nil {
		return m.Iat
	}
	return nil
}

func (m *ReplToken) GetExp() *timestamp.Timestamp {
	if m != nil {
		return m.Exp
	}
	return nil
}

func (m *ReplToken) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *ReplToken) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *ReplToken) GetPersistence() ReplToken_Persistence {
	if m != nil {
		return m.Persistence
	}
	return ReplToken_PERSISTENT
}

type isReplToken_Metadata interface {
	isReplToken_Metadata()
}

type ReplToken_Repl struct {
	Repl *Repl `protobuf:"bytes,7,opt,name=repl,proto3,oneof"`
}

type ReplToken_Id struct {
	Id *ReplToken_ReplID `protobuf:"bytes,8,opt,name=id,proto3,oneof"`
}

type ReplToken_Classroom struct {
	Classroom *ReplToken_ClassroomMetadata `protobuf:"bytes,9,opt,name=classroom,proto3,oneof"`
}

func (*ReplToken_Repl) isReplToken_Metadata() {}

func (*ReplToken_Id) isReplToken_Metadata() {}

func (*ReplToken_Classroom) isReplToken_Metadata() {}

func (m *ReplToken) GetMetadata() isReplToken_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ReplToken) GetRepl() *Repl {
	if x, ok := m.GetMetadata().(*ReplToken_Repl); ok {
		return x.Repl
	}
	return nil
}

func (m *ReplToken) GetId() *ReplToken_ReplID {
	if x, ok := m.GetMetadata().(*ReplToken_Id); ok {
		return x.Id
	}
	return nil
}

// Deprecated: Do not use.
func (m *ReplToken) GetClassroom() *ReplToken_ClassroomMetadata {
	if x, ok := m.GetMetadata().(*ReplToken_Classroom); ok {
		return x.Classroom
	}
	return nil
}

func (m *ReplToken) GetResourceLimits() *ResourceLimits {
	if m != nil {
		return m.ResourceLimits
	}
	return nil
}

func (m *ReplToken) GetFormat() ReplToken_WireFormat {
	if m != nil {
		return m.Format
	}
	return ReplToken_PROTOBUF
}

func (m *ReplToken) GetPresenced() *ReplToken_Presenced {
	if m != nil {
		return m.Presenced
	}
	return nil
}

func (m *ReplToken) GetFlags() []string {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *ReplToken) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReplToken) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReplToken_Repl)(nil),
		(*ReplToken_Id)(nil),
		(*ReplToken_Classroom)(nil),
	}
}

// Metadata for the classroom. This is deprecated and should be removed
// hopefully soon.
type ReplToken_ClassroomMetadata struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Language             string   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplToken_ClassroomMetadata) Reset()         { *m = ReplToken_ClassroomMetadata{} }
func (m *ReplToken_ClassroomMetadata) String() string { return proto.CompactTextString(m) }
func (*ReplToken_ClassroomMetadata) ProtoMessage()    {}
func (*ReplToken_ClassroomMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3, 0}
}

func (m *ReplToken_ClassroomMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplToken_ClassroomMetadata.Unmarshal(m, b)
}
func (m *ReplToken_ClassroomMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplToken_ClassroomMetadata.Marshal(b, m, deterministic)
}
func (m *ReplToken_ClassroomMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplToken_ClassroomMetadata.Merge(m, src)
}
func (m *ReplToken_ClassroomMetadata) XXX_Size() int {
	return xxx_messageInfo_ReplToken_ClassroomMetadata.Size(m)
}
func (m *ReplToken_ClassroomMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplToken_ClassroomMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ReplToken_ClassroomMetadata proto.InternalMessageInfo

func (m *ReplToken_ClassroomMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReplToken_ClassroomMetadata) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

// Metadata for a repl that is only identified by its id.
type ReplToken_ReplID struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// (Optional) See the comment for Repl.sourceRepl.
	SourceRepl           string   `protobuf:"bytes,2,opt,name=sourceRepl,proto3" json:"sourceRepl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplToken_ReplID) Reset()         { *m = ReplToken_ReplID{} }
func (m *ReplToken_ReplID) String() string { return proto.CompactTextString(m) }
func (*ReplToken_ReplID) ProtoMessage()    {}
func (*ReplToken_ReplID) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3, 1}
}

func (m *ReplToken_ReplID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplToken_ReplID.Unmarshal(m, b)
}
func (m *ReplToken_ReplID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplToken_ReplID.Marshal(b, m, deterministic)
}
func (m *ReplToken_ReplID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplToken_ReplID.Merge(m, src)
}
func (m *ReplToken_ReplID) XXX_Size() int {
	return xxx_messageInfo_ReplToken_ReplID.Size(m)
}
func (m *ReplToken_ReplID) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplToken_ReplID.DiscardUnknown(m)
}

var xxx_messageInfo_ReplToken_ReplID proto.InternalMessageInfo

func (m *ReplToken_ReplID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ReplToken_ReplID) GetSourceRepl() string {
	if m != nil {
		return m.SourceRepl
	}
	return ""
}

type ReplToken_Presenced struct {
	BearerID             uint32   `protobuf:"varint,1,opt,name=bearerID,proto3" json:"bearerID,omitempty"`
	BearerName           string   `protobuf:"bytes,2,opt,name=bearerName,proto3" json:"bearerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplToken_Presenced) Reset()         { *m = ReplToken_Presenced{} }
func (m *ReplToken_Presenced) String() string { return proto.CompactTextString(m) }
func (*ReplToken_Presenced) ProtoMessage()    {}
func (*ReplToken_Presenced) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{3, 2}
}

func (m *ReplToken_Presenced) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplToken_Presenced.Unmarshal(m, b)
}
func (m *ReplToken_Presenced) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplToken_Presenced.Marshal(b, m, deterministic)
}
func (m *ReplToken_Presenced) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplToken_Presenced.Merge(m, src)
}
func (m *ReplToken_Presenced) XXX_Size() int {
	return xxx_messageInfo_ReplToken_Presenced.Size(m)
}
func (m *ReplToken_Presenced) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplToken_Presenced.DiscardUnknown(m)
}

var xxx_messageInfo_ReplToken_Presenced proto.InternalMessageInfo

func (m *ReplToken_Presenced) GetBearerID() uint32 {
	if m != nil {
		return m.BearerID
	}
	return 0
}

func (m *ReplToken_Presenced) GetBearerName() string {
	if m != nil {
		return m.BearerName
	}
	return ""
}

// GovalTokenMetadata is information about a goval token, that can be used to
// validate it. It is stored in the footer of the PASETO.
type GovalTokenMetadata struct {
	// The ID of the key that was used to sign the token.
	KeyId                string   `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GovalTokenMetadata) Reset()         { *m = GovalTokenMetadata{} }
func (m *GovalTokenMetadata) String() string { return proto.CompactTextString(m) }
func (*GovalTokenMetadata) ProtoMessage()    {}
func (*GovalTokenMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{4}
}

func (m *GovalTokenMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GovalTokenMetadata.Unmarshal(m, b)
}
func (m *GovalTokenMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GovalTokenMetadata.Marshal(b, m, deterministic)
}
func (m *GovalTokenMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GovalTokenMetadata.Merge(m, src)
}
func (m *GovalTokenMetadata) XXX_Size() int {
	return xxx_messageInfo_GovalTokenMetadata.Size(m)
}
func (m *GovalTokenMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GovalTokenMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GovalTokenMetadata proto.InternalMessageInfo

func (m *GovalTokenMetadata) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.ResourceLimits_Cachability", ResourceLimits_Cachability_name, ResourceLimits_Cachability_value)
	proto.RegisterEnum("api.ReplToken_Persistence", ReplToken_Persistence_name, ReplToken_Persistence_value)
	proto.RegisterEnum("api.ReplToken_WireFormat", ReplToken_WireFormat_name, ReplToken_WireFormat_value)
	proto.RegisterType((*Repl)(nil), "api.Repl")
	proto.RegisterType((*ResourceLimits)(nil), "api.ResourceLimits")
	proto.RegisterType((*Permissions)(nil), "api.Permissions")
	proto.RegisterType((*ReplToken)(nil), "api.ReplToken")
	proto.RegisterType((*ReplToken_ClassroomMetadata)(nil), "api.ReplToken.ClassroomMetadata")
	proto.RegisterType((*ReplToken_ReplID)(nil), "api.ReplToken.ReplID")
	proto.RegisterType((*ReplToken_Presenced)(nil), "api.ReplToken.Presenced")
	proto.RegisterType((*GovalTokenMetadata)(nil), "api.GovalTokenMetadata")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6f, 0xeb, 0x44,
	0x10, 0x8d, 0x9d, 0x8f, 0xc6, 0x93, 0x34, 0x98, 0x85, 0x8b, 0x8c, 0x1f, 0x68, 0xe5, 0x07, 0xa8,
	0x74, 0x91, 0x2b, 0x8a, 0xee, 0x15, 0x12, 0x48, 0xe8, 0xa6, 0x75, 0x9b, 0xa0, 0x36, 0x89, 0x36,
	0xa9, 0x90, 0x78, 0x41, 0x1b, 0x67, 0xeb, 0xac, 0x62, 0xc7, 0xd6, 0xae, 0x03, 0xcd, 0x2f, 0xe0,
	0x9d, 0xdf, 0xc8, 0x0f, 0x41, 0x3b, 0xb6, 0x93, 0x34, 0x45, 0x42, 0xf7, 0x6d, 0x66, 0xf6, 0xec,
	0xec, 0xf1, 0x39, 0x33, 0x86, 0x6e, 0x18, 0x0b, 0xbe, 0xce, 0xfd, 0x4c, 0xa6, 0x79, 0x4a, 0xea,
	0x2c, 0x13, 0xee, 0x59, 0x94, 0xa6, 0x51, 0xcc, 0x2f, 0xb1, 0x34, 0xdf, 0x3c, 0x5d, 0xe6, 0x22,
	0xe1, 0x2a, 0x67, 0x49, 0x56, 0xa0, 0xbc, 0xbf, 0x0d, 0x68, 0x50, 0x9e, 0xc5, 0xa4, 0x07, 0xa6,
	0x58, 0x38, 0xc6, 0xb9, 0x71, 0x61, 0x51, 0x53, 0x2c, 0x88, 0x0b, 0xed, 0x98, 0xad, 0xa3, 0x0d,
	0x8b, 0xb8, 0x63, 0x62, 0x75, 0x97, 0x93, 0x2f, 0xa0, 0x35, 0xdf, 0x84, 0x2b, 0x9e, 0x3b, 0x75,
	0x3c, 0x29, 0x33, 0x42, 0xa0, 0xa1, 0xe2, 0x4d, 0xe4, 0x34, 0xb0, 0x8a, 0xb1, 0xae, 0x6d, 0x14,
	0x97, 0x4e, 0xb3, 0xa8, 0xe9, 0x98, 0x7c, 0x05, 0xa0, 0xd2, 0x8d, 0x0c, 0xb9, 0x7e, 0xd9, 0x69,
	0xe1, 0xc9, 0x41, 0xc5, 0xfb, 0xc7, 0x80, 0x1e, 0xe5, 0x45, 0xe1, 0x5e, 0x24, 0x22, 0x57, 0xc4,
	0x86, 0xfa, 0x9a, 0xe7, 0xc8, 0xaf, 0x4d, 0x75, 0xa8, 0x49, 0x24, 0x3c, 0x49, 0xe5, 0x16, 0xe9,
	0xd5, 0x69, 0x99, 0x11, 0x07, 0x4e, 0xf2, 0xa5, 0xe4, 0x6c, 0xa1, 0x90, 0x9d, 0x41, 0xab, 0x54,
	0xdf, 0x50, 0x4b, 0x26, 0xb9, 0x42, 0x82, 0x06, 0x2d, 0x33, 0x4d, 0x71, 0x21, 0xd4, 0x0a, 0x29,
	0xd6, 0x29, 0xc6, 0xe4, 0x1d, 0x34, 0x43, 0x16, 0x2e, 0x39, 0xb2, 0xeb, 0x5d, 0x9d, 0xf9, 0x2c,
	0x13, 0xfe, 0x4b, 0x4e, 0xfe, 0x35, 0x0b, 0x97, 0x6c, 0x2e, 0x62, 0x91, 0x6f, 0x69, 0x81, 0xf6,
	0xde, 0x42, 0xe7, 0xa0, 0x4a, 0xda, 0xd0, 0x18, 0x8d, 0x47, 0x81, 0x5d, 0xd3, 0xd1, 0xe3, 0x34,
	0xa0, 0xb6, 0xa1, 0x23, 0x1a, 0x4c, 0xee, 0x6d, 0xd3, 0x7b, 0x07, 0x9d, 0x09, 0x97, 0x89, 0x50,
	0x4a, 0xa4, 0x6b, 0x45, 0xbe, 0x86, 0x5e, 0x9e, 0x46, 0x51, 0xcc, 0x3f, 0xc4, 0x7f, 0xb2, 0xad,
	0x1a, 0xaf, 0xcb, 0xaf, 0x3d, 0xaa, 0x7a, 0x7f, 0x9d, 0x80, 0xa5, 0x65, 0x9a, 0xa5, 0x2b, 0xbe,
	0x26, 0xdf, 0x42, 0x5d, 0xb0, 0x42, 0x98, 0xce, 0x95, 0xeb, 0x17, 0x7e, 0xfb, 0x95, 0xdf, 0xfe,
	0xac, 0xf2, 0x9b, 0x6a, 0x98, 0x46, 0xf3, 0xe7, 0x0c, 0x15, 0xfb, 0x1f, 0x34, 0x7f, 0xce, 0xd0,
	0x4f, 0x16, 0x57, 0x2e, 0x63, 0xac, 0xe5, 0x0d, 0xe3, 0x8d, 0xca, 0xb9, 0x2c, 0x6d, 0xae, 0x52,
	0xf2, 0x13, 0x74, 0x32, 0x2e, 0x95, 0x50, 0x39, 0x5f, 0x87, 0x95, 0x70, 0x6e, 0x29, 0x5c, 0x49,
	0xd7, 0x9f, 0xec, 0x11, 0xf4, 0x10, 0x4e, 0xce, 0xa0, 0x21, 0xf5, 0x34, 0x9c, 0x20, 0x35, 0x6b,
	0x77, 0x6d, 0x50, 0xa3, 0x78, 0x40, 0xbe, 0xc1, 0x01, 0x6d, 0xe3, 0xf1, 0x9b, 0xa3, 0xae, 0x3a,
	0x1a, 0xde, 0x0c, 0x6a, 0x38, 0xb9, 0x37, 0x60, 0x85, 0x31, 0x53, 0x4a, 0xa6, 0x69, 0xe2, 0x58,
	0x88, 0x3f, 0x3f, 0xc2, 0x5f, 0x57, 0xe7, 0x0f, 0x3c, 0x67, 0x0b, 0x96, 0xb3, 0xbe, 0xe9, 0x18,
	0x83, 0x1a, 0xdd, 0x5f, 0x24, 0x3f, 0x42, 0x4f, 0xbe, 0xb0, 0xdb, 0x01, 0x6c, 0xf5, 0xd9, 0x7f,
	0x4c, 0x02, 0x3d, 0x82, 0x92, 0xef, 0xa0, 0xf5, 0x94, 0xca, 0x84, 0xe5, 0x4e, 0x17, 0x55, 0xf8,
	0xf2, 0xe8, 0xfd, 0x5f, 0x85, 0xe4, 0xb7, 0x08, 0xa0, 0x25, 0x90, 0xbc, 0x07, 0x2b, 0x93, 0x5c,
	0x69, 0x2d, 0x16, 0xce, 0x29, 0x3e, 0xe5, 0x1c, 0x6b, 0x57, 0x9d, 0xd3, 0x3d, 0x94, 0x7c, 0x0e,
	0xcd, 0xa7, 0x98, 0x45, 0xca, 0xe9, 0x9d, 0xd7, 0x2f, 0x2c, 0x5a, 0x24, 0xe4, 0x0a, 0xbd, 0xa8,
	0x46, 0xcb, 0xf9, 0x04, 0xfb, 0xd9, 0xd8, 0xef, 0x60, 0xe4, 0xe8, 0x21, 0xc8, 0xfd, 0x19, 0x3e,
	0x7d, 0xa5, 0xcb, 0xc7, 0xfc, 0x16, 0xdc, 0x1f, 0xa0, 0x55, 0x18, 0xf1, 0xea, 0xd6, 0xcb, 0x85,
	0x37, 0x8f, 0x17, 0xde, 0xbd, 0x03, 0x6b, 0xf7, 0x71, 0xfa, 0x89, 0x39, 0x67, 0x92, 0xcb, 0xe1,
	0x0d, 0xb6, 0x38, 0xa5, 0xbb, 0x5c, 0x37, 0x2a, 0xe2, 0x11, 0x4b, 0x2a, 0x02, 0x07, 0x15, 0xef,
	0x3d, 0xae, 0xd4, 0x6e, 0xa8, 0x7a, 0x00, 0x93, 0x80, 0x4e, 0x87, 0xd3, 0x59, 0x30, 0x9a, 0xd9,
	0x35, 0x72, 0x0a, 0x56, 0x30, 0x19, 0x04, 0x0f, 0x01, 0xfd, 0x70, 0x5f, 0xac, 0x22, 0xae, 0xa7,
	0xe9, 0x5d, 0x00, 0xec, 0x3d, 0x21, 0x5d, 0x68, 0x4f, 0xe8, 0x78, 0x36, 0xee, 0x3f, 0xde, 0xda,
	0x35, 0xd2, 0x85, 0xc6, 0x2f, 0xd3, 0xf1, 0xc8, 0x36, 0x5c, 0xb3, 0x6d, 0xf4, 0x01, 0xda, 0x49,
	0x29, 0x8e, 0xf7, 0x16, 0xc8, 0x5d, 0xfa, 0x07, 0x2b, 0xec, 0xd9, 0x49, 0xf6, 0x06, 0x5a, 0x2b,
	0xbe, 0xfd, 0x7d, 0x27, 0x40, 0x73, 0xc5, 0xb7, 0xc3, 0x45, 0xbf, 0xf5, 0x5b, 0xe3, 0x92, 0x65,
	0x62, 0xde, 0xc2, 0x6d, 0xfb, 0xfe, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x01, 0x61, 0x44,
	0xae, 0x05, 0x00, 0x00,
}
