package graph

import pb "greeting/internal/grpc/proto"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	pb.GreetServiceClient
}
