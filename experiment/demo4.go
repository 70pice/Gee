package main

import "fmt"

func demo5() {
	testOne := make(map[string]interface{})

	testOne["one"] = "qwe"
	testOne["two"] = "two"
	testTwo := "qwe"
	testThree := []int{1, 2, 3}
	for index, value := range testOne {
		fmt.Println(index)
		fmt.Println(value)
	}
	for index, value := range testTwo {
		fmt.Println(index)
		fmt.Println(value)
	}
	for index, value := range testThree {
		fmt.Println(index)
		fmt.Println(value)
	}
}
