package hyper

// Params returns discretization parameters.
// numBuckets represents number of discretization buckets into
// which all values will fall. Ids of those buckets will be used
// to create hashes.
// min and max are minimum and maximum possible values
// of discretized variable.
// bucketWidth is width of the discretization bucket.
// bucketPct is percentage of bucketWidth to allow for an error
// of discretized variable (a specific value of a discretized
// variable may fall into 2 buckets simultaneosly).
// eps is actual width corresponding to the bucketWidth bucketPct
// on the discretized variable axis.
func Params(
	numBuckets int, min, max, bucketPct float64) (bucketWidth, eps float64) {
	if bucketPct >= 0.5 {
		panic(`Error: bucketPct must be less than 50%.
			Recommendation: decrease numBuckets instead.`)
	}
	bucketWidth = (max - min) / float64(numBuckets)
	eps = bucketPct * bucketWidth
	return bucketWidth, eps
}

// Hypercubes returns a set of hypercubes, which represent
// fuzzy discretization of one n-dimensional vector, as described in
// https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html
// One hupercube is defined by bucket numbers in each dimension.
// The function also returns the central hypercube (in which
// the vector end is located).
// min and max are minimum and maximum possible values of
// the vector components. The assumption is that min and max
// are the same for all dimensions.
// bucketWidth and eps are defined in the Params function.
func Hypercubes(
	vector []float64, min, max, bucketWidth, eps float64) (
	set [][]int, central []int) {

	var (
		bC, bS    int     // Central and side bucket ids.
		setCopy   [][]int // Set copy.
		length    int
		branching bool // Branching flag.
	)

	// For each component of the vector.
	for _, val := range vector {

		bC = int(val / bucketWidth)
		central = append(central, bC)
		branching = false

		// Value is in the lower uncertainty interval.
		if val-float64(bC)*bucketWidth < eps {
			bS = bC - 1
			if val-eps >= min {
				branching = true
			}

			// Value is in the upper uncertainty interval.
		} else if float64(bC+1)*bucketWidth-val < eps {
			bS = bC + 1
			if val+eps <= max {
				branching = true
			}
		}

		if branching {
			setCopy = make([][]int, len(set))
			copy(setCopy, set)

			if len(set) == 0 {
				set = append(set, []int{bC})
			} else {
				length = len(set)
				for i := 0; i < length; i++ {
					set[i] = append(set[i], bC)
				}
			}

			if len(setCopy) == 0 {
				setCopy = append(setCopy, []int{bS})
			} else {
				length = len(setCopy)
				for i := 0; i < length; i++ {
					setCopy[i] = append(setCopy[i], bS)
				}
			}

			set = append(set, setCopy...)

		} else {

			if len(set) == 0 {
				set = append(set, []int{bC})
			} else {
				length = len(set)
				for i := 0; i < length; i++ {
					set[i] = append(set[i], bC)
				}
			}
		}
	}

	// Real use case verification that branching works correctly
	// and no buckets are lost for a very large number of vectors.
	// TODO: Remove once tested.
	length = len(vector)
	for i := 0; i < len(set); i++ {
		if len(set[i]) != length {
			panic(`Number of hypercube coordinates must equal to len(vector).`)
		}
	}

	return set, central
}
