package main

import (
	"fmt"
	"math"
)

func main() {
	var number float32 = -35.6
	fmt.Println(number)

	bits := math.Float32bits(number)
	binary := fmt.Sprintf("%32b", bits)
	fmt.Printf("Bit:%s,%s,%s", binary[0:1], binary[1:9], binary[9:32])

}
