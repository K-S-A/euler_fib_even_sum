package main

import "os"
import "strconv"
import "fmt"

func FibEvenSum(maxBorder uint64) uint64 {
  var n, n1 uint64 = 1, 1

  for !(n1 >= maxBorder && n1 % 2 == 0) {
    n, n1 = n1, n + n1
  }

  return (n - 1) / 2
}

func MaxBorder(args []string) (uint64, error) {
  return strconv.ParseUint(args[1], 10, 64)
}

func main() {
  max, err := MaxBorder(os.Args)

  if err != nil {
    fmt.Printf("Invalid max border %q.\n", os.Args[1])
  } else {
    fmt.Printf("Sum of all even Fibonacci numbers, that less than %d is %d.\n", max, FibEvenSum(max))
  }
}
