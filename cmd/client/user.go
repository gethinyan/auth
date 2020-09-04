package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/gethinyan/auth/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5230"
)

func main() {
	// 建立一个与服务端的连接.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	response, err := client.GetUserInfo(ctx, &pb.Identity{ID: 1})
	if nil != err {
		log.Fatalf("get user info failed, %v", err)
	}
	log.Printf("get user info success, %s", response)
	fmt.Println(response)

	_, err = client.DeleteUser(ctx, &pb.Identity{ID: 1})
	if nil != err {
		log.Fatalf("delete user failed, %v", err)
	}
	log.Printf("delete user success, %s", response)

	defer cancel()
}
