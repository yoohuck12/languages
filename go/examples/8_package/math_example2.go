package math_example

func Add(xs []float64) float64 {
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total
}
