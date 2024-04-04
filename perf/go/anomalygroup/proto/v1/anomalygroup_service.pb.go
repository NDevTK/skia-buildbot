// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: anomalygroup_service.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The action to take on a certain group. It is defined in the Alert config.
type GroupActionType int32

const (
	// No action will be taken. It is used for backward compatibility
	// for the existing Alerts before grouping is introduced.
	GroupActionType_NOACTION GroupActionType = 0
	// File a bug with a list of anomalies.
	GroupActionType_REPORT GroupActionType = 1
	// Launch a bisection job on the most signification anomaly, in order to
	// find the culprit commit.
	GroupActionType_BISECT GroupActionType = 2
)

// Enum value maps for GroupActionType.
var (
	GroupActionType_name = map[int32]string{
		0: "NOACTION",
		1: "REPORT",
		2: "BISECT",
	}
	GroupActionType_value = map[string]int32{
		"NOACTION": 0,
		"REPORT":   1,
		"BISECT":   2,
	}
)

func (x GroupActionType) Enum() *GroupActionType {
	p := new(GroupActionType)
	*p = x
	return p
}

func (x GroupActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_anomalygroup_service_proto_enumTypes[0].Descriptor()
}

func (GroupActionType) Type() protoreflect.EnumType {
	return &file_anomalygroup_service_proto_enumTypes[0]
}

func (x GroupActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupActionType.Descriptor instead.
func (GroupActionType) EnumDescriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{0}
}

// Request object for CreateAnomalyGroup
type CreateAnomalyGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the subscription in sheriff config
	SubscriptionName string `protobuf:"bytes,1,opt,name=subscription_name,json=subscriptionName,proto3" json:"subscription_name,omitempty"`
	// The revision of the subscription
	SubscriptionRevision string `protobuf:"bytes,2,opt,name=subscription_revision,json=subscriptionRevision,proto3" json:"subscription_revision,omitempty"`
	// The domain of the test to group (The value of 'master' in alert config.)
	Domain string `protobuf:"bytes,3,opt,name=domain,proto3" json:"domain,omitempty"`
	// The benchmark of the test to group
	Benchmark string `protobuf:"bytes,4,opt,name=benchmark,proto3" json:"benchmark,omitempty"`
	// The current start commit position of the group
	StartCommit int64 `protobuf:"varint,5,opt,name=start_commit,json=startCommit,proto3" json:"start_commit,omitempty"`
	// The current end commit position of the group
	EndCommit int64 `protobuf:"varint,6,opt,name=end_commit,json=endCommit,proto3" json:"end_commit,omitempty"`
	// The action of the group to take.
	Action GroupActionType `protobuf:"varint,7,opt,name=action,proto3,enum=anomalygroup.v1.GroupActionType" json:"action,omitempty"`
}

func (x *CreateAnomalyGroupRequest) Reset() {
	*x = CreateAnomalyGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnomalyGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnomalyGroupRequest) ProtoMessage() {}

func (x *CreateAnomalyGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnomalyGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateAnomalyGroupRequest) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAnomalyGroupRequest) GetSubscriptionName() string {
	if x != nil {
		return x.SubscriptionName
	}
	return ""
}

func (x *CreateAnomalyGroupRequest) GetSubscriptionRevision() string {
	if x != nil {
		return x.SubscriptionRevision
	}
	return ""
}

func (x *CreateAnomalyGroupRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateAnomalyGroupRequest) GetBenchmark() string {
	if x != nil {
		return x.Benchmark
	}
	return ""
}

func (x *CreateAnomalyGroupRequest) GetStartCommit() int64 {
	if x != nil {
		return x.StartCommit
	}
	return 0
}

func (x *CreateAnomalyGroupRequest) GetEndCommit() int64 {
	if x != nil {
		return x.EndCommit
	}
	return 0
}

func (x *CreateAnomalyGroupRequest) GetAction() GroupActionType {
	if x != nil {
		return x.Action
	}
	return GroupActionType_NOACTION
}

// Response object for CreateAnomalyGroup
type CreateAnomalyGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The newly created anomaly group ID.
	AnomalyGroupId string `protobuf:"bytes,1,opt,name=anomaly_group_id,json=anomalyGroupId,proto3" json:"anomaly_group_id,omitempty"`
}

