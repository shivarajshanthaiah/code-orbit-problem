package server

import (
	"fmt"
	"log"
	"net"

	"github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/handler"
	pb "github.com/shivaraj-shanthaiah/code_orbit_problem/pkg/proto"
	"google.golang.org/grpc"
)

func NewGrpcProblemServer(port string, handlr *handler.ProblemHandler) error {
	log.Println("connecting to gRPC server")
	addr := fmt.Sprintf(":%s", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Printf("error creating listener on %v", addr)
		return err
	}

	grpc := grpc.NewServer()

	pb.RegisterProblemServiceServer(grpc, handlr)
	log.Printf("listening on gRPC server %v", port)
	err = grpc.Serve(lis)
	if err != nil {
		log.Println("error connecting to gRPC server")
		return err
	}
	return nil
}
