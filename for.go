package main

import "fmt"

func main() {
	i:= 1

	for i<= 3 {
		fmt.Println(i)
		i++
	}

	for j := 7; j<= 9; j++ {
		fmt.Println(j)
		// Standard for syntax with colons
	}

	for {
		fmt.Println("loop")
		break // break breaks duh
	}

	for n:=0; n<= 5;  n++ {
		if n%2 ==0 {
			continue
		}
		fmt.Println(n)
	}
}
