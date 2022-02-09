# Hashing float vectors in N-dimensions

Package hyper allows fast approximate search of nearest neighbour vectors in n-dimensional space. It does not have any dependencies.

Package functions discretize a vector and generate a set of hashes, as described [here](https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html) (also as [PDF](https://github.com/vitali-fedulov/research/blob/main/Algorithm%20for%20hashing%20float%20vectors.pdf)).

To use the package follow the sequence of functions/methods:
1) CubeSet or CentralCube, depending which one is used for a database record and which one for a query.
2) HashSet and DecimalHash to get corresponding hash set and central hash from results of (2). If DecimalHash is not suitable because of very large number of buckets or dimensions, use FNV1aHash to get both the hash set and the central hash).

[Example](https://github.com/vitali-fedulov/images3/blob/master/hashes.go) of usage for image comparison.

[Go doc](https://pkg.go.dev/github.com/vitali-fedulov/hyper) for the package.
