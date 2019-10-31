package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] { // - 表示 空标识符 blank identifier
		s += sep + arg
		sep = " - "
	}
	fmt.Println(s)
	fmt.Println(strings.Join(os.Args[1:], " - "))
	fmt.Println(os.Args[1:])
}
