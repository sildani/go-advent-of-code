package one

func Eval(input []int) int {
  for i, focusNum := range input {
    for j, inputNum := range input {
      if i != j && focusNum+inputNum == 2020 {
        return focusNum * inputNum
      }
    }
  }
  return 0
}

func EvalPartTwo(input []int) int {
  for i, focusNum := range input {
    for j, inputNum := range input {
      for k, inputNumTwo := range input {
        if i != j && j != k && focusNum+inputNum+inputNumTwo == 2020 {
          return focusNum * inputNum * inputNumTwo
        }
      }
    }
  }
  return 0
}
