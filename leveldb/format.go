package leveldb

import "github.com/JayLiang/dive-into-leveldb/leveldb/utils"

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

func NewParsedInternalKey(slice Slice, seq uint64, typ int) *ParsedInternalKey {
	return &ParsedInternalKey{
		UserKey: slice,
		SeqNum:  seq,
		ValType: byte(typ),
	}
}

type InternalKey struct {
	data []byte
}

func AppendInternalKey(data []byte, key *ParsedInternalKey) []byte {
	data = append(data, key.UserKey.Data...)
	buf := make([]byte, 8)
	utils.EncodeFixed64(PackSeqNumAndType(key.SeqNum, key.ValType), buf)
	data = append(data, buf[:]...)
	return data
}
func NewInternalKey(slice Slice, seq uint64, typ int) *InternalKey {

}

func (k *InternalKey) Clear() {
	k.data = make([]byte, 0)
}

func (k *InternalKey) DecodeFrom(slice Slice) bool {

}

func (k *InternalKey) Encode() *Slice {

}

func (k *InternalKey) UserKey() *Slice {

}

func (k *InternalKey) String() string {

}
