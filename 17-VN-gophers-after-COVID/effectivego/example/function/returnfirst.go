package example

func processRequest(isPromotion bool) (Order, error) {
	if isPromotion {
		o, err := CreateOrderWithPromotionProduct()
		return o, err
	}

	o, err := CreateOrderWithNormalProduct()
	return o, err
}
