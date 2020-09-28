package problems

import (
	"fmt"
	"github.com/hultan/euler/internal/tools"
	"time"
)

type Euler004 struct {
	Name          string
	Description   string
	Answer        string
	CorrectAnswer string
}

func NewEuler004() *Euler004 {
	e := new(Euler004)
	e.Name = "Largest palindrome product"
	e.Description ="A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.\n\nFind the largest palindrome made from the product of two 3-digit numbers."
	e.CorrectAnswer = "906609"
	return e
}

func (e *Euler004) PrintDescription() {
	fmt.Println()
	fmt.Println(e.Name)
	fmt.Println("----------------------------")
	fmt.Println(e.Description)
	fmt.Println()
}

func (e *Euler004) GetAnswer() string {
	return e.CorrectAnswer
}

func (e *Euler004) Solve() string {
	start:=time.Now()
	var answer int64
	var maxPalindrome int64
	var i,j int64

	t := tools.NewNumberFunctions()

	for i=999;i>2;i-- {
		for j=i-1;j>1;j-- {
			answer = i*j
			if answer>maxPalindrome && t.IsPalindrome(answer) {
				maxPalindrome = answer
				break
			}
		}
	}

	time:= tools.TimeTrack(start)
	e.Answer = fmt.Sprintf("%v (%s)", maxPalindrome, time)

	return e.Answer
}

