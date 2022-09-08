package constant

import "fmt"

func RunConst1() {

	// 定义常量
	const PATH string = "https://www.bilibili.com"
	const PI = 3.14
	fmt.Println(PATH, PI)

	// 定义一组常量
	const C1, C2, C3 = 100, 3.14, "sumin"
	const (
		MALE   = 0
		FEMALE = 1
		UNKNOW = 3
	)

	// 一组常量中 如果某个常量没有初始值 默认和上一行一致
	const (
		a int = 7
		b
		c string = "halo"
	)
	fmt.Println(a, b, c)

	// 枚举类型
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)

}
