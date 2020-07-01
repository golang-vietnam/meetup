package example

func processRequest() {
	o, err := CreateOrder("product id", "customer id", "shipment id", true)
	if err != nil {
		panic(err)
	}
}

func CreateOrder(productID, customerID, shipmentID string, isPromotion bool) (Order, error) {
	return Order{}, nil
}
