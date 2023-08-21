package graph

import (
	"fmt"
	pb "greeting/internal/grpc/proto"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Replace with your React app's URL
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Server() {
	var ser Resolver
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	ser.GreetServiceClient = pb.NewGreetServiceClient(conn)
	serve := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{Resolvers: &mutationResolver{Resolver: &ser}},
		),
	)
	serve.AddTransport(transport.POST{})

	fmt.Println("Listening at localhost:8080/auth")
	http.Handle("/greet", corsMiddleware(serve))
	http.ListenAndServe(":8080", nil)
}
