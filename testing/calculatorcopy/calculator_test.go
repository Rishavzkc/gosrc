package calculatorcopy

import "testing"

func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCal(100, 20)
	amount := calculator.Calculate(150)

	if amount != 130 {
		t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCal(100, 20)
	amount := calculator.Calculate(50)

	if amount != 50 {
		t.Fail()
	}
}
