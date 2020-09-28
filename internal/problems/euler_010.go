package problems

import (
	"fmt"
	"github.com/hultan/euler/internal/tools"
	"math/big"
	"time"
)

type Euler010 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler010() *Euler010 {
	e := new(Euler010)
	e.Name = "Summation of primes"
	e.Description ="The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.\n\nFind the sum of all the primes below two million."
	e.CorrectAnswer = "answer"
	return e
}

func (e *Euler010) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler010) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler010) Solve() string {
	start:=time.Now()
	var answer int64 = 0
	var p int64

	for p=5;p<2000000;p+=2 {
		if p%3 == 0 {
			continue
		}
		if big.NewInt(p).ProbablyPrime(0) {
			answer += p
		}
	}

	answer += 5 // 2 and 3

	time:= tools.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}

