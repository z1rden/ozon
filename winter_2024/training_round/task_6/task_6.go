package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Для робота строится маршрут в левый верхний угол.
Алгоритм:
 1. Есть ли стойка над роботом?
    а. Есть. Тогда робот двигается на одну клетку влево и наверх до конца. Затем налево до места назначения.
    b. Нет. Тогда робот двигается наверх до конца. Затем налево до места назначения.
*/
func goToULC(storage [][]rune, x, y int, c rune) {
	if y-1 != -1 {
		if storage[y-1][x] == '#' {
			x--
			storage[y][x] = c
		}
		for i := y - 1; i != -1; i-- {
			storage[i][x] = c
		}
	}

	for i := x - 1; i != -1; i-- {
		storage[0][i] = c
	}
}

/*
	Для робота строится маршрут в правый нижний угол.
	Алгоритм:
	1. Есть ли стойка над роботом?
		а. Есть. Тогда робот двигается на одну клетку вправо и вниз до конца. Затем направо до места назначения.
		b. Нет. Тогда робот двигается вниз до конца. Затем направо до места назначения.
*/

func goToLRC(storage [][]rune, x, y int, c rune) {
	if y+1 != len(storage) {
		if storage[y+1][x] == '#' {
			x++
			storage[y][x] = c
		}
		for i := y + 1; i != len(storage); i++ {
			storage[i][x] = c
		}
	}
	for i := x + 1; i != len(storage[0]); i++ {
		storage[len(storage)-1][i] = c
	}
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
	in.ReadRune()

	for i := 0; i < count; i++ {
		var n, m, x_A, y_A, x_B, y_B int

		if _, err := fmt.Fscanf(in, "%d %d \n", &n, &m); err != nil {
			panic(err)
		}

		storage := make([][]rune, n)
		for j := 0; j < n; j++ {
			storage[j] = make([]rune, m)
		}

		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				c, _, err := in.ReadRune()
				switch c {
				case 'A':
					y_A = j
					x_A = k
				case 'B':
					y_B = j
					x_B = k
				}
				if err != nil {
					panic(err)
				}
				storage[j][k] = c
			}
			in.ReadRune()
		}

		// Определяется, чей маршрут короче до верхнего левого угла.
		if x_A+y_A < x_B+y_B {
			goToULC(storage, x_A, y_A, 'a')
			goToLRC(storage, x_B, y_B, 'b')
		} else {
			goToULC(storage, x_B, y_B, 'b')
			goToLRC(storage, x_A, y_A, 'a')
		}

		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				fmt.Fprintf(out, "%c", storage[j][k])
			}
			fmt.Fprintln(out)
		}
	}
}
