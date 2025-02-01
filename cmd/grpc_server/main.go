package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	desc "github.com/UraharaKiska/LikeBigtech/pkg/auth_v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50052

type server struct {
	desc.UnimplementedAuthV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Auth id: %d", req.GetId())

	return &desc.GetResponse{
		User: &desc.User{
			Id: req.GetId(),
			Info: &desc.UserInfo{
				Name:            gofakeit.BeerName(),
				Email:           gofakeit.Email(),
				Password:        gofakeit.BeerAlcohol(),
				PasswordConfirm: gofakeit.BeerAlcohol(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create User %+v", req)

	return &desc.CreateResponse{
		Id: int64(gofakeit.Number(1000, 10000)),
	}, nil
}

func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update User %+v", req)

	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete User %d", req.GetId())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	log.Printf("grpc server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
