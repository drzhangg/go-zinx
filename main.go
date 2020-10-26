package main

import (
 "fmt"
 "go-zinx/znet"
)

func main() {
 fmt.Println("hello world")
 //1.创建一个server句柄，使用zinx的api
 s := znet.NewServer("[zinx V0.1]")
 //2.启动server
 s.Serve()
}
