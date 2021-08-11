package utils

const M uint32 = 0xc6a4a793
const R uint32 = 24

func Hash(data []byte, seed uint32) uint32 {
	l := uint32(len(data))
	h := seed ^ (l * M)
	i := uint32(0)
	for ; i+4 <= l; i = i + 4 {
		w := DecodeFixed32(data[i:])
		h += w
		h *= M
		h ^= h >> 16
	}

	switch l - i {
	case 3:
		h += uint32(data[i+2]) << 16
		fallthrough
	case 2:
		h += uint32(data[i+1]) << 8
		fallthrough
	case 1:
		h += uint32(data[i])
		h *= M
		h ^= h >> R
	default:
	}
	return h
}
