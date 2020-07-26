package problems

import (
	"fmt"
	"github.com/hultan/euler/tools"
	"time"
)

type Euler009 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler009() *Euler009 {
	e := new(Euler009)
	e.Name = "Special Pythagorean triplet"
	e.Description = "A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,\na2 + b2 = c2\n\nFor example, 32 + 42 = 9 + 16 = 25 = 52.\n\nThere exists exactly one Pythagorean triplet for which a + b + c = 1000.\nFind the product abc."
	e.CorrectAnswer = "answer"
	return e
}

func (e *Euler009) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler009) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler009) Solve() string {
	start := time.Now()
	var answer int64 = 0

	a, b, c := e.getPythagoranTriplet()
	answer = int64(a) * int64(b) * int64(c)

	time := tools.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}

func (e *Euler009) getPythagoranTriplet() (int, int, int) {
	for a := 1; a < 998; a++ {
		for b := 1; b < 998; b++ {
			c := 1000 - a - b

			if a*a+b*b == c*c {
				return a, b, c
			}
		}
	}
	return -1, -1, -1
}
