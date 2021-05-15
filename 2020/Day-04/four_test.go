package four

import (
  "fmt"
  "testing"
)

func assertInts(expected int, actual int, t *testing.T) {
  if expected != actual {
    t.Errorf("expected %v, got %v\n", expected, actual)
  }
}

func assertPassportMaps(expected map[string]string, actual map[string]string, t *testing.T) {
  if len(expected) != len(actual) {
    t.Errorf("expected %v, got %v\n", len(expected), len(actual))
  }

  for k, e := range expected {
    if e != actual[k] {
      t.Errorf("expected %s, got %s\n", e, actual[k])
    }
  }
}

func assertBools(expected bool, actual bool, t *testing.T) {
  if expected != actual {
    t.Errorf("expected %v, got %v\n", expected, actual)
  }
}

func TestParsePassportData(t *testing.T) {
  dat := []byte(
    "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\r\n" +
      "byr:1937 iyr:2017 cid:147 hgt:183cm\r\n" +
      "\r\n" +
      "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\r\n" +
      "hcl:#cfa07d byr:1929\r\n" +
      "\r\n" +
      "hcl:#ae17e1 iyr:2013\r\n" +
      "eyr:2024\r\n" +
      "ecl:brn pid:760753108 byr:1931\r\n" +
      "hgt:179cm\r\n" +
      "\r\n" +
      "hcl:#cfa07d eyr:2025 pid:166559648\r\n" +
      "iyr:2011 ecl:brn hgt:59in")

  passports := ParsePassportData(dat)
  assertInts(4, len(passports), t)

  expectedPassports := []map[string]string{
    {"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"},
    {"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"},
    {"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"},
    {"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"},
  }
  for i, passport := range passports {
    expected := expectedPassports[i]
    actual := passport
    assertPassportMaps(expected, actual, t)
  }
}

func TestIsPassportValid(t *testing.T) {
  type test struct {
    passport           map[string]string
    ignoreCountryField bool
    fieldValidation    bool
    expectedValidity   bool
  }

  tests := []test{
    {passport: map[string]string{"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"}, ignoreCountryField: true, fieldValidation: false, expectedValidity: true},
    {passport: map[string]string{"iyr": "2013", "ecl": "amb", "cid": "350", "eyr": "2023", "pid": "028048884", "hcl": "#cfa07d", "byr": "1929"}, ignoreCountryField: true, fieldValidation: false, expectedValidity: false},
    {passport: map[string]string{"hcl": "#ae17e1", "iyr": "2013", "eyr": "2024", "ecl": "brn", "pid": "760753108", "byr": "1931", "hgt": "179cm"}, ignoreCountryField: true, fieldValidation: false, expectedValidity: true},
    {passport: map[string]string{"hcl": "#cfa07d", "eyr": "2025", "pid": "166559648", "iyr": "2011", "ecl": "brn", "hgt": "59in"}, ignoreCountryField: true, fieldValidation: false, expectedValidity: false},
    {passport: map[string]string{"ecl": "prp", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd", "byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm"}, ignoreCountryField: true, fieldValidation: true, expectedValidity: false},
    {passport: map[string]string{"eyr": "1972", "cid": "100", "hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926"}, ignoreCountryField: true, fieldValidation: true, expectedValidity: false},
    {passport: map[string]string{"iyr": "2019", "hcl": "#602927", "eyr": "1967", "hgt": "170cm", "ecl": "grn", "pid": "012533040", "byr": "1946"}, ignoreCountryField: true, fieldValidation: true, expectedValidity: false},
    {passport: map[string]string{"pid": "087499704", "hgt": "74in", "ecl": "grn", "iyr": "2012", "eyr": "2030", "byr": "1980", "hcl": "#623a2f"}, ignoreCountryField: true, fieldValidation: true, expectedValidity: true},
    {passport: map[string]string{"eyr": "2029", "ecl": "blu", "cid": "129", "byr": "1989", "iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165cm"}, ignoreCountryField: true, fieldValidation: true, expectedValidity: true},
    {passport: map[string]string{"eyr": "2029", "ecl": "blu", "cid": "129", "byr": "1989", "iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165cm"}, ignoreCountryField: false, fieldValidation: true, expectedValidity: true},
    {passport: map[string]string{"eyr": "2029", "ecl": "blu", "cid": "129", "byr": "1989", "iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165mm"}, ignoreCountryField: false, fieldValidation: true, expectedValidity: false},
  }

  for _, test := range tests {
    fmt.Printf("TestIsValidPassport; Checking Passport: %v; Expect: %v\n", test.passport, test.expectedValidity)
    assertBools(test.expectedValidity, IsPassportValid(test.passport, test.ignoreCountryField, test.fieldValidation), t)
  }

}
