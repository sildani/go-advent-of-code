package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"

	"github.com/sildani/go-advent-of-code/2020"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./2020/input/one.txt")
	check(err)
	stringInput := strings.Split(string(dat), "\n")
	var input []int
	for _, str := range stringInput {
		x, e := strconv.Atoi(strings.Replace(str, "\r","",-1))
		check(e)
		input = append(input, x)
	}
	fmt.Printf("one.Eval: %v\n", one.Eval(input))
}
