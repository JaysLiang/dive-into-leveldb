package utils

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestEncodeVarUint32(t *testing.T) {
	for i := uint32(0); i < 100000; i++ {
		buf1 := make([]byte, 5)
		buf2 := make([]byte, 5)
		len1 := EncodeVarUint32(buf1, i)
		len2 := binary.PutUvarint(buf2, uint64(i))
		if len1 != len2 {
			t.Fatalf("wrong lenth expected %v, but got %v", len2, len1)
		}
		if bytes.Compare(buf1[:len1], buf2[:len2]) != 0 {
			t.Fatalf("implement wrong, src: %v expected %+v, but got %+v", i, buf2[:len2], buf1[:len1])
		}
	}
}

func TestDecodeVarUint32(t *testing.T) {
	for i := uint32(0); i < 100000; i++ {
		buf1 := make([]byte, 5)
		len1 := EncodeVarUint32(buf1, i)
		dst, len2 := DecodeVarUint32(buf1[:len1])
		if len1 != len2 {
			t.Fatalf("wrong len, expected %+v ,but got %+v", len1, len2)
		}
		if dst != i {
			t.Fatalf("wrong decode, Expected %v, but got %v", i, dst)
		}
	}
}
