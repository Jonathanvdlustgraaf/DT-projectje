package helpers

import (
	"log"

	"repository"
	. "types"
)

func CalculateTotal(o Order) float32 {
	var total float32

	for _, l := range o.Lines {
		p, err := repository.GetPizza(l.Pizza)
		if err != nil {
			log.Printf("Pizza %v not found, total price may be incorrect\n", l.Pizza)
		} else {
			total += float32(p.Price * float32(l.Qty))
		}
	}
	return total
}
