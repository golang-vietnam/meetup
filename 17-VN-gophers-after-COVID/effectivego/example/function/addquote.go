package example

func main() {
	o, q, err := CreateOrder("product id", "customer id", "shipment id")
	if err != nil {
		panic(err)
	}
}

// Return Quote to show the quote info in dashboard
func CreateOrder(productID, customerID, shipmentID string) (Order, Quote, error) {
	return Order{}, Quote{}, nil
}
