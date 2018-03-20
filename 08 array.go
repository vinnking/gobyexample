package main

import "fmt"

//数组是具有相同数据类型的元素组成的固定长度的有序集合
//数组只能用 ==  != 比较，不能比较大小
//数组作为参数时，函数内部不会改变数组内部的值，除非是传入数组指针
// 数组的指针： *[3]int ; 指针数组：[2]*int

func main() {
	/*
	* 这里我们创建了一个具有5个元素的整型数组
	* 元素的数据类型和数据长度都是数组的一部分
	* 默认情况下，数组元素都是零值
	 */
	var a [5]int
	fmt.Println("emp:", a)

	//可以用索引来设置数组元素的值
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get", a[4])
	//内置len函数返回数组的长度
	fmt.Println("len:", len(a))

	//同时定义和初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//数组都是一维的，可以吧数组的元素定义为一个数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// test的测试
	var s = new([5]int)
	test(s)
	fmt.Println(s, len(s))

	//不定长数组,关联User struct
	m := [...]User{
		{0, "User0"},
		{8, "User8"},
	}

	n := [...]*User{
		{0, "User0"},
		{8, "User8"},
	}

	fmt.Println(m, len(m))
	fmt.Println(n, len(n))
}

//可以用new创建数组，并返回数组的指针
func test(s *[5]int) {
	s[1] = 5
}

type User struct {
	Id   int
	Name string
}
