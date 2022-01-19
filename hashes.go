package hyper

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

// DecimalHash hashes hypercubes without collisions. IMPORTANT:
// To work correctly, the number of buckets must be
// less than 11 and the number of dimensions less than 20.
// Else at certain unexpected moment you might get a hash
// value overflow.
func (cube Cube) DecimalHash() (h uint64) {
	for _, v := range cube {
		h = h*10 + uint64(v)
	}
	return h
}

// FNV1aHash hashes hypercubes with rare collisions,
// and should be used when Decimal cannot be used
// because of very large number of buckets or dimensions.
func (cube Cube) FNV1aHash() uint64 {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(cube)
	hash := fnv.New64a()
	hash.Write(b.Bytes())
	return hash.Sum64()
}

// HashFunc can be any function (also user-defined).
type HashFunc func(cube Cube) uint64

// Hash64Set returns a set of hashes for a hypercube set
// and a concrete hash function.
func (cubeSet Cubes) HashSet(hashFunc HashFunc) (
	hs []uint64) {
	for i := 0; i < len(cubeSet); i++ {
		hs = append(hs, hashFunc(cubeSet[i]))
	}
	return hs
}
