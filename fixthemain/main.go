package main

import "github.com/01-edu/z01"

type Door struct {
	state bool
}

const (
	OPEN  = false
	CLOSE = true
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptrDoor *Door) {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door closing...")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor Door) bool {
	PrintStr("Door is open ?")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor Door) bool {
	PrintStr("Door is close ?")
	return ptrDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}
