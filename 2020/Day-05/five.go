package five

import (
  "fmt"
  "strings"
)

func ParseData(dat []byte) []string {
  datAsString := string(dat)
  datAsString = strings.ReplaceAll(datAsString, "\r", "")
  datAsStringSlice := strings.Split(datAsString, "\n")
  return datAsStringSlice
}

func DeductSeat(boardingPass string) (int, int, int) {
  var row int
  var col int
  var sid int

  inputRunes := []rune(boardingPass)
  rowInputRunes := inputRunes[0:7]
  colInputRunes := inputRunes[7:10]

  row = deductFromRangeAndRunes([]int{1, 128}, rowInputRunes)
  col = deductFromRangeAndRunes([]int{1, 8}, colInputRunes)
  sid = (row * 8) + col

  fmt.Printf("DeductSeat; boardingPass: %s; row: %v; col: %v; sid: %v;\n", boardingPass, row, col, sid)

  return row, col, sid
}

func FindHighestSeatID(seatList []int) int {
  var highestValue int
  for _, seat := range seatList {
    if seat > highestValue {
      highestValue = seat
    }
  }
  return highestValue
}

func deductFromRangeAndRunes(possiblities []int, inputRunes []rune) int {
  // fmt.Printf("DeductSeat; possiblities: %v;\n", possiblities)
  row := 0
  for _, inputRune := range inputRunes {
    x := string(inputRune)
    rangeCount := possiblities[1] - possiblities[0] + 1
    if x == "F" || x == "L" {
      // fmt.Printf(("DeductSeat; F matched; "))
      newLowerRange := possiblities[0]
      newUpperRange := possiblities[0] + (rangeCount / 2) - 1
      possiblities = []int{newLowerRange, newUpperRange}
    }
    if x == "B" || x == "R" {
      // fmt.Printf(("DeductSeat; B matched; "))
      newLowerRange := possiblities[1] - (rangeCount / 2) + 1
      newUpperRange := possiblities[1]
      possiblities = []int{newLowerRange, newUpperRange}
    }
    // fmt.Printf("inputRune: %s; rangeCount: %v; possiblities: %v;\n", string(inputRune), rangeCount, possiblities)
  }
  row = possiblities[0] - 1
  return row
}
