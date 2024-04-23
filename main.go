package main

import (
	"fmt"
)

// Example of a basic function.
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Example of a function with multiple return values.
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return dividend / divisor, nil
}

// Example of a struct.
type Person struct {
	Name string
	Age  int
}

// Example of a method on a struct.
func (p Person) introduce() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Example of a simple interface.
type Shape interface {
	area() float64
}

// Example of a struct implementing an interface.
type Rectangle struct {
	Width  float64
	Height float64
}

// Implementation of the area method for Rectangle.
func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

// Main function.
func main() {
	// Calling a function.
	greet("World")

	// Calling a function with multiple return values.
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Creating an instance of a struct.
	person := Person{Name: "John", Age: 30}

	// Calling a method on a struct.
	person.introduce()

	// Creating an instance of a struct that implements an interface.
	rectangle := Rectangle{Width: 5, Height: 3}

	// Calling a method on a struct that implements an interface.
	fmt.Println("Area of rectangle:", rectangle.area())

	// Example of a for loop with range.
	names := []string{"Alice", "Bob", "Charlie"}
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}

	// Example of a for loop without range.
	for i := 0; i < 3; i++ {
		fmt.Println("Iteration:", i)
	}
}
