package captcha

import (
	"captcha/operand"
	"captcha/operator"
)

type Captcha interface {
	Text() string
	Answer() int
}

type captcha struct {
	left     operand.Operand
	operator operator.Operator
	right    operand.Operand
}

func New(left operand.Operand, operator operator.Operator, right operand.Operand) captcha {
	return captcha{left, operator, right}
}

func (c captcha) Text() string {
	return c.left.Text() + " " + c.operator.Text() + " " + c.right.Text()
}

func (c captcha) Answer() int {
	return c.operator.Apply(c.left, c.right)
}
