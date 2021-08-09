package table

import "github.com/JayLiang/dive-into-leveldb/leveldb"

type BlockContents struct {
	Data      leveldb.Slice
	CacheAble bool
}

const KMaxEncodedLength int = 20

type BlockHandle struct {
	offset uint64
	size   uint64
}

func (b *BlockHandle) SetOffset(offset uint64) {
	b.offset = offset
}
func (b *BlockHandle) GetOffset() uint64 {
	return b.offset
}
func (b *BlockHandle) SetSize(size uint64) {
	b.size = size
}
func (b *BlockHandle) GetSize() uint64 {
	return b.size
}
