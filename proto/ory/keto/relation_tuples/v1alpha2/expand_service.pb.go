// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: ory/keto/relation_tuples/v1alpha2/expand_service.proto

package rts

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "google.golang.org/genproto/googleapis/api/visibility"
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

type NodeType int32

const (
	NodeType_unspecified           NodeType = 0
	NodeType_NODE_TYPE_UNSPECIFIED NodeType = 0
	// This node expands to a union of all children.
	NodeType_union           NodeType = 1
	NodeType_NODE_TYPE_UNION NodeType = 1
	// Not implemented yet.
	NodeType_exclusion           NodeType = 2
	NodeType_NODE_TYPE_EXCLUSION NodeType = 2
	// Not implemented yet.
	NodeType_intersection           NodeType = 3
	NodeType_NODE_TYPE_INTERSECTION NodeType = 3
	// This node is a leaf and contains no children.
	// Its subject is a `SubjectID` unless `max_depth` was reached.
	NodeType_leaf           NodeType = 4
	NodeType_NODE_TYPE_LEAF NodeType = 4
	// This node is a leaf and contains no children.
	// Its subject is a `SubjectID` unless `max_depth` was reached.
	NodeType_tuple_to_subject_set           NodeType = 5
	NodeType_NODE_TYPE_TUPLE_TO_SUBJECT_SET NodeType = 5
	// This node is a leaf and contains no children.
	// Its subject is a `SubjectID` unless `max_depth` was reached.
	NodeType_computed_subject_set           NodeType = 6
	NodeType_NODE_TYPE_COMPUTED_SUBJECT_SET NodeType = 6
	// This node is a leaf and contains no children.
	// Its subject is a `SubjectID` unless `max_depth` was reached.
	NodeType_not           NodeType = 7
	NodeType_NODE_TYPE_NOT NodeType = 7
)

// Enum value maps for NodeType.
var (
	NodeType_name = map[int32]string{
		0: "unspecified",
		// Duplicate value: 0: "NODE_TYPE_UNSPECIFIED",
		1: "union",
		// Duplicate value: 1: "NODE_TYPE_UNION",
		2: "exclusion",
		// Duplicate value: 2: "NODE_TYPE_EXCLUSION",
		3: "intersection",
		// Duplicate value: 3: "NODE_TYPE_INTERSECTION",
		4: "leaf",
		// Duplicate value: 4: "NODE_TYPE_LEAF",
		5: "tuple_to_subject_set",
		// Duplicate value: 5: "NODE_TYPE_TUPLE_TO_SUBJECT_SET",
		6: "computed_subject_set",
		// Duplicate value: 6: "NODE_TYPE_COMPUTED_SUBJECT_SET",
		7: "not",
		// Duplicate value: 7: "NODE_TYPE_NOT",
	}
	NodeType_value = map[string]int32{
		"unspecified":                    0,
		"NODE_TYPE_UNSPECIFIED":          0,
		"union":                          1,
		"NODE_TYPE_UNION":                1,
		"exclusion":                      2,
		"NODE_TYPE_EXCLUSION":            2,
		"intersection":                   3,
		"NODE_TYPE_INTERSECTION":         3,
		"leaf":                           4,
		"NODE_TYPE_LEAF":                 4,
		"tuple_to_subject_set":           5,
		"NODE_TYPE_TUPLE_TO_SUBJECT_SET": 5,
		"computed_subject_set":           6,
		"NODE_TYPE_COMPUTED_SUBJECT_SET": 6,
		"not":                            7,
		"NODE_TYPE_NOT":                  7,
	}
)

func (x NodeType) Enum() *NodeType {
	p := new(NodeType)
	*p = x
	return p
}

func (x NodeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NodeType) Descriptor() protoreflect.EnumDescriptor {
	return file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_enumTypes[0].Descriptor()
}

func (NodeType) Type() protoreflect.EnumType {
	return &file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_enumTypes[0]
}

func (x NodeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NodeType.Descriptor instead.
func (NodeType) EnumDescriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescGZIP(), []int{0}
}

