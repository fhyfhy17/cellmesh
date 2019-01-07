// Generated by github.com/davyxu/protoplus
// DO NOT EDIT!
package proto

import (
	"github.com/davyxu/protoplus/proto"
	"unsafe"
)

var (
	_ *proto.Buffer
	_ unsafe.Pointer
)

type ResultCode int32

const (
	ResultCode_Result_OK          ResultCode = 0
	ResultCode_Result_NotExists   ResultCode = 1
	ResultCode_Result_AuthRequire ResultCode = 2
)

var (
	ResultCodeMapperValueByName = map[string]int32{
		"Result_OK":          0,
		"Result_NotExists":   1,
		"Result_AuthRequire": 2,
	}

	ResultCodeMapperNameByValue = map[int32]string{
		0: "Result_OK",
		1: "Result_NotExists",
		2: "Result_AuthRequire",
	}
)

func (self ResultCode) String() string {
	return ResultCodeMapperNameByValue[int32(self)]
}

type SetValueREQ struct {
	Key     string
	Value   []byte
	SvcName string
}

func (self *SetValueREQ) String() string { return proto.CompactTextString(self) }

func (self *SetValueREQ) Size() (ret int) {

	ret += proto.SizeString(0, self.Key)

	ret += proto.SizeBytes(1, self.Value)

	ret += proto.SizeString(2, self.SvcName)

	return
}

func (self *SetValueREQ) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Key)

	proto.MarshalBytes(buffer, 1, self.Value)

	proto.MarshalString(buffer, 2, self.SvcName)

	return nil
}

func (self *SetValueREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Key)
	case 1:
		return proto.UnmarshalBytes(buffer, wt, &self.Value)
	case 2:
		return proto.UnmarshalString(buffer, wt, &self.SvcName)

	}

	return proto.ErrUnknownField
}

type SetValueACK struct {
	Code ResultCode
}

func (self *SetValueACK) String() string { return proto.CompactTextString(self) }

func (self *SetValueACK) Size() (ret int) {

	ret += proto.SizeInt32(0, int32(self.Code))

	return
}

func (self *SetValueACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt32(buffer, 0, int32(self.Code))

	return nil
}

func (self *SetValueACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&self.Code))

	}

	return proto.ErrUnknownField
}

type GetValueREQ struct {
	Key string
}

func (self *GetValueREQ) String() string { return proto.CompactTextString(self) }

func (self *GetValueREQ) Size() (ret int) {

	ret += proto.SizeString(0, self.Key)

	return
}

func (self *GetValueREQ) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Key)

	return nil
}

func (self *GetValueREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Key)

	}

	return proto.ErrUnknownField
}

type GetValueACK struct {
	Code  ResultCode
	Key   string
	Value []byte
}

func (self *GetValueACK) String() string { return proto.CompactTextString(self) }

func (self *GetValueACK) Size() (ret int) {

	ret += proto.SizeInt32(0, int32(self.Code))

	ret += proto.SizeString(1, self.Key)

	ret += proto.SizeBytes(2, self.Value)

	return
}

func (self *GetValueACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt32(buffer, 0, int32(self.Code))

	proto.MarshalString(buffer, 1, self.Key)

	proto.MarshalBytes(buffer, 2, self.Value)

	return nil
}

func (self *GetValueACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&self.Code))
	case 1:
		return proto.UnmarshalString(buffer, wt, &self.Key)
	case 2:
		return proto.UnmarshalBytes(buffer, wt, &self.Value)

	}

	return proto.ErrUnknownField
}

type DeleteValueREQ struct {
	Key string
}

func (self *DeleteValueREQ) String() string { return proto.CompactTextString(self) }

func (self *DeleteValueREQ) Size() (ret int) {

	ret += proto.SizeString(0, self.Key)

	return
}

func (self *DeleteValueREQ) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Key)

	return nil
}

func (self *DeleteValueREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Key)

	}

	return proto.ErrUnknownField
}

type DeleteValueACK struct {
	Code ResultCode
	Key  string
}

func (self *DeleteValueACK) String() string { return proto.CompactTextString(self) }

func (self *DeleteValueACK) Size() (ret int) {

	ret += proto.SizeInt32(0, int32(self.Code))

	ret += proto.SizeString(1, self.Key)

	return
}

func (self *DeleteValueACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt32(buffer, 0, int32(self.Code))

	proto.MarshalString(buffer, 1, self.Key)

	return nil
}

