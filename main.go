package main

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) movePoint(x, y int) Point {
	p.X += x
	p.Y += y
	return p // don't use this
}

func (p *Point) movePointPtr(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)
	fmt.Println(p.movePoint(1, 1))
	fmt.Println(p)
	p.movePointPtr(1, 1)
	fmt.Println(p)

	v := &p
	v.movePoint(1, 1)
	fmt.Println(p)
	v.movePointPtr(2, 3)
	fmt.Println(p)
}
