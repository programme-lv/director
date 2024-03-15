// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: msg/director.proto

package msg

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

type EvaluationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Submission     string               `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"`
	Language       *ProgrammingLanguage `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Limits         *RuntimeLimits       `protobuf:"bytes,3,opt,name=limits,proto3" json:"limits,omitempty"`
	EvalType       string               `protobuf:"bytes,4,opt,name=evalType,proto3" json:"evalType,omitempty"`
	Tests          []*Test              `protobuf:"bytes,5,rep,name=tests,proto3" json:"tests,omitempty"`
	TestlibChecker string               `protobuf:"bytes,6,opt,name=testlibChecker,proto3" json:"testlibChecker,omitempty"`
}

func (x *EvaluationRequest) Reset() {
	*x = EvaluationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationRequest) ProtoMessage() {}

func (x *EvaluationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationRequest.ProtoReflect.Descriptor instead.
func (*EvaluationRequest) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{0}
}

func (x *EvaluationRequest) GetSubmission() string {
	if x != nil {
		return x.Submission
	}
	return ""
}

func (x *EvaluationRequest) GetLanguage() *ProgrammingLanguage {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *EvaluationRequest) GetLimits() *RuntimeLimits {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *EvaluationRequest) GetEvalType() string {
	if x != nil {
		return x.EvalType
	}
	return ""
}

func (x *EvaluationRequest) GetTests() []*Test {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *EvaluationRequest) GetTestlibChecker() string {
	if x != nil {
		return x.TestlibChecker
	}
	return ""
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InSha256       string  `protobuf:"bytes,2,opt,name=inSha256,proto3" json:"inSha256,omitempty"`
	InDownloadUrl  *string `protobuf:"bytes,3,opt,name=inDownloadUrl,proto3,oneof" json:"inDownloadUrl,omitempty"`
	InContent      *string `protobuf:"bytes,4,opt,name=inContent,proto3,oneof" json:"inContent,omitempty"`
	AnsSha256      string  `protobuf:"bytes,5,opt,name=ansSha256,proto3" json:"ansSha256,omitempty"`
	AnsDownloadUrl *string `protobuf:"bytes,6,opt,name=ansDownloadUrl,proto3,oneof" json:"ansDownloadUrl,omitempty"`
	AnsContent     *string `protobuf:"bytes,7,opt,name=ansContent,proto3,oneof" json:"ansContent,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{1}
}

func (x *Test) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Test) GetInSha256() string {
	if x != nil {
		return x.InSha256
	}
	return ""
}

func (x *Test) GetInDownloadUrl() string {
	if x != nil && x.InDownloadUrl != nil {
		return *x.InDownloadUrl
	}
	return ""
}

func (x *Test) GetInContent() string {
	if x != nil && x.InContent != nil {
		return *x.InContent
	}
	return ""
}

func (x *Test) GetAnsSha256() string {
	if x != nil {
		return x.AnsSha256
	}
	return ""
}

func (x *Test) GetAnsDownloadUrl() string {
	if x != nil && x.AnsDownloadUrl != nil {
		return *x.AnsDownloadUrl
	}
	return ""
}

func (x *Test) GetAnsContent() string {
	if x != nil && x.AnsContent != nil {
		return *x.AnsContent
	}
	return ""
}

type ProgrammingLanguage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CodeFilename     string  `protobuf:"bytes,3,opt,name=codeFilename,proto3" json:"codeFilename,omitempty"`
	CompileCmd       *string `protobuf:"bytes,4,opt,name=compileCmd,proto3,oneof" json:"compileCmd,omitempty"`
	CompiledFilename *string `protobuf:"bytes,5,opt,name=compiledFilename,proto3,oneof" json:"compiledFilename,omitempty"`
	ExecuteCmd       string  `protobuf:"bytes,6,opt,name=executeCmd,proto3" json:"executeCmd,omitempty"`
}

func (x *ProgrammingLanguage) Reset() {
	*x = ProgrammingLanguage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProgrammingLanguage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgrammingLanguage) ProtoMessage() {}

func (x *ProgrammingLanguage) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProgrammingLanguage.ProtoReflect.Descriptor instead.
func (*ProgrammingLanguage) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{2}
}

func (x *ProgrammingLanguage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ProgrammingLanguage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProgrammingLanguage) GetCodeFilename() string {
	if x != nil {
		return x.CodeFilename
	}
	return ""
}

func (x *ProgrammingLanguage) GetCompileCmd() string {
	if x != nil && x.CompileCmd != nil {
		return *x.CompileCmd
	}
	return ""
}

