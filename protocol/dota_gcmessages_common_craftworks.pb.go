// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: dota_gcmessages_common_craftworks.proto

package protocol

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

type ECraftworksAuditAction int32

const (
	ECraftworksAuditAction_k_eInvalid           ECraftworksAuditAction = 0
	ECraftworksAuditAction_k_eRecipeCrafted     ECraftworksAuditAction = 1
	ECraftworksAuditAction_k_eMatchRewards      ECraftworksAuditAction = 2
	ECraftworksAuditAction_k_eMatchRewardsTurbo ECraftworksAuditAction = 3
)

// Enum value maps for ECraftworksAuditAction.
var (
	ECraftworksAuditAction_name = map[int32]string{
		0: "k_eInvalid",
		1: "k_eRecipeCrafted",
		2: "k_eMatchRewards",
		3: "k_eMatchRewardsTurbo",
	}
	ECraftworksAuditAction_value = map[string]int32{
		"k_eInvalid":           0,
		"k_eRecipeCrafted":     1,
		"k_eMatchRewards":      2,
		"k_eMatchRewardsTurbo": 3,
	}
)

func (x ECraftworksAuditAction) Enum() *ECraftworksAuditAction {
	p := new(ECraftworksAuditAction)
	*p = x
	return p
}

func (x ECraftworksAuditAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ECraftworksAuditAction) Descriptor() protoreflect.EnumDescriptor {
	return file_dota_gcmessages_common_craftworks_proto_enumTypes[0].Descriptor()
}

func (ECraftworksAuditAction) Type() protoreflect.EnumType {
	return &file_dota_gcmessages_common_craftworks_proto_enumTypes[0]
}

func (x ECraftworksAuditAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *ECraftworksAuditAction) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = ECraftworksAuditAction(num)
	return nil
}

// Deprecated: Use ECraftworksAuditAction.Descriptor instead.
func (ECraftworksAuditAction) EnumDescriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_craftworks_proto_rawDescGZIP(), []int{0}
}

type CMsgCraftworksComponents struct {
	state               protoimpl.MessageState                               `protogen:"open.v1"`
	ComponentQuantities []*CMsgCraftworksComponents_ComponentQuantitiesEntry `protobuf:"bytes,1,rep,name=component_quantities,json=componentQuantities" json:"component_quantities,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *CMsgCraftworksComponents) Reset() {
	*x = CMsgCraftworksComponents{}
	mi := &file_dota_gcmessages_common_craftworks_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CMsgCraftworksComponents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgCraftworksComponents) ProtoMessage() {}

func (x *CMsgCraftworksComponents) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_craftworks_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgCraftworksComponents.ProtoReflect.Descriptor instead.
func (*CMsgCraftworksComponents) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_craftworks_proto_rawDescGZIP(), []int{0}
}

func (x *CMsgCraftworksComponents) GetComponentQuantities() []*CMsgCraftworksComponents_ComponentQuantitiesEntry {
	if x != nil {
		return x.ComponentQuantities
	}
	return nil
}

type CMsgCraftworksQuestReward struct {
	state            protoimpl.MessageState    `protogen:"open.v1"`
	QuestId          *uint32                   `protobuf:"varint,1,opt,name=quest_id,json=questId" json:"quest_id,omitempty"`
	RewardComponents *CMsgCraftworksComponents `protobuf:"bytes,2,opt,name=reward_components,json=rewardComponents" json:"reward_components,omitempty"`
	StatValue        *uint32                   `protobuf:"varint,3,opt,name=stat_value,json=statValue" json:"stat_value,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *CMsgCraftworksQuestReward) Reset() {
	*x = CMsgCraftworksQuestReward{}
	mi := &file_dota_gcmessages_common_craftworks_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CMsgCraftworksQuestReward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgCraftworksQuestReward) ProtoMessage() {}

func (x *CMsgCraftworksQuestReward) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_craftworks_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgCraftworksQuestReward.ProtoReflect.Descriptor instead.
func (*CMsgCraftworksQuestReward) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_craftworks_proto_rawDescGZIP(), []int{1}
}

func (x *CMsgCraftworksQuestReward) GetQuestId() uint32 {
	if x != nil && x.QuestId != nil {
		return *x.QuestId
	}
	return 0
}

func (x *CMsgCraftworksQuestReward) GetRewardComponents() *CMsgCraftworksComponents {
	if x != nil {
		return x.RewardComponents
	}
	return nil
}

func (x *CMsgCraftworksQuestReward) GetStatValue() uint32 {
	if x != nil && x.StatValue != nil {
		return *x.StatValue
	}
	return 0
}

