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

func TestBuckets(t *testing.T) {
	numBuckets, min, max, bucketPct := 10, 0.0, 255.0, 0.25
	values := []float64{25.5, 0.01, 210.3, 93.9, 6.6, 9.1, 254.9}
	bucketWidth, eps := Params(numBuckets, min, max, bucketPct)
	got := Buckets(values, min, max, bucketWidth, eps)
	want := [][]int{{0, 0, 7, 3, 0, 0, 9}, {1, 0, 7, 3, 0, 0, 9},
		{0, 0, 8, 3, 0, 0, 9}, {1, 0, 8, 3, 0, 0, 9}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`Got %v, want %v. Number of buckets is %v.`, got, want, numBuckets)
	}

	values = []float64{0.01, bucketWidth * 2 * 0.999, bucketWidth * 2 * 1.001}
	got = Buckets(values, min, max, bucketWidth, eps)
	want = [][]int{{0, 1, 1}, {0, 2, 1}, {0, 1, 2}, {0, 2, 2}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf(`Got %v, want %v. Number of buckets is %v.`, got, want, numBuckets)
	}
}
