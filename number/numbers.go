//Package evaluate the operations with numbers and types
package number

import (
	"math"
)

// evaluate if two number are multiples theyself
// true if are multiples or false in other case
func IsMultiple(num, bas int) bool {
	return num%bas == 0
}

// evaluate if a number is even or not
// true if is even or false
func IsEven(num int) bool {
	return IsMultiple(num, 2)
}

// descompose the number in your values that are multiples and summarize
func descompose(num int) int {
	res := 0
	max := (num / 2) + 1
	for i := 1; i < max; i++ {
		if IsMultiple(num, i) {
			res += i
		}
	}
	return res
}

// evaluate if a number is prime(alone is divisible by 1 and yourself)
// true if is prime or false in other case
func IsPrime(num int) bool {
	return descompose(num) == 1
}

func IsEmirP(num int) bool {
	return IsPrime(num) && IsPrime(ReverseNumber(num))
}

// evaluate number by number the multiple in the first match return false
// in case the not found return true
func FastPrime(num int) bool {
	max := (num / 2) + 1
	for i := 2; i < max; i++ {
		if IsMultiple(num, i) {
			return false
		}
	}
	return true
}

// descompose a number used the prime factors and stored in the slice
func DescomposePrime(num int) []int {
	var digits []int
	div := 2
	for num >= 1 {
		if IsMultiple(num, div) {
			digits = append(digits, div)
			num /= div
		} else {
			div++
			if div > num {
				break
			}
		}
	}
	return digits
}

// evaluate if a number is perfect the summarize the yours divider are equal to
// yourself
func IsPerfect(num int) bool {
	return descompose(num) == num
}

// evaluate if two number are friends, the summarize the one is equal to other
// and reverse
func AreFriends(num1, num2 int) bool {
	return (descompose(num1) == num2) && (descompose(num2) == num1)
}

// count the number de digits under base 10
func CountDigits(num int) int {
	var cont int = 0
	for num > 0 {
		cont++
		num /= 10
	}
	return cont
}

// descompose a number in your digits and return in a array
func DescomposeDigits(num int) []int {
	var digits []int
	for num >= 1 {
		digits = append(digits, num%10)
		num /= 10
	}
	return digits
}

//evaluate a number and define if number is perfect,abundant or reflective
func Type(num int) string {
	var value = descompose(num)
	if num == value {
		return "perfect"
	} else if num > value {
		return "abundant"
	} else {
		return "reflective"
	}
}

//Calculate a perfect number based in equation (2^n)*((2^n)-1)
func CalculatePerfect(num int) int {
	return int(math.Pow(2, float64(num)) * (math.Pow(2, float64(num)) - 1))
}

//Reverse a number ej 12 return 21
func ReverseNumber(num int) int {
	var digits []int = DescomposeDigits(num)
	var max int = CountDigits(num) - 1
	var inv int = 0
	for i, v := range digits {
		inv += v * (int(math.Pow(10, float64(max-i))))
	}

	return inv
}

//Evaluate if a number and your reverse are equals
func IsPalyndrom(num int) bool {
	return num == ReverseNumber(num)
}
