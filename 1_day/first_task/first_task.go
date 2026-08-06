package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func readInt(reader *bufio.Reader) int {
	res := 0
	sign := 1
	b, _ := reader.ReadByte()
	for b < '0' || b > '9' {
		if b == '-' {
			sign = -1
		}
		b, _ = reader.ReadByte()
	}
	for b >= '0' && b <= '9' {
		res = res*10 + int(b-'0')
		b, _ = reader.ReadByte()
	}
	return res * sign
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)

	answer := rand.Intn(100) + 1

	for {
		guess := readInt(reader)

		if guess < answer {
			fmt.Println("Too low!")
		} else if guess > answer {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Correct!")
			break
		}
	}

}
