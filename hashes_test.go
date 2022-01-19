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
	want := uint64(1659788114117494335)
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
		6172277127052188606,
		3265650857171344968,
		13730239218993256724,
		6843127655045710906}
	if !reflect.DeepEqual(hashSet, want) {
		t.Errorf(`Got %v, want %v.`, hashSet, want)
	}
}
