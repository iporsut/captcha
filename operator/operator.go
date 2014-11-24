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

type binaryOperator func(left, right int) int

var (
	textOperators []string = []string{"+", "-", "x"}
)

var (
	realOperators []binaryOperator = []binaryOperator{add, sub, mul}
)

func (o operator) Text() string {
	return textOperators[int(o)]
}

func (o operator) Apply(left, right operand.Operand) int {
	realOperator := realOperators[int(o)]
	return realOperator(left.Value(), right.Value())
}

func add(left, right int) int {
	return left + right
}

func sub(left, right int) int {
	return left - right
}

func mul(left, right int) int {
	return left * right
}
