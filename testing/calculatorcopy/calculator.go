package calculatorcopy

type Discount struct {
	minimumPurchaseAmount int
	discountAmount        int
}

func NewDiscountCal(minimumPurchaseAmount int, discountAmount int) *Discount {
	return &Discount{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}
}

func (c *Discount) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {
		return purchaseAmount - c.discountAmount
	}
	return purchaseAmount
}
