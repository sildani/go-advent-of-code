package two

import (
  "fmt"
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
