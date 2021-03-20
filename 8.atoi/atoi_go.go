package main

import (
	"log"
	"math"
)

func myAtoi(s string, base int) (n int) {
	s = trimLeadingWhitespace(s)
	var negative bool
	if len(s) > 1 {
		first := s[0]
		if first == '-' || first == '+' {
			if first == '-' {
				negative = true
			}
			s = s[1:]
		}
	}
	for _, r := range s {
		if r == '-' || r == '+' {
			return 0
		}
		if r >= '0' && r <= '9' {
			digit := r - '0'
			n *= base
			if negative {
				n -= int(digit)
			} else {
				n += int(digit)
			}
			continue
		}
		break
	}
	clamp(&n)
	return
}

func trimLeadingWhitespace(s string) string {
	last := 0
	for i, r := range s {
		if r != ' ' {
			break
		}
		last = i + 1
	}
	return s[last:]
}

func clamp(n *int) {
	min, max := int(-math.Pow(2, 31)), int(math.Pow(2, 31)-1)
	if *n < min {
		*n = min
		return
	}
	if *n > max {
		*n = max
	}
}

func main() {
	n := myAtoi("   -45290xfge", 10)
	log.Println(n)
	n = myAtoi("     -101011101 ", 2)
	log.Println(n)
	n = myAtoi("9223372036854775808", 10)
	log.Println(n)
}
