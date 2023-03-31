package main

import (
	"context"
	"errors"
	"fmt"
	pb "ginDemo/gPRC/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
)

// 服务端编写
// 1、创建gRPC Server对象，理解为Server端的抽象对象
// 2、将server（其包含需要被调用的服务端接口）注册到gRPC Server的内部注册中心。
// 这样可以在接受到请求时，通过内部的服务发现，发现该服务端接口并转接进行逻辑处理
// 3、创建Listen，监听TCP端口
// 4、给RPC Server 开始lis。Accept，直到stop

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 获取元数据的信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("未传输token")
		return nil, errors.New("未传输token")
	}
	var appId string
	var appKey string
	if v, ok := md["appid"]; ok {
		appId = v[0]
	}
	if v, ok := md["appkey"]; ok {
		appKey = v[0]
	}
	if appId != "lwk" || appKey != "123456" {
		fmt.Println("token不正确")
		return nil, errors.New("token不正确")
	}

	fmt.Println("hello" + req.RequestName)
	return &pb.HelloResponse{ResponseMsg: "hello" + req.RequestName}, nil
}

func main() {
	//// TSL认证
	//// 两个参数分别是cretFile， keyFile
	//// 自签名证书文件和私钥文件
	creds, _ := credentials.NewServerTLSFromFile("D:\\Work Space\\Go_WorkSpace\\src\\ginDemo\\gPRC\\key\\test.pem",
		"D:\\Work Space\\Go_WorkSpace\\src\\ginDemo\\gPRC\\key\\test.key")
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	//grpcServer := grpc.NewServer()
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	// 服务端中去注册编写的服务
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Println("failed to serve: %v", err)
		return
	}
}
