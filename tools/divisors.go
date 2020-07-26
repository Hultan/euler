package tools

type Divisors struct {
}

func NewDivisors() *Divisors {
	return new(Divisors)
}

// Usage :
// Calculate number of divisors of a given number
//   factors := GetPrimeFactors(36)			returns {2,2,3,2}
func (d *Divisors) NumberOfDivisors(number uint64) int {
	tool := NewPrimeFactors()
	factors := tool.GetPrimeFactorCount(number)
	//fmt.Printf("Number : %v\n", number)
	//fmt.Println("Factors : ", factors)
	num := 1
	for _, factor := range factors {
		num *= int(factor[1]) + 1
	}
	//fmt.Printf("NumberOfDivisors : %v\n", num)

	return num
}
