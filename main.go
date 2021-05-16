package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "strings"

  one "github.com/sildani/go-advent-of-code/2020/Day-01"
  two "github.com/sildani/go-advent-of-code/2020/Day-02"
  three "github.com/sildani/go-advent-of-code/2020/Day-03"
  four "github.com/sildani/go-advent-of-code/2020/Day-04"
  five "github.com/sildani/go-advent-of-code/2020/Day-05"
  six "github.com/sildani/go-advent-of-code/2020/Day-06"
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
  slopes := [][]int{
    {1, 1},
    {3, 1},
    {5, 1},
    {7, 1},
    {1, 2},
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

func dayFour() {
  dat, err := ioutil.ReadFile("./2020/Day-04/resources/input.txt")
  check(err)
  passports := four.ParsePassportData(dat)

  validPassportCount := 0
  for _, passport := range passports {
    if four.IsPassportValid(passport, true, false) {
      validPassportCount++
    }
  }

  validPassportStrictRulesCount := 0
  for _, passport := range passports {
    if four.IsPassportValid(passport, true, true) {
      validPassportStrictRulesCount++
    }
  }

  fmt.Printf("dayThree Part One; treeCount: %v;\n", validPassportCount)
  fmt.Printf("dayThree Part Two; treeCount: %v;\n", validPassportStrictRulesCount)
}

func dayFive() {
  dat, err := ioutil.ReadFile("./2020/Day-05/resources/input.txt")
  check(err)
  boardingPassses := five.ParseData(dat)

  var seatIDs []int
  for _, boardingPass := range boardingPassses {
    _, _, seatID := five.DeductSeat(boardingPass)
    seatIDs = append(seatIDs, seatID)
  }

  highestSeatID := five.FindHighestSeatID(seatIDs)
  missingSeatID := five.FindMissingSeatID(seatIDs)
  
  fmt.Printf("dayFive Part One; Highest seat ID: %v;\n", highestSeatID)
  fmt.Printf("dayFive Part Two; Missing seat ID: %v;\n", missingSeatID)
}

func daySix() {
  dat, err := ioutil.ReadFile("./2020/Day-06/resources/input.txt")
  check(err)
  answersList := six.ParseData(dat, "Part One")
  sumOfUniqueAnswers := six.GetSumOfUniqueAnswers(answersList)
  answersList = six.ParseData(dat, "Part Two")
  sumOfConsistentAnswers := six.GetSumOfConsistentAnswers(answersList)
  fmt.Printf("daySix Part One; Sum of unique answers: %v;\n", sumOfUniqueAnswers)
  fmt.Printf("daySix Part Two; Sum of consistent answers: %v;\n", sumOfConsistentAnswers)
}

func main() {
  dayOne()
  dayTwo()
  dayThree()
  dayFour()
  dayFive()
  daySix()
}
