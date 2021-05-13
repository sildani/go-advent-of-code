package two

import (
  "fmt"
  "strings"
  "regexp"
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
  return fmt.Sprintf("%s{%s,%s}", x[1], y[0], y[1])
}

func PasswordAudit(policy string, password string) bool {
  policyRegex := GeneratePolicyRegex(policy)
  matched, _ := regexp.MatchString(policyRegex, password)
  fmt.Printf("PasswordAudit; policyRegex: %s; password: %s; matched: %v;\n", policyRegex, password, matched)
  return matched
}
