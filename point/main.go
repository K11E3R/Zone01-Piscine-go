package main

import "github.com/01-edu/z01"

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
	z01.PrintRune([]rune("x = %d, y = %d\n", points.x, points.y))
}
