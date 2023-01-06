package main

import "fmt"

func main() {
    var x string = "Hello World"
    fmt.Println(x)

    x = "hello"
    var y string = "world"
    fmt.Println(x == y)

    var i = 5  //변수 타입 없이 선언
    z := 5  // 변수 타입 및 var 없이 선언
    fmt.Println(i,z)

    const cx string = "Hello World"
    fmt.Println(cx)

    var (
        a = 5
        b = 10
        c = 15
    )
    fmt.Println(a,b,c)

    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input * 2
    fmt.Println(output)
}
