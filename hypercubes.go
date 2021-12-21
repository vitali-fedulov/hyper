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

// clone makes a totally independent copy of a 2D slice.
func clone(src [][]int) (dst [][]int) {
	dst = make([][]int, len(src))
	for i := range src {
		dst[i] = append([]int{}, src[i]...)
	}
	return dst
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
		bC         int     // Central bucket number.
		bL, bR     int     // Left and right bucket number.
		setL, setR [][]int // Set copies.
		branching  bool    // Branching flag.
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
			branching = true
		}

	branchingCheck:

		if branching {

			setL = clone(set)
			setR = clone(set)

			if len(setL) == 0 {
				setL = append(setL, []int{bL})
			} else {
				for i := range setL {
					setL[i] = append(setL[i], bL)
				}
			}

			if len(setR) == 0 {
				setR = append(setR, []int{bR})
			} else {
				for i := range setR {
					setR[i] = append(setR[i], bR)
				}
			}

			set = append(setL, setR...)

		} else { // No branching.
			if len(set) == 0 {
				set = append(set, []int{bC})
			} else {
				for i := range set {
					set[i] = append(set[i], bC)
				}
			}
		}
	}

	// Real use case verification that branching works correctly
	// and no buckets are lost for a very large number of vectors.
	// TODO: Remove once tested.
	for i := 0; i < len(set); i++ {
		if len(set[i]) != len(vector) {
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
