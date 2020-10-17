package main

import "fmt"

func genDisplaceFn(a, v0, s0 float64) func(float64) float64 {
  return func(t float64) float64 {
    return 1/2.0*a*t*t + v0*t + s0
  }
}

func main() {
  fn1 := genDisplaceFn(10, 2, 1)
  fmt.Println(fn1(1))

  fn2 := genDisplaceFn(5, 5, 3)
  fmt.Println(fn2(1))
}
