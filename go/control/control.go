package main

import "fmt"

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i += 1
    }

    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "짝수")
        } else {
            fmt.Println(i, "홀수")
        }
    }
}
