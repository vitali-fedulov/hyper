package hyper

// Package hyper allows fast approximate search of nearest
// neighbour vectors in n-dimensional space.
// Package functions discretize a vector and generate a set
// of fuzzy hashes, as described in the following document:
// https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html

// To use the package follow the sequence:
// 1) Params, 2) CubeSet or CentralCube, depending which one
// is used for a database record and which one for a query,
// 3) HashSet and Decimal to get corresponding hash set
// and central hash from results of (2). If Decimal hash
// is not suitable because of very large number of buckets
// or dimensions,  use FNV1a to get both the hash set and
// the central hash).
