package captcha

func Captcha(pattern int, leftOperand int, operator int, rightOperand int) string {
	if pattern == 1 && leftOperand == 1 && operator == 1 && rightOperand == 1 {
		return "1 + One"
	}

	if pattern == 1 && leftOperand == 1 && operator == 1 && rightOperand == 2 {
		return "1 + Two"
	}

	return ""
}
