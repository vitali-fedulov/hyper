package hyper

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

type Hash64 func(buckets []int) uint64

// Fnva64 returns a FVN-1a hash for a slice of bucket numbers.
func Fnva64(buckets []int) uint64 {
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
