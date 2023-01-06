package main

import "fmt"
import "examples/examples/8_package"

func main() {
    xs := []float64{1,2,3,4}
    fmt.Println(math_example.Average(xs))
    fmt.Println(math_example.Add(xs))
}
