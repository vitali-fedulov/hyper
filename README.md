# Hashing float vectors in N-dimensions

Package hyper allows fast approximate search of nearest neighbour vectors in n-dimensional space.

**This is an early beta version**. Description below will be improved (TODO). See tests for examples.

Package functions discretize a vector and generate a set of hashes, as described in the following document: https://similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html

To use the package follow the sequence of functions/methods:
1) CubeSet or CentralCube, depending which one is used for a database record and which one for a query.
2) HashSet and DecimalHash to get corresponding hash set and central hash from results of (2). If DecimalHash is not suitable because of very large number of buckets or dimensions, use FNV1aHash to get both the hash set and the central hash).
