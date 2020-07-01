package example

func main() {
	o, q, t, i, err := CreateOrder("product id", "customer id", "shipment id")
	if err != nil {
		panic(err)
	}
}

// Return Invoice to show the invoice info in dashboard
func CreateOrder(productID, customerID, shipmentID string) (Order, Quote, Transaction, Invoice, error) {
	return Order{}, Quote{}, Transaction{}, Invoice{}, nil
}
