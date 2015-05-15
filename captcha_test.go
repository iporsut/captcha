package captcha

import (
	"lab/kata/captcha/operand"
	"lab/kata/captcha/operator"
	"testing"
)

func TestShowCaptchaTextOneAddNumberTwo(t *testing.T) {
	captcha := New(operand.TEXT_ONE, operator.ADD, operand.TWO)
	show := captcha.Text()
	expected := "ONE + 2"

	if show != expected {
		t.Errorf("Expect ONE + 2 but %s", show)
	}
}

func TestShowCaptchaTextNineAddNumberTwo(t *testing.T) {
	captcha := New(operand.TEXT_NINE, operator.ADD, operand.TWO)
	show := captcha.Text()
	expected := "NINE + 2"

	if show != expected {
		t.Errorf("Expect NINE + 2 but %s", show)
	}
}

func TestAnswerCaptchaTextNineAddNumberTwo(t *testing.T) {
	captcha := New(operand.TEXT_NINE, operator.ADD, operand.TWO)
	answer := captcha.Answer()
	expected := 11

	if answer != expected {
		t.Errorf("Expect 11 but %d", answer)
	}
}
