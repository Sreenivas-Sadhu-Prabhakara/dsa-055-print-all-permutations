package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermutations(t *testing.T) {
	got := Permutations([]int{1, 2, 3})
	if len(got) != 6 {
		t.Fatalf("expected 6 permutations, got %d", len(got))
	}
	sort.Slice(got, func(i, j int) bool {
		for k := range got[i] {
			if got[i][k] != got[j][k] {
				return got[i][k] < got[j][k]
			}
		}
		return false
	})
	want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("permutation set = %v, want %v", got, want)
	}

	if n := len(Permutations([]int{1, 2, 3, 4})); n != 24 {
		t.Errorf("expected 24, got %d", n)
	}
	if got := Permutations([]int{7}); !reflect.DeepEqual(got, [][]int{{7}}) {
		t.Errorf("single element: %v", got)
	}
	if got := Permutations([]int{}); !reflect.DeepEqual(got, [][]int{{}}) {
		t.Errorf("empty input should give one empty permutation, got %v", got)
	}
}
