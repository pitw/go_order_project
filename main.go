package main

import (
	"fmt"
	"log"

	"github.com/axnd/goaxnd/shared"
)

func main() {

	/// Open Database Connection (with Autoclose)
	shared.Connect()

	newOrder := shared.Order{
		Price:       2,
		Tax:         23,
		ProductType: shared.ProductType{ProductType: "PC", Category: "Computer"},
	}

	/// Create a new Order
	db := shared.SQL.Model(&shared.Order{}).Create(&newOrder)

	if db.Error != nil {
		log.Print(db.Error)
	}

	/// Create a massive batch order

	for i := 1; i <= 21; i++ {

		tempOrder := shared.Order{Price: float64(i) + 2.0, Tax: (float64(i) + 2.0) / (100.0 * 7.7)}
		db := shared.SQL.Model(&shared.Order{}).Create(&tempOrder)
		if db.Error != nil {
			log.Print(db.Error)
		}

	}

	// Select all Orders

	orders := []shared.Order{}
	dbSearch := shared.SQL.Model(&shared.Order{}).Preload("ProductType").Find(&orders)
	if dbSearch.Error != nil {
		log.Print(db.Error)
	}

	// Print all found Orders
	for _, order := range orders {

		fmt.Printf("Price: %v Tax: %v FinalPrice: %v ID: %v \n", order.Price, order.Tax, order.FinalPrice, order.ID)
		if order.ProductType.ID > 0 {
			fmt.Printf(">>>>>>>>>> Order has also a ProductTye! %v", order.ProductType.ProductType)
		}

	}

	/* 	dbDelete := shared.SQL.Model(&shared.Order{}).Where("id > 0").Delete(&shared.Order{})
	   	if dbDelete.Error != nil {
	   		log.Print(db.Error)
	   	} */
}
