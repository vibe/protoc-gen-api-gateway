// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: webhooks/v1/request.proto

package v1beta

import (
	_ "github.com/ductone/protoc-gen-apigw/apigw/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// version contains the constant value "v1". Future versions of the Webhook body will use a different string.
	//
	// This value will match the "Webhook-Version" header.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Unique ID for this Webhook. Your receiver should only process this ID once.
	//
	// This value will match the "Webhook-Id" header.
	WebhookId string `protobuf:"bytes,2,opt,name=webhook_id,json=webhookId,proto3" json:"webhook_id,omitempty"`
	// If your receiver returns HTTP Status Code 202 Accepted, it MUST send its response to this URL as a POST
	// message body.
	//
	// If your receiver returns any other status code, it is expected to not use the callback url.
	//
	// This value will match the "Webhook-Callback-Url" header.
	CallbackUrl string `protobuf:"bytes,3,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	// The type of event that triggered this Webhook.
	//
	// This value will match the "Webhook-Event" header.
	//
	// The value will be one of:
	// - "webhooks.v1.PayloadTest"
	// - "webhooks.v1.PayloadPolicyApprovalStep"
	// - "webhooks.v1.PayloadPolicyPostAction"
	// - "webhooks.v1.PayloadProvisionStep"
	Event string `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	// The Payload of the webhook.
	// The payload will be different depending on the event type.
	//
	// The @type will be one of:
	// - "type.googleapis.com/webhooks.v1.PayloadTest"
	// - "type.googleapis.com/webhooks.v1.PayloadPolicyApprovalStep"
	// - "type.googleapis.com/webhooks.v1.PayloadPolicyPostAction"
	// - "type.googleapis.com/webhooks.v1.PayloadProvisionStep"
	//
	// And map to the associated payload message type.
	Payload *anypb.Any `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhooks_v1_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_webhooks_v1_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_webhooks_v1_request_proto_rawDescGZIP(), []int{0}
}

func (x *Body) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Body) GetWebhookId() string {
	if x != nil {
		return x.WebhookId
	}
	return ""
}

func (x *Body) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

func (x *Body) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Body) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

type PayloadTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PayloadTest) Reset() {
	*x = PayloadTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhooks_v1_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadTest) ProtoMessage() {}

func (x *PayloadTest) ProtoReflect() protoreflect.Message {
	mi := &file_webhooks_v1_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadTest.ProtoReflect.Descriptor instead.
func (*PayloadTest) Descriptor() ([]byte, []int) {
	return file_webhooks_v1_request_proto_rawDescGZIP(), []int{1}
}

type TaskView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId     string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	OtherField string `protobuf:"bytes,2,opt,name=other_field,json=otherField,proto3" json:"other_field,omitempty"`
}

func (x *TaskView) Reset() {
	*x = TaskView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhooks_v1_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskView) ProtoMessage() {}

func (x *TaskView) ProtoReflect() protoreflect.Message {
	mi := &file_webhooks_v1_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskView.ProtoReflect.Descriptor instead.
func (*TaskView) Descriptor() ([]byte, []int) {
	return file_webhooks_v1_request_proto_rawDescGZIP(), []int{2}
}

func (x *TaskView) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskView) GetOtherField() string {
	if x != nil {
		return x.OtherField
	}
	return ""
}

type PayloadPolicyApprovalStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A view of the task, contains the serialized task and paths to objects referenced by the task
	TaskView *TaskView `protobuf:"bytes,1,opt,name=task_view,json=taskView,proto3" json:"task_view,omitempty"`
	// List of serialized related objects.
	Expanded []*anypb.Any `protobuf:"bytes,2,rep,name=expanded,proto3" json:"expanded,omitempty"`
}

func (x *PayloadPolicyApprovalStep) Reset() {
	*x = PayloadPolicyApprovalStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhooks_v1_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadPolicyApprovalStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadPolicyApprovalStep) ProtoMessage() {}

func (x *PayloadPolicyApprovalStep) ProtoReflect() protoreflect.Message {
	mi := &file_webhooks_v1_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadPolicyApprovalStep.ProtoReflect.Descriptor instead.
func (*PayloadPolicyApprovalStep) Descriptor() ([]byte, []int) {
	return file_webhooks_v1_request_proto_rawDescGZIP(), []int{3}
}

func (x *PayloadPolicyApprovalStep) GetTaskView() *TaskView {
	if x != nil {
		return x.TaskView
	}
	return nil
}

func (x *PayloadPolicyApprovalStep) GetExpanded() []*anypb.Any {
	if x != nil {
		return x.Expanded
	}
	return nil
}

type PayloadPolicyPostAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A view of the task, contains the serialized task and paths to objects referenced by the task
	TaskView *TaskView `protobuf:"bytes,1,opt,name=task_view,json=taskView,proto3" json:"task_view,omitempty"`
	// List of serialized related objects.
	Expanded []*anypb.Any `protobuf:"bytes,2,rep,name=expanded,proto3" json:"expanded,omitempty"`
}

func (x *PayloadPolicyPostAction) Reset() {
	*x = PayloadPolicyPostAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhooks_v1_request_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadPolicyPostAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadPolicyPostAction) ProtoMessage() {}

func (x *PayloadPolicyPostAction) ProtoReflect() protoreflect.Message {
	mi := &file_webhooks_v1_request_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadPolicyPostAction.ProtoReflect.Descriptor instead.
func (*PayloadPolicyPostAction) Descriptor() ([]byte, []int) {
	return file_webhooks_v1_request_proto_rawDescGZIP(), []int{4}
}

func (x *PayloadPolicyPostAction) GetTaskView() *TaskView {
	if x != nil {
		return x.TaskView
	}
	return nil
}

func (x *PayloadPolicyPostAction) GetExpanded() []*anypb.Any {
	if x != nil {
		return x.Expanded
	}
	return nil
}

type PayloadProvisionStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A view of the task, contains the serialized task and paths to objects referenced by the task
	TaskView *TaskView `protobuf:"bytes,1,opt,name=task_view,json=taskView,proto3" json:"task_view,omitempty"`
	// List of serialized related objects.
	Expanded []*anypb.Any `protobuf:"bytes,2,rep,name=expanded,proto3" json:"expanded,omitempty"`
}

func (x *PayloadProvisionStep) Reset() {
	*x = PayloadProvisionStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhooks_v1_request_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadProvisionStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadProvisionStep) ProtoMessage() {}

func (x *PayloadProvisionStep) ProtoReflect() protoreflect.Message {
	mi := &file_webhooks_v1_request_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadProvisionStep.ProtoReflect.Descriptor instead.
func (*PayloadProvisionStep) Descriptor() ([]byte, []int) {
	return file_webhooks_v1_request_proto_rawDescGZIP(), []int{5}
}

func (x *PayloadProvisionStep) GetTaskView() *TaskView {
	if x != nil {
		return x.TaskView
	}
	return nil
}

func (x *PayloadProvisionStep) GetExpanded() []*anypb.Any {
	if x != nil {
		return x.Expanded
	}
	return nil
}

var File_webhooks_v1_request_proto protoreflect.FileDescriptor

var file_webhooks_v1_request_proto_rawDesc = []byte{
	0x0a, 0x19, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x61, 0x70, 0x69, 0x67, 0x77, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc6, 0x01, 0x0a, 0x04, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x3a, 0x1c, 0xaa, 0xde, 0x03, 0x18, 0x0a, 0x16, 0x18, 0x01, 0x22, 0x12,
	0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x22, 0x17, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x65, 0x73,
	0x74, 0x3a, 0x08, 0xaa, 0xde, 0x03, 0x04, 0x0a, 0x02, 0x18, 0x01, 0x22, 0x44, 0x0a, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x22, 0x8b, 0x01, 0x0a, 0x19, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x12,
	0x32, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x56,
	0x69, 0x65, 0x77, 0x12, 0x30, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x65, 0x78, 0x70,
	0x61, 0x6e, 0x64, 0x65, 0x64, 0x3a, 0x08, 0xaa, 0xde, 0x03, 0x04, 0x0a, 0x02, 0x18, 0x01, 0x22,
	0x89, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x50, 0x6f, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x12,
	0x30, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65,
	0x64, 0x3a, 0x08, 0xaa, 0xde, 0x03, 0x04, 0x0a, 0x02, 0x18, 0x01, 0x22, 0x86, 0x01, 0x0a, 0x14,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x65, 0x70, 0x12, 0x32, 0x0a, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x69, 0x65,
	0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x08,
	0x74, 0x61, 0x73, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x12, 0x30, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x61,
	0x6e, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x08, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x65, 0x64, 0x3a, 0x08, 0xaa, 0xde, 0x03, 0x04,
	0x0a, 0x02, 0x18, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x75, 0x63, 0x74, 0x6f, 0x6e, 0x65, 0x2f, 0x63, 0x31, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_webhooks_v1_request_proto_rawDescOnce sync.Once
	file_webhooks_v1_request_proto_rawDescData = file_webhooks_v1_request_proto_rawDesc
)

func file_webhooks_v1_request_proto_rawDescGZIP() []byte {
	file_webhooks_v1_request_proto_rawDescOnce.Do(func() {
		file_webhooks_v1_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_webhooks_v1_request_proto_rawDescData)
	})
	return file_webhooks_v1_request_proto_rawDescData
}

var file_webhooks_v1_request_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_webhooks_v1_request_proto_goTypes = []any{
	(*Body)(nil),                      // 0: webhooks.v1.Body
	(*PayloadTest)(nil),               // 1: webhooks.v1.PayloadTest
	(*TaskView)(nil),                  // 2: webhooks.v1.TaskView
	(*PayloadPolicyApprovalStep)(nil), // 3: webhooks.v1.PayloadPolicyApprovalStep
	(*PayloadPolicyPostAction)(nil),   // 4: webhooks.v1.PayloadPolicyPostAction
	(*PayloadProvisionStep)(nil),      // 5: webhooks.v1.PayloadProvisionStep
	(*anypb.Any)(nil),                 // 6: google.protobuf.Any
}
var file_webhooks_v1_request_proto_depIdxs = []int32{
	6, // 0: webhooks.v1.Body.payload:type_name -> google.protobuf.Any
	2, // 1: webhooks.v1.PayloadPolicyApprovalStep.task_view:type_name -> webhooks.v1.TaskView
	6, // 2: webhooks.v1.PayloadPolicyApprovalStep.expanded:type_name -> google.protobuf.Any
	2, // 3: webhooks.v1.PayloadPolicyPostAction.task_view:type_name -> webhooks.v1.TaskView
	6, // 4: webhooks.v1.PayloadPolicyPostAction.expanded:type_name -> google.protobuf.Any
	2, // 5: webhooks.v1.PayloadProvisionStep.task_view:type_name -> webhooks.v1.TaskView
	6, // 6: webhooks.v1.PayloadProvisionStep.expanded:type_name -> google.protobuf.Any
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_webhooks_v1_request_proto_init() }
func file_webhooks_v1_request_proto_init() {
	if File_webhooks_v1_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_webhooks_v1_request_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Body); i {
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
		file_webhooks_v1_request_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadTest); i {
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
		file_webhooks_v1_request_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TaskView); i {
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
		file_webhooks_v1_request_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadPolicyApprovalStep); i {
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
		file_webhooks_v1_request_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadPolicyPostAction); i {
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
		file_webhooks_v1_request_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PayloadProvisionStep); i {
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
			RawDescriptor: file_webhooks_v1_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_webhooks_v1_request_proto_goTypes,
		DependencyIndexes: file_webhooks_v1_request_proto_depIdxs,
		MessageInfos:      file_webhooks_v1_request_proto_msgTypes,
	}.Build()
	File_webhooks_v1_request_proto = out.File
	file_webhooks_v1_request_proto_rawDesc = nil
	file_webhooks_v1_request_proto_goTypes = nil
	file_webhooks_v1_request_proto_depIdxs = nil
}
