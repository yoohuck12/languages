package main

import (
    "fmt"
    "time"
)

func pinger(c1 chan<- string) {
    for i := 0; ; i++ {
        c1 <- "ping"
        time.Sleep(time.Second * 2)
    }
}

func ponger(c2 chan<- string) {
    for i := 0; ; i++ {
        c2 <- "pong"
        time.Sleep(time.Second * 3)
    }
}

//See improving code in
// https://etloveguitar.tistory.com/53
func printer(c1, c2 <-chan string) {
    for {
        select {
            case msg := <- c1:
                fmt.Println("c1", msg)
            case msg := <- c2:
                fmt.Println("c2", msg)
            case <- time.After(time.Second):
                fmt.Println("timeout")
        }
    }
}

func main() {
    var c1 chan string = make(chan string)
    var c2 chan string = make(chan string)

    go pinger(c1)
    go ponger(c2)
    go printer(c1, c2)

    var input string
    fmt.Scanln(&input)
}
