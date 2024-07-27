package main

import (
	"context"
	"net"
	"os"
	"testing"

	damagepb "pokecalc/pkg/grpc/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	damagepb.RegisterDamageCalcServer(s, NewMyServer())
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAttack(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()
	client := damagepb.NewDamageCalcClient(conn)

	req := &damagepb.DamageCalcRequest{
		Name: "test",
	}
	resp, err := client.Attack(ctx, req)
	if err != nil {
		t.Fatalf("Attack failed: %v", err)
	}

	if resp.MaxDamage != 100 {
		t.Errorf("Expected MaxDamage=100, got %v", resp.MaxDamage)
	}
	if resp.MinDamage != 50 {
		t.Errorf("Expected MinDamage=50, got %v", resp.MinDamage)
	}
}

func TestMain(m *testing.M) {
	go main() // Run the main function in a separate goroutine to test the server

	// Run the tests
	code := m.Run()

	// Clean up any resources if necessary
	os.Exit(code)
}
