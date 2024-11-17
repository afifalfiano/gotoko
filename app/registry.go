package app

import "github.com/afifalfiano/gotoko/app/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
		{Model: models.Product{}},
		{Model: models.ProductImage{}},
		{Model: models.Section{}},
		{Model: models.Category{}},
		{Model: models.Order{}},
		{Model: models.OrderItem{}},
		{Model: models.OrderCustomer{}},
		{Model: models.Shipment{}},
		{Model: models.Payment{}},
	}
}
