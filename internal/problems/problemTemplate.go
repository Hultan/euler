package euler

import (
	"fmt"
<<<<<<< HEAD:internal/euler/problemTemplate.go
	tools2 "github.com/hultan/euler/internal/tools"
=======
	"github.com/hultan/euler/internal/tools"
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/problemTemplate.go
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

<<<<<<< HEAD:internal/euler/problemTemplate.go
	time:= tools2.TimeTrack(start)
=======
	time:= tools.TimeTrack(start)
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/problemTemplate.go
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}

