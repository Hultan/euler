package main

import (
	"github.com/hultan/euler/problems"
)
import "fmt"

func main() {
	runAllProblems := false
	solvedProblems := []int{1,2,3,4,5,6,7,8}
	currentProblem := 8

	if runAllProblems {
		for _, value := range solvedProblems {
			runProblem(value)
		}
	} else {
		runProblem(currentProblem)
	}
}

func getProblem(number int) problems.Problem  {
	switch number {
	case 1:
		return problems.NewEuler001()
	case 2:
		return problems.NewEuler002()
	case 3:
		return problems.NewEuler003()
	case 4:
		return problems.NewEuler004()
	case 5:
		return problems.NewEuler005()
	case 6:
		return problems.NewEuler006()
	case 7:
		return problems.NewEuler007()
	case 8:
		return problems.NewEuler008()
	default:
		return nil
	}
}

func runProblem(number int) {
	euler := getProblem(number)
	euler.PrintDescription()
	fmt.Printf("Result         : %v\n", euler.Solve())
	fmt.Printf("Correct answer : %v\n", euler.GetAnswer())
}