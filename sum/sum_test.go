package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	checkSums(t, got, want)

}

func TestSumAllTails(t *testing.T) {
	t.Run("sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{0, 9, 2, 4, 6, 8})
		want := []int{7, 29}

		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9, 2, 4, 6, 8})
		want := []int{0, 29}

		checkSums(t, got, want)
	})

}

func checkSums(t testing.TB, got, want []int) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
