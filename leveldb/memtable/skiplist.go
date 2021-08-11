package memtable

type SkipList interface {
	Insert(key interface{})
	Contain(key interface{})
}

type Iterator interface {
	Valid() bool
	Key() interface{}
	Next()
	Prev()
	Seek(key interface{})
	SeekToFirst()
}