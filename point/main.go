package main

import "github.com/01-edu/z01.PrintRune"

type point struct {
	x int
	y int
}

func setPoint(ptr point) point {
	ptr.x = 42
	ptr.y = 21
	return ptr
}

func main() {
	points := point{}
	points = setPoint(points)
	PrintRune([]rune("x = %d, y = %d\n", points.x, points.y))
}
 