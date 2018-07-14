package gcws

// Cuckoo Filter implementation in Go
/*
 */

type CuckooFilter struct {
}

func New() *CuckooFilter {
	return new(CuckooFilter)
}

func (cf *CuckooFilter) Insert(item []byte) bool {
	return false
}

func (fc *CuckooFilter) Lookup(item []byte) bool {
	return false
}

func (cf *CuckooFilter) Delete(item []byte) bool {
	return false
}

func (cf *CuckooFilter) Count() uint {
	return 0
}
