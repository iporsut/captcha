package operand

type Operand interface {
	Text() string
	Value() int
}

type textOperand int

const (
	TEXT_ZERO textOperand = iota
	TEXT_ONE
	TEXT_TWO
	TEXT_THREE
	TEXT_FOUR
	TEXT_FIVE
	TEXT_SIX
	TEXT_SEVEN
	TEXT_EIGHT
	TEXT_NINE
)

type numberOperand int

const (
	ZERO numberOperand = iota
	ONE
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
)

func (o textOperand) Text() string {
	return textOperands[o]
}

var textOperands []string = []string{"ZERO", "ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}

func (o textOperand) Value() int {
	return int(o)
}

func (o numberOperand) Text() string {
	return numberOperands[o]
}

var numberOperands []string = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func (o numberOperand) Value() int {
	return int(o)
}
