# go 공부
- https://codingnuri.com/golang-book/2.html

## 주의하기
- 같은 폴더 내에 package main 이 있고, main 함수가 두 개이상 들어있는
파일이 있으면 `main redeclared in this block previous declaation at X`
에러가 발생함.
- 같은 폴더 내에 여러 package 를 추가할 수 없음. 같은 이름의 패키지로는
여러 함수가 추가되나, 다른 패키지가 같은 폴더 안에 있으면 안됨.

## 명령어 모음
```bash
go run hello.go
go build hello.go

go install golang.org/x/lint/golint@latest
golint hello.go

go doc fmt.Println

go install golang.org/x/tools/cmd/godoc
godoc -http=:6060

# directory 를 기준으로 package 를 만든다.
go mod init <directory>

# doc 출력. 함수 위의 주석을 보여줌.
go doc examples/8_package/math_example.go Average 
go doc examples/8_package

# testing
# sub directory
go test ./...  
# 현재 directory
go test  
```

