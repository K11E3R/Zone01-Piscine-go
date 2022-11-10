package main

import "fmt"

type donnie struct {
	Name     string
	Life     float32
	Age      int
	Aircraft int
}

func main() {
	const AIRCRAFT1 = 1
	donnie := donnie{"Donnie", 100.0, 24, AIRCRAFT1}
	fmt.Println(donnie)
}
