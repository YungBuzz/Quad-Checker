package main

import "github.com/01-edu/z01"

func main() {
	QuadD(5, 3)
	QuadD(5, 1)
	QuadD(1, 1)
	QuadD(1, 5)
}

func QuadD(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for c := 1; c <= y; c++ {
		for i := 1; i <= x; i++ {
			if c == 1 && i == 1 {
				z01.PrintRune('A')
			} else if c == 1 && i == x {
				z01.PrintRune('C')
			} else if c == y && i == 1 {
				z01.PrintRune('A')
			} else if c == y && i == x {
				z01.PrintRune('C')
			} else if c == 1 && i > 1 && i < x || c == y && i > 1 && i < x {
				z01.PrintRune('B')
			} else if c > 1 && i == 1 && c < y || c > 1 && i == x && c < y {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
