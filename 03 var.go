package main

import "fmt"

func main() {
	//var关键字可以用来定义一个或者多个变量
	var a string = "Hello"
	fmt.Println(a)

	//一次可以定义多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	//Go会推断出具有初始值的变量的类型
	var d = true
	fmt.Println(d)

	//定义变量的时候，如果没有给出初始值，那么变量默认初始化为0值
	var e int
	fmt.Println(e)

	//":="语法是同时定义和初始化变量的方式，只能用于函数内部
	f := "short"
	fmt.Println(f)
}
