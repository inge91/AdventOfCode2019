package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compute(sl []int) {
	for i := 0; i < len(sl)-4; i += 4 {
		if sl[i] == 99 {
			return
		}
		pos1 := sl[i+1]
		pos2 := sl[i+2]
		pos3 := sl[i+3]
		if pos3 >= len(sl) || pos2 >= len(sl) || pos1 >= len(sl) {
			return
		}
		switch sl[i] {
		case 1:
			sl[pos3] = sl[pos1] + sl[pos2]
		case 2:
			sl[pos3] = sl[pos1] * sl[pos2]
		}
	}
}

func compute2(sl []int) {
	tmp := make([]int, len(sl))
	for i := 0; i < 100; i++ {
		for j := 0; j < 00; j++ {
			copy(tmp, sl)

			tmp[1] = i
			tmp[2] = j
			compute(tmp)
			if tmp[0] == 19690720 {
				fmt.Println(tmp)
				fmt.Println((100 * i) + j)
				break
			}
		}
	}
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

	tmp := make([]int, len(sl))
	copy(tmp, sl)
	compute(sl)
	fmt.Println(sl)
	compute2(tmp)

}
