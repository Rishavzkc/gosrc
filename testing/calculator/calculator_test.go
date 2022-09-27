package calculator

import "testing"

func TestCalculator(t *testing.T) {
	type testcase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		expectedAmount        int
	}

	testcases := []testcase{
		{name: "test 1", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, expectedAmount: 130},
		{name: "test 2", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 200, expectedAmount: 160},
		{name: "test 3", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 350, expectedAmount: 290},
		{name: "test 4", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 50, expectedAmount: 50},
		//{name: "mimimum amount is zero", minimumPurchaseAmount: 0, discount: 20, purchaseAmount: 50, expectedAmount: 90},
	}

	//to run specfic test = (go test -run="TestCalculator/test 1")

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			calculator, err := NewDiscountCal(tc.minimumPurchaseAmount, tc.discount)
			if err != nil {
				t.Errorf("could not instantiate %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)
			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
			}

		})
		/*	calculator := NewDiscountCal(tc.minimumPurchaseAmount, tc.discount)
			amount := calculator.Calculate(tc.purchaseAmount)
			if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
			}
		} */
	}
}

//for error when minimum amount is zero
func TestCalculatorErrror(t *testing.T) {
	_, err := NewDiscountCal(0, 20)
	if err == nil {
		t.Fatalf("should not calculate with zero purchase amount")
	}
}

/*
func TestDiscountApplied(t *testing.T) {
	calculator := NewDiscountCal(100, 20)
	amount := calculator.Calculate(150)

	if amount != 130 {
		t.Errorf("expected 160, got %v", amount) //Error =log +fail
		//	t.Fail()
	}
}

func TestDiscountMultiplied(t *testing.T) {
	calculator := NewDiscountCal(100, 20)
	amount := calculator.Calculate(200)

	if amount != 160 {
		t.Errorf("expected 160, got %v", amount) //Error =log +fail
		//	t.Fail()
	}
}

func TestDiscountNotApplied(t *testing.T) {
	calculator := NewDiscountCal(100, 20)
	amount := calculator.Calculate(50)

	if amount != 50 {
		t.Logf("expected 50, got %v", amount)
		t.Fail()
		//	t.Errorf("expected 50, got %v", amount) //Error =log +fail
	}
} */