func (x *ProgrammingLanguage) GetCompiledFilename() string {
	if x != nil && x.CompiledFilename != nil {
		return *x.CompiledFilename
	}
	return ""
}

func (x *ProgrammingLanguage) GetExecuteCmd() string {
	if x != nil {
		return x.ExecuteCmd
	}
	return ""
}

type RuntimeLimits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CPUTimeMillis int64 `protobuf:"varint,1,opt,name=CPUTimeMillis,proto3" json:"CPUTimeMillis,omitempty"`
	MemKibiBytes  int64 `protobuf:"varint,2,opt,name=memKibiBytes,proto3" json:"memKibiBytes,omitempty"`
}

func (x *RuntimeLimits) Reset() {
	*x = RuntimeLimits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuntimeLimits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuntimeLimits) ProtoMessage() {}

func (x *RuntimeLimits) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuntimeLimits.ProtoReflect.Descriptor instead.
func (*RuntimeLimits) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{3}
}

func (x *RuntimeLimits) GetCPUTimeMillis() int64 {
	if x != nil {
		return x.CPUTimeMillis
	}
	return 0
}

func (x *RuntimeLimits) GetMemKibiBytes() int64 {
	if x != nil {
		return x.MemKibiBytes
	}
	return 0
}

type EvaluationFeedback struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to FeedbackTypes:
	//
	//	*EvaluationFeedback_Start
	FeedbackTypes isEvaluationFeedback_FeedbackTypes `protobuf_oneof:"feedback_types"`
}

func (x *EvaluationFeedback) Reset() {
	*x = EvaluationFeedback{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluationFeedback) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluationFeedback) ProtoMessage() {}

func (x *EvaluationFeedback) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluationFeedback.ProtoReflect.Descriptor instead.
func (*EvaluationFeedback) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{4}
}

func (m *EvaluationFeedback) GetFeedbackTypes() isEvaluationFeedback_FeedbackTypes {
	if m != nil {
		return m.FeedbackTypes
	}
	return nil
}

func (x *EvaluationFeedback) GetStart() *StartEvaluation {
	if x, ok := x.GetFeedbackTypes().(*EvaluationFeedback_Start); ok {
		return x.Start
	}
	return nil
}

type isEvaluationFeedback_FeedbackTypes interface {
	isEvaluationFeedback_FeedbackTypes()
}

type EvaluationFeedback_Start struct {
	Start *StartEvaluation `protobuf:"bytes,1,opt,name=start,proto3,oneof"`
}

func (*EvaluationFeedback_Start) isEvaluationFeedback_FeedbackTypes() {}

type StartEvaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartEvaluation) Reset() {
	*x = StartEvaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartEvaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartEvaluation) ProtoMessage() {}

func (x *StartEvaluation) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartEvaluation.ProtoReflect.Descriptor instead.
func (*StartEvaluation) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{5}
}

