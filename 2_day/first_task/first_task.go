package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readWord(reader *bufio.Reader, startCh chan int) string {
	start := <-startCh
	_ = start

}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1<<20)
	fullLine, _ := reader.ReadString('\n')
	_ = fullLine
	startCh := make(chan int, 1)
	startCh <- 0
}
