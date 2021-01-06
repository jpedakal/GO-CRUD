package main 

import "fmt"

func add(a, b int) int {
  return a+b
}

func main() {
  sum := add(4, 5)
  fmt.Printf("Hello, Gopher! Your result is %d\n", sum)
}