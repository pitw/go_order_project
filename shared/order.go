package shared

import (
	"gorm.io/gorm"
)

// like a class
type Order struct {
	ID               int `json:"ID" gorm:"unique;primary_key;auto_increment:true"`
	Price            float64
	Tax              float64
	FinalPrice       float64
	ProductTypeRefer int
	ProductType      ProductType `gorm:"foreignKey:ProductTypeRefer"`
}

type ProductType struct {
	ID          int `json:"ID" gorm:"unique;primary_key;auto_increment:true"`
	ProductType string
	Category    string
}

// method for creating new orders

// *Order / &Order -- has to do with pointers and will be explained at the end of this lesson

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.FinalPrice = o.Price + o.Tax
	tx.Model(&Order{}).Where("id = ?", o.ID).Updates(&o)

	return
}
