package problems

import (
	"fmt"
	"github.com/hultan/euler/internal/tools"
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

	primeFactor := tools.NewPrimeFactors()
	factors := primeFactor.GetPrimeFactors(600851475143)
	answer = factors[len(factors)-1]

	time:= tools.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}

