package rpc

import (
	"context"
	"github.com/azd1997/go-frame/config"
	"github.com/azd1997/go-frame/process/controller"
	"github.com/azd1997/go-frame/process/rpc/server"
	"google.golang.org/grpc"
	"net"
)

type Server struct {

}

// 启动RPC服务
func StartRPCServer(conf config.RpcConfig) {
	listener, err := net.Listen("tcp", conf.Addr)
	if err != nil {

	}

	s := grpc.NewServer()
	// 将RPCServer注册到服务server
	server.RegisterServerServer(s, &Server{})
	err = s.Serve(listener)
	if err != nil {

	}
}

func (s *Server) GetServerTime(ctx context.Context, req *server.ServerTimeRequest) (*server.ServerTimeResponse, error) {
	code, msg, data := controller.GetServerTime()
	respData := &server.ServerTimeResponseData{
		ServerTime:           uint64(data.ServerTime),
	}
	resp := &server.ServerTimeResponse{
		Code:                 uint32(code),
		Msg:                  msg,
		Data:                 respData,
	}
	return resp, nil
}