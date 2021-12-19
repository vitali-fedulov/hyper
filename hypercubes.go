package hyper

// Params helps with discretization parameters.
// numBuckets is number of buckets per dimension.
// min and max are value limits per dimension.
// epsPercent is the uncertainty interval expressed as fraction
// of bucketWidth.
// eps is the absolute value of the uncertainty interval epsilon.
func Params(
	numBuckets int, min, max, epsPercent float64) (
	bucketWidth, eps float64) {
	if epsPercent >= 0.5 {
		panic(`Error: epsPercent must be less than 50%.
			Recommendation: decrease numBuckets instead.`)
	}
	bucketWidth = (max - min) / float64(numBuckets)
	eps = epsPercent * bucketWidth
	return bucketWidth, eps
}

// CubeSet returns a set of hypercubes, which represent
// fuzzy discretization of one n-dimensional vector,
// as described in
// https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html
// One hupercube is defined by bucket numbers in each dimension.
// min and max are minimum and maximum possible values of
// the vector components. The assumption is that min and max
// are the same for all dimensions.
// bucketWidth and eps are defined in the Params function.
func CubeSet(
	vector []float64, min, max, bucketWidth, eps float64) (
	set [][]int) {

	var (
		bC, bS    int     // Central and side bucket ids.
		setCopy   [][]int // Set copy.
		length    int
		branching bool // Branching flag.
	)

	// For each component of the vector.
	for _, val := range vector {

		bC = int(val / bucketWidth)
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
			panic(`Number of hypercube coordinates must equal
			to len(vector).`)
		}
	}

	return set
}

// CentralCube returns the hypercube containing the vector end.
// Arguments are the same as for the CubeSet function.
func CentralCube(
	vector []float64, min, max, bucketWidth, eps float64) (
	central []int) {

	var bC int // Central bucket ids.

	// For each component of the vector.
	for _, val := range vector {
		bC = int(val / bucketWidth)
		central = append(central, bC)
	}

	return central
}
