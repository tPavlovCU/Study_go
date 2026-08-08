package main

import (
	"bufio"
	"fmt"
	"os"
)

func StartsWith(s, part string) bool {
	i := 0
	j := len(part)
	if s[i:j] == part {
		return true
	}
	return false
}

func ParserLogFile(filepath string) (int, error) {
	file, err := os.Open(filepath)
	cnt := 0
	if err != nil {
		fmt.Println("ошибка открытия файла", err)
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		if StartsWith(line, "[ERROR]") {
			cnt++
		}
	}
	return cnt, nil

}
func main() {
	fmt.Println(ParserLogFile("test_file.txt"))

}
