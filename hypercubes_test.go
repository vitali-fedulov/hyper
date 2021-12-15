package hyper

import (
	"reflect"
	"testing"
)

func TestParams(t *testing.T) {
	numBuckets, min, max, bucketPct := 10, 0.0, 255.0, 0.25
	bucketWidth, eps := Params(numBuckets, min, max, bucketPct)
	wantBucketWidth, wantEps := 25.5, 6.375
	if bucketWidth != wantBucketWidth {
		t.Errorf(`Got bucketWidth %v, want %v`, bucketWidth, wantBucketWidth)
	}
	if eps != wantEps {
		t.Errorf(`Got eps %v, want %v`, eps, wantEps)
	}
}

func TestParamsPanic(t *testing.T) {
	defer func() { recover() }()
	// Intentionally forbiden value for bucketPct.
	numBuckets, min, max, bucketPct := 10, 0.0, 255.0, 0.51
	_, _ = Params(numBuckets, min, max, bucketPct)
	// Never reaches here if Params panics.
	t.Errorf("Params did not panic on bucketPct > 0.5")
}

func TestHypercubes(t *testing.T) {
	numBuckets, min, max, bucketPct := 10, 0.0, 255.0, 0.25
	values := []float64{25.5, 0.01, 210.3, 93.9, 6.6, 9.1, 254.9}
	bucketWidth, eps := Params(numBuckets, min, max, bucketPct)
	gotCubes, gotCentral := Hypercubes(values, min, max, bucketWidth, eps)
	wantCubes := [][]int{{1, 0, 8, 3, 0, 0, 9}, {0, 0, 8, 3, 0, 0, 9},
		{1, 0, 7, 3, 0, 0, 9}, {0, 0, 7, 3, 0, 0, 9}}
	wantCentral := []int{1, 0, 8, 3, 0, 0, 9}
	if !reflect.DeepEqual(gotCubes, wantCubes) {
		t.Errorf(`Got %v, want %v.`, gotCubes, wantCubes)
	}
	if !reflect.DeepEqual(gotCentral, wantCentral) {
		t.Errorf(`Got %v, want %v.`, gotCentral, wantCentral)
	}
	if centralIsNotInTheSet(gotCubes, gotCentral) {
		t.Errorf(`Central %v is not in the set %v.`, gotCentral, gotCubes)
	}

	values = []float64{0.01, bucketWidth * 2 * 0.999, bucketWidth * 2 * 1.001}
	gotCubes, gotCentral = Hypercubes(values, min, max, bucketWidth, eps)
	wantCubes = [][]int{{0, 1, 2}, {0, 2, 2}, {0, 1, 1}, {0, 2, 1}}
	wantCentral = []int{0, 1, 2}
	if !reflect.DeepEqual(gotCubes, wantCubes) {
		t.Errorf(`Got %v, want %v.`, gotCubes, wantCubes)
	}
	if !reflect.DeepEqual(gotCentral, wantCentral) {
		t.Errorf(`Got %v, want %v.`, gotCentral, wantCentral)
	}
	if centralIsNotInTheSet(gotCubes, wantCentral) {
		t.Errorf(`Central %v is not in the set %v.`, gotCentral, gotCubes)
	}
}

func centralIsNotInTheSet(set [][]int, central []int) bool {
	for _, cube := range set {
		counter := 0
		for i, c := range central {
			if cube[i] == c {
				counter++
			}
		}
		if counter == len(central) {
			return false
		}
	}
	return true
}
