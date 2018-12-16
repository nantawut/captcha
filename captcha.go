package captcha

import "strconv"

func Captcha(pattern int, leftOperand int, operator int, rightOperand int) string {

	numWords := [9]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	if pattern == 1 && operator == 1 {
		return strconv.Itoa(leftOperand) + " + " + numWords[rightOperand-1]
	}

	if pattern == 1 && operator == 2 {
		return strconv.Itoa(leftOperand) + " - " + numWords[rightOperand-1]
	}

	return ""
}
