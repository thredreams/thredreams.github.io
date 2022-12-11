package main

import (
	"errors"
	"fmt"
	"math"
)

func circleArea1(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func circleArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}
func main() {
	radius := -20.0
	//area, err := circleArea1(radius)
	area, err := circleArea2(radius)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Area of circle %0.2f", area)

}
