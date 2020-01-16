package main

import "fmt"

var home = []homeShop{
	{
		id:                 1,
		name:               "2-комн. квартира, 8 этаж, 15 м²",
		distanceFromCenter: 20,
		price:              10_000,
		district:           "Рудаки",
	},
	{
		id:                 2,
		name:               "3-комн. квартира, 4 этаж, 20 м²",
		distanceFromCenter: 5,
		price:              20_000,
		district:           "Фирдавси",
	},
	{
		id:                 3,
		name:               "3-комн. квартира, 2 этаж, 20 м²",
		distanceFromCenter: 2,
		price:              25_000,
		district:           "Фирдавси",
	},
	{
		id:                 4,
		name:               "4-комн. квартира, 5 этаж, 25 м²",
		distanceFromCenter: 0,
		price:              30_000,
		district:           "Сомони",
	},
}

func ExampleSortByCheapPrice() {
	result := sortByCheapPrice(home)
	fmt.Println(result)
	//Output:[{1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки} {2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси} {4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони}]
}

func ExampleSortByExpensivePrice() {
	result := sortByExpensivePrice(home)
	fmt.Println(result)
	//Output:[{4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси} {2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки}]
}
func ExampleSortByNearToCenter() {
	result := sortByNearToCenter(home)
	fmt.Println(result)
	//Output:[{4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси} {2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки}]
}
func ExampleSortByFarFromCenter() {
	result := sortByFarFromCenter(home)
	fmt.Println(result)
	//Output:[{1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки} {2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси} {4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони}]
}
func ExampleSearchByMaxPrice_NoResult() {
	result := searchByMaxPrice(home, 2_000)
	fmt.Println(result)
	//Output:[]
}
func ExampleSearchByMaxPrice_OneResult() {
	result := searchByMaxPrice(home, 10_000)
	fmt.Println(result)
	//Output:[{1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки}]
}
func ExampleSearchByMaxPrice_TwoAndMoreResults() {
	result := searchByMaxPrice(home, 30_000)
	fmt.Println(result)
	//Output:[{1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки} {2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси} {4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони}]
}
func ExampleSearchByLimitPrice_NoResult() {
	result := searchByLimitPrice(home, 2_300, 2_600)
	fmt.Println(result)
	//Output:[]
}
func ExampleSearchByLimitPrice_OneResult() {
	result := searchByLimitPrice(home, 10_000, 15_000)
	fmt.Println(result)
	//Output:[{1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки}]
}
func ExampleSearchByLimitPrice_TwoAndMoreResults() {
	result := searchByLimitPrice(home, 10_000, 30_000)
	fmt.Println(result)
	//Output:[{1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки} {2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси} {4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони}]
}
func ExampleSearchByDistrict_NoResult() {
	result := searchByDistrict(home, "Сино")
	fmt.Println(result)
	//Output:[]
}
func ExampleSearchByDistrict_OneResult() {
	result := searchByDistrict(home, "Сомони")
	fmt.Println(result)
	//Output:[{4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони}]
}
func ExampleSearchByDistrict_TwoAndMoreResults() {
	result := searchByDistrict(home, "Фирдавси")
	fmt.Println(result)
	//Output:[{2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси}]
}
func ExampleSearchByDistricts_NoResult() {
	result := searchByDistricts(home, []string{"Сино", "Шоҳтемур"})
	fmt.Println(result)
	//Output:[]
}
func ExampleSearchByDistricts_OneResult() {
	result := searchByDistricts(home, []string{"Сино", "Сомони"})
	fmt.Println(result)
	//Output:[{4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони}]
}
func ExampleSearchByDistricts_TwoAndMoreResults() {
	result := searchByDistricts(home, []string{"Фирдавси", "Сомони", "Рудаки"})
	fmt.Println(result)
	//Output:[{1 2-комн. квартира, 8 этаж, 15 м² 20 10000 Рудаки} {2 3-комн. квартира, 4 этаж, 20 м² 5 20000 Фирдавси} {3 3-комн. квартира, 2 этаж, 20 м² 2 25000 Фирдавси} {4 4-комн. квартира, 5 этаж, 25 м² 0 30000 Сомони}]
}