// The request for an ExpandService.Expand RPC.
// Expands the given subject set.
type ExpandRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The subject to expand.
	Subject *Subject `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	// The maximum depth of tree to build.
	//
	// If the value is less than 1 or greater than the global
	// max-depth then the global max-depth will be used instead.
	//
	// It is important to set this parameter to a meaningful
	// value. Ponder how deep you really want to display this.
	MaxDepth int32 `protobuf:"varint,2,opt,name=max_depth,json=max-depth,proto3" json:"max_depth,omitempty"`
	// This field is not implemented yet and has no effect.
	// <!--
	// Optional. Like reads, a expand is always evaluated at a
	// consistent snapshot no earlier than the given snaptoken.
	//
	// Leave this field blank if you want to expand
	// based on eventually consistent ACLs, benefiting from very
	// low latency, but possibly slightly stale results.
	//
	// If the specified token is too old and no longer known,
	// the server falls back as if no snaptoken had been specified.
	//
	// If not specified the server tries to build the tree
	// on the best snapshot version where it is very likely that
	// ACLs had already been replicated to all availability zones.
	// -->
	Snaptoken string `protobuf:"bytes,3,opt,name=snaptoken,proto3" json:"snaptoken,omitempty"`
	// The namespace of the object and relation
	// referenced in this subject set.
	//
	// Deprecated: Do not use.
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The object related by this subject set.
	//
	// Deprecated: Do not use.
	Object string `protobuf:"bytes,5,opt,name=object,proto3" json:"object,omitempty"`
	// The relation between the object and the subjects.
	//
	// Deprecated: Do not use.
	Relation string `protobuf:"bytes,6,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *ExpandRequest) Reset() {
	*x = ExpandRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandRequest) ProtoMessage() {}

func (x *ExpandRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandRequest.ProtoReflect.Descriptor instead.
func (*ExpandRequest) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescGZIP(), []int{0}
}

func (x *ExpandRequest) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *ExpandRequest) GetMaxDepth() int32 {
	if x != nil {
		return x.MaxDepth
	}
	return 0
}

func (x *ExpandRequest) GetSnaptoken() string {
	if x != nil {
		return x.Snaptoken
	}
	return ""
}

// Deprecated: Do not use.
func (x *ExpandRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// Deprecated: Do not use.
func (x *ExpandRequest) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

// Deprecated: Do not use.
func (x *ExpandRequest) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

// The response for a ExpandService.Expand RPC.
type ExpandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tree the requested subject set expands to.
	// The requested subject set is the subject of the root.
	//
	// This field can be nil in some circumstances.
	Tree *SubjectTree `protobuf:"bytes,1,opt,name=tree,proto3" json:"tree,omitempty"`
}

func (x *ExpandResponse) Reset() {
	*x = ExpandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpandResponse) ProtoMessage() {}

func (x *ExpandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpandResponse.ProtoReflect.Descriptor instead.
func (*ExpandResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExpandResponse) GetTree() *SubjectTree {
	if x != nil {
		return x.Tree
	}
	return nil
}

type SubjectTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of the node.
	NodeType NodeType `protobuf:"varint,1,opt,name=node_type,json=type,proto3,enum=ory.keto.relation_tuples.v1alpha2.NodeType" json:"node_type,omitempty"`
	// The subject this node represents.
	// Deprecated: More information is now available in the tuple field.
	//
	// Deprecated: Do not use.
	Subject *Subject `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// The relation tuple this node represents.
	Tuple *RelationTuple `protobuf:"bytes,4,opt,name=tuple,proto3" json:"tuple,omitempty"`
	// The children of this node.
	//
	// This is never set if `node_type` == `NODE_TYPE_LEAF`.
	Children []*SubjectTree `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *SubjectTree) Reset() {
	*x = SubjectTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubjectTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubjectTree) ProtoMessage() {}

func (x *SubjectTree) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubjectTree.ProtoReflect.Descriptor instead.
func (*SubjectTree) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescGZIP(), []int{2}
}

func (x *SubjectTree) GetNodeType() NodeType {
	if x != nil {
		return x.NodeType
	}
	return NodeType_unspecified
}

// Deprecated: Do not use.
func (x *SubjectTree) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *SubjectTree) GetTuple() *RelationTuple {
	if x != nil {
		return x.Tuple
	}
	return nil
}

func (x *SubjectTree) GetChildren() []*SubjectTree {
	if x != nil {
		return x.Children
	}
	return nil
}

var File_ory_keto_relation_tuples_v1alpha2_expand_service_proto protoreflect.FileDescriptor

