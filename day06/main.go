package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name   string
	parent *Node
}

func findAllOrbitees(n *Node) int {
	if n.parent == nil {
		return 0
	}
	return findAllOrbitees(n.parent) + 1
}

func getPath(node *Node) []string {
	if node.parent == nil {
		return []string{}
	}
	s := []string{node.parent.name}
	return append(getPath(node.parent), s...)
}

func createOrGet(n string, m map[string]*Node) *Node {
	v, ok := m[n]
	if ok {
		return v
	}
	newN := Node{name: n, parent: nil}
	m[n] = &newN
	return &newN
}

func main() {
	f, _ := os.Open("input")
	s := bufio.NewScanner(f)
	// Map going from orbittee to orbitter
	m := make(map[string]*Node)
	for s.Scan() {
		relation := s.Text()
		s := strings.Split(relation, ")")
		fmt.Println(s)
		orbiteeN, orbiterN := s[0], s[1]
		orbitee := createOrGet(orbiteeN, m)
		orbiter := createOrGet(orbiterN, m)
		orbiter.parent = orbitee
	}

	// Part 1
	v := 0
	// Go through the tree to find out all orbitees
	for n, orbiter := range m {
		all := findAllOrbitees(orbiter)
		fmt.Printf("%s for %d\n", n, all)
		v += all
	}
	fmt.Println(v)

	// Part 2
	p1 := getPath(m["YOU"])
	p2 := getPath(m["SAN"])
	ind := 0
	for i := range p1 {
		if p1[i] == p2[i] {
			ind = i
		} else {
			break
		}
	}
	// Calculate the steps between orbits (-2 to not include "SAN and YOU")
	out := len(p1) - ind + len(p2) - ind - 2
	fmt.Println(out)
}
