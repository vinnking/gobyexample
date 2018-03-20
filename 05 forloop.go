package main

import "fmt"

// for有三种基本的循环类型

func main() {
	//1、单一条件循环，类似其他语言的while循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//2、经典的循环条件初始化、条件判断、循环后条件变化用法
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//3、无条件的for循环是死循环，除非使用break跳出循环
	//或者使用return从函数返回
	for {
		fmt.Println("loop")
		break
	}

}