type FinishWithInernalServerError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg string `protobuf:"bytes,1,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *FinishWithInernalServerError) Reset() {
	*x = FinishWithInernalServerError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishWithInernalServerError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishWithInernalServerError) ProtoMessage() {}

func (x *FinishWithInernalServerError) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishWithInernalServerError.ProtoReflect.Descriptor instead.
func (*FinishWithInernalServerError) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{6}
}

func (x *FinishWithInernalServerError) GetErrorMsg() string {
	if x != nil {
		return x.ErrorMsg
	}
	return ""
}

type FinishEvaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FinishEvaluation) Reset() {
	*x = FinishEvaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishEvaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishEvaluation) ProtoMessage() {}

func (x *FinishEvaluation) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishEvaluation.ProtoReflect.Descriptor instead.
func (*FinishEvaluation) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{7}
}

type StartCompilation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartCompilation) Reset() {
	*x = StartCompilation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_director_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartCompilation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartCompilation) ProtoMessage() {}

func (x *StartCompilation) ProtoReflect() protoreflect.Message {
	mi := &file_msg_director_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartCompilation.ProtoReflect.Descriptor instead.
func (*StartCompilation) Descriptor() ([]byte, []int) {
	return file_msg_director_proto_rawDescGZIP(), []int{8}
}

var File_msg_director_proto protoreflect.FileDescriptor

var file_msg_director_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x73, 0x67, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xfa, 0x01, 0x0a, 0x11, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x34, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d,
	0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x76, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x69, 0x62, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x6c, 0x69, 0x62, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x22, 0xb2, 0x02, 0x0a, 0x04, 0x54, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x6e, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x29, 0x0a, 0x0d, 0x69,
	0x6e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x69, 0x6e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09, 0x69, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x6e, 0x73,
	0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e,
	0x73, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x2b, 0x0a, 0x0e, 0x61, 0x6e, 0x73, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x0e, 0x61, 0x6e, 0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x61, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0a, 0x61, 0x6e, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x69, 0x6e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x61, 0x6e,
	0x73, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x61, 0x6e, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xf7, 0x01, 0x0a, 0x13,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x69, 0x6e, 0x67, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x64, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x6f, 0x64, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x43, 0x6d, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x43, 0x6d, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x2f, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x10, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6d, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x43, 0x6d,
	0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x43, 0x6d, 0x64,
	0x42, 0x13, 0x0a, 0x11, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x59, 0x0a, 0x0d, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x43, 0x50, 0x55, 0x54, 0x69, 0x6d,
	0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x43,
	0x50, 0x55, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x12, 0x22, 0x0a, 0x0c,
	0x6d, 0x65, 0x6d, 0x4b, 0x69, 0x62, 0x69, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x6d, 0x65, 0x6d, 0x4b, 0x69, 0x62, 0x69, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x22, 0x54, 0x0a, 0x12, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65,
	0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x42, 0x10, 0x0a, 0x0e, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x0a, 0x1c, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x57, 0x69, 0x74, 0x68, 0x49, 0x6e, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x12, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x55, 0x0a,
	0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x49, 0x0a, 0x12, 0x45, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b,
	0x22, 0x00, 0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x65, 0x2d, 0x6c, 0x76, 0x2f,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_director_proto_rawDescOnce sync.Once
	file_msg_director_proto_rawDescData = file_msg_director_proto_rawDesc
)

func file_msg_director_proto_rawDescGZIP() []byte {
	file_msg_director_proto_rawDescOnce.Do(func() {
		file_msg_director_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_director_proto_rawDescData)
	})
	return file_msg_director_proto_rawDescData
}

var file_msg_director_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_msg_director_proto_goTypes = []interface{}{
	(*EvaluationRequest)(nil),            // 0: msg.EvaluationRequest
	(*Test)(nil),                         // 1: msg.Test
	(*ProgrammingLanguage)(nil),          // 2: msg.ProgrammingLanguage
	(*RuntimeLimits)(nil),                // 3: msg.RuntimeLimits
	(*EvaluationFeedback)(nil),           // 4: msg.EvaluationFeedback
	(*StartEvaluation)(nil),              // 5: msg.StartEvaluation
	(*FinishWithInernalServerError)(nil), // 6: msg.FinishWithInernalServerError
	(*FinishEvaluation)(nil),             // 7: msg.FinishEvaluation
	(*StartCompilation)(nil),             // 8: msg.StartCompilation
}
var file_msg_director_proto_depIdxs = []int32{
	2, // 0: msg.EvaluationRequest.language:type_name -> msg.ProgrammingLanguage
	3, // 1: msg.EvaluationRequest.limits:type_name -> msg.RuntimeLimits
	1, // 2: msg.EvaluationRequest.tests:type_name -> msg.Test
	5, // 3: msg.EvaluationFeedback.start:type_name -> msg.StartEvaluation
	0, // 4: msg.Director.EvaluateSubmission:input_type -> msg.EvaluationRequest
	4, // 5: msg.Director.EvaluateSubmission:output_type -> msg.EvaluationFeedback
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_msg_director_proto_init() }
func file_msg_director_proto_init() {
	if File_msg_director_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_director_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluationRequest); i {
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
		file_msg_director_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_msg_director_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProgrammingLanguage); i {
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
		file_msg_director_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuntimeLimits); i {
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
		file_msg_director_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvaluationFeedback); i {
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
		file_msg_director_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartEvaluation); i {
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
		file_msg_director_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishWithInernalServerError); i {
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
		file_msg_director_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishEvaluation); i {
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
		file_msg_director_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartCompilation); i {
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
	file_msg_director_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_msg_director_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_msg_director_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*EvaluationFeedback_Start)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_director_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_director_proto_goTypes,
		DependencyIndexes: file_msg_director_proto_depIdxs,
		MessageInfos:      file_msg_director_proto_msgTypes,
	}.Build()
	File_msg_director_proto = out.File
	file_msg_director_proto_rawDesc = nil
	file_msg_director_proto_goTypes = nil
	file_msg_director_proto_depIdxs = nil
}