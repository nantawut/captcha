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

func TestCaptcha_GivePattern_2_LeftOperand_9_Operator_3_RightOperand_9_ReturnNineMultiple9(t *testing.T) {
	pattern := 2
	leftOperand := 9
	operator := 3
	rightOperand := 9

	expected := "Nine x 9"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptcha_GivePattern_3_ReturnsErrorPatternNotFound(t *testing.T) {
	pattern := 3
	leftOperand := 1
	operator := 1
	rightOperand := 1

	expected := "Pattern not found"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptch_GiveOperator_4_ReturnsErrorOperatorNotFound(t *testing.T) {
	pattern := 1
	leftOperand := 1
	operator := 4
	rightOperand := 1

	expected := "Operator not found"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptch_GiveLeftOperand_0_ReturnsLeftOperandNotFound(t *testing.T) {
	pattern := 1
	leftOperand := 0
	operator := 1
	rightOperand := 1

	expected := "Left operand not found"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptch_GiveLeftOperand_10_ReturnsLeftOperandNotFound(t *testing.T) {
	pattern := 1
	leftOperand := 10
	operator := 1
	rightOperand := 1

	expected := "Left operand not found"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptch_GiveLeftOperand_11_ReturnsLeftOperandNotFound(t *testing.T) {
	pattern := 1
	leftOperand := 11
	operator := 1
	rightOperand := 1

	expected := "Left operand not found"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}

func TestCaptch_GiveOperator_0_ReturnsOperatorNotFound(t *testing.T) {
	pattern := 1
	leftOperand := 1
	operator := 0
	rightOperand := 1

	expected := "Operator not found"

	actual := Captcha(pattern, leftOperand, operator, rightOperand)

	if expected != actual {
		t.Errorf("It should be %q but get to %q", expected, actual)
	}
}