# Hashing float vectors in N-dimensions

Search nearest neighbour vectors in n-dimensional space with hashes. There are no dependencies in this package.

Its functions discretize a vector and generate a set of hashes, as described [here](https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html) (also as [PDF](https://github.com/vitali-fedulov/research/blob/main/Algorithm%20for%20hashing%20float%20vectors.pdf)).

Usage sequence:
1) CubeSet or CentralCube, depending which one is used for a database record and which one for a query.
2) HashSet and DecimalHash to get corresponding hash set and central hash from results of (2). If DecimalHash is not suitable because of very large number of buckets or dimensions, use FNV1aHash to get both the hash set and the central hash).

[Example](https://github.com/vitali-fedulov/images3/blob/master/hashes.go) of usage for image comparison.

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/hyper) for code documentation.
