package tools

type NumberFunctions struct {
}

func NewNumberFunctions() *NumberFunctions {
	return new(NumberFunctions)
}

func (n *NumberFunctions) IsPalindrome(number int64) bool {
	var temp int64
	var reverse int64

	temp = number

	// For Loop used in format of While Loop
	for {
		reverse = reverse*10 + number % 10
		number /= 10

		if number == 0 {
			break // Break Statement used to exit from loop
		}
	}

	return temp==reverse
}
