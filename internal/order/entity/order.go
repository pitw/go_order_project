package entity

import "errors"

// like a class
type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

// method for creating new orders

// *Order / &Order -- has to do with pointers and will be explained at the end of this lesson

// a new order will be validated before it is created, avoiding e.g. price -100.00 - (*Order, error) and the function returns an order or an error
func NewOrder(id string, price float64, tax float64) (*Order, error) {
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
	if o.ID == "" {
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

// how to calculate the final price
func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
