package dvariable

import "fmt"

func RunVar1() {

	/*
		变量: dvariable
		概念: 一小块内存 用于存储数据 在运行过程中数值可以改变

		使用:
			step1: 变量的声明
			step2: 变量的访问
	*/
	// 第一种: 先定义变量 然后进行赋值
	var num1 int
	num1 = 27
	fmt.Printf("num1 = %d\n", num1)
	// 也可以写在一行
	var num2 int = 28
	fmt.Printf("num2 = %d\n", num2)

	// 第二种: 类型推断
	var name1 = "sumin"
	fmt.Printf("name1 is %s\n", name1)

	// 第三种: 简短声明
	sum1 := 100
	fmt.Printf("sum1 = %d\n", sum1)

	// 多变量声明
	// 第一种: 以","分隔 声明与赋值分开 若不赋值 存在默认值
	var a, b, c int
	a = 1
	b = 2
	fmt.Println(a, b, c)

	// 第二种: 直接赋值 可以是不同类型
	var m, n, o = 100, 100.2, "hello"
	fmt.Println(m, n, o)

	// 第三种: 集合类型
	var (
		username = "sumin"
		age      = 17
		sex      = "woman"
	)
	fmt.Printf("username = %s, age = %d, sex = %s\n", username, age, sex)

}
