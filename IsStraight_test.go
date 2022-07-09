package main

import "testing"

func TestAlgorithm(t *testing.T){
	result1 := IsStraight([]int{2, 3, 4 ,5, 6})
	want1 := true

	if result1 != want1 {
		t.Errorf("got %t, wanted %t", result1, want1)
	}

	result2 := IsStraight([]int{14, 5, 4 ,2, 3})
	want2 := true

	if result2 != want2 {
		t.Errorf("got %t, wanted %t", result2, want2)
	}

	result3 := IsStraight([]int{7, 7, 12 ,11, 3, 4, 14})
	want3 := false

	if result3 != want3 {
		t.Errorf("got %t, wanted %t", result3, want3)
	}

	result4 := IsStraight([]int{7, 3, 2})
	want4 := false

	if result4 != want4 {
		t.Errorf("got %t, wanted %t", result4, want4)
	}
}
