package main

import "fmt"

func main() {
    var x [5]float64
    x[4] = 100
    fmt.Println(x)

    var total float64 = 0
    for i := 0; i < len(x); i++ {
        total += x[i]
    }
    fmt.Println(total / float64(len(x)))

    var total2 float64 = 0
    for _, value := range x {
        total2 += value
    }
    fmt.Println(total2 / float64(len(x)))

    var dx []float64 //dynamic x
    ndx := make([]float64, 5, 10)  //new dx
    slice := x[0:3]
    fmt.Println(dx, ndx, slice)

    slice1 := []int{1,2,3}
    slice2 := append(slice1, 4, 5)
    fmt.Println(slice1, slice2) //[123], [12345]

    slice1 = []int{1,2,3}
    slice2 = make([]int, 2)
    copy(slice2, slice1)
    fmt.Println(slice1, slice2)  //[123] [12]
}
