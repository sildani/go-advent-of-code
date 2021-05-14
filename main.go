package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"

  "github.com/sildani/go-advent-of-code/2020/Day-01"
  "github.com/sildani/go-advent-of-code/2020/Day-02"
  "github.com/sildani/go-advent-of-code/2020/Day-03"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func convertByteSliceToIntSlice(bytes []byte) []int {
  stringInput := strings.Split(string(bytes), "\n")
  var retval []int
  for _, str := range stringInput {
    x, e := strconv.Atoi(strings.Replace(str, "\r", "", -1))
    check(e)
    retval = append(retval, x)
  }
  return retval
}

func dayOne() {
  dat, err := ioutil.ReadFile("./2020/Day-01/resources/input.txt")
  check(err)
  input := convertByteSliceToIntSlice(dat)
  fmt.Printf("one.Eval: %v\n", one.Eval(input))
  fmt.Printf("one.EvalPartTwo: %v\n", one.EvalPartTwo(input))
}

func dayTwo() {
  var partOneRulesValidPasswords int
  var partTwoRulesValidPasswords int
  var passwordsChecked int
  dat, err := ioutil.ReadFile("./2020/Day-02/resources/input.txt")
  check(err)
  parsedData := two.ParseData(dat)
  for policy, passwordSlice := range parsedData {
    for _, password := range passwordSlice {
      passwordsChecked++
      if two.IsPasswordValid(policy, password, two.PartOneRuleSet) {
        partOneRulesValidPasswords++
      }
      if two.IsPasswordValid(policy, password, two.PartTwoRuleSet) {
        partTwoRulesValidPasswords++
      }
    }
  }
  fmt.Printf("dayTwo; partOneRulesValidPasswords: %v; passwordsChecked: %v;\n", partOneRulesValidPasswords, passwordsChecked)
  fmt.Printf("dayTwo; partTwoRulesValidPasswords: %v; passwordsChecked: %v;\n", partTwoRulesValidPasswords, passwordsChecked)
}

func dayThree() {
  dat, err := ioutil.ReadFile("./2020/Day-03/resources/input.txt")
  check(err)
  terrain := strings.Split(string(dat), "\n")
  treeCount := three.FindTreesInTerrain(terrain, []int{3, 1})
  fmt.Printf("dayThree Part One; treeCount: %v;\n", treeCount)

  treeCount = 0
  slopes := [][]int {
    {1,1,},
    {3,1,},
    {5,1,},
    {7,1,},
    {1,2,},
  }
  for _, slope := range slopes {
    if treeCount == 0 {
      treeCount = three.FindTreesInTerrain(terrain, slope)
    } else {
      treeCount = treeCount * three.FindTreesInTerrain(terrain, slope)
    }
  }
  fmt.Printf("dayThree Part Two; treeCount: %v;\n", treeCount)
}

func main() {
  dayOne()
  dayTwo()
  dayThree()
}
