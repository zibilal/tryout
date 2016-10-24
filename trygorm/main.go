package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/zibilal/tryout/trygorm/model"
	"fmt"
)

func main() {
	db, err := gorm.Open("mysql", "root:bilalmuh@tcp(localhost:3306)/mmdbv3?parseTime=true")
	if err != nil {
		panic("Failed to open the database")
	}

	defer db.Close()

	fmt.Println("Querying all the sales order")
	salesOrders := [] model.SalesOrder{}
	db.Find(&salesOrders)

	fmt.Println("The sales orders", salesOrders)
}
