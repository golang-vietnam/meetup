package example

func main() {
	o, q, t, err := CreateOrder("product id", "customer id", "shipment id")
	if err != nil {
		panic(err)
	}
}

// Return Transaction to show the transaction info in dashboard
func CreateOrder(productID, customerID, shipmentID string) (Order, Quote, Transaction, error) {
	return Order{}, Quote{}, Transaction{}, nil
}
