package main

import (
	"context"
	"log"
	"net"

	"github.com/gethinyan/auth/models"
	pb "github.com/gethinyan/auth/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5230"
)

type user struct{}

func (u *user) GetUserInfo(ctx context.Context, in *pb.Identity) (*pb.UserInfo, error) {
	userInfo, err := models.GetUserByID(uint(in.ID))
	if err != nil {
		return &pb.UserInfo{}, err
	}
	return &pb.UserInfo{
		ID:       int64(userInfo.ID),
		Phone:    userInfo.Phone,
		Username: userInfo.Username,
		Address:  userInfo.Address,
		Birth:    userInfo.Birth.String(),
		Email:    userInfo.Email,
		Gender:   1,
	}, nil
}

func main() {
	log.Printf("begin to start rpc server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &user{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
