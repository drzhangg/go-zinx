package znet

import (
	"fmt"
	"go-zinx/ziface"
	"net"
)

type Server struct {
	Name      string //服务器的名称
	IPVersion string //服务器绑定的版本
	IP        string // 服务器监听的IP
	Port      int    //服务器监听的端口
}

func (s *Server) Start() {

	fmt.Printf("[Start] Server Listenner at IP:%s, Port %d, is starting\n", s.IP, s.Port)
	go func() {
		//1.获取一个TCP的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tp addr error:", err)
			return
		}

		//2.监听服务器的地址
		listener,err := net.ListenTCP(s.IPVersion,addr)
		if err != nil{
			fmt.Println("listen ",s.IPVersion," err:",err)
			return
		}
		fmt.Println("start Zinx server success, ",s.Name," succ,listenning...")

		//3.阻塞的等待客户端链接，处理客户端链接业务（d读写）
		for {
			//如果有客户端链接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err:", err)
				continue
			}

			//已经与客户端建立连接，做一些业务，做一个最基本的最大512字节长度的回显
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err:", err)
						continue
					}

					//回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err:", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()

	//阻塞状态
	select {}
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
