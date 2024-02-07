# Hashing N-dimensional float vectors

Search nearest neighbour vectors in n-dimensional space with hashes. There are no dependencies in this package.

Each vestor is discretized into a set of hashes, as described [here](https://vitali-fedulov.github.io/similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html) (also as [PDF](https://github.com/vitali-fedulov/research/blob/main/Algorithm%20for%20hashing%20float%20vectors.pdf)).

## How to use

1) Provided a float vector []float64, use `CubeSet` and `CentralCube` functions to generate hypercube coordinates []int. The difference between the two functions is that one corresponds to hash-table record and the other to a query or vice versa, depending on performance/memory preference.
2) `HashSet` and `DecimalHash`/`FNV1aHash` are used to get corresponding hash set and central hash from the hypercube coordinates above. There are 2 alternative hash functions: DecimalHash and FNV1aHash. DecimalHash does not have collisions, but is not suitable for cases with large number of buckets or dimensions. FNV1aHash is applicable for all cases.

[Example](https://github.com/vitali-fedulov/imagehash/blob/master/hashes.go) for similar image search/clustering.

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/hyper) for full code documentation.
