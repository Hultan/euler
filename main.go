package main

import "github.com/hultan/euler/problems"
import "fmt"

func main() {
	runAllProblems := false
	var euler problems.Problem

	euler = problems.NewEuler003()
	euler.PrintDescription()
	fmt.Printf("Result: %v\n", euler.Solve())
	fmt.Printf("Correct answer: %v\n", euler.GetAnswer())

	// All problems
	if runAllProblems {
		euler = problems.NewEuler001()
		euler.PrintDescription()
		fmt.Printf("Result: %v\n", euler.Solve())
		fmt.Printf("Correct answer: %v\n", euler.GetAnswer())

		euler = problems.NewEuler002()
		euler.PrintDescription()
		fmt.Printf("Result: %v\n", euler.Solve())
		fmt.Printf("Correct answer: %v\n", euler.GetAnswer())
	}
}
