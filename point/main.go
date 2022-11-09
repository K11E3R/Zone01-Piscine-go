package main

import "github.com/01-edu/z01"

func itoa(a int) string {
	var result string
	if a < 0 {
		a = -a
	}
	for a/10 != 0 && a%10 != 0 {
		count := '0'
		for i := a % 10; i > 0; i-- {
			count++
		}
		result = string(count) + result
		a = a / 10
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
	str := "x = " + itoa(points.x) + ", y = " + itoa(points.y) + "\n"
	for _, elem := range str {
		z01.PrintRune(elem)
	}
}
