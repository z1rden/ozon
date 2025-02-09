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
		var message string
		if _, err := fmt.Fscan(in, &message); err != nil {
			panic(err)
		}

		if len(message) == 1 || message[0] != 'M' || message[len(message)-1] != 'D' {
			// Если длина сообщения равна 1 (один символ), первый символ не M, а последний не D, то ответ "NO".
			fmt.Fprintln(out, "NO")
		} else {
			var j int
		loop:
			for j = 0; j < len(message)-1; j++ {
				/*
					Для каждого текущего символа есть только один символ, который может следовать далее.
					Если следующий символ не совпадает с условием задачи, то ответ "NO".
				*/
				switch message[j] {
				case 'R':
					if message[j+1] != 'C' {
						fmt.Fprintln(out, "NO")
						break loop
					}
				case 'C':
					if message[j+1] != 'M' {
						fmt.Fprintln(out, "NO")
						break loop
					}
				case 'D':
					if message[j+1] != 'M' {
						fmt.Fprintln(out, "NO")
						break loop
					}
				case 'M':
					if message[j+1] == 'M' {
						fmt.Fprintln(out, "NO")
						break loop
					}
				}
			}

			// Если проблем выше не возникло, то все хорошо.
			if j == len(message)-1 {
				fmt.Fprintln(out, "YES")
			}
		}
	}
}
