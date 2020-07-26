package tools

import "math"

type PrimeFactor struct {
}

func NewPrimeFactors() *PrimeFactor {
	return new(PrimeFactor)
}

// Usage :
//   factors := GetPrimeFactors(36)			returns {2,2,3,2}
func (p *PrimeFactor) GetPrimeFactors(number uint64) []uint64 {
	var result []uint64

	for number%2==0 {
		result = append(result, 2)
		number /= 2
	}
	factor := uint64(3)
	for factor*factor <=number {
		if number%factor ==0 {
			result = append(result, factor)
			number /= factor
		} else {
			factor += 2
		}
	}
	if number!=1 {
		result = append(result, number)
	}

	return result
}

// Usage:
//   factors := GetPrimeFactors(36)    		returns {{2,2}, {3,2}}
func (p *PrimeFactor) GetPrimeFactorCount(number uint64) [][2]uint64 {
	factors := p.GetPrimeFactors(number)
	var currentFactor, currentCount uint64
	var result [][2]uint64

	for _, value := range factors {
		if value == currentFactor {
			currentCount ++
		} else {
			if currentFactor!=0 {
				result = append(result, [2]uint64{currentFactor, currentCount})
			}
			currentFactor = value
			currentCount = 1
		}
	}
	result = append(result, [2]uint64{currentFactor, currentCount})

	return result
}

// Usage:
//   factors := GetPrimeFactors(36)    		returns {{2,2}, {3,2}}
//   factor := GetPrimeFactor(factors, 3) 	returns {3,2}
func (p *PrimeFactor) GetPrimeFactor(factors [][2]uint64, factor uint64) (index int,resultFactor [2]uint64) {
	for index,value := range factors {
		if value[0]==factor {
			return index, value
		}
	}
	return -1, [2]uint64{}
}

func (p *PrimeFactor) CalculatePrimeFactors(factors [][2]uint64) (result uint64) {
	result = 1
	for _,value := range factors {
		result = result * uint64(math.Pow(float64(value[0]), float64(value[1])))
	}

	return result
}
