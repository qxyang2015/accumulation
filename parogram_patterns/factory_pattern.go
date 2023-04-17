package parogram_patterns

import (
	"fmt"
)

/*
简单工厂模式
*/
type Person1 struct {
	Name string
	Age  int
}

func (p Person1) Greet() {
	fmt.Printf("Hi! My name is %s", p.Name)
}

func NewPerson1(name string, age int) *Person1 {
	return &Person1{
		Name: name,
		Age:  age,
	}
}

/*
抽象工厂
*/

type Person2 interface {
	Greet()
}

type person2 struct {
	name string
	age  int
}

func (p person2) Greet() {
	fmt.Printf("Hi! My name is %s", p.name)
}

// Here, NewPerson returns an interface, and not the person struct itself
func NewPerson(name string, age int) Person2 {
	return person2{
		name: name,
		age:  age,
	}
}

/*
在简单工厂模式中，依赖于唯一的工厂对象，如果我们需要实例化一个产品，就要向工厂中传入一个参数，获取对应的对象；如果要增加一种产品，就要在工厂中修改创建产品的函数。这会导致耦合性过高，这时我们就可以使用工厂方法模式。在
工厂方法模式中，依赖工厂函数，我们可以通过实现工厂函数来创建多种工厂，将对象创建从由一个对象负责所有具体类的实例化，变成由一群子类来负责对具体类的实例化，从而将过程解耦。
*/

/*
抽象工厂
*/

type Person struct {
	name string
	age  int
}

func NewPersonFactory(age int) func(name string) Person {
	return func(name string) Person {
		return Person{
			name: name,
			age:  age,
		}
	}
}

/*
可以使用此功能来创建具有默认年龄的工厂
*/
func TestNewPersonFactory() {
	newBaby := NewPersonFactory(1)
	baby := newBaby("john")

	newTeenager := NewPersonFactory(16)
	teen := newTeenager("jill")
	fmt.Println(baby, teen)
}
