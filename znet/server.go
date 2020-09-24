package znet

import "go-zinx/ziface"

type Server struct {
	Name      string //服务器的名称
	IPVersion string //服务器绑定的版本
	IP        string // 服务器监听的IP
	Port      int    //服务器监听的端口
}

func (s *Server) Start() {

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
