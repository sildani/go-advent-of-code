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

func IsPasswordValid(policy string, password string) bool {
  policyParts := strings.Split(policy, " ")

  countRange := strings.Split(policyParts[0], "-")
  lowerRange, _ := strconv.Atoi(countRange[0])
  upperRange, _ := strconv.Atoi(countRange[1])

  letter := policyParts[1]

  count := strings.Count(password, letter)
  valid := count >= lowerRange && count <= upperRange

  fmt.Printf("PasswordAudit; policy: %v; password: %v; valid: %v;\n", policy, password, valid)
  
  return valid
}
