package main

import "fmt"

type People struct {
}

func (p *People) showA() {
	fmt.Println("peo showA")
	p.showB()
}

func (p *People) showB() {
	fmt.Println("peo showB")
}

type Teacher struct {
	*People
}

func (t *Teacher) showB() {
	fmt.Println("teacher showB")
}

type Student struct {
	P People
}

func (s *Student) showB() {
	fmt.Println("stu showB")
}

type SpecialTeacher struct {
	Teacher
}

func main() {
	t := &Teacher{}
	t.showA()
	t.showB()

	fmt.Printf("%T", t)

	fmt.Println(t)

	//var x int
	//var y int
	//var z *int
	//var c *int
	//x = 1
	//y = 2
	//z = &x
	//c = &y
	//fmt.Println(&x)
	//fmt.Println(*z)
	//fmt.Printf("%T", z)
	//fmt.Println(&y)
	//fmt.Println(*c)
}
