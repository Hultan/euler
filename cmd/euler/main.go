package main

import (
	"github.com/hultan/euler/internal/problems"
)
import "fmt"

func main() {
	runAllProblems := false
	solvedProblems := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	currentProblem := 12

	if runAllProblems {
		for _, value := range solvedProblems {
			runProblem(value)
		}
	} else {
		runProblem(currentProblem)
	}
}

func getProblem(number int) problems.Problem {
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
	case 9:
		return problems.NewEuler009()
	case 10:
		return problems.NewEuler010()
	case 11:
		return problems.NewEuler011()
	case 12:
		return problems.NewEuler012()
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