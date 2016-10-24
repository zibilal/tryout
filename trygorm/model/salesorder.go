package model

import "time"

type SalesOrder struct {
	Id uint64
	OrderId uint64
	SoStoreNumber string
	StoreId uint64
	TotalPrice uint64
	MasterSalesOrderStatusId uint64
	SoDate time.Time
	DeliveredTime time.Time
	CollectionTime time.Time
}

func NewSalesOrder() *SalesOrder {
	return &SalesOrder{}
}
