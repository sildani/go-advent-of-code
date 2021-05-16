package six

import "testing"

func TestGetUniqueList(t *testing.T) {
  inputList := []string{"A","A","B","B","C","D"}
  expectedList := []string{"A","B","C","D"}
  actualList := getUniqueList(inputList)

  expectedListSize := len(expectedList)
  actualListSize := len(actualList)
  if expectedListSize != actualListSize {
    t.Errorf("TestGetUniqueList; expected list size: %v; actual list size: %v;\n", expectedListSize, actualListSize)
  }

  for _, expected := range expectedList {
    found := false
    for _, actual := range actualList {
      if !found {
        found = expected == actual
      }
    }
    if !found {
      t.Errorf("TestGetUniqueList; expected to find \"%v\" in the actual list; list: %v;\n", expected, actualList)
    }
  }
}

func TestParseData(t *testing.T) {
  dat := []byte(
    "abc\r\n"+
    "\r\n"+
    "a\r\n"+
    "b\r\n"+
    "c\r\n"+
    "\r\n"+
    "ab\r\n"+
    "ac\r\n"+
    "\r\n"+
    "a\r\n"+
    "a\r\n"+
    "a\r\n"+
    "a\r\n"+
    "\r\n"+
    "b",
  )

  expectedList := [][]string{
    {"a","b","c"},
    {"a","b","c"},
    {"a","b","a","c"},
    {"a","a","a","a"},
    {"b"},
  }
  actualList := ParseData(dat, "Part One")
  checkParsedData(expectedList, actualList, t)

  expectedList = [][]string{
    {"abc"},
    {"a","b","c"},
    {"ab","ac"},
    {"a","a","a","a"},
    {"b"},
  }
  actualList = ParseData(dat, "Part Two")
  checkParsedData(expectedList, actualList, t)
}

func checkParsedData(expectedList [][]string, actualList [][]string, t *testing.T) {
  if len(expectedList) != len(actualList) {
    t.Errorf("TestParseData; number of answer lists didn't match expected; expected: %v; actual: %v;\n", expectedList, actualList)
  }

  for i := range expectedList {
    for j := range expectedList[i] {
      if expectedList[i][j] != actualList[i][j] {
        t.Errorf("TestParseData; items didn't match; expected: %v; actual: %v;\n", expectedList[i][j], actualList[i][j])
      }
    }
  }
}

func TestGetSumOfUniqueAnswers(t *testing.T) {
  inputList := [][]string{
    {"a","b","c"},
    {"a","b","c"},
    {"a","b","a","c"},
    {"a","a","a","a"},
    {"b"},
  }
  expected := 11
  actual := GetSumOfUniqueAnswers(inputList)
  if expected != actual {
    t.Errorf("TestGetSumOfUniqueAnsers; did not get expected sum; expected: %v; actual: %v;\n", expected, actual)
  }
}

func TestGetConsistentList(t *testing.T) {
  type test struct {
    inputList    []string
    expectedList []string
  }

  tests := []test{
    {inputList: []string{"AB","AC","AD","A"}, expectedList: []string{"A"}},
    {inputList: []string{"AB","C","D","EF"}, expectedList: []string{}},
    {inputList: []string{"F","FA"}, expectedList: []string{"F"}},
  }

  for _, test := range tests {
    actualList := getConsistentList(test.inputList)
    expectedListSize := len(test.expectedList)
    actualListSize := len(actualList)
    if expectedListSize != actualListSize {
      t.Errorf("TestGetUniqueList; expected list size: %v; actual list size: %v;\n", expectedListSize, actualListSize)
    }
    for _, expected := range test.expectedList {
      found := false
      for _, actual := range actualList {
        if !found {
          found = expected == actual
        }
      }
      if !found {
        t.Errorf("TestGetUniqueList; expected to find \"%v\" in the actual list; list: %v;\n", expected, actualList)
      }
    }
  }
}

func TestGetSumOfConsistentAnswers(t *testing.T) {
  inputList := [][]string{
    {"abc"},
    {"a","b","c"},
    {"ab","ac"},
    {"a","a","a","a"},
    {"b"},
  }
  expected := 6
  actual := GetSumOfConsistentAnswers(inputList)
  if expected != actual {
    t.Errorf("TestGetSumOfConsistentAnswers; did not get expected sum; expected: %v; actual: %v;\n", expected, actual)
  }
}