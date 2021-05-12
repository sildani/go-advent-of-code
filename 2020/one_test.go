package one

import "testing"

func TestEval(t *testing.T) {
	// find the product of the two numbers in the list that add up to 2020

	type test struct {
    expected   int
		actual     int
  }

  tests := []test{
    {expected: 514579, actual: Eval([]int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		})},
		{expected: 0, actual: Eval([]int{0,1,})},
  }

  for _, tc := range tests {
    if tc.expected != tc.actual {
      t.Errorf("expected '%v', got '%v'", tc.expected, tc.actual)
    }
  }
}

func TestEvalPartTwo(t *testing.T) {
	type test struct {
    expected   int
		actual     int
  }

  tests := []test{
    {expected: 241861950, actual: EvalPartTwo([]int{
			1721,
			979,
			366,
			299,
			675,
			1456,
		})},
		{expected: 0, actual: EvalPartTwo([]int{0,1,})},
  }

  for _, tc := range tests {
    if tc.expected != tc.actual {
      t.Errorf("expected '%v', got '%v'", tc.expected, tc.actual)
    }
  }
}