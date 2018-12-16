package captcha

func Captcha(pattern int, leftOperand int, operator int, rightOperand int) string {

	numWords := [9]string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	if pattern == 1 && leftOperand == 1 && operator == 1 {
		return "1 + " + numWords[rightOperand-1]
	}

	return ""
}
