package main

import "fmt"

func main() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	power, exists := lookup["goku"]
	// prints 0, false
	// 0 is the default value for an integer
	fmt.Println(power, exists)
	delete(lookup, "goku")
	fmt.Println(lookup)
}

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

type Conqueror struct {
	*Saiyan
	ConquerState bool
}