func (x *CreateAnomalyGroupResponse) Reset() {
	*x = CreateAnomalyGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAnomalyGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAnomalyGroupResponse) ProtoMessage() {}

func (x *CreateAnomalyGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAnomalyGroupResponse.ProtoReflect.Descriptor instead.
func (*CreateAnomalyGroupResponse) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAnomalyGroupResponse) GetAnomalyGroupId() string {
	if x != nil {
		return x.AnomalyGroupId
	}
	return ""
}

// Request object for ReadAnomalyGroup
type ReadAnomalyGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the anomaly group to read from
	AnomalyGroupId string `protobuf:"bytes,1,opt,name=anomaly_group_id,json=anomalyGroupId,proto3" json:"anomaly_group_id,omitempty"`
}

func (x *ReadAnomalyGroupRequest) Reset() {
	*x = ReadAnomalyGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAnomalyGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAnomalyGroupRequest) ProtoMessage() {}

func (x *ReadAnomalyGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAnomalyGroupRequest.ProtoReflect.Descriptor instead.
func (*ReadAnomalyGroupRequest) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReadAnomalyGroupRequest) GetAnomalyGroupId() string {
	if x != nil {
		return x.AnomalyGroupId
	}
	return ""
}

// Response object for ReadAnomalyGroup
type ReadAnomalyGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The anomaly group object
	AnomalyGroup *AnomalyGroup `protobuf:"bytes,1,opt,name=anomaly_group,json=anomalyGroup,proto3" json:"anomaly_group,omitempty"`
}

func (x *ReadAnomalyGroupResponse) Reset() {
	*x = ReadAnomalyGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadAnomalyGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadAnomalyGroupResponse) ProtoMessage() {}

func (x *ReadAnomalyGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadAnomalyGroupResponse.ProtoReflect.Descriptor instead.
func (*ReadAnomalyGroupResponse) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReadAnomalyGroupResponse) GetAnomalyGroup() *AnomalyGroup {
	if x != nil {
		return x.AnomalyGroup
	}
	return nil
}

