package core

import (
	"fmt"
)


// defined for fecund n
func ThreeChild(n int) int {
	r58 := n % 54
	var m int
	switch r58 {
	case 34:
		m = 2
	case 52:
		m = 2
	case 4:
		m = 4
	case 40:
		m = 4
	case 16:
		m = 8
	default: // case 4
		m = 16
	}
	c := m * (n - 1) / 3

	fmt.Println(c)
	return c
}

// defined for fecund n
func TwoChild(n int) int {
	var c int = n * (n % 18)

	return c

}

