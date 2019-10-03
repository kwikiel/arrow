package main

import (
	"fmt"
	"math"
)

const s string  = "constant"

func main() {
	fmt.Println(s)
	const n = 5000
	fmt.Println(math.Sin(n))
}
