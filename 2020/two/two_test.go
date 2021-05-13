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

func assertStrings(expected string, actual string, t *testing.T) {
  if expected != actual {
    t.Errorf("expected %v, got %v\n", expected, actual)
  }
}

func assertBooleans(expected bool, actual bool, t *testing.T) {
  if expected != actual {
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

func TestParseDataWithLineFeedAtEnd(t *testing.T) {
  data := []byte("1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc\n")
  parsedData := ParseData(data)

  assertInts(3, len(parsedData), t)
  assertStringSlices([]string{"abcde"}, parsedData["1-3 a"], t)
  assertStringSlices([]string{"cdefg"}, parsedData["1-3 b"], t)
  assertStringSlices([]string{"ccccccccc"}, parsedData["2-9 c"], t)
}

func TestParseDataSameRuleTwice(t *testing.T) {
  data := []byte("1-3 a: abcde\n1-3 a: cdefg")
  parsedData := ParseData(data)

  assertInts(1, len(parsedData), t)
  assertStringSlices([]string{"abcde", "cdefg"}, parsedData["1-3 a"], t)
}

func TestGeneratePolicyRegex(t *testing.T) {
  assertStrings("a{1,3}", GeneratePolicyRegex("1-3 a"), t)
  assertStrings("b{1,3}", GeneratePolicyRegex("1-3 b"), t)
  assertStrings("c{2,9}", GeneratePolicyRegex("2-9 c"), t)
}

func TestPasswordValid(t *testing.T) {
  assertBooleans(true, PasswordAudit("1-3 a", "abcde"), t)
  assertBooleans(false, PasswordAudit("1-3 b", "cdefg"), t)
  assertBooleans(true, PasswordAudit("2-9 c", "ccccccccc"), t)
}
