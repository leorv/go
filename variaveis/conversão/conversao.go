package main

import "fmt"

var a int = 20

func main() {
	var c float64 = float64(a)

	fmt.Printf("O valor de \"c\" é %g, e o valor de \"a\" é %d", c, a)

}
