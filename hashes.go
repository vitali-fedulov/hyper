package hyper

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// Hash64 can be any function of this kind.
type Hash64 func(buckets []int) uint64

// FVN1a is the default hash in this package.
func FVN1a(buckets []int) uint64 {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(buckets)
	hash := fnv.New64a()
	hash.Write(b.Bytes())
	return hash.Sum64()
}

// Hashes64 returns a set of hashes for a tree of bucket ids.
func Hashes64(tree [][]int, hash Hash64) (hs []uint64) {
	for i := 0; i < len(tree); i++ {
		hs = append(hs, hash(tree[i]))
	}
	return hs
}
