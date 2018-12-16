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

func TestCaptcha_GivePattern_1_LeftOperand_1_Operator_1_RightOperand_3_Retruns1PlusTwo(t *testing.T) {
	pattern := 1
	leftOperand := 1
	operator := 1
	rightOperand := 3

	expected := "1 + Three"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptcha_GivePattern_1_LeftOperand_2_Operator_2_RightOperand_4_Return2SubtractFour(t *testing.T) {
	pattern := 1
	leftOperand := 2
	operator := 2
	rightOperand := 4

	expected := "2 - Four"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}
