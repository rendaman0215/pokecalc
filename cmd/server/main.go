package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	damagepb "pokecalc/pkg/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	damagepb.UnimplementedDamageCalcServer
}

func (s *myServer) Attack(ctx context.Context, req *damagepb.DamageCalcRequest) (*damagepb.DamageCalcResponse, error) {
	return &damagepb.DamageCalcResponse{
		MaxDamage: 100,
		MinDamage: 50,
	}, nil
}

func (s *myServer) mustEmbedUnimplementedDamageCalcServer() {}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	// grpcサーバを作成
	s := grpc.NewServer()

	damagepb.RegisterDamageCalcServer(s, NewMyServer())

	// サーバリフレクションの設定（リクエスト時にシリアライズの情報を返す設定）
	reflection.Register(s)

	// 作成したgrpcサーバをリスナーに登録
	go func() {
		log.Printf("start grpc server on port %d", port)
		s.Serve(listener)
	}()

	// サーバを停止する
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
