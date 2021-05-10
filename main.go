/*
* PROJECT: go-project-10-05-2021
* DESCRIPTION: Given an integer X, return how many bits are active in the binary form of X.
* AUTHOR: Vahin Sharma
*/

package main

import "fmt"

func countActiveBits(input int) (activeBits int) {
  for _, bit := range fmt.Sprintf("%b", input) {
    if bit == '1' {activeBits++}
  }
  return
}

func main() {
  tests := map[int]int {
    0: 0,
    4: 1,
    7: 3,
    9: 2,
    10: 2,
    1234: 5,
    
    20042000202000: 19,
    1099511627776: 1,
    3824: 7,
    3776: 5,
  }
  score := 0
  
  for q, a := range tests {
    fmt.Println("Input:", q)
    if returnedA := countActiveBits(q); returnedA != a {
      fmt.Printf("Expected %d, instead got %d\n\n", a, returnedA)
    } else {
      score++
      fmt.Println("Passed\n")
    }
  }
  
  fmt.Printf("Score: %d out of %d\n", score, len(tests))
}
