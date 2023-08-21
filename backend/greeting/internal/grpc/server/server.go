package server

import (
	"greeting/internal/db"
	pb "greeting/internal/grpc/proto"
	"greeting/internal/model"
	"log"
	"net"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var database string = "postgresql://root:secret@localhost:5432/greet?sslmode=disable"

func GrpcServer() {
	var ser Server
	db, err := db.NewConnection(database, model.User{})
	if err != nil {
		log.Fatalf("Error in database connection: %s", err)
	}
	ser.DB = &db
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	log.Printf("Listning to: %s", addr)
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &ser)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}
}
