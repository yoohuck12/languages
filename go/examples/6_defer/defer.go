//일반적으로 패닉은 프로그래밍 오류를 나타내거나(가령 범위를
//벗어난 배열 인덱스에 접근하려고 시도하거나 맵을 초기화하는
//것을 잊어버리는 등) 손쉽게 복구할 수 없는 예외적인 상황을
//나타낸다(그래서 "패닉"이라고 하는 것이다).
package main

import "fmt"

func main() {
    defer func() {
        str := recover()
        fmt.Println(str)
    }()
    panic("PANIC")
}
