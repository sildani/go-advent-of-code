package five

import "testing"

func TestParseData(t *testing.T) {
  dat := []byte(
    "BFFFBBFRRR\r\n"+
    "FFFBBBFRRR\r\n"+
    "BBFFBBFRLL")
  expected := []string{
    "BFFFBBFRRR",
    "FFFBBBFRRR",
    "BBFFBBFRLL",
  }
  actual := ParseData(dat)
  for i := range expected {
    if expected[i] != actual[i] {
      t.Errorf("TestParseData; Slices did not match; expected: %s; actual: %s;\n", expected[i], actual[i])
    }
  }
}

func TestDeductSeat(t *testing.T) {
  type test struct {
    input          string
    expectedRow    int
    expectedCol    int
    expectedSeatID int
  }

  testCases := []test{
    {input: "BFFFBBFRRR", expectedRow: 70,  expectedCol: 7, expectedSeatID: 567},
    {input: "FFFBBBFRRR", expectedRow: 14,  expectedCol: 7, expectedSeatID: 119},
    {input: "BBFFBBFRLL", expectedRow: 102, expectedCol: 4, expectedSeatID: 820},
  }

  for _, testCase := range testCases {
    actualRow, actualCol, actualSeatID := DeductSeat(testCase.input)
    if testCase.expectedRow != actualRow {
      t.Errorf("TestDeductSeat; Row did not match; expected: %v; actual: %v; input: %v\n", testCase.expectedRow, actualRow, testCase.input)
    }
    if testCase.expectedCol != actualCol {
      t.Errorf("TestDeductSeat; Column did not match; expected: %v; actual: %v; input: %v\n", testCase.expectedCol, actualCol, testCase.input)
    }
    if testCase.expectedSeatID != actualSeatID {
      t.Errorf("TestDeductSeat; Seat ID did not match; expected: %v; actual: %v; input: %v\n", testCase.expectedSeatID, actualSeatID, testCase.input)
    }
  }
}

func TestFindHighestSeatID(t *testing.T) {
  input    := []int {567,119,820}
  expected := 820
  actual   := FindHighestSeatID(input)
  if expected != actual {
    t.Errorf("TestFindHighestSeatID; expected: %v; actual: %v; input: %v;\n", expected, actual, input)
  }
}

func TestFindMissingSeatID(t *testing.T) {
  input    := []int {10,15,8,14,7,13,11,9}
  expected := 12
  actual   := FindMissingSeatID(input)
  if expected != actual {
    t.Errorf("TestFindMissingSeatID; expected: %v; actual: %v; input: %v;\n", expected, actual, input)
  }
}