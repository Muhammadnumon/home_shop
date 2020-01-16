package main

import (
	"sort"
)

type homeShop struct {
	id                 int64
	name               string
	distanceFromCenter int64
	price              int64
	district           string
}

func sortByExpensivePrice(homeShops []homeShop) []homeShop {
	result := make([]homeShop, len(homeShops))
	copy(result, homeShops)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price > result[j].price
	})
	return result
}

func sortByCheapPrice(homeShops []homeShop) []homeShop {
	result := make([]homeShop, len(homeShops))
	copy(result, homeShops)
	sort.Slice(result, func(i, j int) bool {
		return result[i].price < result[j].price
	})
	return result
}

func searchByMaxPrice(homeShops []homeShop, max int64) []homeShop {
	result := make([]homeShop, 0)
	for _, home := range homeShops {
		if home.price <= max {
			result = append(result, home)
		}
	}
	return result
}

func searchByLimitPrice(homeShops []homeShop, min int64, max int64) []homeShop {
	result := make([]homeShop, 0)
	for _, home := range homeShops {
		if home.price <= max && home.price >= min {
			result = append(result, home)
		}
	}
	return result
}

func sortByNearToCenter(homeShops []homeShop) []homeShop {
	result := make([]homeShop, len(homeShops))
	copy(result, homeShops)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromCenter < result[j].distanceFromCenter
	})
	return result
}

func sortByFarFromCenter(homeShops []homeShop) []homeShop {
	result := make([]homeShop, len(homeShops))
	copy(result, homeShops)
	sort.Slice(result, func(i, j int) bool {
		return result[i].distanceFromCenter > result[j].distanceFromCenter
	})
	return result
}

func searchByDistrict(homeShops []homeShop, district string) []homeShop {
	result := make([]homeShop, 0)
	for _, home := range homeShops {
		if home.district == district {
			result = append(result, home)
		}
	}
	return result
}

func searchByDistricts(homeShops []homeShop, districts []string) []homeShop {
	result := make([]homeShop, 0)
	for _, home := range homeShops {
		for _, district := range districts {
			if home.district == district {
				result = append(result, home)
			}
		}
	}
	return result
}
func main() {}
