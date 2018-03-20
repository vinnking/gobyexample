package main

import (
	"fmt"
	"time"
)

//判断分支太多的时候，可以使用switch语句来优化逻辑

func main() {

	//基础switch用法
	i := 2
	fmt.Print("write ", i, " as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//你可以使用逗号在case中分开多个条件，还可以使用default语句
	//当上面的case都没有满足的时候执行default语句
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend")
	default:
		fmt.Println("It is a weekday")
	}

	//当switch没有跟表达式的时候，功能和if/else相同，
	//可以看到case后面的表达式不一定是常量
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}
}
