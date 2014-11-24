package operator

import (
	"captcha/operand"
	"testing"
)

func TestShowOperatorAdd(t *testing.T) {
	if ADD.Text() != "+" {
		t.Errorf("Expect Operator ADD Show + but %s", ADD.Text())
	}
}

func TestShowOperatorSub(t *testing.T) {
	if SUB.Text() != "-" {
		t.Errorf("Expect Operator SUB Show - but %s", SUB.Text())
	}
}

func TestShowOperatorMul(t *testing.T) {
	if MUL.Text() != "x" {
		t.Errorf("Expect Operator MUL Show x but %s", MUL.Text())
	}
}

func TestApplyAdd(t *testing.T) {
	result := ADD.Apply(operand.ONE, operand.TWO)
	if result != 3 {
		t.Errorf("Expect result 3 but %d", result)
	}
}

func TestApplySub(t *testing.T) {
	result := SUB.Apply(operand.TWO, operand.ONE)
	if result != 1 {
		t.Errorf("Expect result 1 but %d", result)
	}
}

func TestApplyMul(t *testing.T) {
	result := MUL.Apply(operand.TWO, operand.TWO)
	if result != 4 {
		t.Errorf("Expect result 4 but %d", result)
	}
}
