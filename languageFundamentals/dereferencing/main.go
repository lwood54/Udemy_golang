package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // memory address in hex

	var b *int = &a
	fmt.Println(b)  // memory address in hex
	fmt.Println(*b) // 43

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as "dereferencing"
	// the * is an operator in this case

	*b = 42        // b says, "The value at this address, change it to 42"
	fmt.Println(a) // 42

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}
