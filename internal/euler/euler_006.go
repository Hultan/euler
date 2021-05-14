package euler

import (
	"fmt"
	tools2 "github.com/hultan/euler/internal/tools"
	"time"
)

type Euler006 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler006() *Euler006 {
	e := new(Euler006)
	e.Name = "Sum square difference"
	e.Description = "The sum of the squares of the first ten natural numbers is,\n\nThe square of the sum of the first ten natural numbers is,\n\nHence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is\n\n.\n\nFind the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum."
	e.CorrectAnswer = "25164150"
	return e
}

func (e *Euler006) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler006) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler006) Solve() string {
	start := time.Now()
	var answer uint64 = 0

	var i uint64
	var sumOfSquares uint64
	var squareOfSums uint64

	for i = 1; i <= 100; i++ {
		sumOfSquares += i * i
		squareOfSums += i
	}
	squareOfSums *= squareOfSums
	answer = squareOfSums - sumOfSquares

	time := tools2.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}
