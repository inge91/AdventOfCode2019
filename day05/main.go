package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverse(p []int) {
	for i := 0; i < len(p)/2; i++ {
		p[i], p[len(p)-i-1] = p[len(p)-i-1], p[i]
	}
}

func toPaddedSlice(v int, l int) []int {
	p := make([]int, l)
	i := 0
	for v > 0 {
		e := v % 10
		v /= 10
		p[i] = e
		i++
	}
	reverse(p)
	return p
}

func compute(sl []int, io int) int {
	for i := 0; i < len(sl); {
		fmt.Println(sl)
		s := toPaddedSlice(sl[i], 4)
		t := s[len(s)-1] // Instruction type
		// Instructions without argment
		if t == 9 {
			return io
		}

		// single argument instructions
		a1 := sl[i+1]
		// 2 argument instructions
		switch t {
		case 3:
			{
				sl[a1] = io
				i += 2
				continue
			}
		case 4:
			{
				v := 0
				if s[1] == 0 {
					v = sl[a1]
				} else {
					v = a1
				}
				io = v
				i += 2
				continue
			}
		}
		a2 := sl[i+2]
		v1 := 0
		v2 := 0
		// First element mode:
		if s[1] == 0 {
			v1 = sl[a1] // position mode
		} else {
			v1 = a1 // immediate mode
		}
		if s[0] == 0 {
			v2 = sl[a2] // position mode
		} else {
			v2 = a2 // immediate mode
		}

		// Instruction type
		switch t {
		case 5:
			if v1 != 0 {
				i = v2
			} else {
				i += 3
			}
			continue
		case 6:
			if v1 == 0 {
				i = v2
			} else {
				i += 3
			}
			continue
		}

		// 3 arguments
		writePos := sl[i+3]
		switch t {
		case 1:
			sl[writePos] = v1 + v2
			i += 4
		case 2:
			sl[writePos] = v1 * v2
			i += 4
		case 7:
			if v1 < v2 {
				sl[writePos] = 1
			} else {
				sl[writePos] = 0
			}
			i += 4
		case 8:
			if v1 == v2 {
				sl[writePos] = 1
			} else {
				sl[writePos] = 0
			}
			i += 4
		}
	}
	return io
}

func main() {
	f, _ := os.Open("input")
	s := bufio.NewScanner(f)
	s.Scan()
	t := s.Text()
	l := strings.Split(t, ",")
	sl := make([]int, len(l))
	for i := range l {
		sl[i], _ = strconv.Atoi(l[i])
	}

	//fmt.Print(compute(sl, 1))

	fmt.Print(compute(sl, 5))

}
