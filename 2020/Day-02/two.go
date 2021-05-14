package two

import (
  "fmt"
  "strconv"
  "strings"
)

func ParseData(data []byte) map[string][]string {
  retval := make(map[string][]string)
  stringData := strings.Split(string(data), "\n")
  for _, line := range stringData {
    if len(line) > 0 {
      lineData := strings.Split(line, ": ")
      retval[lineData[0]] = append(retval[lineData[0]], lineData[1])
    }
  }
  fmt.Printf("retval: %v;\n", retval)
  return retval
}

func GeneratePolicyRegex(policy string) string {
  x := strings.Split(policy, " ")
  y := strings.Split(x[0], "-")
  fmt.Printf("GeneratePolicyRegex; x: %s; y: %s;\n", x, y)
  return fmt.Sprintf("[%s]{%s,%s}", x[1], y[0], y[1])
}

func IsPasswordValid(policy string, password string, ruleset RuleSet) bool {
  var valid bool

  policyParts := strings.Split(policy, " ")
  countRange := strings.Split(policyParts[0], "-")
  lowerRange, _ := strconv.Atoi(countRange[0])
  upperRange, _ := strconv.Atoi(countRange[1])

  letter := policyParts[1]

  count := strings.Count(password, letter)

  if ruleset == PartOneRuleSet {
    valid = count >= lowerRange && count <= upperRange
  }

  if ruleset == PartTwoRuleSet {
    var letterAtFirstEligiblePosition bool
    var letterAtSecondEligiblePosition bool
    if len(password) >= lowerRange {
      letterAtFirstEligiblePosition = string(password[lowerRange-1]) == letter
    }
    if len(password) >= upperRange {
      letterAtSecondEligiblePosition = string(password[upperRange-1]) == letter
    }
    valid = (letterAtFirstEligiblePosition || letterAtSecondEligiblePosition) && (letterAtFirstEligiblePosition != letterAtSecondEligiblePosition)
  }

  fmt.Printf("PasswordAudit; policy: %v; password: %v; ruleset: %v; valid: %v;\n", policy, password, ruleset, valid)
  return valid
}
