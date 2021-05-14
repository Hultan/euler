package euler

import (
	"fmt"
	tools2 "github.com/hultan/euler/internal/tools"
	"time"
)

type Euler000 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler000() *Euler000 {
	e := new(Euler000)
	e.Name = "name"
	e.Description ="description"
	e.CorrectAnswer = "answer"
	return e
}

func (e *Euler000) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler000) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler000) Solve() string {
	start:=time.Now()
	var answer int64 = 0

	// Problem here

	time:= tools2.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}

