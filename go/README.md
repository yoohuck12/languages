# go 공부
- https://codingnuri.com/golang-book/2.html

## 주의하기
- 같은 폴더 내에 package main 이 있고, main 함수가 두 개이상 들어있는
파일이 있으면 `main redeclared in this block previous declaation at X`
에러가 발생함.

## 명령어 모음
```bash
go run hello.go
go build hello.go

go install golang.org/x/lint/golint@latest
golint hello.go

go doc fmt.Println

go install golang.org/x/tools/cmd/godoc
godoc -http=:6060
```

