package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {

	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	// 把container的值转换为空接口：interface{}(container)
	// 判断前者的类型是否为切片类型： .([]string)
	value, ok := interface{}(container).([]string)

	if ok == true {

		// 那么被判断的值将会被自动转换为[]string类型的值，并赋给变量value
	} else {

		// 否则value将被赋予nil（即“空”）。
	}

	fmt.Println(value, ok)

	fmt.Printf("The element is %q.\n", container[1])
}
