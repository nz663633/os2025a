package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	i, _ := r.ReadString('\n') // ignore error
	fmt.Println(i)
}
