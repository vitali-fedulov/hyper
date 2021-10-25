package hyper

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// For a specific hashing function to be (re)defined.
type Hash func(buckets []int) uint64

// Fnva64 is a specific hash implementation, which returns
// a FVN-1a hash for a slice of bucket numbers.
func Fnva64(buckets []int) uint64 {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(buckets)
	hash := fnv.New64a()
	hash.Write(b.Bytes())
	return hash.Sum64()
}

// HashSet returns a slice of hashes for a tree of bucket ids.
func HashSet(tree [][]int, hash Hash) (hs []uint64) {
	for i := 0; i < len(tree); i++ {
		hs = append(hs, hash(tree[i]))
	}
	return hs
}
