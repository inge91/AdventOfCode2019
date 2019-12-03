package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type point struct {
	X int
	Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(wires [][]string) {
	m1 := getDirection(wires[0])
	m2 := getDirection(wires[1])
	m := make(map[point]int)
	for k, _ := range m1 {
		m[k]++
	}
	for k, _ := range m2 {
		m[k]++
	}

	s := make([]point, 0)
	for key, el := range m {
		if el > 1 {
			s = append(s, key)
		}
	}
	sort.Slice(s, func(i, j int) bool {
		return abs(s[i].X)+abs(s[i].Y) < abs(s[j].X)+abs(s[j].Y)
	})

	answer := abs(s[0].X) + abs(s[0].Y)
	fmt.Println(answer)
}

func getIncrement(d byte) point {
	switch d {
	case 'U':
		{
			return point{0, 1}
		}
	case 'D':
		{
			return point{0, -1}
		}
	case 'L':
		{
			return point{-1, 0}
		}
	case 'R':
		{
			return point{1, 0}
		}
	}
	return point{0, 0}
}

func getDirection(wire []string) map[point]int {
	m := make(map[point]int)
	c := 1 //Global count of all steps taken
	last := point{0, 0}
	for _, direction := range wire {
		d := getIncrement(direction[0])
		amount, _ := strconv.Atoi(direction[1:])
		for i := 1; i <= amount; i++ {
			newPoint := point{last.X + d.X, last.Y + d.Y}
			_, ok := m[newPoint]
			if !ok {
				m[newPoint] = c
			}
			last = newPoint
			c++
		}
	}
	return m
}

func part2(wires [][]string) {
	m1 := getDirection(wires[0])
	m2 := getDirection(wires[1])

	minV := math.MaxInt64
	for key := range m1 {
		if m2[key] > 0 {
			v := m2[key] + m1[key]
			if v < minV {
				minV = v
			}
		}
	}
	fmt.Println(minV)
}

func main() {
	wires := make([][]string, 0)
	f, _ := os.Open("input")
	s := bufio.NewScanner(f)
	for s.Scan() {
		wireRoute := s.Text()
		wires = append(wires, strings.Split(wireRoute, ","))
	}
	part1(wires)
	part2(wires)
}
