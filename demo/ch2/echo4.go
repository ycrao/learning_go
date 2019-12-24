// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

/*
var 变量名字 类型 = 表达式
go 采用零值初始化，意味着 var s string ，没有赋值情况下，s 的值为 ""（空字符串）

var i, j, k int                 // int, int, int
var b, f, s = true, 2.3, "four" // bool, float64, string

请记住“:=”是一个变量声明语句，而“=”是一个变量赋值操作。也不要混淆多个变量的声明和元组的多重赋值（§2.4.1），后者是将右边各个的表达式值赋值给左边对应位置的各个变量：

i, j = j, i // 交换 i 和 j 的值

*/
// ./echo4 -s / a bc def
func main() {

	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"

	var x1, y1 int
	fmt.Println(&x1 == &x1, &x1 == &y1, &x1 == nil) // "true false false"
	
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
