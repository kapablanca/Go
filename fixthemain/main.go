package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

const (
	OPEN  = "OPEN"
	CLOSE = "CLOSE"
)

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n') // Assume we need a newline after printing the string
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

func IsDoorOpen(door *Door) bool {
	PrintStr("is the Door opened ?")
	return door.state == OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{state: CLOSE} // Initialize door as closed

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}