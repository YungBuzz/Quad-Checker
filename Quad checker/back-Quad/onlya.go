package main

func only1(s string) int {
	count := 0
	for _, c := range s {
		if c == '1' {
			count++
		}
	}
	if count == 1 {
		return 1
	}
	return 0

}
