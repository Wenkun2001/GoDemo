package main

import (
	"context"
	"fmt"
	pb "ginDemo/gPRC/hello-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// Token认证
// gRPC提供PerRPCCredentials接口，接口有两个方法，接口位于credentials包下，这个接口需要客户端来实现

//type PerRPCCredentials interface {
//	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
//	RequireTransportSecurity() bool
//}

// 第一个方法作用是获取元数据信息，也就是客户端提供的key-value对，context用于控制超时和取消，uri是请求入口处的uri
// 第二个方法作用是否需要基于TLS认证进行安全传输，如果返回值是true，则必须加上TLS验证，返回值是false则不用

type ClientTokenAuth struct {
}

func (c ClientTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appId":  "lwk",
		"appKey": "123456",
	}, nil
}

func (c ClientTokenAuth) RequireTransportSecurity() bool {
	return true
}

// 客户端编写
// 1、创建与给定目标（服务端）的连接交互
// 2、创建server的客户端对象
// 3、发送RPC请求，等待同步响应，得到回调后返回响应结果
// 4、输出响应结果
func main() {
	creds, _ := credentials.NewClientTLSFromFile("D:\\Work Space\\Go_WorkSpace\\src\\ginDemo\\gPRC\\key\\test.pem",
		"*.lwk.com")

	// 连接到server端，此处禁用安全传输，没有加密和验证
	var opts []grpc.DialOption
	//opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithPerRPCCredentials(new(ClientTokenAuth)))

	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial("127.0.0.1:9090", opts...)
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)
	// 执行rpc调用（这个方法在服务器端来实现并返回结果）
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{RequestName: "lwk"})

	fmt.Println(resp.GetResponseMsg())
}
