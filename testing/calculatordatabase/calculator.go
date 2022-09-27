package calculatordatabase

import (
	"errors"
	"testing/database"
)

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountRepository    database.Discount
}

func NewDiscountCal(minimumPurchaseAmount int, discountRepository database.Discount) (*DiscountCalculator, error) {
	if minimumPurchaseAmount == 0 {
		return &DiscountCalculator{}, errors.New("Minimum purchanse amount can not be zero")
	}

	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountRepository:    discountRepository,
	}, nil
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {

		multiplier := purchaseAmount / c.minimumPurchaseAmount
		discount := c.discountRepository.FindCurrrentDiscount()

		return purchaseAmount - (discount * multiplier)
	}
	return purchaseAmount
}
