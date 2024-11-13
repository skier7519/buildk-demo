FROM golang:1.18.0
ENV hello=/hello/hello
run go build -o hello/hello hello/hello.go
EXPOSE 8080
