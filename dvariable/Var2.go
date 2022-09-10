package dvariable

import "fmt"

func RunVar2() {
	var num int
	num = 100
	fmt.Printf("num = %d, &num = %p\n", num, &num)

	num = 99
	fmt.Printf("num = %d, &num = %p\n", num, &num)
}
