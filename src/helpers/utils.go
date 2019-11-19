package helpers

import (
	"sort"

	"repository"
)



func GetOrderNo() int {
	oo, _ := repository.LoadOrders()
	if len(oo) == 0 {
		return 1
	}
	sort.Slice(oo, func(i, j int) bool {
		return oo[i].No < oo[j].No
	})

	return oo[len(oo)-1].No + 1
}