// Request object for UpdateAnomalyGroup
type UpdateAnomalyGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the anomaly group to update
	AnomalyGroupId string `protobuf:"bytes,1,opt,name=anomaly_group_id,json=anomalyGroupId,proto3" json:"anomaly_group_id,omitempty"`
	// The anomaly ID to append to the anomaly list.
	AnomalyId string `protobuf:"bytes,2,opt,name=anomaly_id,json=anomalyId,proto3" json:"anomaly_id,omitempty"`
	// The biesction ID to add to the anomaly group.
	// This should be populated only when the action value is BISECT.
	BisectionId string `protobuf:"bytes,3,opt,name=bisection_id,json=bisectionId,proto3" json:"bisection_id,omitempty"`
	// The issue ID to add to the anomaly group.
	// This should be populated only when the action value is REPORT.
	IssueId string `protobuf:"bytes,4,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	// The culprit IDs correlated to the group.
	// Culprits are found by a bisection job. This should be populated
	// only when the action value is BISECT and the bisection_id exists.
	CulpritIds []string `protobuf:"bytes,5,rep,name=culprit_ids,json=culpritIds,proto3" json:"culprit_ids,omitempty"`
}

func (x *UpdateAnomalyGroupRequest) Reset() {
	*x = UpdateAnomalyGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnomalyGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnomalyGroupRequest) ProtoMessage() {}

func (x *UpdateAnomalyGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnomalyGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateAnomalyGroupRequest) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAnomalyGroupRequest) GetAnomalyGroupId() string {
	if x != nil {
		return x.AnomalyGroupId
	}
	return ""
}

func (x *UpdateAnomalyGroupRequest) GetAnomalyId() string {
	if x != nil {
		return x.AnomalyId
	}
	return ""
}

func (x *UpdateAnomalyGroupRequest) GetBisectionId() string {
	if x != nil {
		return x.BisectionId
	}
	return ""
}

func (x *UpdateAnomalyGroupRequest) GetIssueId() string {
	if x != nil {
		return x.IssueId
	}
	return ""
}

func (x *UpdateAnomalyGroupRequest) GetCulpritIds() []string {
	if x != nil {
		return x.CulpritIds
	}
	return nil
}

// Response object for UpdateAnomalyGroup
type UpdateAnomalyGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateAnomalyGroupResponse) Reset() {
	*x = UpdateAnomalyGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnomalyGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnomalyGroupResponse) ProtoMessage() {}

func (x *UpdateAnomalyGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnomalyGroupResponse.ProtoReflect.Descriptor instead.
func (*UpdateAnomalyGroupResponse) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{5}
}

// Request object for FindExistingGroups
type FindExistingGroupsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subscription name from the alert config, from which the
	// anomaly is detected.
	SubscriptionName string `protobuf:"bytes,1,opt,name=subscription_name,json=subscriptionName,proto3" json:"subscription_name,omitempty"`
	// The revision of the subscription
	SubscriptionRevision string `protobuf:"bytes,2,opt,name=subscription_revision,json=subscriptionRevision,proto3" json:"subscription_revision,omitempty"`
	// The action value from the alert config.
	Action GroupActionType `protobuf:"varint,3,opt,name=action,proto3,enum=anomalygroup.v1.GroupActionType" json:"action,omitempty"`
	// The previous commit position before the anomaly's data point.
	StartCommit int64 `protobuf:"varint,4,opt,name=start_commit,json=startCommit,proto3" json:"start_commit,omitempty"`
	// The commit position before the anomaly's data point.
	EndCommit int64 `protobuf:"varint,5,opt,name=end_commit,json=endCommit,proto3" json:"end_commit,omitempty"`
	// The test path from the anomaly.
	TestPath string `protobuf:"bytes,6,opt,name=test_path,json=testPath,proto3" json:"test_path,omitempty"`
}

func (x *FindExistingGroupsRequest) Reset() {
	*x = FindExistingGroupsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindExistingGroupsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindExistingGroupsRequest) ProtoMessage() {}

func (x *FindExistingGroupsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindExistingGroupsRequest.ProtoReflect.Descriptor instead.
func (*FindExistingGroupsRequest) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{6}
}

func (x *FindExistingGroupsRequest) GetSubscriptionName() string {
	if x != nil {
		return x.SubscriptionName
	}
	return ""
}

func (x *FindExistingGroupsRequest) GetSubscriptionRevision() string {
	if x != nil {
		return x.SubscriptionRevision
	}
	return ""
}

func (x *FindExistingGroupsRequest) GetAction() GroupActionType {
	if x != nil {
		return x.Action
	}
	return GroupActionType_NOACTION
}

func (x *FindExistingGroupsRequest) GetStartCommit() int64 {
	if x != nil {
		return x.StartCommit
	}
	return 0
}

func (x *FindExistingGroupsRequest) GetEndCommit() int64 {
	if x != nil {
		return x.EndCommit
	}
	return 0
}

func (x *FindExistingGroupsRequest) GetTestPath() string {
	if x != nil {
		return x.TestPath
	}
	return ""
}

// Response object for FindExistingGroups
type FindExistingGroupsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of anomaly groups which an anomaly can be added to.
	AnomalyGroups []*AnomalyGroup `protobuf:"bytes,1,rep,name=anomaly_groups,json=anomalyGroups,proto3" json:"anomaly_groups,omitempty"`
}

func (x *FindExistingGroupsResponse) Reset() {
	*x = FindExistingGroupsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindExistingGroupsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindExistingGroupsResponse) ProtoMessage() {}

func (x *FindExistingGroupsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindExistingGroupsResponse.ProtoReflect.Descriptor instead.
func (*FindExistingGroupsResponse) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{7}
}

func (x *FindExistingGroupsResponse) GetAnomalyGroups() []*AnomalyGroup {
	if x != nil {
		return x.AnomalyGroups
	}
	return nil
}

// Simplified format for an anomaly group, which should be sufficient
// in the following use cases:
//  1. provide a list of anomalies for filing a bug.
//  2. provide the most significant anomaly to launch a bisection.
//  3. for the new anomaly to be added in, and decide whether the new anomaly
//     needs to be added to an existing bug.
type AnomalyGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the anomaly group.
	GroupId string `protobuf:"bytes,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// The action to take for the anomaly group.
	GroupAction GroupActionType `protobuf:"varint,2,opt,name=group_action,json=groupAction,proto3,enum=anomalygroup.v1.GroupActionType" json:"group_action,omitempty"`
	// The anomalies added to this group.
	AnomalyIds []string `protobuf:"bytes,3,rep,name=anomaly_ids,json=anomalyIds,proto3" json:"anomaly_ids,omitempty"`
	// The culprits associated to this group.
	CulpritIds []string `protobuf:"bytes,4,rep,name=culprit_ids,json=culpritIds,proto3" json:"culprit_ids,omitempty"`
	// The reported issue associated to this group.
	ReportedIssueId int64 `protobuf:"varint,5,opt,name=reported_issue_id,json=reportedIssueId,proto3" json:"reported_issue_id,omitempty"`
}

