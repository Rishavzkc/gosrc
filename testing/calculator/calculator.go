package calculator

import "errors"

type Discount struct {
	minimumPurchaseAmount int
	discountAmount        int
}

func NewDiscountCal(minimumPurchaseAmount int, discountAmount int) (*Discount, error) {
	if minimumPurchaseAmount == 0 {
		return &Discount{}, errors.New("Minimum purchanse amount can not be zero")
	}

	return &Discount{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	},nil
}

func (c *Discount) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {

		multiplier := purchaseAmount / c.minimumPurchaseAmount

		return purchaseAmount - (c.discountAmount)*multiplier
	}
	return purchaseAmount
}
