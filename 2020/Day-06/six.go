package six

import "strings"

func getUniqueList(inputList []string) []string {
  exists := struct{}{}
  set := make(map[string]struct{})
  for _, item := range inputList {
    set[item] = exists
  }
  var uniqueList []string
  for item := range set {
    uniqueList = append(uniqueList, item)
  }
  return uniqueList
}

func ParseData(dat []byte, part string) [][]string {
  datAsString := string(dat)
  datAsString = strings.ReplaceAll(datAsString, "\r", "")
  datAsStringSlice := strings.Split(datAsString, "\n")

  var parsedData [][]string
  var currentSlice []string
  for _, line := range datAsStringSlice {
    if line == "" && len(currentSlice) > 0 {
      parsedData = append(parsedData, currentSlice)
      currentSlice = []string{}
    } else {
      if part == "Part One" {
        for _, rune := range line {
          currentSlice = append(currentSlice, string(rune))
        }
      }
      if part == "Part Two" {
        currentSlice = append(currentSlice, line)
      }
    }
  }
  if len(currentSlice) > 0 {
    parsedData = append(parsedData, currentSlice)
  }

  return parsedData
}

func GetSumOfUniqueAnswers(answersList [][]string) int {
  var sum int
  for _, answers := range answersList {
    uniqueAnswers := getUniqueList(answers)
    // fmt.Printf("answers: %v; uniqueAnswers: %v;\n", answers, uniqueAnswers)
    sum = sum + len(uniqueAnswers)
  }
  return sum
}

func getConsistentList(inputList []string) []string {
  consistentList := []string{}

  tracker := make(map[string]int)
  for _, item := range inputList {
    for _, rune := range item {
      tracker[string(rune)]++
    }
  }
  // fmt.Printf("tracker: %v;\n", tracker)
  for k, v := range tracker {
    if v == len(inputList) {
      consistentList = append(consistentList, k)
    }
  }

  return consistentList
}

func GetSumOfConsistentAnswers(answersList [][]string) int {
  var sum int
  for _, answers := range answersList {
    consistentAnswers := getConsistentList(answers)
    // fmt.Printf("answers: %v; uniqueAnswers: %v;\n", answers, uniqueAnswers)
    sum = sum + len(consistentAnswers)
  }
  return sum
}
