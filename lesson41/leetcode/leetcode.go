package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	intMax := int(math.Pow(2, 31) - 1)
	intMin := int(math.Pow(-2, 31))

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	r := 1
	var isminus int

	if s[0] == '-' {
		r = -1
		isminus = 1
	} else if s[0] == '+' {
		isminus = 1
	}

	var a int
	for i := isminus; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		intv, err := strconv.Atoi(string(s[i]))
		if err != nil {
			break
		}
		a = a*10 + intv

		if r*a > intMax {
			return intMax
		}
		if r*a < intMin {
			return intMin
		}
	}

	a = a * r
	return a
}

func main() {
	fmt.Println(myAtoi("-91283472332"))
}
