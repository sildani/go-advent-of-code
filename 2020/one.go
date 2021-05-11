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
