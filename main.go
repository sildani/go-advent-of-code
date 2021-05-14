package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sildani/go-advent-of-code/2020/one"
	"github.com/sildani/go-advent-of-code/2020/three"
	"github.com/sildani/go-advent-of-code/2020/two"
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

func dayTwo() {
	var partOneRulesValidPasswords int
	var partTwoRulesValidPasswords int
	var passwordsChecked int
	dat, err := ioutil.ReadFile("./2020/two/input/two.txt")
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
	dat, err := ioutil.ReadFile("./2020/three/input/three.txt")
	check(err)
	terrain := strings.Split(string(dat), "\n")
	treeCount := three.FindTreesInTerrain(terrain, []int{3, 1})
	fmt.Printf("dayThree Part One; treeCount: %v;\n", treeCount)
}

func main() {
	dayOne()
	dayTwo()
	dayThree()
}
