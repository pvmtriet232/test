package main

import (
	"github.com/pvmtriet232/test/db"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
