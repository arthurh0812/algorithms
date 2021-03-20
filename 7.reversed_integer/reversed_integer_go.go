package main

import (
	"log"
	"math"
)

func reverse(x int, base int) (reversed int) {
	for x != 0 {
		d := pop(&x, base)
		push(&reversed, base, d)
	}
	if isOutOfBounds(reversed) {
		return 0
	}
	return reversed
}

func isOutOfBounds(rev int) bool {
	min, max := -math.Pow(2, 31), math.Pow(2, 31)-1
	if rev < int(min) || rev > int(max) {
		return true
	}
	return false
}

func pop(x *int, base int) (popped int) {
	d := *x % base
	*x /= base
	return d
}

func push(x *int, base int, d int) {
	*x *= base
	*x += d
}

func main() {
	rev := reverse(64, 2)
	log.Println(rev)
}
