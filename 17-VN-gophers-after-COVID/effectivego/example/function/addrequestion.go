package example

type Order struct {
	ID string
}

func processRequest() {
	// What is this true stand for, why i put true value in here????
	o, err := CreateOrder("product id", "customer id", "shipment id", true)
	if err != nil {
		panic(err)
	}
}

func CreateOrder(productID, customerID, shipmentID string, isPromotion bool) (Order, error) {
	return Order{}, nil
}
