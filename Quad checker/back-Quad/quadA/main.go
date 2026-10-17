package main

import "github.com/01-edu/z01"

func main() {
	QuadA(5, 3)
	QuadA(5, 1)
	QuadA(1, 1)
	QuadA(1, 5)
}

func QuadA(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for c := 1; c <= y; c++ {
		for i := 1; i <= x; i++ {
			if c == 1 && i == 1 {
				z01.PrintRune('o')
			} else if c == 1 && i == x {
				z01.PrintRune('o')
			} else if c == y && i == 1 {
				z01.PrintRune('o')
			} else if c == y && i == x {
				z01.PrintRune('o')
			} else if c == 1 && i > 1 && i < x || c == y && i > 1 && i < x {
				z01.PrintRune('-')
			} else if c > 1 && i == 1 && c < y || c > 1 && i == x && c < y {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
