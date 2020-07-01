package example

func main() {
	o, err := CreateOrder("product id", "customer id", "shipment id")
	if err != nil {
		panic(err)
	}
}

func CreateOrder(pID string, cID string, sID string) (Order, error) {
	return Order{}, nil
}
