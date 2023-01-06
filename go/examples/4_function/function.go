package main

import "fmt"

func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}

func f() (int, int) {
    return 5, 6
}

func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}

func addf(args ...float64) float64 {
    #NOTE(yoo): slice conversion 더 좋은 방법은??
    argsInt := make([]int, len(args))
    for i, v := range args {
        argsInt[i] = int(v)
    }
    total := add(argsInt...)
    return float64(total)
}

func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))

    x, y := f()
    fmt.Println(x, y)

    fmt.Println(add(1,2,3))
    fmt.Println(addf(xs...))
}
