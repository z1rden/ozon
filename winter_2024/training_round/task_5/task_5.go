package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

/*
Функция сопоставления заказа и машины.
Если время окончания погрузки грузовой машины меньше времени прибытия заказа, то эта машина не понадобится для всех остальных заказов,
так как заказы отсортированы от меньшего времени к большему. Аналогично с вместимостью.
*/
func comparison(orders []Order, trucks []Truck) []Order {
	for j := 0; j < len(orders); j++ {
		for len(trucks) > 0 && (trucks[0].end < orders[j].arr || trucks[0].cap == 0) {
			trucks = trucks[1:]
		}
		if len(trucks) == 0 || !(trucks[0].start <= orders[j].arr && orders[j].arr <= trucks[0].end) {
			orders[j].indexTruck = -1
			continue
		}
		orders[j].indexTruck = trucks[0].index
		trucks[0].cap -= 1
	}
	return orders
}

/*
Структура для заказа:
1. index - индекс согласно вводу;
2. arr - время прибытия заказа;
3. indexTruck - индекс грузовой машины, в которую помещен.
*/
type Order struct {
	index      int
	arr        int
	indexTruck int
}

/*
Структура для грузовой машины:
1. index - индекс согласно вводу;
2. start - время начала погрузки;
3. end - время окончания погрузки;
4. cap - вместимость.
*/
type Truck struct {
	index int
	start int
	end   int
	cap   int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	if _, err := fmt.Fscan(in, &count); err != nil {
		panic(err)
	}

	for i := 0; i < count; i++ {
		var countOrders int
		if _, err := fmt.Fscan(in, &countOrders); err != nil {
			panic(err)
		}
		orders := make([]Order, countOrders)
		for j := 0; j < countOrders; j++ {
			if _, err := fmt.Fscan(in, &orders[j].arr); err != nil {
				panic(err)
			}
			orders[j].index = j
		}

		var countTrucks int
		if _, err := fmt.Fscan(in, &countTrucks); err != nil {
			panic(err)
		}
		trucks := make([]Truck, countTrucks)
		for j := 0; j < countTrucks; j++ {
			if _, err := fmt.Fscan(in, &trucks[j].start); err != nil {
				panic(err)
			}
			if _, err := fmt.Fscan(in, &trucks[j].end); err != nil {
				panic(err)
			}
			if _, err := fmt.Fscan(in, &trucks[j].cap); err != nil {
				panic(err)
			}
			trucks[j].index = j + 1
		}

		// Сортировка заказов по возрастанию времени прибытия.
		sort.Slice(orders, func(i, j int) bool {
			return orders[i].arr < orders[j].arr
		})
		// Сортировка грузовых машин по времени начала погрузки от меньшего к большему. Если одинаковое время, то по индексу.
		sort.Slice(trucks, func(i, j int) bool {
			return trucks[i].start < trucks[j].start ||
				(trucks[i].start == trucks[j].start && trucks[i].index < trucks[j].index)
		})

		orders = comparison(orders, trucks)

		sort.Slice(orders, func(i, j int) bool {
			return orders[i].index < orders[j].index
		})
		for j := 0; j < countOrders; j++ {
			fmt.Fprint(out, orders[j].indexTruck)
			fmt.Fprint(out, " ")
		}
		fmt.Fprintln(out)
	}
}
