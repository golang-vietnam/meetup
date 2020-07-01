package example

func main() {
	order, err := CreateOrder("product id", "customer id", "shipment id")
	if err != nil {
		panic(err)
	}

	quote, err := CreateQuote(order.id)
	if err != nil {
		panic(err)
	}

	transaction, err := CreateTransaction(order.id)
	if err != nil {
		panic(err)
	}

	invoice, err := CreateInvoice(order.id)
	if err != nil {
		return nil, err
	}
}
