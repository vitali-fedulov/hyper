package hyper

// Package hyper allows fast approximate search of nearest
// neighbour vectors in n-dimensional space.
// Package functions discretize a vector and generate a set
// of fuzzy hashes, as described in the following paper:
// https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html

// A typical sequence of functions when using the package is:
// 1) Params, 2) Hypercubes, 3) FVN1a to get the central hash,
// and Hashes64 with FVN1a as the hash argument to get
// the full hash set.

// You can also define own function for hashing hypercubes.
