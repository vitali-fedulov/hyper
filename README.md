# Hashing N-dimensional float vectors

Search nearest neighbour vectors in n-dimensional space with hashes. There are no dependencies in this package.

The algorithm is based on the assumption that two real numbers can be considered equal within certain equality distance. Then quantization is used for comparison. To make sure points near or at quantization borders are also comparable, a vector can be discretized into more than one hash, as described [here](https://vitali-fedulov.github.io/similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html) (also as [PDF](https://github.com/vitali-fedulov/research/blob/main/Algorithm%20for%20hashing%20float%20vectors.pdf)). The method indirectly clusters given vectors by hypercubes and their neighbourhoods.


## How to use

1) Normalize values of your input vectors to same min/max value range. Use min/max values in the parameters settings.
2) Provided a float vector []float64, use `CentralCube` and `CubeSet` functions to generate hypercube coordinates []int and [][]int.
3) Generate a `DecimalHash`/`FNV1aHash` and `HashSet` for corresponding central hash and hash set from the hypercube coordinates above. The difference between one hash and a hash set is that one corresponds to a hash-table record and the other to a query or vice versa, depending on performance/memory preference. There are 2 alternative hash functions: DecimalHash and FNV1aHash. DecimalHash does not have collisions, but is not suitable for cases with large number of buckets or dimensions. FNV1aHash is applicable for all cases.

[Example](https://github.com/vitali-fedulov/imagehash2/blob/main/hashes.go) for similar image search and clustering.

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/hyper) for full code documentation.
