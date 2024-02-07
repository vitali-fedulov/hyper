package hyper

import (
	"reflect"
	"testing"
)

func TestDecimalHash(t *testing.T) {
	cube := Cube{3, 2, 0, 1, 1, 4, 1, 0}
	hash := cube.DecimalHash()
	want := uint64(32011410)
	if hash != want {
		t.Errorf(`Got %v, want %v.`, hash, want)
	}
}

func TestFNV1aHash(t *testing.T) {
	cube := Cube{5, 59, 255, 9, 7, 12, 22, 31}
	hash := cube.FNV1aHash()
	want := uint64(6267598672213710911)
	if hash != want {
		t.Errorf(`Got %v, want %v.`, hash, want)
	}
}

func TestHashSet(t *testing.T) {
	cubes := Cubes{
		{0, 0, 7, 3, 0, 0, 9},
		{1, 0, 7, 3, 0, 0, 9},
		{0, 0, 8, 3, 0, 0, 9},
		{1, 0, 8, 3, 0, 0, 9}}
	hashSet := cubes.HashSet((Cube).FNV1aHash)
	want := []uint64{
		9211138565158515574,
		6304441926533466432,
		5296875461196147964,
		13706017245957046114}
	if !reflect.DeepEqual(hashSet, want) {
		t.Errorf(`Got %v, want %v.`, hashSet, want)
	}
}
