package math_example

// float64 array 를 받아서 Average 를 구하는 함수.
func Average(xs []float64) float64{
    total := float64(0)
    for _, x := range xs {
        total += x
    }
    return total/float64(len(xs))
}
