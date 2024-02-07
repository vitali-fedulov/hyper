package hyper

// Hypercube is represented by a slice of its coordinates.
type Cube []int
type Cubes []Cube

// Parameters of space discretization.
type Params struct {
	// Value limits per dimension. For example 0, 255 for pixel values.
	Min, Max float64
	// Uncertainty interval expressed as a fraction of bucketWidth
	// (for example 0.25 for eps = 1/4 of bucketWidth).
	EpsPercent float64
	// Number of buckets per dimension.
	NumBuckets int
}

// CubeSet returns a set of hypercubes, which represent
// fuzzy discretization of one n-dimensional vector,
// as described in
// https://vitali-fedulov.github.io/similar.pictures/algorithm-for-hashing-high-dimensional-float-vectors.html
// One hupercube is defined by bucket numbers in each dimension.
// min and max are minimum and maximum possible values of
// the vector components. The assumption is that min and max
// are the same for all dimensions.
func CubeSet(vector []float64, params Params) (set Cubes) {

	if params.EpsPercent >= 0.5 {
		panic(`Error: EpsPercent must be less than 0.5.`)
	}

	var (
		bC         int   // Central bucket number.
		bL, bR     int   // Left and right bucket number.
		setL, setR Cubes // Set clones (for Left and Right).
		branching  bool  // Branching flag.
	)

	// Rescaling vector to avoid potential mistakes with
	// divisions and offsets later on.
	rescaled := rescale(vector, params)
	// After the rescale value range of the vector are
	// [0, numBuckets], and not [min, max].

	// min = 0.0 from now on.
	max := float64(params.NumBuckets)

	for _, val := range rescaled {

		branching = false

		bL = int(val - params.EpsPercent)
		bR = int(val + params.EpsPercent)

		// Get extreme values out of the way.
		if val-params.EpsPercent <= 0.0 { // This means that val >= 0.
			bC = bR
			goto branchingCheck // No branching.
		}

		// Get extreme values out of the way.
		if val+params.EpsPercent >= max { // This means that val =< max.
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
func CentralCube(vector []float64, params Params) (central Cube) {

	if params.EpsPercent >= 0.5 {
		panic(`Error: EpsPercent must be less than 0.5.`)
	}

	var bC int // Central bucket numbers.

	// Rescaling vector to avoid potential mistakes with
	// divisions and offsets later on.
	rescaled := rescale(vector, params)
	// After the rescale value range of the vector are
	// [0, numBuckets], and not [min, max].

	// min = 0.0 from now on.
	max := float64(params.NumBuckets)

	for _, val := range rescaled {
		bC = int(val)
		if val-params.EpsPercent <= 0.0 { //  This means that val >= 0.
			bC = int(val + params.EpsPercent)
		}
		if val+params.EpsPercent >= max { // Meaning val =< max.
			bC = int(val - params.EpsPercent)
		}
		central = append(central, bC)
	}
	return central
}

// rescale is a helper function to offset and rescale all values
// to [0, numBuckets] range.
func rescale(vector []float64, params Params) []float64 {
	rescaled := make([]float64, len(vector))
	amp := params.Max - params.Min
	for i := range vector {
		// Offset to zero and rescale to [0, numBuckets] range.
		rescaled[i] =
			(vector[i] - params.Min) * float64(params.NumBuckets) / amp
	}
	return rescaled
}

// clone makes an unlinked copy of a 2D slice.
func clone(src Cubes) (dst Cubes) {
	dst = make(Cubes, len(src))
	for i := range src {
		dst[i] = append(Cube{}, src[i]...)
	}
	return dst
}
