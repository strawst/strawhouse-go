// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: proto/driver/metadata.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type File struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Directory     string                 `protobuf:"bytes,2,opt,name=directory,proto3" json:"directory,omitempty"`
	Checksum      string                 `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
	Size          int64                  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Mtime         int64                  `protobuf:"varint,5,opt,name=mtime,proto3" json:"mtime,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *File) Reset() {
	*x = File{}
	mi := &file_proto_driver_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_proto_driver_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_proto_driver_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

func (x *File) GetChecksum() string {
	if x != nil {
		return x.Checksum
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetMtime() int64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

type Directory struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path          string                 `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Directory) Reset() {
	*x = Directory{}
	mi := &file_proto_driver_metadata_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Directory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Directory) ProtoMessage() {}

func (x *Directory) ProtoReflect() protoreflect.Message {
	mi := &file_proto_driver_metadata_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Directory.ProtoReflect.Descriptor instead.
func (*Directory) Descriptor() ([]byte, []int) {
	return file_proto_driver_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *Directory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Directory) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DirectoryListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Directory     string                 `protobuf:"bytes,1,opt,name=directory,proto3" json:"directory,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DirectoryListRequest) Reset() {
	*x = DirectoryListRequest{}
	mi := &file_proto_driver_metadata_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DirectoryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryListRequest) ProtoMessage() {}

func (x *DirectoryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_driver_metadata_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryListRequest.ProtoReflect.Descriptor instead.
func (*DirectoryListRequest) Descriptor() ([]byte, []int) {
	return file_proto_driver_metadata_proto_rawDescGZIP(), []int{2}
}

func (x *DirectoryListRequest) GetDirectory() string {
	if x != nil {
		return x.Directory
	}
	return ""
}

type DirectoryListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Files         []*File                `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
	Directories   []*Directory           `protobuf:"bytes,3,rep,name=directories,proto3" json:"directories,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DirectoryListResponse) Reset() {
	*x = DirectoryListResponse{}
	mi := &file_proto_driver_metadata_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DirectoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DirectoryListResponse) ProtoMessage() {}

func (x *DirectoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_driver_metadata_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DirectoryListResponse.ProtoReflect.Descriptor instead.
func (*DirectoryListResponse) Descriptor() ([]byte, []int) {
	return file_proto_driver_metadata_proto_rawDescGZIP(), []int{3}
}

func (x *DirectoryListResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DirectoryListResponse) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *DirectoryListResponse) GetDirectories() []*Directory {
	if x != nil {
		return x.Directories
	}
	return nil
}

var File_proto_driver_metadata_proto protoreflect.FileDescriptor

var file_proto_driver_metadata_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d,
	0x74, 0x69, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x34, 0x0a, 0x14, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x22,
	0x82, 0x01, 0x0a, 0x15, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x12, 0x32, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x32, 0x5c, 0x0a, 0x0e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x4a, 0x0a, 0x0d, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_proto_driver_metadata_proto_rawDescOnce sync.Once
	file_proto_driver_metadata_proto_rawDescData []byte
)

func file_proto_driver_metadata_proto_rawDescGZIP() []byte {
	file_proto_driver_metadata_proto_rawDescOnce.Do(func() {
		file_proto_driver_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_driver_metadata_proto_rawDesc), len(file_proto_driver_metadata_proto_rawDesc)))
	})
	return file_proto_driver_metadata_proto_rawDescData
}

var file_proto_driver_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_driver_metadata_proto_goTypes = []any{
	(*File)(nil),                  // 0: proto.File
	(*Directory)(nil),             // 1: proto.Directory
	(*DirectoryListRequest)(nil),  // 2: proto.DirectoryListRequest
	(*DirectoryListResponse)(nil), // 3: proto.DirectoryListResponse
}
var file_proto_driver_metadata_proto_depIdxs = []int32{
	0, // 0: proto.DirectoryListResponse.files:type_name -> proto.File
	1, // 1: proto.DirectoryListResponse.directories:type_name -> proto.Directory
	2, // 2: proto.DriverMetadata.DirectoryList:input_type -> proto.DirectoryListRequest
	3, // 3: proto.DriverMetadata.DirectoryList:output_type -> proto.DirectoryListResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_driver_metadata_proto_init() }
func file_proto_driver_metadata_proto_init() {
	if File_proto_driver_metadata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_driver_metadata_proto_rawDesc), len(file_proto_driver_metadata_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_driver_metadata_proto_goTypes,
		DependencyIndexes: file_proto_driver_metadata_proto_depIdxs,
		MessageInfos:      file_proto_driver_metadata_proto_msgTypes,
	}.Build()
	File_proto_driver_metadata_proto = out.File
	file_proto_driver_metadata_proto_goTypes = nil
	file_proto_driver_metadata_proto_depIdxs = nil
}
