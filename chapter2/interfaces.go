package main

import "fmt"

type Fruit interface {
	String() string
}

type Apple struct {
	Variety string
}

func (a Apple) String() string {
	return fmt.Sprintf("A %s apple", a.Variety)
}

type Orange struct {
	Size string
}

func (o Orange) String() string {
	return fmt.Sprintf("A %s orange", o.Size)
}

func PrintFruit(fruit Fruit) {
	fmt.Println("I have this fruit:", fruit.String())
}

// Error is a go interface and is any type which implements the Error method
type ErrInvalidStatusCode int
func (code ErrInvalidStatusCode) Error() string {
	return fmt.Sprintf("Expected 200 ok, but got %d", code)
}

func main() {
	apple := Apple{"Golden Delicious"}
	orange := Orange{"large"}

	PrintFruit(apple)
	PrintFruit(orange)

}