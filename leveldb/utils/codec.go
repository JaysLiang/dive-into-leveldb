package utils

func EncodeVarUint32(buf []byte, v uint32) int {
	i := 0
	for v >= 0x80 {
		buf[i] = byte(v) | 0x80
		v = v >> 7
		i++
	}
	buf[i] = byte(v)
	return i + 1
}

func DecodeVarUint32(buf []byte) (uint32, int) {
	i := 0
	var dst uint32 = 0
	s := 0
	for (buf[i] & 0x80) == 0x80 {
		dst = (uint32(buf[i]&0x7f) << s) | dst
		s += 7
		i++
	}
	dst = (uint32(buf[i]&0x7f) << s) | dst
	return dst, i + 1
}

func DecodeFixed32(data []byte) uint32 {
	return uint32(data[0]) | uint32(data[1])<<8 |
		uint32(data[2])<<16 | uint32(data[3])<<24
}

func EncodeFixed32(src uint32, dst []byte) {
	dst[0] = byte(src)
	dst[1] = byte(src >> 8)
	dst[2] = byte(src >> 16)
	dst[3] = byte(src >> 24)
}

func EncodeFixed64(src uint64, dst []byte) {
	dst[0] = byte(src)
	dst[1] = byte(src >> 8)
	dst[2] = byte(src >> 16)
	dst[3] = byte(src >> 24)
	dst[4] = byte(src >> 32)
	dst[5] = byte(src >> 40)
	dst[6] = byte(src >> 48)
	dst[7] = byte(src >> 56)
}
func DecodeFixed64(data []byte) uint64 {
	return uint64(data[0]) | uint64(data[1])<<8 |
		uint64(data[2])<<16 | uint64(data[3])<<24 |
		uint64(data[4])<<32 | uint64(data[5])<<40 |
		uint64(data[6])<<48 | uint64(data[7])<<56
}
