// 测试RPC的客户端
package main

import (
	"context"
	"github.com/azd1997/go-frame/process/rpc/server"
	"google.golang.org/grpc"
	"log"
	"time"
)

const Addr = "127.0.0.1:8081"

func main() {
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v\n", err)
	}
	defer conn.Close()

	client := server.NewServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, _ := client.GetServerTime(ctx, &server.ServerTimeRequest{})

	log.Printf("ServerTime: %v\n", resp.Data.ServerTime)
}
