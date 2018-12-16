package captcha

import "strconv"

func Captcha(pattern int, leftOperand int, operator int, rightOperand int) string {

	numWords := [9]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	operators := [3]string{" + ", " - ", " x "}

	if pattern == 1 {
		return strconv.Itoa(leftOperand) + operators[operator-1] + numWords[rightOperand-1]
	}

	return ""
}
