package hyper

import (
	"reflect"
	"testing"
)

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

func TestRescale(t *testing.T) { // Testing panic.
	numBuckets, min, max, _ := 10, 0.0, 255.0, 0.25
	vector := []float64{25.5, 0.01, 210.3, 93.9, 6.6, 9.1, 255.0}
	rescaled := rescale(vector, numBuckets, min, max)
	got := rescaled
	want := []float64{
		1, 0.0003921568627450981, 8.24705882352941,
		3.6823529411764704, 0.25882352941176473,
		0.3568627450980392, 10}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`Got %v, want %v.`, got, want)
	}
}

func TestCubeSet1(t *testing.T) { // Testing panic.
	defer func() { recover() }()
	// Intentionally forbiden value for epsPercent.
	values := []float64{25.5, 0.01, 210.3, 93.9, 6.6, 9.1, 254.9}
	min, max, epsPercent, numBuckets := 0.0, 255.0, 0.51, 10
	_ = CubeSet(values, min, max, epsPercent, numBuckets)
	// Never reaches here if Params panics.
	t.Errorf("Params did not panic on epsPercent > 0.5")
}

func TestCubeSet2(t *testing.T) {
	numBuckets, min, max, epsPercent := 10, 0.0, 255.0, 0.25
	values := []float64{25.5, 0.01, 210.3, 93.9, 6.6, 9.1, 254.9}
	gotCubes := CubeSet(values, min, max, epsPercent, numBuckets)
	gotCentral := CentralCube(values, min, max, epsPercent, numBuckets)
	wantCubes := [][]int{{0, 0, 7, 3, 0, 0, 9}, {1, 0, 7, 3, 0, 0, 9},
		{0, 0, 8, 3, 0, 0, 9}, {1, 0, 8, 3, 0, 0, 9}}
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
}

// Testing bucket borders.
func TestCubeSet3(t *testing.T) {
	numBuckets, min, max, epsPercent := 4, 0.0, 4.0, 0.25
	values := []float64{0.01, 2 * 0.999, 2 * 1.001}
	gotCubes := CubeSet(values, min, max, epsPercent, numBuckets)
	gotCentral := CentralCube(values, min, max, epsPercent, numBuckets)
	wantCubes := [][]int{{0, 1, 1}, {0, 2, 1}, {0, 1, 2}, {0, 2, 2}}
	wantCentral := []int{0, 1, 2}
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

// Testing extreme buckets.
func TestCubeSet4(t *testing.T) {
	values := []float64{255.0, 0.0, 255.0, 0.0, 255.0, 0.0, 255.0}
	numBuckets, min, max, epsPercent := 4, 0.0, 255.0, 0.25
	gotCubes := CubeSet(values, min, max, epsPercent, numBuckets)
	wantCubes := [][]int{{3, 0, 3, 0, 3, 0, 3}}
	if !reflect.DeepEqual(gotCubes, wantCubes) {
		t.Errorf(`Got %v, want %v.`, gotCubes, wantCubes)
	}
}

var vector = []float64{
	0, 183, 148, 21, 47, 16, 69, 45, 151, 64, 181}

func TestCubeSet5(t *testing.T) {
	numBuckets, min, max, epsPercent := 4, 0.0, 255.0, 0.25
	gotCubes := CubeSet(vector, min, max, epsPercent, numBuckets)
	wantCubes := [][]int{
		{0, 2, 2, 0, 0, 0, 0, 0, 2, 0, 2}, {0, 3, 2, 0, 0, 0, 0, 0, 2, 0, 2},
		{0, 2, 2, 0, 0, 0, 1, 0, 2, 0, 2}, {0, 3, 2, 0, 0, 0, 1, 0, 2, 0, 2},
		{0, 2, 2, 0, 0, 0, 0, 0, 2, 1, 2}, {0, 3, 2, 0, 0, 0, 0, 0, 2, 1, 2},
		{0, 2, 2, 0, 0, 0, 1, 0, 2, 1, 2}, {0, 3, 2, 0, 0, 0, 1, 0, 2, 1, 2},
		{0, 2, 2, 0, 0, 0, 0, 0, 2, 0, 3}, {0, 3, 2, 0, 0, 0, 0, 0, 2, 0, 3},
		{0, 2, 2, 0, 0, 0, 1, 0, 2, 0, 3}, {0, 3, 2, 0, 0, 0, 1, 0, 2, 0, 3},
		{0, 2, 2, 0, 0, 0, 0, 0, 2, 1, 3}, {0, 3, 2, 0, 0, 0, 0, 0, 2, 1, 3},
		{0, 2, 2, 0, 0, 0, 1, 0, 2, 1, 3}, {0, 3, 2, 0, 0, 0, 1, 0, 2, 1, 3}}
	if !reflect.DeepEqual(gotCubes, wantCubes) {
		t.Errorf(`Got %v, want %v.`, gotCubes, wantCubes)
	}
}
