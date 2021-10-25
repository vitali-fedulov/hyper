package hyper

import (
	"reflect"
	"testing"
)

func TestFnva64(t *testing.T) {
	buckets := []int{5, 59, 255, 9, 7, 12, 22, 31}
	hash := Fnva64(buckets)
	want := uint64(13992349377752315208)
	if hash != want {
		t.Errorf(`Got %v, want %v`, hash, want)
	}
}

func TestHashSet(t *testing.T) {
	tree := [][]int{
		{0, 0, 7, 3, 0, 0, 9},
		{1, 0, 7, 3, 0, 0, 9},
		{0, 0, 8, 3, 0, 0, 9},
		{1, 0, 8, 3, 0, 0, 9}}
	hs := HashSet(tree, Fnva64)
	want := []uint64{
		14647827280143437043,
		17530493565529410009,
		7065940388079601005,
		13953051952027146823}
	if !reflect.DeepEqual(hs, want) {
		t.Errorf(`Got %v, want %v`, hs, want)
	}
}