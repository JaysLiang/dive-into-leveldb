package leveldb

import (
	"bytes"
	"testing"
)

func TestNewEmptySlice(t *testing.T) {
	slice := NewEmptySlice()
	if slice.Size() != 0 {
		t.Fatalf("empty size error, expected 0, but got %d", slice.Size())
	}
}

func TestNewSlice(t *testing.T) {
	data := []byte("hello world")
	slice := NewSlice(data)
	if bytes.Compare(slice.Data, data) != 0 {
		t.Fatalf("Error: slice'data is changed")
	}
	slice.RemovePrefix(6)
	if bytes.Compare(slice.Data, []byte("world")) != 0 {
		t.Fatalf("remove prefix error, expected %+v, but got %+v", string(slice.Data), string(slice.Data))
	}
}
