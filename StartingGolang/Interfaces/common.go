package main

import "fmt"

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

func Println(s Stringify){
	fmt.Println(s.ToString())
}

// fmt.Printlnをオーバーライド
type T struct {
	Id int
	Name string
}
func (t *T) String() string {
	return fmt.Sprintf("<<%d, %s>>", t.Id, t.Name)
}

func main() {
	vs := []Stringify{
		&Person{ Name: "Taro", Age: 21 },
		&Car{Number: "XXX-0123", Model: "Tesla-X"},
	}
	for _, v := range vs {
		fmt.Println("v.ToString() =>", v.ToString())
	}
	fmt.Println("\n==================================================================\n")

	Println(&Person{Name: "Okamoto", Age: 99})
	Println(&Car{Number: "32-312", Model: "Pro Retina"})
	fmt.Println("\n==================================================================\n")

	// fmt.Printlnをオーバーライド
	t := &T{Id: 77, Name: "John Doe"}
	fmt.Println(t)
	fmt.Println("\n==================================================================\n")
}
