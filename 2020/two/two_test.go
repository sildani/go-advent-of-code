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
  assertStrings("[a]{1,3}", GeneratePolicyRegex("1-3 a"), t)
  assertStrings("[b]{1,3}", GeneratePolicyRegex("1-3 b"), t)
  assertStrings("[c]{2,9}", GeneratePolicyRegex("2-9 c"), t)
}

func TestIsPasswordValid(t *testing.T) {
  assertBooleans(true,  IsPasswordValid("1-3 a", "abcde", PartOneRuleSet), t)
  assertBooleans(false, IsPasswordValid("1-3 b", "cdefg", PartOneRuleSet), t)
  assertBooleans(true,  IsPasswordValid("2-9 c", "ccccccccc", PartOneRuleSet), t)
  assertBooleans(false, IsPasswordValid("8-9 c", "cccccccwn", PartOneRuleSet), t)
  assertBooleans(false, IsPasswordValid("6-15 z", "zznzzzzzzzzzzzzzzzzz", PartOneRuleSet), t)

  assertBooleans(true,  IsPasswordValid("1-3 a", "abcde", PartTwoRuleSet), t)
  assertBooleans(false,  IsPasswordValid("10-11 a", "abcde", PartTwoRuleSet), t)
  assertBooleans(false, IsPasswordValid("1-3 b", "cdefg", PartTwoRuleSet), t)
  assertBooleans(false,  IsPasswordValid("2-9 c", "ccccccccc", PartTwoRuleSet), t)
  assertBooleans(true,  IsPasswordValid("1-7 c", "acccccccc", PartTwoRuleSet), t)
}
