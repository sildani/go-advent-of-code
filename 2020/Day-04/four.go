package four

import (
  "fmt"
  "strconv"
  "strings"
)

func ParsePassportData(dat []byte) []map[string]string {
  datAsString := string(dat)
  datAsString = strings.ReplaceAll(datAsString, "\r", "")
  datAsStringSlice := strings.Split(datAsString, "\n")
  var passportData []map[string]string
  passport := make(map[string]string)
  done := false
  for _, line := range datAsStringSlice {
    done = len(line) == 0
    if !done {
      keyValuePairs := strings.Split(line, " ")
      for _, kvp := range keyValuePairs {
        kvpSliced := strings.Split(kvp, ":")
        key := kvpSliced[0]
        value := kvpSliced[1]
        passport[key] = value
      }
    } else {
      passportData = append(passportData, passport)
      passport = make(map[string]string)
    }
  }
  if len(passport) > 0 {
    passportData = append(passportData, passport)
  }
  fmt.Printf("ParsePassportData; passportData: %v;\n", passportData)
  return passportData
}

func IsPassportValid(passport map[string]string, ignoreCountryField bool, fieldValidation bool) bool {
  valid := true

  fields := map[string]string{
    "byr": "Birth Year",
    "iyr": "Issue Year",
    "eyr": "Expiration Date",
    "hgt": "Height",
    "hcl": "Hair Color",
    "ecl": "Eye Color",
    "pid": "Passport ID",
    "cid": "Country ID",
  }

  for field := range fields {
    if valid {
      if !(ignoreCountryField && field == "cid") {
        if len(passport[field]) == 0 {
          valid = false
        }
        if fieldValidation {
          switch field {
          case "byr":
            valid = isValidYear(passport[field], 1920, 2002)
          case "iyr":
            valid = isValidYear(passport[field], 2010, 2020)
          case "eyr":
            valid = isValidYear(passport[field], 2020, 2030)
          case "hgt":
            valid = isValidHeight(passport[field])
          case "hcl":
            valid = isValidColor(passport[field])
          case "ecl":
            valid = isValidEyeColor(passport[field])
          case "pid":
            valid = isValidPassportID(passport[field])
          case "cid":
            if !ignoreCountryField {
              valid = isValidCountryID(passport[field])
            }
          }
        }
      }
    }
    fmt.Printf("IsPassportValid; field: %v; valid: %v;\n", field, valid)
  }
  fmt.Printf("IsPassportValid; passport: %v; valid: %v;\n", passport, valid)
  return valid
}

func isValidYear(value string, lowerBound int, upperBound int) bool {
  valid := true
  valueAsInt, _ := strconv.Atoi(value)
  if len(value) != 4 || valueAsInt < lowerBound || valueAsInt > upperBound {
    valid = false
  }
  // fmt.Printf("isValidYear; value: %v; valid: %v;\n", value, valid)
  return valid
}

func isValidHeight(value string) bool {
  valid := len(value) > 2
  if valid {
    runes := []rune(value)
    height, _ := strconv.Atoi(string(runes[0 : len(value)-2]))
    unit := string(runes[len(value)-2 : len(value)])
    switch unit {
    case "cm":
      valid = height >= 150 && height <= 193
    case "in":
      valid = height >= 59 && height <= 76
    default:
      valid = false
    }
  }
  // fmt.Printf("isValidHeight; value: %v; valid: %v;\n", value, valid)
  return valid
}

func isValidColor(value string) bool {
  valid := len(value) == 7
  if valid {
    runes := []rune(value)
    hex := string(runes[1:len(value)])
    _, err := strconv.ParseUint(hex, 16, 64)
    valid = err == nil
  }
  // fmt.Printf("isValidColor; value: %v; valid: %v;\n", value, valid)
  return valid
}

func isValidEyeColor(value string) bool {
  valid := len(value) == 3
  if valid {
    validColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
    valid = false // false until proven otherwise
    for _, validColor := range validColors {
      if !valid {
        valid = value == validColor
      }
    }
  }
  fmt.Printf("isValidEyeColor; value: %v; valid: %v;\n", value, valid)
  return valid
}

func isValidPassportID(value string) bool {
  valid := len(value) == 9
  if valid {
    _, err := strconv.Atoi(value)
    valid = err == nil
  }
  // fmt.Printf("isValidPassportID; value: %v; valid: %v;\n", value, valid)
  return valid
}

func isValidCountryID(value string) bool {
  valid := true
  // fmt.Printf("isValidCountryID; value: %v; valid: %v;\n", value, valid)
  return valid
}
