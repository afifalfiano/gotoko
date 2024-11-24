package models

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: User{}},
		{Model: Address{}},
		{Model: Product{}},
		{Model: ProductImage{}},
		{Model: Section{}},
		{Model: Category{}},
		{Model: Order{}},
		{Model: OrderItem{}},
		{Model: OrderCustomer{}},
		{Model: Shipment{}},
		{Model: Payment{}},
	}
}
