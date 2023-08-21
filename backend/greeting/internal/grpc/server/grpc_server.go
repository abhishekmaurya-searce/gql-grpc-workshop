package server

import (
	"context"
	"fmt"
	"greeting/internal/db"
	pb "greeting/internal/grpc/proto"
	"greeting/internal/model"

	"github.com/google/uuid"
)

type Server struct {
	pb.GreetServiceServer
	*db.DB
}

func (s *Server) Greet(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	usr := model.User{
		Id:    uuid.NewString(),
		First: in.Firstname,
		Last:  in.Lastname,
	}
	res, err := s.DB.Insert(usr)
	if err != nil {
		return nil, fmt.Errorf("error in inserting: %s", err)
	}
	return &pb.Res{Greeting: res}, nil
}
