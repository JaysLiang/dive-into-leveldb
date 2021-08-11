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
		dst = (uint32(buf[i]&0x7f)<<s) | dst
		s += 7
		i++
	}
	dst = (uint32(buf[i]&0x7f)<<s) | dst
	return dst, i + 1
}
