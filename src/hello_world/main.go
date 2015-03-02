package main // package declaration.

import "fmt" //importing code from another package.

func main() {
	// multiple ways to declare/assign variables
	var x string = "Hello World"
	y := "Woo"
	z := "Doggy"
	sum := 0

	// Array of even numbers
	var evens  [5]int

    fmt.Println(x)
	fmt.Println("And now for some maths:")
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1 + 1 =", 1.0 + 1.0)
	fmt.Println(x[1])
	fmt.Println(x + y + z)
	fmt.Println(true || false)

	// Even numbers
	n := 0
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			evens[n] = i
			n += 1
		}
	}
	fmt.Println("Evens: ", evens)

	// Sum of even numbers
	for _, value := range evens {
		sum += value
	}
	fmt.Println("Sum of evens: ", sum)

	// Slicing arrays.
	fmt.Println("A slice of evens: ", evens[1:4])

	// Maps (like python dicts)
	// maps have to be initialized before they can be used; so this doesn't work:
	// var ages = map[string]int
	ages := make(map[string]int)  // initializes the map first.
	ages["brad"] = 36
	ages["dakota"] = 11
	ages["cheyenne"] = 9
	fmt.Println(ages)
	fmt.Println("There are ", len(ages), "items in the map")
	ages["louie"] = 5
	fmt.Println("There are now ", len(ages), "items in the map")

	// attempt a lookup and do something if it was successful.
	if age, ok := ages["brad"]; ok {
		fmt.Println("Brad is: ", age)
		delete(ages, "brad")
	}
	fmt.Println(len(ages), "items in the map: ", ages)
}
