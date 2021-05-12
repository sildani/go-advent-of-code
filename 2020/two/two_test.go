package two

import (
  "reflect"
  "testing"
)

func assertInts(expected int, actual int, t *testing.T) {
  if expected != actual {
    t.Errorf("expected %v, got %v\n", expected, actual)
  }
}

func assertStringSlices(expected []string, actual []string, t *testing.T) {
  if !reflect.DeepEqual(expected, actual) {
    t.Errorf("expected %v, got %v\n", expected, actual)
  }
}
func TestParseData(t *testing.T) {
  data := []byte("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc")
  parsedData := ParseData(data)
  
  assertInts(3, len(parsedData), t)
  assertStringSlices([]string{"abcde"}, parsedData["1-3 a"], t)
  assertStringSlices([]string{"cdefg"}, parsedData["1-3 b"], t)
  assertStringSlices([]string{"ccccccccc"}, parsedData["2-9 c"], t)
}