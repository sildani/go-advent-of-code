package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sildani/go-advent-of-code/2020/one"
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
	dat, err := ioutil.ReadFile("./2020/one/input/one.txt")
	check(err)
	input := convertByteSliceToIntSlice(dat)
	fmt.Printf("one.Eval: %v\n", one.Eval(input))
	fmt.Printf("one.EvalPartTwo: %v\n", one.EvalPartTwo(input))
}

func main() {
	dayOne()
}
