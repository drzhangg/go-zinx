package znet

import "go-zinx/ziface"

type Server struct {
	Name      string //服务器的名称
	IPVersion string //服务器绑定的版本
	IP        string // 服务器监听的IP
	Port      int    //服务器监听的端口
}

func (s *Server) Start() {
	//1.获取一个TCP的addr

	//2.监听服务器的地址

	//3.阻塞的等待客户端链接，处理客户端链接业务（d读写）

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {

}

//初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
