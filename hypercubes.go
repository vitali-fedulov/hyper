package hyper

// rescale is a helper function to offset and rescale all values
// to [0, numBuckets] range.
func rescale(vector []float64, numBuckets int, min, max float64) []float64 {
	rescaled := make([]float64, len(vector))
	amp := max - min
	for i := range vector {
		// Offset to zero and rescale to [0, numBuckets] range.
		rescaled[i] = (vector[i] - min) * float64(numBuckets) / amp
	}
	return rescaled
}

// CubeSet returns a set of hypercubes, which represent
// fuzzy discretization of one n-dimensional vector,
// as described in
// https://vitali-fedulov.github.io/algorithm-for-hashing-high-dimensional-float-vectors.html
// One hupercube is defined by bucket numbers in each dimension.
// min and max are minimum and maximum possible values of
// the vector components. The assumption is that min and max
// are the same for all dimensions.
// numBuckets is number of buckets per dimension.
// min and max are value limits per dimension.
// epsPercent is the uncertainty interval expressed as
// a fraction of bucketWidth (for example 0.25 for eps = 1/4
// of bucketWidth).
func CubeSet(vector []float64, min, max, epsPercent float64,
	numBuckets int) (set [][]int) {

	if epsPercent >= 0.5 {
		panic(`Error: epsPercent must be less than 0.5.`)
	}

	var (
		bC, bS    int     // Central and side bucket number.
		bL, bR    int     // Left and right bucket number.
		setCopy   [][]int // Set copy.
		length    int
		branching bool // Branching flag.
	)

	// Rescaling vector to avoid potential mistakes with
	// divisions and offsets later on.
	rescaled := rescale(vector, numBuckets, min, max)
	// After the rescale value range of the vector are
	// [0, numBuckets], and not [min, max].

	// min = 0.0 from now on.
	max = float64(numBuckets)

	for _, val := range rescaled {

		branching = false

		bL = int(val - epsPercent)
		bR = int(val + epsPercent)

		// Get extreme values out of the way.
		if val-epsPercent <= 0.0 { // This means that val >= 0.
			bC = bR
			goto branchingCheck // No branching.
		}

		// Get extreme values out of the way.
		if val+epsPercent >= max { // This means that val =< max.
			// Above max = numBuckets.
			bC = bL
			goto branchingCheck // No branching.
		}

		if bL == bR {
			bC = bL
			goto branchingCheck // No branching.

		} else { // Meaning bL != bR and not any condition above.
			bC = int(val)
			if bL == bC {
				bS = bR // So we have bC, have not lost bL, and get bR.
			} else { // That is when bL != bC
				bS = bL // So we have bC, have bL, and since can only have
				// 2 buckets possible, bC is our bR (bR not lost).
			}
			branching = true
		}

	branchingCheck:
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
func CentralCube(vector []float64, min, max, epsPercent float64,
	numBuckets int) (central []int) {

	if epsPercent >= 0.5 {
		panic(`Error: epsPercent must be less than 0.5.`)
	}

	var bC int // Central bucket numbers.

	// Rescaling vector to avoid potential mistakes with
	// divisions and offsets later on.
	rescaled := rescale(vector, numBuckets, min, max)
	// After the rescale value range of the vector are
	// [0, numBuckets], and not [min, max].

	// min = 0.0 from now on.
	max = float64(numBuckets)

	for _, val := range rescaled {
		bC = int(val)
		if val-epsPercent <= 0.0 { //  This means that val >= 0.
			bC = int(val + epsPercent)
		}
		if val+epsPercent >= max { // Meaning val =< max.
			bC = int(val - epsPercent)
		}
		central = append(central, bC)
	}
	return central
}
