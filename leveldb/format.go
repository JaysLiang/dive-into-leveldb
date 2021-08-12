package leveldb

import (
	"fmt"
	"github.com/JayLiang/dive-into-leveldb/leveldb/utils"
)

const (
	TypeKey      = 1
	TypeDeletion = 0
)

type ParsedInternalKey struct {
	UserKey Slice
	SeqNum  uint64
	ValType byte
}

func PackSeqNumAndType(seqNum uint64, typ byte) uint64 {
	return seqNum<<8 | uint64(typ)
}

func ParseInternKey(data []byte) (*ParsedInternalKey, bool) {
	if len(data) < 8 {
		return nil, false
	}
	fixed64 := utils.DecodeFixed64(data[len(data)-8:])
	return &ParsedInternalKey{
		UserKey: Slice{
			Data: data[:len(data)-8],
		},
		SeqNum:  fixed64 >> 8,
		ValType: byte(fixed64),
	}, true
}

func NewParsedInternalKey(slice Slice, seq uint64, typ int) *ParsedInternalKey {
	return &ParsedInternalKey{
		UserKey: slice,
		SeqNum:  seq,
		ValType: byte(typ),
	}
}

func (p *ParsedInternalKey) String() string {
	return fmt.Sprintf("key: %s, seq: %v type: %v\n", string(p.UserKey.Data), p.SeqNum, p.ValType)
}

type InternalKey struct {
	data []byte
}

func AppendInternalKey(data []byte, key *ParsedInternalKey) []byte {
	data = append(data, key.UserKey.Data...)
	buf := make([]byte, 8)
	utils.EncodeFixed64(PackSeqNumAndType(key.SeqNum, key.ValType), buf)
	data = append(data, buf...)
	return data
}

func ExtractUserKey(data []byte) *Slice {
	return &Slice{Data: data[:len(data)-8]}
}

func NewEmptyInternalKey() *InternalKey {
	return &InternalKey{}
}
func NewInternalKey(slice Slice, seq uint64, typ int) *InternalKey {
	return &InternalKey{
		data: AppendInternalKey(nil, NewParsedInternalKey(slice, seq, typ)),
	}
}

func (k *InternalKey) Clear() {
	k.data = make([]byte, 0)
}

func (k *InternalKey) DecodeFrom(slice Slice) bool {
	k.data = slice.Data
	return len(k.data) != 0
}

func (k *InternalKey) Encode() *Slice {
	return &Slice{k.data}
}

func (k *InternalKey) UserKey() *Slice {
	return ExtractUserKey(k.data)
}

func (k *InternalKey) String() string {
	parsedInternKey, ok := ParseInternKey(k.data)
	if !ok {
		return "bad"
	}
	return parsedInternKey.String()
}

func (k *InternalKey) Valid() bool {
	return len(k.data) != 0
}

type LookupKey struct {
	Start    int
	KeyStart int
	End      int
	data     []byte
}

func NewLookupKey() &L
