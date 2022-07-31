package main

import (
	"flag"
	gofoundpb "gofound-grpc/api/gen/v1"
	"gofound-grpc/global"
	"gofound-grpc/initialize"
	"gofound-grpc/interceptor"
	"gofound-grpc/internal/service"
	"log"
	"net"

	middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

func main() {
	// 配置地址，默认路径:./config/config.yaml
	var config string
	flag.StringVar(&config, "conf", "./config/config.yaml", "choose config file.")
	flag.Parse()

	// 初始化解析器
	initialize.InitViper(config)

	// 启动服务
	lis, err := net.Listen("tcp", global.CONFIG.GRPC.Addr)
	if err != nil {
		log.Fatal("cannot listen:", err)
	}

	in := make([]grpc.UnaryServerInterceptor, 0, 3)
	in = append(in, interceptor.GrpcRecover(), interceptor.Validator())
	if global.CONFIG.Auth.EnableAdmin {
		in = append(in, interceptor.VerifyAuth())
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(in...)),
	}

	s := grpc.NewServer(opts...)
	gofoundpb.RegisterGofoundServiceServer(s, service.NewGofoundService())
	log.Println("server started addr", global.CONFIG.GRPC.Addr)
	log.Fatal(s.Serve(lis))
}
