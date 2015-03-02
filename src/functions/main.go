package main // package declaration.

import "fmt" //importing code from another package.


// A function!
// function_name (param_name param_type, ...) return_type
func average(xs []float64) float64 {
	//panic("Not Implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// A Variadic Function
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// A function that uses a closure to create an even-number generator.
func makeEvenGenerator() func() uint  {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// A function that makes an odd-number generator.
func makeOddGenerator() func() uint  {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		if i > 0 {
			i += 2
		} else {
			i += 1
		}
		return
	}
}

// Factorial via recursion.
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}


// Defer, panic, and Recover.
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}


func main() {
	xs := []float64{98, 93, 77, 82, 83}
	ys := []float64{101, 23, 45, 42, 128}
	fmt.Println(average(xs))
	fmt.Println(average(ys))

	fmt.Println("Sum of 1-5: ", add(1, 2, 3, 4, 5))
	fmt.Println("Sum of 1st 3 odd: ", add(1, 3, 5))

	s := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("Sum of a slice: ", add(s...))

	// Function defined within a function (main)
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println("add(1, 1) = ", add(1, 1))

	// A closure; due to it being defined in main with access to main's scope.
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println("increment() = ", increment())
	fmt.Println("increment() = ", increment())
	fmt.Println("increment() = ", increment())

	// Using a closure to generate even numbers
	nextEven := makeEvenGenerator()
	fmt.Println("Even:", nextEven())
	fmt.Println("Even:", nextEven())
	fmt.Println("Even:", nextEven())
	fmt.Println("Even:", nextEven())

	nextOdd := makeOddGenerator()
	fmt.Println("Odd:", nextOdd())
	fmt.Println("Odd:", nextOdd())
	fmt.Println("Odd:", nextOdd())
	fmt.Println("Odd:", nextOdd())
	fmt.Println("Odd:", nextOdd())

	// FActorial.
	fmt.Println("Factorial 10: ", factorial(10))
	fmt.Println("Factorial 25: ", factorial(25))

	// Defer panic and recover examples.
	defer second() // defer this function to the end.
	first()

	// Define, then defer execution; it should recover after the panic.
	defer func() {
		str := recover()
		fmt.Println("recovering from: ", str)
	}()
	panic("PANIC")
}
