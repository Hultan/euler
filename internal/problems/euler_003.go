package euler

import (
	"fmt"
<<<<<<< HEAD:internal/euler/euler_003.go
	tools2 "github.com/hultan/euler/internal/tools"
=======
	"github.com/hultan/euler/internal/tools"
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/euler_003.go
	"time"
)

type Euler003 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler003() *Euler003 {
	e := new(Euler003)
	e.Name = "003 : Largest prime factor"
	e.Description ="The prime factors of 13195 are 5, 7, 13 and 29.\n\nWhat is the largest prime factor of the number 600851475143?"
	e.CorrectAnswer = "6857"
	return e
}

func (e *Euler003) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler003) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler003) Solve() string {
	start:=time.Now()
	var answer uint64 = 0

	primeFactor := tools2.NewPrimeFactors()
	factors := primeFactor.GetPrimeFactors(600851475143)
	answer = factors[len(factors)-1]

<<<<<<< HEAD:internal/euler/euler_003.go
	time:= tools2.TimeTrack(start)
=======
	time:= tools.TimeTrack(start)
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/euler_003.go
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}

