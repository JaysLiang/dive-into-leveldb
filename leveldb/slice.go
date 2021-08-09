package leveldb

import "bytes"

type Slice struct {
	Data []byte
}

func NewEmptySlice() *Slice {
	return &Slice{}
}

func NewSlice(data []byte) *Slice {
	return &Slice{Data: data}
}

func (s *Slice) GetData() []byte {
	return s.Data
}

func (s *Slice) Size() int {
	return len(s.Data)
}

func (s *Slice) Empty() bool {
	return s.Size() == 0
}

func (s *Slice) Clear() {
	s.Data = nil
}

func (s *Slice) RemovePrefix(n int) {
	if s.Size() <= n {
		panic("beyond max prefix")
	}
	s.Data = s.Data[n:]
}
func (s *Slice) Compare(other *Slice) int {
	return bytes.Compare(s.GetData(), other.GetData())
}
