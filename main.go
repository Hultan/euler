package main

import "github.com/hultan/euler/problems"
import "fmt"

func main() {
	e := problems.NewEuler001()
	e.Start()
	fmt.Println(e.Sum.Value())
//	e.Start2()
//	fmt.Println(e.Sum.Value())
}
