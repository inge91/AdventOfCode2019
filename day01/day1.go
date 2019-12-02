package main

import "fmt"
import "os"
import "bufio"
import "strconv"

// GetRequiredFull Calculates the fuel required given the input mass
func GetRequiredFull(m int) int {
  return m / 3 - 2
}

// GetRecursiveRequiredFull calculates the fuel required given
// the input mass, and the fuel mass itself.
func GetRecursiveRequiredFull(m int) int {
  totFuel := 0
  f := GetRequiredFull(m)
  for f > 0 {
    totFuel += f
    f = GetRequiredFull(f)
  }
  return totFuel
}

func part1(s []int) int{
  totFuel := 0
  for _, mass := range s {
    totFuel += GetRequiredFull(mass)
  }
  return totFuel
}

 func part2(s []int) int{
   totFuel := 0
   for _, mass := range s {
     totFuel += GetRecursiveRequiredFull(mass)
   }
   return totFuel
}

func main() {
  sl := make([]int, 0)
  f, _ := os.Open("input")
  s := bufio.NewScanner(f)
  for s.Scan() {
    mass := s.Text()
    massI, _ := strconv.Atoi(mass)
    sl = append(sl, massI)
  }

  fmt.Println(part1(sl))
  fmt.Println(part2(sl))
}
