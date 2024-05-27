package example_test

import (
	"fmt"
	example "main/jobs"
	"testing"
	"time"
)

func TestExamples(t *testing.T) {
	engine := example.NewEngine(
		1000,
		4,
		"gas",
	)

	engine.Start()

	time.Sleep(time.Second * 3)

	engine.Stop()
}

func TestShapes(t *testing.T) {
	// create an array of shapes
	shapes := []example.Shape{
		example.Circle{Radius: 5},
		example.Square{Side: 4},
	}

	// loop through the array and print the perimeter and area
	for _, shape := range shapes {
		fmt.Printf("Perimeter: %.2f\n", shape.CalcPerimeter())
		fmt.Printf("Area: %.2f\n\n", shape.CalcArea())
	}
}
