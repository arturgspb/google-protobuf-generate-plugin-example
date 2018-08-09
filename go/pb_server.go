package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "test-grpc/example"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

const (
	port = ":9090"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	//value := ctx.Value("UserID")
	//if value != nil {
	//	print(value)
	//} else {
	//	print("NOT DECODED")
	//}
	//userID := strconv.Itoa(ctx.Value("UserID").(int))
	userID := "UID-xxxx"
	return &pb.StringMessage{Value: "Hello! " + in.Value + " | " + userID}, nil
}

func parseToken(token string) (string, error) {
	return token, nil
}

func userClaimFromToken(string) string {
	return "foobar"
}

func main() {
	exampleAuthFunc := func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		tokenInfo, err := parseToken(token)
		if err != nil {
			return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}
		grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))
		print(tokenInfo)
		newCtx := context.WithValue(ctx, "UserID", tokenInfo)
		return newCtx, nil
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(

		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(exampleAuthFunc)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(exampleAuthFunc)),
		//grpc.StreamInterceptor(
		//	grpc_middleware.ChainStreamServer(
		//		grpc_auth.StreamServerInterceptor(exampleAuthFunc)
		//) )
	)

	pb.RegisterYourServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
