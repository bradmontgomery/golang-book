package main

import ("fmt"; "math")

/* -----------------------------
This is tedious Without Structs.
----------------------------- */

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectagleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}


// How about an Interface. Explicitly requires that structs have the listed
// methods implemented.
type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

// A Circle Method.
// (c *Circle) is a "receiver", and allows us to call the method using dot
// notation: c.area()
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	// aka circumference: 2 * pi * r
	return math.Pi * 2 * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	return r.x1 + r.x2 + r.y1 + r.y2
}

// Using the Interface in a Function Definition; This can work with both
// Circles and Rectangles.
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.perimeter()
	}
	return area
}

// Embedded Types (aka anonymous files; is-a relationships, rather than has-a)
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("hi, myname is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {

	/* -------------------------------------------------
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectagleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))
	------------------------------------------------- */

	c := Circle{x: 0, y: 0, r: 5}
	//fmt.Println("Circle Area: ", circleArea(c))
	fmt.Println("Circle Area:", c.area())
	fmt.Println("Circle Perimeter:", c.perimeter())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Rectangle Area:", r.area())
	fmt.Println("Rectangle Perimeter:", r.perimeter())

	fmt.Println("Total Area:", totalArea(&c, &r))
	fmt.Println("Total Perimeter:", totalPerimeter(&c, &r))

	// Using embedded types
	a := new(Android)
	a.Name = "Jim"
	a.Person.Talk()
	// or...
	a.Talk()
}
