# Golang中的继承

简单来说，结构体中给了名字就是，那么就当做属性使用即可，如果没有给属性，就是继承

```go
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
	People
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
	t := Teacher{}
	t.showA()
	t.People.showA()
	t.showB()
	t.People.showB()

	s := Student{}
	s.P.showA()

	st := SpecialTeacher{}
	st.showA()
	st.showB()
}
/*
peo showA
peo showB    
peo showA    
peo showB    
teacher showB
peo showB    
peo showA    
peo showB    
peo showA    
peo showB    
teacher showB

*/
```

