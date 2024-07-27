package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	damagepb "pokecalc/pkg/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	scanner *bufio.Scanner
	client  damagepb.DamageCalcClient
)

func main() {
	fmt.Println("start gRPC Client.")

	// 1. 標準入力から文字列を受け取るスキャナを用意
	scanner = bufio.NewScanner(os.Stdin)

	// 2. gRPCサーバーとのコネクションを確立
	address := "localhost:8080"
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = damagepb.NewDamageCalcClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
		case "1":
			Attack()

		case "2":
			fmt.Println("bye.")
			goto M
		}
	}
M:
}

func Attack() {
	fmt.Println("Please enter pokemon name.")
	scanner.Scan()
	name := scanner.Text()

	req := &damagepb.DamageCalcRequest{
		Name: name,
	}
	res, err := client.Attack(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMaxDamage())
		fmt.Println(res.GetMinDamage())
	}
}
