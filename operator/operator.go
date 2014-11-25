package operator

import (
	"captcha/operand"
)

type Operator int

const (
	ADD Operator = iota
	SUB
	MUL
)

func (o Operator) Text() string {
	return textOperators[o]
}

var (
	textOperators []string = []string{"+", "-", "x"}
)

func (o Operator) Apply(left, right operand.Operand) int {
	realOperator := realOperators[o]
	return realOperator(left.Value(), right.Value())
}

var (
	realOperators []binaryOperator = []binaryOperator{add, sub, mul}
)

type binaryOperator func(left, right int) int

func add(left, right int) int {
	return left + right
}

func sub(left, right int) int {
	return left - right
}

func mul(left, right int) int {
	return left * right
}