var file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDesc = []byte{
	0x0a, 0x36, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x2f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65,
	0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74,
	0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65,
	0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa0, 0x02, 0x0a, 0x0d, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x58, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42,
	0x12, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47,
	0x47, 0x45, 0x52, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6d, 0x61, 0x78, 0x2d, 0x64, 0x65, 0x70, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x09, 0x73, 0x6e,
	0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0xfa,
	0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47, 0x47, 0x45,
	0x52, 0x52, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x05, 0x18, 0x01, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x12, 0x21, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x05, 0x18, 0x01, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0e, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54,
	0x72, 0x65, 0x65, 0x52, 0x04, 0x74, 0x72, 0x65, 0x65, 0x22, 0xb6, 0x02, 0x0a, 0x0b, 0x53, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6f,
	0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x46,
	0x0a, 0x05, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x52,
	0x05, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b,
	0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x65, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72,
	0x65, 0x6e, 0x2a, 0x86, 0x04, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x2d, 0x0a, 0x15, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x12, 0xfa, 0xd2, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47, 0x47, 0x45, 0x52, 0x12,
	0x09, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x0f, 0x4e, 0x4f,
	0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x1a,
	0x12, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47,
	0x47, 0x45, 0x52, 0x12, 0x0d, 0x0a, 0x09, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x6f, 0x6e,
	0x10, 0x02, 0x12, 0x2b, 0x0a, 0x13, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x45, 0x58, 0x43, 0x4c, 0x55, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x1a, 0x12, 0xfa, 0xd2, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47, 0x47, 0x45, 0x52, 0x12,
	0x10, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x03, 0x12, 0x2e, 0x0a, 0x16, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x4e, 0x54, 0x45, 0x52, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x1a, 0x12, 0xfa,
	0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47, 0x47, 0x45,
	0x52, 0x12, 0x08, 0x0a, 0x04, 0x6c, 0x65, 0x61, 0x66, 0x10, 0x04, 0x12, 0x26, 0x0a, 0x0e, 0x4e,
	0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x45, 0x41, 0x46, 0x10, 0x04, 0x1a,
	0x12, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57, 0x41, 0x47,
	0x47, 0x45, 0x52, 0x12, 0x18, 0x0a, 0x14, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x5f,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x10, 0x05, 0x12, 0x36, 0x0a,
	0x1e, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x55, 0x50, 0x4c, 0x45,
	0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x45, 0x54, 0x10,
	0x05, 0x1a, 0x12, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f, 0x53, 0x57,
	0x41, 0x47, 0x47, 0x45, 0x52, 0x12, 0x18, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x64, 0x5f, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x10, 0x06, 0x12,
	0x36, 0x0a, 0x1e, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d,
	0x50, 0x55, 0x54, 0x45, 0x44, 0x5f, 0x53, 0x55, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x45,
	0x54, 0x10, 0x06, 0x1a, 0x12, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f,
	0x53, 0x57, 0x41, 0x47, 0x47, 0x45, 0x52, 0x12, 0x07, 0x0a, 0x03, 0x6e, 0x6f, 0x74, 0x10, 0x07,
	0x12, 0x25, 0x0a, 0x0d, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f,
	0x54, 0x10, 0x07, 0x1a, 0x12, 0xfa, 0xd2, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x4e, 0x4f, 0x5f,
	0x53, 0x57, 0x41, 0x47, 0x47, 0x45, 0x52, 0x1a, 0x02, 0x10, 0x01, 0x32, 0xa7, 0x03, 0x0a, 0x0d,
	0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x03,
	0x0a, 0x06, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x30, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b,
	0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x45, 0x78, 0x70,
	0x61, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x72, 0x79,
	0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x45,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa5, 0x02,
	0x92, 0x41, 0xfc, 0x01, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2a, 0x11, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x21, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x78, 0x2d, 0x77, 0x77, 0x77, 0x2d, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x75, 0x72, 0x6c, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4a, 0xb7, 0x01, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0xaf,
	0x01, 0x0a, 0x79, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x20,
	0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2c, 0x20, 0x69, 0x6e, 0x20, 0x63, 0x61, 0x73, 0x65,
	0x20, 0x69, 0x74, 0x20, 0x69, 0x73, 0x20, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x2e, 0x20,
	0x46, 0x6f, 0x72, 0x20, 0x74, 0x68, 0x65, 0x20, 0x60, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x60, 0x20, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2c, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x20, 0x69, 0x73,
	0x20, 0x61, 0x6c, 0x77, 0x61, 0x79, 0x73, 0x20, 0x32, 0x30, 0x30, 0x2e, 0x12, 0x32, 0x0a, 0x30,
	0x1a, 0x2e, 0x2e, 0x6f, 0x72, 0x79, 0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x72, 0x65, 0x65,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x62, 0x04, 0x74, 0x72, 0x65, 0x65, 0x12, 0x17, 0x2f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65,
	0x78, 0x70, 0x61, 0x6e, 0x64, 0x42, 0xc3, 0x01, 0x0a, 0x24, 0x73, 0x68, 0x2e, 0x6f, 0x72, 0x79,
	0x2e, 0x6b, 0x65, 0x74, 0x6f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x12,
	0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6f, 0x72, 0x79, 0x2f, 0x6b, 0x65, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x3b, 0x72, 0x74, 0x73, 0xaa, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x2e, 0x4b, 0x65, 0x74, 0x6f,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x20, 0x4f, 0x72, 0x79, 0x5c, 0x4b,
	0x65, 0x74, 0x6f, 0x5c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x75, 0x70, 0x6c,
	0x65, 0x73, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescOnce sync.Once
	file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescData = file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDesc
)

