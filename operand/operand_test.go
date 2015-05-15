package operand

import (
	"testing"
)

func TestTextOfOperandTextZero(t *testing.T) {
	assertShow(t, "TEXT_ZERO", "ZERO", TEXT_ZERO.Text())
}

func TestTextOfOperandTextOne(t *testing.T) {
	assertShow(t, "TEXT_ONE", "ONE", TEXT_ONE.Text())
}

func TestTextOfOperandTextTwo(t *testing.T) {
	assertShow(t, "TEXT_TWO", "TWO", TEXT_TWO.Text())
}

func TestTextOfOperandTextThree(t *testing.T) {
	assertShow(t, "TEXT_THREE", "THREE", TEXT_THREE.Text())
}

func TestTextOfOperandTextFour(t *testing.T) {
	assertShow(t, "TEXT_FOUR", "FOUR", TEXT_FOUR.Text())
}

func TestTextOfOperandTextFive(t *testing.T) {
	assertShow(t, "TEXT_FIVE", "FIVE", TEXT_FIVE.Text())
}

func TestTextOfOperandTextSix(t *testing.T) {
	assertShow(t, "TEXT_SIX", "SIX", TEXT_SIX.Text())
}

func TestTextOfOperandTextSeven(t *testing.T) {
	assertShow(t, "TEXT_SEVEN", "SEVEN", TEXT_SEVEN.Text())
}

func TestTextOfOperandTextEight(t *testing.T) {
	assertShow(t, "TEXT_EIGHT", "EIGHT", TEXT_EIGHT.Text())
}

func TestTextOfOperandTextNine(t *testing.T) {
	assertShow(t, "TEXT_NINE", "NINE", TEXT_NINE.Text())
}

func TestValueOfOperandTextZero(t *testing.T) {
	assertValue(t, "TEXT_ZERO", 0, TEXT_ZERO.Value())
}

func TestValueOfOperandTextOne(t *testing.T) {
	assertValue(t, "TEXT_ONE", 1, TEXT_ONE.Value())
}

func TestValueOfOperandTextTwo(t *testing.T) {
	assertValue(t, "TEXT_TWO", 2, TEXT_TWO.Value())
}

func TestValueOfOperandTextThree(t *testing.T) {
	assertValue(t, "TEXT_THREE", 3, TEXT_THREE.Value())
}

func TestValueOfOperandTextFour(t *testing.T) {
	assertValue(t, "TEXT_FOUR", 4, TEXT_FOUR.Value())
}

func TestValueOfOperandTextFive(t *testing.T) {
	assertValue(t, "TEXT_FIVE", 5, TEXT_FIVE.Value())
}

func TestValueOfOperandTextSix(t *testing.T) {
	assertValue(t, "TEXT_SIX", 6, TEXT_SIX.Value())
}

func TestValueOfOperandTextSeven(t *testing.T) {
	assertValue(t, "TEXT_SEVEN", 7, TEXT_SEVEN.Value())
}

func TestValueOfOperandTextEight(t *testing.T) {
	assertValue(t, "TEXT_EIGHT", 8, TEXT_EIGHT.Value())
}

func TestValueOfOperandTextNine(t *testing.T) {
	assertValue(t, "TEXT_NINE", 9, TEXT_NINE.Value())
}

func TestTextOfOperandZero(t *testing.T) {
	assertShow(t, "ZERO", "0", ZERO.Text())
}

func TestTextOfOperandOne(t *testing.T) {
	assertShow(t, "ONE", "1", ONE.Text())
}

func TestTextOfOperandTwo(t *testing.T) {
	assertShow(t, "TWO", "2", TWO.Text())
}

func TestTextOfOperandThree(t *testing.T) {
	assertShow(t, "THREE", "3", THREE.Text())
}

func TestTextOfOperandFour(t *testing.T) {
	assertShow(t, "FOUR", "4", FOUR.Text())
}

func TestTextOfOperandFive(t *testing.T) {
	assertShow(t, "FIVE", "5", FIVE.Text())
}

func TestTextOfOperandSix(t *testing.T) {
	assertShow(t, "SIX", "6", SIX.Text())
}

func TestTextOfOperandSeven(t *testing.T) {
	assertShow(t, "SEVEN", "7", SEVEN.Text())
}

func TestTextOfOperandEight(t *testing.T) {
	assertShow(t, "EIGHT", "8", EIGHT.Text())
}

func TestTextOfOperandNine(t *testing.T) {
	assertShow(t, "NINE", "9", NINE.Text())
}

func TestValueOfOperandZero(t *testing.T) {
	assertValue(t, "ZERO", 0, ZERO.Value())
}

func TestValueOfOperandOne(t *testing.T) {
	assertValue(t, "ONE", 1, ONE.Value())
}

func TestValueOfOperandTwo(t *testing.T) {
	assertValue(t, "TWO", 2, TWO.Value())
}

func TestValueOfOperandThree(t *testing.T) {
	assertValue(t, "THREE", 3, THREE.Value())
}

func TestValueOfOperandFour(t *testing.T) {
	assertValue(t, "FOUR", 4, FOUR.Value())
}

func TestValueOfOperandFive(t *testing.T) {
	assertValue(t, "FIVE", 5, FIVE.Value())
}

func TestValueOfOperandSix(t *testing.T) {
	assertValue(t, "SIX", 6, SIX.Value())
}

func TestValueOfOperandSeven(t *testing.T) {
	assertValue(t, "SEVEN", 7, SEVEN.Value())
}

func TestValueOfOperandEight(t *testing.T) {
	assertValue(t, "EIGHT", 8, EIGHT.Value())
}

func TestValueOfOperandNine(t *testing.T) {
	assertValue(t, "NINE", 9, NINE.Value())
}
func assertValue(t *testing.T, description string, expected, actual int) {
	if actual != expected {
		t.Errorf("Expect %s show %s but %s ", description, expected, actual)
	}
}

func assertShow(t *testing.T, description, expected, actual string) {
	if actual != expected {
		t.Errorf("Expect %s show %s but %s ", description, expected, actual)
	}
}
