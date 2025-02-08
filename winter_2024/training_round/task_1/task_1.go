package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		var number_as_string string
		if _, err := fmt.Fscan(in, &number_as_string); err != nil {
			panic(err)
		}

		// Если число одноразрядное, то ответ 0.
		if len(number_as_string) == 1 {
			fmt.Fprintln(out, 0)
		} else {
			var j int
			// Если текущая цифра меньше следующей, то необходимо удалить ее.
			for j = 0; j < len(number_as_string)-1; j++ {
				if number_as_string[j] < number_as_string[j+1] {
					fmt.Fprintln(out, number_as_string[:j]+number_as_string[j+1:])
					break
				}
			}

			// Если цифра была не найдена, то удаляется последняя.
			if j == len(number_as_string)-1 {
				fmt.Fprintln(out, number_as_string[:j])
			}
		}

	}
}
