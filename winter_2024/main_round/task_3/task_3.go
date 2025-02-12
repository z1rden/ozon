package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция, разбивающая товар из списка товаров, на имя товара и его цену.
func productInputProcess(s string) (name string, cost int) {
	product := strings.Split(s[:len(s)-1], " ")
	name = product[0]
	cost, err := strconv.Atoi(product[1])
	if err != nil {
		panic(err)
	}

	return
}

/*
Функция для проверки каждого товара из проверяемой строки:
1. Есть ли у товара цена?
2. Содержит ли цена ведущие нули?
3. Соответствует ли формату int цена?
*/
func productOutputProcess(product []string) (name string, cost int, flag bool) {
	if len(product) != 2 {
		return
	}
	name = product[0]
	if len(product[1]) > 1 && product[1][0] == '0' {
		return
	}
	cost, err := strconv.Atoi(product[1])
	if err != nil {
		return
	}

	flag = true
	return
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

loop:
	for i := 0; i < count; i++ {
		var countProducts int
		if _, err := fmt.Fscan(in, &countProducts); err != nil {
			panic(err)
		}
		in.ReadRune()

		// Список возможных продуктов представляется как map[цена]=[имена_товаров].
		listProducts := make(map[int][]string)
		for j := 0; j < countProducts; j++ {
			s, err := in.ReadString('\n')
			if err != nil {
				panic(err)
			}

			name, cost := productInputProcess(s)
			listProducts[cost] = append(listProducts[cost], name)
		}

		s, err := in.ReadString('\n')
		if err != nil {
			panic(err)
		}

		// Строка, которая проверяется,  представляется как map[цена]=имя_товара.
		listProductsOut := make(map[int]string)
		productsWithColon := strings.Split(s[:len(s)-1], ",")
		for _, productWithColon := range productsWithColon {
			product := strings.Split(productWithColon, ":")

			name, cost, flag := productOutputProcess(product)
			if !flag {
				fmt.Fprintln(out, "NO")
				continue loop
			}

			_, ok := listProductsOut[cost]
			// Если товар с такой уже существует, то ответ "no".
			if ok {
				fmt.Fprintln(out, "NO")
				continue loop
			} else {
				listProductsOut[cost] = name
			}
		}

		/*
			Идет проверка наличия товаров из выходной строки с товарами из списка товаров. Если произошел match,
			то удаляем такую цену у обоих map.
		*/
		for keyLPO, valueLPO := range listProductsOut {
			value, ok := listProducts[keyLPO]
			if !ok {
				fmt.Fprintln(out, "NO")
				continue loop
			}

			for _, v := range value {
				if v == valueLPO {
					delete(listProducts, keyLPO)
					delete(listProductsOut, keyLPO)
					break
				}
			}
		}

		if len(listProductsOut) == 0 && len(listProducts) == 0 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

}
