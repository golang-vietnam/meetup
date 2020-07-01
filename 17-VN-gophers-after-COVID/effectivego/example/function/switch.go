package example

func processRequest() (Order, error) {
	if isPromotion {
		order, err := CreateOrderWithPromotionProduct()
		return order, err
	} else {
		order, err := CreateOrderWithNormalProduct()
		return order, err
	}
}
