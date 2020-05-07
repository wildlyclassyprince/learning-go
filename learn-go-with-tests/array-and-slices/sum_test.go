package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{3, 7, 9, 1, 6, 4, 2, 8, 5, 5}
		got := Sum(numbers)
		want := 50
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
