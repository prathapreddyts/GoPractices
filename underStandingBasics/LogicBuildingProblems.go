package underStandingBasics

// SumOfTwoNumbers takes two integers as input and returns their sum.
func SumOfTwoNumbers(a int, b int) int {
	return a + b
}
func SubstractTwoNumber(a int, b int) int {
	return a - b
}
func MultiplicationOfTwoNumber(a int, b int) int {
	return a * b
}
func QuotientOfTwoNumber(a int, b int) int {
	return a / b
}
func RemainderOfTwoNumber(a int, b int) int {
	return a % b
}
func SwapTwoNumber(a int, b int) (int, int) {
	var temp int
	temp = a
	a = b
	b = temp
	return a, b
}
func IsEvenNumber(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}
func ISOddNumber(a int) bool {
	if a%2 != 0 {
		return true
	}
	return false
}
func IsPositiveOrNegativeNumber(a int) string {
	if a < 0 {
		return "negative"
	}
	return "positive"
}

func FindSumOfFirstNNaturalNumbers(a int) int {
	var sum int
	for j := 0; j <= a; j++ {
		sum += j
	}
	return sum
}
func FindSumOfFirstNOddNumbers(a int) (int, int) {
	var evenSum int
	var oddSum int
	for i := 1; i <= a; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}
	return evenSum, oddSum
}
func reverseNumber(a int) int {
	var revNumber int
	for a != 0 {
		rem := a % 10
		revNumber = revNumber*10 + rem
		a = a / 10
	}
	return revNumber
}

func SumOfDigitsInNumber(a int) int {
	var digitSum int
	for a != 0 {
		rem := a % 10
		digitSum += rem
		a = a / 10
	}
	return digitSum
}
func IsPalindromeNumber(a int) bool {
	var revNumber int
	temp := a
	for a != 0 {
		rem := a % 10
		revNumber = revNumber*10 + rem
		a = a / 10
	}
	return temp == revNumber
}
