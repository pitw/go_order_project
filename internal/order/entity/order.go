package entity

import (
	"errors"

	"gorm.io/gorm"
)

// like a class
type Order struct {
	ID         int `json:"ID" gorm:"unique;primary_key;auto_increment:true"`
	Price      float64
	Tax        float64
	FinalPrice float64
}

// method for creating new orders

// *Order / &Order -- has to do with pointers and will be explained at the end of this lesson

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.FinalPrice = o.Price + o.Tax
	tx.Model(&Order{}).Where("id = ?", o.ID).Updates(&o)

	return
}

// a new order will be validated before it is created, avoiding e.g. price -100.00 - (*Order, error) and the function returns an order or an error
func NewOrder(id int, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil

	// return &Order{
	// 	ID:    id,
	// 	Price: price,
	// 	Tax:   tax,
	// }, nil
}

// the (the *Order) is an Order method and can be called from Order struct
func (o *Order) IsValid() error {
	if o.ID == 0 {
		return errors.New("invalid id")
	}
	if o.Price <= 0 {
		return errors.New("invalid price")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

// Ckeck the final price
func (o *Order) CheckFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}

// Calculacte Final Price
func (o Order) CalculateFinalPrice() Order {
	o.FinalPrice = o.Price + o.Tax

	return o
}
