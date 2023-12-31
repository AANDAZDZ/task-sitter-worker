// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: schedule.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ScheduleTaskReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskName      string `protobuf:"bytes,1,opt,name=TaskName,proto3" json:"TaskName,omitempty"`
	TaskType      string `protobuf:"bytes,2,opt,name=TaskType,proto3" json:"TaskType,omitempty"`
	TaskContext   string `protobuf:"bytes,3,opt,name=TaskContext,proto3" json:"TaskContext,omitempty"`
	BeginTime     int64  `protobuf:"varint,4,opt,name=BeginTime,proto3" json:"BeginTime,omitempty"`
	EndTime       int64  `protobuf:"varint,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	RetryNum      int32  `protobuf:"varint,6,opt,name=RetryNum,proto3" json:"RetryNum,omitempty"`
	RetryInterval int64  `protobuf:"varint,7,opt,name=RetryInterval,proto3" json:"RetryInterval,omitempty"`
	BizName       string `protobuf:"bytes,8,opt,name=BizName,proto3" json:"BizName,omitempty"`
	RateInterval  int64  `protobuf:"varint,9,opt,name=RateInterval,proto3" json:"RateInterval,omitempty"`
}

func (x *ScheduleTaskReq) Reset() {
	*x = ScheduleTaskReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleTaskReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleTaskReq) ProtoMessage() {}

func (x *ScheduleTaskReq) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleTaskReq.ProtoReflect.Descriptor instead.
func (*ScheduleTaskReq) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *ScheduleTaskReq) GetTaskName() string {
	if x != nil {
		return x.TaskName
	}
	return ""
}

func (x *ScheduleTaskReq) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *ScheduleTaskReq) GetTaskContext() string {
	if x != nil {
		return x.TaskContext
	}
	return ""
}

func (x *ScheduleTaskReq) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *ScheduleTaskReq) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ScheduleTaskReq) GetRetryNum() int32 {
	if x != nil {
		return x.RetryNum
	}
	return 0
}

func (x *ScheduleTaskReq) GetRetryInterval() int64 {
	if x != nil {
		return x.RetryInterval
	}
	return 0
}

func (x *ScheduleTaskReq) GetBizName() string {
	if x != nil {
		return x.BizName
	}
	return ""
}

func (x *ScheduleTaskReq) GetRateInterval() int64 {
	if x != nil {
		return x.RateInterval
	}
	return 0
}

type ScheduleTaskResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ScheduleTaskResp) Reset() {
	*x = ScheduleTaskResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleTaskResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleTaskResp) ProtoMessage() {}

func (x *ScheduleTaskResp) ProtoReflect() protoreflect.Message {
	mi := &file_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleTaskResp.ProtoReflect.Descriptor instead.
func (*ScheduleTaskResp) Descriptor() ([]byte, []int) {
	return file_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *ScheduleTaskResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_schedule_proto protoreflect.FileDescriptor

var file_schedule_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x02, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x52, 0x65, 0x74, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x52, 0x65, 0x74, 0x72, 0x79, 0x4e, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x74,
	0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x42, 0x69, 0x7a, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x42, 0x69, 0x7a, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x52, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x2c, 0x0a,
	0x10, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4b, 0x0a, 0x08, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x42, 0x1e, 0x5a, 0x1c, 0x74, 0x61, 0x73, 0x6b,
	0x2d, 0x73, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2d, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schedule_proto_rawDescOnce sync.Once
	file_schedule_proto_rawDescData = file_schedule_proto_rawDesc
)

func file_schedule_proto_rawDescGZIP() []byte {
	file_schedule_proto_rawDescOnce.Do(func() {
		file_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_schedule_proto_rawDescData)
	})
	return file_schedule_proto_rawDescData
}

var file_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_schedule_proto_goTypes = []interface{}{
	(*ScheduleTaskReq)(nil),  // 0: proto.ScheduleTaskReq
	(*ScheduleTaskResp)(nil), // 1: proto.ScheduleTaskResp
}
var file_schedule_proto_depIdxs = []int32{
	0, // 0: proto.Schedule.ScheduleTask:input_type -> proto.ScheduleTaskReq
	1, // 1: proto.Schedule.ScheduleTask:output_type -> proto.ScheduleTaskResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schedule_proto_init() }
func file_schedule_proto_init() {
	if File_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleTaskReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleTaskResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schedule_proto_goTypes,
		DependencyIndexes: file_schedule_proto_depIdxs,
		MessageInfos:      file_schedule_proto_msgTypes,
	}.Build()
	File_schedule_proto = out.File
	file_schedule_proto_rawDesc = nil
	file_schedule_proto_goTypes = nil
	file_schedule_proto_depIdxs = nil
}
