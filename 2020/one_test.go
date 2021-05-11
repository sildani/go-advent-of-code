package one

import "testing"

func TestEval(t *testing.T) {
	// find the product of the two numbers in the list that add up to 2020

	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	expected := 514579
	actual := Eval(input)

	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}

	// TODO: test for error when no combination matches
}