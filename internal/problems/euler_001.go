package euler

import (
	"fmt"
<<<<<<< HEAD:internal/euler/euler_001.go
	tools2 "github.com/hultan/euler/internal/tools"
=======
	"github.com/hultan/euler/internal/tools"
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/euler_001.go
	"time"
)

type Euler001 struct {
	Name          string
	Description   string
	CorrectAnswer string
	Answer        string
}

func NewEuler001() *Euler001 {
	e := new(Euler001)
	e.Name = "001 : Multiples of 3 and 5"
	e.Description ="If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.\n\nFind the sum of all the multiples of 3 or 5 below 1000."
	e.CorrectAnswer = "233168"
	return e
}

func (e *Euler001) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler001) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler001) Solve() string {
	start := time.Now()
	var answer int64 =0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			answer += int64(i)
		}
	}
	time := tools2.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}
