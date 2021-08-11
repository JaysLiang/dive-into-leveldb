package leveldb

type Iterator interface {
	Valid() bool
	SeekToFirst()
	SeekToLast()
	Seek(s Slice)
	Next()
	Prev()
	Key() Slice
	Value() Slice
	Status() Status
}
