package operator

import (
	"captcha/operand"
)

type Operator interface {
	Apply(left, right operand.Operand) int
	Text() string
}

type operator int

const (
	ADD operator = iota
	SUB
	MUL
)

func (o operator) Text() string {
	return textOperators[o]
}

var (
	textOperators []string = []string{"+", "-", "x"}
)

func (o operator) Apply(left, right operand.Operand) int {
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