func (self *DeleteValueACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&self.Code))
	case 1:
		return proto.UnmarshalString(buffer, wt, &self.Key)

	}

	return proto.ErrUnknownField
}

type ValueChangeNotifyACK struct {
	Key     string
	Value   []byte
	SvcName string
}

func (self *ValueChangeNotifyACK) String() string { return proto.CompactTextString(self) }

func (self *ValueChangeNotifyACK) Size() (ret int) {

	ret += proto.SizeString(0, self.Key)

	ret += proto.SizeBytes(1, self.Value)

	ret += proto.SizeString(2, self.SvcName)

	return
}

func (self *ValueChangeNotifyACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Key)

	proto.MarshalBytes(buffer, 1, self.Value)

	proto.MarshalString(buffer, 2, self.SvcName)

	return nil
}

func (self *ValueChangeNotifyACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Key)
	case 1:
		return proto.UnmarshalBytes(buffer, wt, &self.Value)
	case 2:
		return proto.UnmarshalString(buffer, wt, &self.SvcName)

	}

	return proto.ErrUnknownField
}

type ValueDeleteNotifyACK struct {
	Key     string
	SvcName string
}

func (self *ValueDeleteNotifyACK) String() string { return proto.CompactTextString(self) }

func (self *ValueDeleteNotifyACK) Size() (ret int) {

	ret += proto.SizeString(0, self.Key)

	ret += proto.SizeString(1, self.SvcName)

	return
}

func (self *ValueDeleteNotifyACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Key)

	proto.MarshalString(buffer, 1, self.SvcName)

	return nil
}

func (self *ValueDeleteNotifyACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Key)
	case 1:
		return proto.UnmarshalString(buffer, wt, &self.SvcName)

	}

	return proto.ErrUnknownField
}

type AuthREQ struct {
	Token string
}

func (self *AuthREQ) String() string { return proto.CompactTextString(self) }

func (self *AuthREQ) Size() (ret int) {

	ret += proto.SizeString(0, self.Token)

	return
}

func (self *AuthREQ) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Token)

	return nil
}

func (self *AuthREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Token)

	}

	return proto.ErrUnknownField
}

type AuthACK struct {
	Token string
}

func (self *AuthACK) String() string { return proto.CompactTextString(self) }

func (self *AuthACK) Size() (ret int) {

	ret += proto.SizeString(0, self.Token)

	return
}

func (self *AuthACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalString(buffer, 0, self.Token)

	return nil
}

func (self *AuthACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalString(buffer, wt, &self.Token)

	}

	return proto.ErrUnknownField
}

type ClearSvcREQ struct {
}

func (self *ClearSvcREQ) String() string { return proto.CompactTextString(self) }

func (self *ClearSvcREQ) Size() (ret int) {

	return
}

func (self *ClearSvcREQ) Marshal(buffer *proto.Buffer) error {

	return nil
}

func (self *ClearSvcREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {

	}

	return proto.ErrUnknownField
}

type ClearSvcACK struct {
	Code ResultCode
}

func (self *ClearSvcACK) String() string { return proto.CompactTextString(self) }

func (self *ClearSvcACK) Size() (ret int) {

	ret += proto.SizeInt32(0, int32(self.Code))

	return
}

func (self *ClearSvcACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt32(buffer, 0, int32(self.Code))

	return nil
}

func (self *ClearSvcACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&self.Code))

	}

	return proto.ErrUnknownField
}

type ClearKeyREQ struct {
}

func (self *ClearKeyREQ) String() string { return proto.CompactTextString(self) }

func (self *ClearKeyREQ) Size() (ret int) {

	return
}

func (self *ClearKeyREQ) Marshal(buffer *proto.Buffer) error {

	return nil
}

func (self *ClearKeyREQ) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {

	}

	return proto.ErrUnknownField
}

type ClearKeyACK struct {
	Code ResultCode
}

func (self *ClearKeyACK) String() string { return proto.CompactTextString(self) }

func (self *ClearKeyACK) Size() (ret int) {

	ret += proto.SizeInt32(0, int32(self.Code))

	return
}

func (self *ClearKeyACK) Marshal(buffer *proto.Buffer) error {

	proto.MarshalInt32(buffer, 0, int32(self.Code))

	return nil
}

func (self *ClearKeyACK) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&self.Code))

	}

	return proto.ErrUnknownField
}
