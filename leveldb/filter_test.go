package leveldb

import "testing"

func TestNewBloomFilterPolicy(t *testing.T) {
	policy := NewBloomFilterPolicy(64)
	s := make([]Slice, 0)
	s = append(s, Slice{
		Data: []byte("hello"),
	}, Slice{Data: []byte("world")})
	filter := policy.CreateFilter(s, nil)
	match := policy.KeyMayMatch(Slice{Data: []byte("hello")}, filter)
	if !match {
		t.Fatalf("not match")
	}
	if !policy.KeyMayMatch(Slice{Data: []byte("world")}, filter) {
		t.Fatalf("not match")
	}
	if policy.KeyMayMatch(Slice{Data: []byte("sensetime")}, filter) {
		t.Fatalf("error, should not be matched")
	}
}
