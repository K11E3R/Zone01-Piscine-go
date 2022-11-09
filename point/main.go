package main

import "github.com/01-edu/z01"

func itoa(a int) string {
	var result string
	for a/10 != 0 {
		count := '0'
		for i := a % 10; i > 0; i-- {
			count++
		}
		result = string(count) + result
	}
	return result
}

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
	str := "x = " + itoa(points.x) + "y = " + itoa(points.y) + "\n"
	z01.PrintRune([]rune(str))
}
