package captcha

import (
	"math/rand"
	"strconv"
)

func Captcha(pattern int, leftOperand int, operator int, rightOperand int) string {

	result := ""
	numWords := [9]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	operators := [3]string{" + ", " - ", " x "}

	if pattern == 3 {
		return "Pattern not found"
	}

	if operator == 0 || operator > len(operators) {
		return "Operator not found"
	}

	if leftOperand == 0 || leftOperand > len(numWords) {
		return "Left operand not found"
	}

	if pattern == 1 {
		result = strconv.Itoa(leftOperand) + operators[operator-1] + numWords[rightOperand-1]
	}

	if pattern == 2 {
		result = numWords[leftOperand-1] + operators[operator-1] + strconv.Itoa(rightOperand)
	}

	return result
}

func RandomCaptcha() string {
	rdPattern := rand.Intn(2) + 1
	rdLeftOperand := rand.Intn(9) + 1
	rdOperator := rand.Intn(3) + 1
	rdRightOperand := rand.Intn(9) + 1

	return Captcha(rdPattern, rdLeftOperand, rdOperator, rdRightOperand)
}
