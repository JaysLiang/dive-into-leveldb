package leveldb

import "github.com/JayLiang/dive-into-leveldb/leveldb/utils"

type FilterPolicy interface {
	Name() string
	CreateFilter(sliceArray []Slice, dst []byte) []byte
	KeyMayMatch(slice Slice, filter []byte) bool
}

const BLOOMSEED uint32 = 0xbc9f1d34

func NewBloomFilterPolicy(bitsPerKey uint64) FilterPolicy {
	k := int(float64(bitsPerKey) * 0.69)
	if k < 1 {
		k = 1
	}
	if k > 30 {
		k = 30
	}
	return &BloomFilterPolicy{
		bitsPerKey: bitsPerKey,
		k:          uint64(k),
	}
}

type BloomFilterPolicy struct {
	bitsPerKey uint64
	k          uint64
}

func (b *BloomFilterPolicy) hash(data []byte) uint32 {
	return utils.Hash(data, BLOOMSEED)
}

func (b BloomFilterPolicy) Name() string {
	return "leveldb.BuiltinBloomFilter2"
}

func (b BloomFilterPolicy) CreateFilter(sliceArray []Slice, dst []byte) []byte {
	l := len(sliceArray)
	bits := b.bitsPerKey * uint64(l)
	if bits < 64 {
		bits = 64
	}
	bytes := (bits + 7) / 8
	bits = bytes * 8
	filter := make([]byte, bytes)
	for i := 0; i < l; i++ {
		hash := b.hash(sliceArray[i].Data)
		delta := hash>>17 | hash<<15
		for j := uint64(0); j < b.k; j++ {
			bitPos := hash % uint32(bits)
			filter[bitPos/8] |= 1 << (bitPos % 8)
			hash += delta
		}
	}

	dst = append(dst, filter...)
	dst = append(dst, byte(b.k))
	return dst
}

func (b BloomFilterPolicy) KeyMayMatch(slice Slice, filter []byte) bool {
	l := uint32(len(filter))
	if l < 2 {
		return false
	}
	k := uint64(filter[l-1])
	bits := (l - 1) * 8

	if k > 30 {
		return true
	}
	hash := b.hash(slice.Data)
	delta := hash>>17 | hash<<15
	for j := uint64(0); j < k; j++ {
		bitPos := hash % bits
		if filter[bitPos/8]&(1<<(bitPos%8)) == 0 {
			return false
		}
		hash += delta
	}
	return true
}