func file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescGZIP() []byte {
	file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescOnce.Do(func() {
		file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescData)
	})
	return file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDescData
}

var file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_goTypes = []interface{}{
	(NodeType)(0),          // 0: ory.keto.relation_tuples.v1alpha2.NodeType
	(*ExpandRequest)(nil),  // 1: ory.keto.relation_tuples.v1alpha2.ExpandRequest
	(*ExpandResponse)(nil), // 2: ory.keto.relation_tuples.v1alpha2.ExpandResponse
	(*SubjectTree)(nil),    // 3: ory.keto.relation_tuples.v1alpha2.SubjectTree
	(*Subject)(nil),        // 4: ory.keto.relation_tuples.v1alpha2.Subject
	(*RelationTuple)(nil),  // 5: ory.keto.relation_tuples.v1alpha2.RelationTuple
}
var file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_depIdxs = []int32{
	4, // 0: ory.keto.relation_tuples.v1alpha2.ExpandRequest.subject:type_name -> ory.keto.relation_tuples.v1alpha2.Subject
	3, // 1: ory.keto.relation_tuples.v1alpha2.ExpandResponse.tree:type_name -> ory.keto.relation_tuples.v1alpha2.SubjectTree
	0, // 2: ory.keto.relation_tuples.v1alpha2.SubjectTree.node_type:type_name -> ory.keto.relation_tuples.v1alpha2.NodeType
	4, // 3: ory.keto.relation_tuples.v1alpha2.SubjectTree.subject:type_name -> ory.keto.relation_tuples.v1alpha2.Subject
	5, // 4: ory.keto.relation_tuples.v1alpha2.SubjectTree.tuple:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTuple
	3, // 5: ory.keto.relation_tuples.v1alpha2.SubjectTree.children:type_name -> ory.keto.relation_tuples.v1alpha2.SubjectTree
	1, // 6: ory.keto.relation_tuples.v1alpha2.ExpandService.Expand:input_type -> ory.keto.relation_tuples.v1alpha2.ExpandRequest
	2, // 7: ory.keto.relation_tuples.v1alpha2.ExpandService.Expand:output_type -> ory.keto.relation_tuples.v1alpha2.ExpandResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_init() }
func file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_init() {
	if File_ory_keto_relation_tuples_v1alpha2_expand_service_proto != nil {
		return
	}
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandRequest); i {
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
		file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpandResponse); i {
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
		file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubjectTree); i {
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
			RawDescriptor: file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_goTypes,
		DependencyIndexes: file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_depIdxs,
		EnumInfos:         file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_enumTypes,
		MessageInfos:      file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_msgTypes,
	}.Build()
	File_ory_keto_relation_tuples_v1alpha2_expand_service_proto = out.File
	file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_rawDesc = nil
	file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_goTypes = nil
	file_ory_keto_relation_tuples_v1alpha2_expand_service_proto_depIdxs = nil
}
