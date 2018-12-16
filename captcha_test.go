package captcha

import "testing"

func TestCaptcha_GivePattern1LeftOperand1Operator1RightOperand1_Returns1PlusOne(t *testing.T) {
	pattern := 1
	leftOperand := 1
	operator := 1
	rightOperand := 1

	expected := "1 + One"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptcha_GivePattern1LeftOperand1Operator1RightOperand2_Retruns1PlusTwo(t *testing.T) {
	pattern := 1
	leftOperand := 1
	operator := 1
	rightOperand := 2

	expected := "1 + Two"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}