func (x *AnomalyGroup) Reset() {
	*x = AnomalyGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_anomalygroup_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnomalyGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnomalyGroup) ProtoMessage() {}

func (x *AnomalyGroup) ProtoReflect() protoreflect.Message {
	mi := &file_anomalygroup_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnomalyGroup.ProtoReflect.Descriptor instead.
func (*AnomalyGroup) Descriptor() ([]byte, []int) {
	return file_anomalygroup_service_proto_rawDescGZIP(), []int{8}
}

func (x *AnomalyGroup) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *AnomalyGroup) GetGroupAction() GroupActionType {
	if x != nil {
		return x.GroupAction
	}
	return GroupActionType_NOACTION
}

func (x *AnomalyGroup) GetAnomalyIds() []string {
	if x != nil {
		return x.AnomalyIds
	}
	return nil
}

func (x *AnomalyGroup) GetCulpritIds() []string {
	if x != nil {
		return x.CulpritIds
	}
	return nil
}

func (x *AnomalyGroup) GetReportedIssueId() int64 {
	if x != nil {
		return x.ReportedIssueId
	}
	return 0
}

var File_anomalygroup_service_proto protoreflect.FileDescriptor

var file_anomalygroup_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x22, 0xaf, 0x02,
	0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x46, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x17, 0x52, 0x65, 0x61, 0x64, 0x41,
	0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x18,
	0x52, 0x65, 0x61, 0x64, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x61, 0x6e, 0x6f, 0x6d,
	0x61, 0x6c, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0c,
	0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0xc3, 0x01, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c,
	0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x69, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x69, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x49,
	0x64, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d,
	0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x96, 0x02, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b,
	0x0a, 0x11, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x15, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x73, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x62, 0x0a, 0x1a, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0e, 0x61, 0x6e, 0x6f, 0x6d, 0x61,
	0x6c, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x0d,
	0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0xdc, 0x01,
	0x0a, 0x0c, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0c, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x0b, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x49, 0x64, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x6c, 0x70, 0x72, 0x69, 0x74, 0x49, 0x64, 0x73,
	0x12, 0x2a, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x49, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x2a, 0x37, 0x0a, 0x0f,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x53,
	0x45, 0x43, 0x54, 0x10, 0x02, 0x32, 0xd3, 0x03, 0x0a, 0x13, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x2a, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d,
	0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69,
	0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x28, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61,
	0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x61, 0x64, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x2a, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x6e,
	0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6f, 0x0a, 0x12, 0x46, 0x69,
	0x6e, 0x64, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x2a, 0x2e, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61,
	0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x45, 0x78, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x6f, 0x2e, 0x73, 0x6b, 0x69, 0x61, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x2f, 0x70, 0x65, 0x72, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x6f, 0x6d, 0x61, 0x6c, 0x79,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_anomalygroup_service_proto_rawDescOnce sync.Once
	file_anomalygroup_service_proto_rawDescData = file_anomalygroup_service_proto_rawDesc
)

func file_anomalygroup_service_proto_rawDescGZIP() []byte {
	file_anomalygroup_service_proto_rawDescOnce.Do(func() {
		file_anomalygroup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_anomalygroup_service_proto_rawDescData)
	})
	return file_anomalygroup_service_proto_rawDescData
}

