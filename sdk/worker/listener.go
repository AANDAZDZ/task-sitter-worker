package worker

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"task-sitter-worker/sdk/config"
	"task-sitter-worker/sdk/rpc/proto"
	"task-sitter-worker/sdk/rpc/server"
)

// Init 初始化并启动 rpc 服务
func Init() {
	err := config.InitConfig("")
	if err != nil {
	}
	err = Listener()
	if err != nil {
	}
}

// Listener 监听 rpc 请求
func Listener() (err error) {
	quit := make(chan error)
	s := grpc.NewServer()
	var listen net.Listener
	go func() {
		addr := fmt.Sprintf(":%d", config.Cfg.GRpcServerConf.Port)
		listen, err = net.Listen(config.Cfg.GRpcServerConf.Protocol, addr)
		if err != nil {
			quit <- err
			return
		}
		proto.RegisterScheduleServer(s, &rpc.ScheduleServer{})
		err = s.Serve(listen)
		if err != nil {
			quit <- err
		}
	}()
	select {
	case <-quit:
		s.Stop()
		err = listen.Close()
		return err
	}
}
