package main

import (
	"github.com/pvmtriet232/test/models"
)

func main() {
	goku := &models.Saiyan{"goku", 9001, true}
	goku.DescribeSaiyan()
}