var file_anomalygroup_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_anomalygroup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_anomalygroup_service_proto_goTypes = []interface{}{
	(GroupActionType)(0),               // 0: anomalygroup.v1.GroupActionType
	(*CreateAnomalyGroupRequest)(nil),  // 1: anomalygroup.v1.CreateAnomalyGroupRequest
	(*CreateAnomalyGroupResponse)(nil), // 2: anomalygroup.v1.CreateAnomalyGroupResponse
	(*ReadAnomalyGroupRequest)(nil),    // 3: anomalygroup.v1.ReadAnomalyGroupRequest
	(*ReadAnomalyGroupResponse)(nil),   // 4: anomalygroup.v1.ReadAnomalyGroupResponse
	(*UpdateAnomalyGroupRequest)(nil),  // 5: anomalygroup.v1.UpdateAnomalyGroupRequest
	(*UpdateAnomalyGroupResponse)(nil), // 6: anomalygroup.v1.UpdateAnomalyGroupResponse
	(*FindExistingGroupsRequest)(nil),  // 7: anomalygroup.v1.FindExistingGroupsRequest
	(*FindExistingGroupsResponse)(nil), // 8: anomalygroup.v1.FindExistingGroupsResponse
	(*AnomalyGroup)(nil),               // 9: anomalygroup.v1.AnomalyGroup
}
var file_anomalygroup_service_proto_depIdxs = []int32{
	0, // 0: anomalygroup.v1.CreateAnomalyGroupRequest.action:type_name -> anomalygroup.v1.GroupActionType
	9, // 1: anomalygroup.v1.ReadAnomalyGroupResponse.anomaly_group:type_name -> anomalygroup.v1.AnomalyGroup
	0, // 2: anomalygroup.v1.FindExistingGroupsRequest.action:type_name -> anomalygroup.v1.GroupActionType
	9, // 3: anomalygroup.v1.FindExistingGroupsResponse.anomaly_groups:type_name -> anomalygroup.v1.AnomalyGroup
	0, // 4: anomalygroup.v1.AnomalyGroup.group_action:type_name -> anomalygroup.v1.GroupActionType
	1, // 5: anomalygroup.v1.AnomalyGroupService.CreateAnomalyGroup:input_type -> anomalygroup.v1.CreateAnomalyGroupRequest
	3, // 6: anomalygroup.v1.AnomalyGroupService.ReadAnomalyGroup:input_type -> anomalygroup.v1.ReadAnomalyGroupRequest
	5, // 7: anomalygroup.v1.AnomalyGroupService.UpdateAnomalyGroup:input_type -> anomalygroup.v1.UpdateAnomalyGroupRequest
	7, // 8: anomalygroup.v1.AnomalyGroupService.FindExistingGroups:input_type -> anomalygroup.v1.FindExistingGroupsRequest
	2, // 9: anomalygroup.v1.AnomalyGroupService.CreateAnomalyGroup:output_type -> anomalygroup.v1.CreateAnomalyGroupResponse
	4, // 10: anomalygroup.v1.AnomalyGroupService.ReadAnomalyGroup:output_type -> anomalygroup.v1.ReadAnomalyGroupResponse
	6, // 11: anomalygroup.v1.AnomalyGroupService.UpdateAnomalyGroup:output_type -> anomalygroup.v1.UpdateAnomalyGroupResponse
	8, // 12: anomalygroup.v1.AnomalyGroupService.FindExistingGroups:output_type -> anomalygroup.v1.FindExistingGroupsResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_anomalygroup_service_proto_init() }
func file_anomalygroup_service_proto_init() {
	if File_anomalygroup_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_anomalygroup_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnomalyGroupRequest); i {
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
		file_anomalygroup_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAnomalyGroupResponse); i {
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
		file_anomalygroup_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAnomalyGroupRequest); i {
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
		file_anomalygroup_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadAnomalyGroupResponse); i {
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
		file_anomalygroup_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnomalyGroupRequest); i {
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
		file_anomalygroup_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnomalyGroupResponse); i {
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
		file_anomalygroup_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindExistingGroupsRequest); i {
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
		file_anomalygroup_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindExistingGroupsResponse); i {
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
		file_anomalygroup_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnomalyGroup); i {
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
			RawDescriptor: file_anomalygroup_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_anomalygroup_service_proto_goTypes,
		DependencyIndexes: file_anomalygroup_service_proto_depIdxs,
		EnumInfos:         file_anomalygroup_service_proto_enumTypes,
		MessageInfos:      file_anomalygroup_service_proto_msgTypes,
	}.Build()
	File_anomalygroup_service_proto = out.File
	file_anomalygroup_service_proto_rawDesc = nil
	file_anomalygroup_service_proto_goTypes = nil
	file_anomalygroup_service_proto_depIdxs = nil
}
