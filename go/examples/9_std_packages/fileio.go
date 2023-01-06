package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        // 오류를 처리
        return
    }
    defer file.Close()

    // 파일의 크기를 구함
    stat, err := file.Stat()
    if err != nil {
        return
    }

    // 파일을 읽음
    bs := make([]byte, stat.Size())
    _, err = file.Read(bs)
    if err != nil {
        return
    }

    str := string(bs)
    fmt.Println(str)
}
