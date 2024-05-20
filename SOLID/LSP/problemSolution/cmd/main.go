package main

import (
	"fmt"

	"github.com/douglasandradeee/go-expert-course/SOLID/LSP/problemSolution/quadrilateral"
)

func main() {
	square := &quadrilateral.Square{}
	rectangle := &quadrilateral.Rectangle{}

	square.SetSize(2.5)
	rectangle.SetWidth(2.5)
	rectangle.SetHeight(3.7)

	fmt.Printf("Get size square: %.2f\n", square.GetSize())
	fmt.Printf("Get width rectangle: %.2f\n", rectangle.GetWidth())
	fmt.Printf("Get Height rectangle: %.2f\n", rectangle.GetHeight())

	fmt.Printf("Square area: %.2f\n", square.GetArea())
	fmt.Printf("Rectangle area: %.2f\n", rectangle.GetArea())

}
