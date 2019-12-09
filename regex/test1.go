package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	//1、过正则来判断字符串是否匹配
	if ok, _ := regexp.Match("^[0-9a-zA-Z_]+$", []byte("hello")); ok {
		fmt.Println("ok")
	}

	//上面的例子也可以通过MatchString实现
	if ok, _ := regexp.MatchString("^[0-9a-zA-Z_]+$", "hello"); ok {
		fmt.Println("ok")
	}

	//2、解析正则表达式
	//正则表达式如果合法，Compile会返回一个Regexp对象指针，通过该指针可以在任意字符串上进行操作
	re, _ := regexp.Compile("[0-9a-zA-Z_]+")

	//3、查找正则匹配的字符串
	data := "I am a good man"

	//Find函数返回匹配的第一个字符串
	one := re.Find([]byte(data))
	fmt.Println("one is " + string(one))

	//FindAll函数返回匹配的所有字符串，n小于0返回全部字符串，否则返回指定长度
	all := re.FindAll([]byte(data), 2)
	//all为长度为2的slice
	fmt.Println("all[0] is " + string(all[0]))
	fmt.Println("all[1] is " + string(all[1]))

	//FindIndex查找匹配的开始位置和结束位置
	ix := re.FindIndex([]byte(data))
	fmt.Print("FindIndex ")
	fmt.Println(ix)

	//FindAllIndex查找所有匹配的开始位置和结束位置
	//n小于0返回全部，否则返回指定长度
	all_ix := re.FindAllIndex([]byte(data), -1)
	fmt.Print("FindAllIndex: ")
	fmt.Println(all_ix)
	re2, _ := regexp.Compile("a(.*)g(.*)")

	//FindSubmatch查找子匹配项
	sub := re2.FindSubmatch([]byte(data))
	//第一个匹配的是全部元素
	fmt.Println("sub[0] " + string(sub[0]))
	//第二个匹配的是第一个()里面的
	fmt.Println("sub[1] " + string(sub[1]))
	//第三个匹配的是第二个()里面的
	fmt.Println("sub[2] " + string(sub[2]))

	//FindAllSubmatch查找所有子匹配项
	all_sub := re2.FindAllSubmatch([]byte(data), 2)
	fmt.Println(all_sub)
	fmt.Println("all_sub[0][0] " + string(all_sub[0][0]))
	fmt.Println("all_sub[0][1] " + string(all_sub[0][1]))
	fmt.Println("all_sub[0][2] " + string(all_sub[0][2]))

	//FindSubmatchIndex用于查找子匹配项的开始位置和结束位置
	sub_ix := re2.FindSubmatchIndex([]byte(data))
	fmt.Print("sub_ix: ")
	fmt.Println(sub_ix)

	//FindAllSubmatchIndex查找所有子匹配项的开始位置和结束位置
	//n小于0返回全部，否则返回指定长度
	all_sub_ix := re2.FindAllSubmatchIndex([]byte(data), -1)
	fmt.Print("all_sub_ix: ")
	fmt.Println(all_sub_ix)

	//4、正则替换
	//通过函数进行替换
	re3, _ := regexp.Compile("a")
	rep := re3.ReplaceAllStringFunc(data, replFunc)
	// rep := re3.ReplaceAllStringFunc(data, strings.ToUpper)
	fmt.Print("rep: ")
	fmt.Println(rep)

	//把匹配的所有字符a替换成b
	rep2 := re3.ReplaceAllString(data, "b")
	fmt.Print("rep2: ")
	fmt.Println(rep2)

	multi_line_text := `I am a good man!  
			hello   world.`
	re4 := regexp.MustCompile(`(?i)\s+`)
	rep3 := re4.ReplaceAll([]byte(multi_line_text), []byte(" "))
	fmt.Print("rep3: ")
	// fmt.Println(rep3)
	fmt.Println(string(rep3))
}

func replFunc(s string) string {
	return strings.Repeat(s, 3)
}