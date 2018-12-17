package captcha

import "strconv"

func Captcha(pattern int, leftOperand int, operator int, rightOperand int) string {

	result := ""
	numWords := [9]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	operators := [3]string{" + ", " - ", " x "}

	if pattern == 1 {
		result = strconv.Itoa(leftOperand) + operators[operator-1] + numWords[rightOperand-1]
	}

	if pattern == 2 {
		result = numWords[leftOperand-1] + operators[operator-1] + strconv.Itoa(rightOperand)
	}

	if pattern == 3 {
		result = "Pattern not found"
	}

	return result
}
