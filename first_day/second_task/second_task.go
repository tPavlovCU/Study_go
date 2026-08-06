package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(reader *bufio.Reader) int {
	res := 0
	sign := 1
	b, err := reader.ReadByte()
	for (b < '0' || b > '9') && err == nil {
		if b == '-' {
			sign = -1
		}
		if err != nil {
			break
		}
		b, err = reader.ReadByte()
		fmt.Println("111")
	}
	for b >= '0' && b <= '9' && err == nil {
		res = res*10 + int(b-'0')
		b, err = reader.ReadByte()
		if err != nil {
			break
		}
	}
	fmt.Println("res", res)
	return res * sign
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	first := readInt(reader)
	sign, _ := reader.ReadByte()
	second := readInt(reader)

	switch sign {
	case '+':
		fmt.Println(first + second)
	case '-':
		fmt.Println(first - second)
	case '*':
		fmt.Println(first * second)
	case '/':
		if second == 0 {
			fmt.Println("Division by zero!")
		} else {
			fmt.Println(first / second)
		}
	default:
		fmt.Println("Unknown operator")
	}

}
