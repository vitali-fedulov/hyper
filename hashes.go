package hyper

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// Decimal hashes hypercubes without collisions. For that
// it assumes that number of buckets is 10 or less
// and number of dimensions is 19 or less.
func Decimal(cube []int, numBuckets int) (h uint64) {
	if numBuckets > 10 {
		panic(`Decimal hash can only be used if
		numBuckets <= 10. FVN1a can be used instead.`)
	}
	// Max uint64 equals 18446744073709551615,
	// therefore larger number of dimensions will overflow.
	if len(cube) > 19 {
		panic(`Decimal hash can only be used if
		number of dimensions is less than 20.
		FVN1a hash can be used instead.`)
	}
	for _, v := range cube {
		h = h*10 + uint64(v)
	}
	return h
}

// FVN1a hashes hypercubes with rare collisions,
// and should be used when Decimal cannot be used
// because of very large number of buckets or dimensions.
func FVN1a(cube []int) uint64 {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(cube)
	hash := fnv.New64a()
	hash.Write(b.Bytes())
	return hash.Sum64()
}

// HashFunc can be any function (also user-defined).
type HashFunc func(hypercube []int) uint64

// Hash64Set returns a set of hashes for a hypercube set
// and a concrete hash function.
func HashSet(cubeSet [][]int, hashFunc HashFunc) (
	hs []uint64) {
	for i := 0; i < len(cubeSet); i++ {
		hs = append(hs, hashFunc(cubeSet[i]))
	}
	return hs
}
