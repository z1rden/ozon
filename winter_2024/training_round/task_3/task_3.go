package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func read_int_numbers(in *bufio.Reader, array []int, i int, n int) {
	if n == 0 {
		return
	}
	if _, err := fmt.Fscan(in, &array[i]); err != nil {
		panic(err)
	}
	read_int_numbers(in, array, i+1, n-1)
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count_arrays int
	if _, err := fmt.Fscan(in, &count_arrays); err != nil {
		panic(err)
	}

loop:
	for i := 0; i < count_arrays; i++ {
		var len_array int
		if _, err := fmt.Fscan(in, &len_array); err != nil {
			panic(err)
		}

		array := make([]int, len_array)
		// Рекурсивно считывается массив заведомо верных чисел.
		read_int_numbers(in, array, 0, len_array)

		// Пропуск символа конца строки, оставшегося от прошлого считывания.
		in.ReadString('\n')

		s, _ := in.ReadString('\n')
		s = s[:len(s)-1]

		words := strings.Split(s, " ")
		if len(words) != len_array {
			fmt.Fprintln(out, "no")
			continue
		}

		sort_array := make([]int, len_array)
		// Если в массиве, который должен получиться присутствуют не целые числа, то ответ "no".
		for i, value := range words {
			value_int, err := strconv.Atoi(value)
			if err != nil || value != strconv.Itoa(value_int) {
				fmt.Fprintln(out, "no")
				continue loop
			}
			sort_array[i] = value_int
		}

		sort.Ints(array)

		if slices.Equal(array, sort_array) {
			fmt.Fprintln(out, "yes")
		} else {
			fmt.Fprintln(out, "no")
		}
	}
}
