package euler

import (
	"fmt"
<<<<<<< HEAD:internal/euler/euler_005.go
	tools2 "github.com/hultan/euler/internal/tools"
=======
	"github.com/hultan/euler/internal/tools"
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/euler_005.go
	"time"
)

type Euler005 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler005() *Euler005 {
	e := new(Euler005)
	e.Name = "Smallest multiple"
	e.Description ="2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.\n\nWhat is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?"
	e.CorrectAnswer = "232792560"
	return e
}

func (e *Euler005) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler005) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler005) Solve() string {
	start:=time.Now()
	var answer uint64 = 0
	var result = [][2]uint64{{2,1}}
	var i uint64

	tool := tools2.NewPrimeFactors()

	for i=3;i<=20;i++ {
		wantFactors := tool.GetPrimeFactorCount(i)
		for _, wantFactor := range wantFactors {
			index, haveFactor := tool.GetPrimeFactor(result, wantFactor[0])

			if index>=0 && haveFactor[1] < wantFactor[1] {
				result[index][1] = wantFactor[1]
			} else if index<0 {
				result = append(result, wantFactor)
			}
		}
	}

	answer = tool.CalculatePrimeFactors(result)
<<<<<<< HEAD:internal/euler/euler_005.go
	time:= tools2.TimeTrack(start)
=======
	time:= tools.TimeTrack(start)
>>>>>>> bfd95881f1c854e1634d0acbc087d5e4994b7147:internal/problems/euler_005.go
	e.Answer = fmt.Sprintf("%v (%s)", answer, time)

	return e.Answer
}
