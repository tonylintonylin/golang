package main

import (
	"fmt"
)

var o float32 = 44
var actorName string = "tony lin"

func main() {
	i := 42
	var j float32 = 27
	k := 99.

	fmt.Printf("%v, %T", o, o)
	fmt.Println(actorName)
	fmt.Printf("%v, %T", i, i)
	fmt.Printf("%v, %T", j, j)
	fmt.Printf("%v, %T", k, k)

	// fmt.Println(quote.Hello())
	// fmt.Println(quote.Opt())
}
