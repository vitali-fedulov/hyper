package hyper

// Package hyper allows fast approximate search of nearest
// neighbour vectors in n-dimensional space.
// Package functions discretize a vector and generate a set
// of fuzzy hashes, as described in the following document:
// https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html

// A typical sequence of functions when using the package is:
// 1) Params, 2) CubeSet or CentralCube, depending which one
// is used for a database record and which one for a query,
// 3) HashSet or CentralHash to get corresponding hashes
// from results of (2).

// It is possible to define own hashing function instead of
// using the default one.
