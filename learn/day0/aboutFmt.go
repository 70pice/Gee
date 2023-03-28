package main

import "fmt"

func main() {
	s1 := fmt.Sprint("70pice nb")
	s2 := fmt.Sprintf("70pice is %s", "nb")

	fmt.Print(s1, s2)
}
