package main

import "fmt"

type Person struct {
  name string
}

// SetName receives a pointer to Foo so it can modify it.
func (p *Person) SetName(name string) {
	p.name = name
}

// GetName receives a copy of Foo since it doesn't need to modify it.
func (p Person) GetName() string {
	return p.name
}

func main() {

	p := Person{}
	p.SetName("Abc")
	name := p.GetName()
	fmt.Println(name)
}
