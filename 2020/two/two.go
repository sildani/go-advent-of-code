package two

import (
  "strings"
  "fmt"
)

func ParseData(data []byte) map[string][]string {
  retval := make(map[string][]string)
  stringData := strings.Split(string(data), "\n")
  for _, line := range stringData {
    lineData := strings.Split(line, ": ")
    retval[lineData[0]] = append(retval[lineData[0]], lineData[1])
  }
  fmt.Printf("retval: %v;\n", retval)
  return retval
}