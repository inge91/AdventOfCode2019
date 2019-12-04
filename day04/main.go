package main

import "fmt"

func numsIncrease(p []int) bool {
	for i := range p {
		if i == 0 {
			continue
		}
		if p[i-1] > p[i] {
			return false
		}
	}
	return true
}

func containsAdjacentDouble(p []int) bool {
	currentValue := p[0]
	for i, v := range p {
		if i == 0 {
			continue
		}
		if v == currentValue {
			return true
		}
		currentValue = v
	}
	return false
}

func isPassword(p []int) bool {
	return numsIncrease(p) && containsAdjacentDouble(p)
}

func reverse(p []int) {
	for i := 0; i < len(p)/2; i++ {
		p[i], p[len(p)-i-1] = p[len(p)-i-1], p[i]
	}
}

func toSlice(v int) []int {
	p := make([]int, 0)
	for v > 0 {
		e := v % 10
		v /= 10
		p = append(p, e)
	}
	reverse(p)
	return p
}

func onlySingleAdjacent(p []int) bool {
	currentValue := p[0]
	count := 1
	for i, v := range p {
		if i == 0 {
			continue
		}
		if v == currentValue {
			count++
			continue
		}
		// We just left from a double
		if count == 2 {
			return true
		}
		currentValue = v
		count = 1
	}
	return count == 2
}

func main() {
	lowBound := 171309
	upBound := 643603
	countPart1 := 0
	countPart2 := 0
	for i := lowBound; i <= upBound; i++ {
		s := toSlice(i)
		if isPassword(s) {
			countPart1++
			if onlySingleAdjacent(s) {
				countPart2++
			}
		}
	}
	fmt.Println(countPart1)
	fmt.Println(countPart2)
}
