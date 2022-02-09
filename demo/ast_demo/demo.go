package main

type Animal interface {
	Name(n string)
	Eat()
	Sleep()
}

type Duck interface {
	Name(n string)
	Eat()
	Sleep()
	Color() string
}
