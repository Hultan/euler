package euler

import (
	"fmt"
<<<<<<< HEAD:internal/euler/euler_007.go
	tools2 "github.com/hultan/euler/internal/tools"
=======
	"github.com/hultan/euler/internal/tools"
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/euler_007.go
	"math/big"
	"time"
)

type Euler007 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler007() *Euler007 {
	e := new(Euler007)
	e.Name = "10001st prime"
	e.Description ="By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.\n\nWhat is the 10 001st prime number?"
	e.CorrectAnswer = "104743"
	return e
}

func (e *Euler007) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler007) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler007) Solve() string {
	start:=time.Now()
	var answer int64 = 0
	var i int64
	var numberOfPrimes int

	for i=1;numberOfPrimes<10001;i++ {
		if big.NewInt(i).ProbablyPrime(0) {
			answer=i
			numberOfPrimes++
		}
	}

<<<<<<< HEAD:internal/euler/euler_007.go
	time:= tools2.TimeTrack(start)
=======
	time:= tools.TimeTrack(start)
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/euler_007.go
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}


