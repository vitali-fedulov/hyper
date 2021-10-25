package hyper

// Params returns discretization parameters.
// numBuckets represents number of discretization buckets into which all values
// will fall. Ids of those buckets will be used to create hashes.
// min and max are minimum and maximum possible values of discretized variable.
// bucketWidth is width of the discretization bucket.
// bucketPct is percentage of bucketWidth to allow for an error of discretized
// variable (a specific value of a discretized variable may fall into 2 buckets
// simultaneosly).
// eps is actual width corresponding to the bucketWidth bucketPct on the discretized
// variable axis.
func Params(numBuckets int, min, max, bucketPct float64) (bucketWidth, eps float64) {
	if bucketPct >= 0.5 {
		panic("Error: bucketPct must be less than 50%. Recommendation: decrease numBuckets instead.")
	}
	bucketWidth = (max - min) / float64(numBuckets)
	eps = bucketPct * bucketWidth
	return bucketWidth, eps
}

// Buckets generates a set of slices of all possible bucket ids
// as permutations based on n-dimensional space discretization.
// point are values for each of those n dimensions.
// min and max are minimum and maximum possible values of discretized
// point components. The assumption is that min and max are the same for all
// dimensions (in the context of the Buckets function).
// bucketWidth and eps are defined in the Params function.
func Buckets(point []float64, min, max, bucketWidth, eps float64) (tree [][]int) {

	// Bucket ids. Default bucket is b.
	var (
		val      float64 // Sample value (one axis of n-space).
		bL, bR   int     // Left and right bucket ids.
		treeCopy [][]int // Bucket tree copy.
		length   int
	)

	// For each component of the point.
	for k := 0; k < len(point); k++ {
		val = point[k]

		bL = int((val - eps) / bucketWidth)
		bR = int((val + eps) / bucketWidth)

		if val-eps < min { // No bucket for smaller than min.
			bL = bR
		} else if val+eps > max { // No bucket for larger than max.
			bR = bL
		}

		if bL == bR { // No branching.
			if len(tree) == 0 {
				tree = append(tree, []int{bL})
			} else {
				length = len(tree)
				for i := 0; i < length; i++ {
					// Constructing buckets set.
					tree[i] = append(tree[i], bL)
				}
			}

		} else { // Branching.
			treeCopy = make([][]int, len(tree))
			copy(treeCopy, tree)

			if len(tree) == 0 {
				tree = append(tree, []int{bL})
			} else {
				length = len(tree)
				for i := 0; i < length; i++ {
					tree[i] = append(tree[i], bL)
				}
			}

			if len(treeCopy) == 0 {
				treeCopy = append(treeCopy, []int{bR})
			} else {
				length = len(treeCopy)
				for i := 0; i < length; i++ {
					treeCopy[i] = append(treeCopy[i], bR)
				}
			}

			tree = append(tree, treeCopy...)
		}

	}

	// Verification that branching works correctly and no buckets are lost.
	// TODO: Disable once whole package got tested on large image sets.
	length = len(point)
	for i := 0; i < len(tree); i++ {
		if len(tree[i]) != length {
			panic(`Buckets slice length must be equal to len(point).`)
		}
	}

	return tree
}