type CMsgCraftworksComponents_ComponentQuantitiesEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           *uint32                `protobuf:"varint,1,opt,name=key" json:"key,omitempty"`
	Value         *uint32                `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CMsgCraftworksComponents_ComponentQuantitiesEntry) Reset() {
	*x = CMsgCraftworksComponents_ComponentQuantitiesEntry{}
	mi := &file_dota_gcmessages_common_craftworks_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CMsgCraftworksComponents_ComponentQuantitiesEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CMsgCraftworksComponents_ComponentQuantitiesEntry) ProtoMessage() {}

func (x *CMsgCraftworksComponents_ComponentQuantitiesEntry) ProtoReflect() protoreflect.Message {
	mi := &file_dota_gcmessages_common_craftworks_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CMsgCraftworksComponents_ComponentQuantitiesEntry.ProtoReflect.Descriptor instead.
func (*CMsgCraftworksComponents_ComponentQuantitiesEntry) Descriptor() ([]byte, []int) {
	return file_dota_gcmessages_common_craftworks_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CMsgCraftworksComponents_ComponentQuantitiesEntry) GetKey() uint32 {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return 0
}

func (x *CMsgCraftworksComponents_ComponentQuantitiesEntry) GetValue() uint32 {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return 0
}

var File_dota_gcmessages_common_craftworks_proto protoreflect.FileDescriptor

var file_dota_gcmessages_common_craftworks_proto_rawDesc = string([]byte{
	0x0a, 0x27, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x67, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x63, 0x72, 0x61, 0x66, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x1a, 0x13, 0x73, 0x74, 0x65, 0x61, 0x6d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x64, 0x6f, 0x74, 0x61, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x67, 0x63, 0x73, 0x64, 0x6b, 0x5f, 0x67, 0x63, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x18, 0x43, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x61, 0x66, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x43, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x6e, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x43, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x61, 0x66, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x43, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65,
	0x6e, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x42, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x19, 0x43,
	0x4d, 0x73, 0x67, 0x43, 0x72, 0x61, 0x66, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x51, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x11, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x43, 0x4d, 0x73, 0x67, 0x43, 0x72,
	0x61, 0x66, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x10, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x2a, 0x6d, 0x0a, 0x16, 0x45, 0x43, 0x72, 0x61, 0x66, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x73, 0x41, 0x75, 0x64, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x0a, 0x6b, 0x5f, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x6b, 0x5f, 0x65, 0x52, 0x65, 0x63, 0x69, 0x70, 0x65, 0x43, 0x72, 0x61, 0x66, 0x74, 0x65,
	0x64, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x6b, 0x5f, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x6b, 0x5f, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x54, 0x75, 0x72, 0x62, 0x6f,
	0x10, 0x03, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
})

var (
	file_dota_gcmessages_common_craftworks_proto_rawDescOnce sync.Once
	file_dota_gcmessages_common_craftworks_proto_rawDescData []byte
)

func file_dota_gcmessages_common_craftworks_proto_rawDescGZIP() []byte {
	file_dota_gcmessages_common_craftworks_proto_rawDescOnce.Do(func() {
		file_dota_gcmessages_common_craftworks_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dota_gcmessages_common_craftworks_proto_rawDesc), len(file_dota_gcmessages_common_craftworks_proto_rawDesc)))
	})
	return file_dota_gcmessages_common_craftworks_proto_rawDescData
}

var file_dota_gcmessages_common_craftworks_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dota_gcmessages_common_craftworks_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_dota_gcmessages_common_craftworks_proto_goTypes = []any{
	(ECraftworksAuditAction)(0),                               // 0: protocol.ECraftworksAuditAction
	(*CMsgCraftworksComponents)(nil),                          // 1: protocol.CMsgCraftworksComponents
	(*CMsgCraftworksQuestReward)(nil),                         // 2: protocol.CMsgCraftworksQuestReward
	(*CMsgCraftworksComponents_ComponentQuantitiesEntry)(nil), // 3: protocol.CMsgCraftworksComponents.ComponentQuantitiesEntry
}
var file_dota_gcmessages_common_craftworks_proto_depIdxs = []int32{
	3, // 0: protocol.CMsgCraftworksComponents.component_quantities:type_name -> protocol.CMsgCraftworksComponents.ComponentQuantitiesEntry
	1, // 1: protocol.CMsgCraftworksQuestReward.reward_components:type_name -> protocol.CMsgCraftworksComponents
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dota_gcmessages_common_craftworks_proto_init() }
func file_dota_gcmessages_common_craftworks_proto_init() {
	if File_dota_gcmessages_common_craftworks_proto != nil {
		return
	}
	file_steammessages_proto_init()
	file_dota_shared_enums_proto_init()
	file_gcsdk_gcmessages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dota_gcmessages_common_craftworks_proto_rawDesc), len(file_dota_gcmessages_common_craftworks_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dota_gcmessages_common_craftworks_proto_goTypes,
		DependencyIndexes: file_dota_gcmessages_common_craftworks_proto_depIdxs,
		EnumInfos:         file_dota_gcmessages_common_craftworks_proto_enumTypes,
		MessageInfos:      file_dota_gcmessages_common_craftworks_proto_msgTypes,
	}.Build()
	File_dota_gcmessages_common_craftworks_proto = out.File
	file_dota_gcmessages_common_craftworks_proto_goTypes = nil
	file_dota_gcmessages_common_craftworks_proto_depIdxs = nil
}
