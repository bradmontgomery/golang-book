package main // package declaration.

import "fmt" //importing code from another package.


// Pass by reference using a Pointer.
func zero(xPtr *int) {
	*xPtr = 0  // Dereferencing; access the value to which the pointer, points.
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func swap(xPtr *int, yPtr *int) {
	tPtr := new(int)
	*tPtr = *xPtr
	*xPtr = *yPtr
	*yPtr = *tPtr
}

func main() {
	x := 5
	fmt.Println("x =", x) // x is 5
	zero(&x) // & address of x
	fmt.Println("x =", x) // x is 0

	// Use `new` to get a Pointer.
	xPtr := new(int)
	fmt.Println(*xPtr)  // appears to be 0; is that normal?
	one(xPtr)
	fmt.Println(*xPtr)  // dereference to get it's value.

	// Square
	y := 1.5
	square(&y)
	fmt.Println("Square of 1.5 =", y)

	// Swapping values
	a := 1
	b := 2
	fmt.Println("a =", a, "; b =", b)
	swap(&a, &b)
	fmt.Println("a =", a, "; b =", b)
}
