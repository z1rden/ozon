package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Метод, который считает все зараженные файлы в текущей Folder и в subFolder.
func (folder Folder) countFiles() (count_hack_files int) {
	count_hack_files += len(folder.Files)
	for _, sub_folder := range folder.Folders {
		count_hack_files += sub_folder.countFiles()
	}
	return
}

// Метод, который проверяет наличие файла с hack.
func (folder Folder) isHacked() bool {
	for _, file := range folder.Files {
		if strings.HasSuffix(file, ".hack") {
			return true
		}
	}
	return false
}

// Метод, который считает общее количество зараженных файлов.
func (folder Folder) countHackedFiles() int {
	if folder.isHacked() {
		return folder.countFiles()
	}

	count := 0
	for _, sub_folder := range folder.Folders {
		count += sub_folder.countHackedFiles()
	}

	return count
}

// Функция, которая из входного JSON делает структуру.
func createStructFromJson(data []byte) (folder Folder) {
	err := json.Unmarshal(data, &folder)
	if err != nil {
		panic(err)
	}

	return
}

// Unmarshal может не записывать данные в переменную: имена полей структуры должны начинаться с больших букв.
type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
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
		var count_strings int
		if _, err := fmt.Fscan(in, &count_strings); err != nil {
			panic(err)
		}
		in.ReadString('\n')

		var data string
		for j := 0; j < count_strings; j++ {
			line, err := in.ReadString('\n')
			if err != nil {
				panic(err)
			}
			data += line
		}

		folder := createStructFromJson([]byte(data))
		fmt.Fprintln(out, folder.countHackedFiles())
	}
}
