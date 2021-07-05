package main

import (
<<<<<<< HEAD:main.go
	euler2 "github.com/hultan/euler/internal/euler"
=======
	"github.com/hultan/euler/internal/problems"
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:cmd/euler/main.go
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

<<<<<<< HEAD:main.go
func getProblem(number int) euler2.Problem {
=======
func getProblem(number int) problems.Problem {
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:cmd/euler/main.go
	switch number {
	case 1:
		return euler2.NewEuler001()
	case 2:
		return euler2.NewEuler002()
	case 3:
		return euler2.NewEuler003()
	case 4:
		return euler2.NewEuler004()
	case 5:
		return euler2.NewEuler005()
	case 6:
		return euler2.NewEuler006()
	case 7:
		return euler2.NewEuler007()
	case 8:
		return euler2.NewEuler008()
	case 9:
		return euler2.NewEuler009()
	case 10:
		return euler2.NewEuler010()
	case 11:
		return euler2.NewEuler011()
	case 12:
		return euler2.NewEuler012()
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