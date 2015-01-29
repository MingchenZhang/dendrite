// Code generated by protoc-gen-go.
// source: dtable.proto
// DO NOT EDIT!

/*
Package dtable is a generated protocol buffer package.

It is generated from these files:
	dtable.proto

It has these top-level messages:
	PBDTableGet
	PBDTableGetResp
*/
package dtable

import proto "code.google.com/p/goprotobuf/proto"
import math "math"
import dendrite "github.com/fastfn/dendrite"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type PBDTableGet struct {
	Dest             *dendrite.PBProtoVnode `protobuf:"bytes,1,req,name=dest" json:"dest,omitempty"`
	Key              []byte                 `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *PBDTableGet) Reset()         { *m = PBDTableGet{} }
func (m *PBDTableGet) String() string { return proto.CompactTextString(m) }
func (*PBDTableGet) ProtoMessage()    {}

func (m *PBDTableGet) GetDest() *dendrite.PBProtoVnode {
	if m != nil {
		return m.Dest
	}
	return nil
}

func (m *PBDTableGet) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type PBDTableGetResp struct {
	Found            *bool  `protobuf:"varint,1,req,name=found" json:"found,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PBDTableGetResp) Reset()         { *m = PBDTableGetResp{} }
func (m *PBDTableGetResp) String() string { return proto.CompactTextString(m) }
func (*PBDTableGetResp) ProtoMessage()    {}

func (m *PBDTableGetResp) GetFound() bool {
	if m != nil && m.Found != nil {
		return *m.Found
	}
	return false
}

func (m *PBDTableGetResp) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
}