package entity

import "sync"

type Book struct {
	Order        []*Order
	Transactions []*Transaction
	OrdersChan   chan *Order
	OrderChanOut chan *Order
	Wg           *sync.WaitGroup
}

func NewBook(orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		Order:        []*Order{},
		Transactions: []*Transaction{},
		OrdersChan:   orderChan,
		OrderChanOut: orderChanOut,
		Wg:           wg,
	}
}
