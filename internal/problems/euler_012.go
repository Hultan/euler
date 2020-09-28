package problems

import (
	"fmt"
	"github.com/hultan/euler/internal/tools"
	"time"
)

type Euler012 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler012() *Euler012 {
	e := new(Euler012)
	e.Name = "name"
	e.Description = "description"
	e.CorrectAnswer = "answer"
	return e
}

func (e *Euler012) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler012) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler012) Solve() string {
	start := time.Now()
	var answer uint64 = 0
	var count uint64 = 0
	var triangleNumber uint64 = 0
	var divisors int = 0
	divisorTool := tools.NewDivisors()

	for divisors < 500 {
		count++
		triangleNumber += count
		divisors = divisorTool.NumberOfDivisors(triangleNumber)
	}

	answer = triangleNumber
	time := tools.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}
