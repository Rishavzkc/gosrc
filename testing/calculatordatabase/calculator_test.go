package calculatordatabase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DiscountRepositoryMock struct {
	mock.Mock
}

func (drm *DiscountRepositoryMock) FindCurrrentDiscount() int {
	args := drm.Called()

	return args.Int(0)
}

func TestCalculatator(t *testing.T) {
	type testcase struct {
		name                  string
		minimumPurchaseAmount int
		purchaseAmount        int
		discount              int
		expectedAmount        int
	}
	testcases := []testcase{
		{
			name:                  "test 1",
			minimumPurchaseAmount: 100,
			purchaseAmount:        150,
			discount:              20,
			expectedAmount:        130,
		},
		{
			name:                  "test 2",
			minimumPurchaseAmount: 100,
			purchaseAmount:        200,
			discount:              20,
			expectedAmount:        160,
		},
		{
			name:                  "test 3",
			minimumPurchaseAmount: 100,
			purchaseAmount:        350,
			discount:              20,
			expectedAmount:        290,
		},
		{
			name:                  "test 4",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              20,
			expectedAmount:        50,
		},
		{
			name:                  "test 5",
			minimumPurchaseAmount: 100,
			purchaseAmount:        50,
			discount:              0,
			expectedAmount:        50,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := DiscountRepositoryMock{}

			discountRepositoryMock.On("FindCurrentDiscount").Return(tc.discount)
			calculator, err := NewDiscountCal(tc.minimumPurchaseAmount, &discountRepositoryMock)
			if err != nil {
				t.Errorf("could not instantiate %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)

			assert.Equal(t, tc.expectedAmount, amount)

			//in place of if statement we use assert
			/*	if amount != tc.expectedAmount {
				t.Errorf("expected %v, got %v", tc.expectedAmount, amount)
			} */

		})
	}
}

func TestCalculatorErrror(t *testing.T) {
	discountRepositoryMock := DiscountRepositoryMock{}

	_, err := NewDiscountCal(0, &discountRepositoryMock)
	if err == nil {
		t.Fatalf("should not calculate with zero purchase amount")
	}
}
