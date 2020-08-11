package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vasu81in/simple-grpc/hellopb"
	"github.com/vasu81in/simple-grpc/internal/graphlib"
	"google.golang.org/grpc"
)

type server struct {
}

// Every request will carry a clientID that server will
// use to authorize.

// CreateGraph ... Creates a Graph and returns a GraphID.
func (*server) CreateGraph(ctx context.Context,
	request *hellopb.CreateGraphRequest) (*hellopb.CreateGraphResponse, error) {
	size := request.Size
	clientID := request.ClientId
	graphID := graphlib.NewCreateGraph(clientID, int(size))

	response := &hellopb.CreateGraphResponse{
		GraphId: graphID,
	}

	return response, nil
}

// Every request will carry a clientID that server will
// use to authorize.

// DeleteGraph ... Deletes the graph from server if found.
func (*server) DeleteGraph(ctx context.Context,
	request *hellopb.DeleteGraphRequest) (*hellopb.DeleteGraphResponse, error) {
	clientID := request.ClientId
	graphID := request.GraphId

	success := graphlib.DeleteGraph(graphID, clientID)
	response := &hellopb.DeleteGraphResponse{
		Success: success,
	}

	return response, nil
}

// Every request will carry a clientID that server will
// use to authorize.

// AddEdgeToGraph ... Adds edge to the graph if found.
func (*server) AddEdgeToGraph(ctx context.Context,
	request *hellopb.AddEdgeRequest) (*hellopb.AddEdgeResponse, error) {
	clientID := request.ClientId
	graphID := request.GraphId
	vertexA := request.VertexA
	vertexB := request.VertexB

	success := graphlib.AddEdgeToGraph(graphID, clientID, int(vertexA), int(vertexB))
	response := &hellopb.AddEdgeResponse{
		Success: success,
	}

	return response, nil
}

// Every request will carry a clientID that server will
// use to authorize.

// GetSPFFromGraph ... returns SPF distance if found.
func (*server) GetSPFFromGraph(ctx context.Context,
	request *hellopb.GetSPFRequest) (*hellopb.GetSPFResponse, error) {
	clientID := request.ClientId
	graphID := request.GraphId
	vertexA := request.VertexA
	vertexB := request.VertexB

	distance := graphlib.NewShortestPath(graphID, clientID, int(vertexA), int(vertexB))
	response := &hellopb.GetSPFResponse{
		Distance: distance,
	}

	return response, nil
}

func main() {
	address := "0.0.0.0:50051"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	hellopb.RegisterCreateGraphServiceServer(s, &server{})
	hellopb.RegisterDeleteGraphServiceServer(s, &server{})
	hellopb.RegisterAddEdgeGraphServiceServer(s, &server{})
	hellopb.RegisterGetSPFGraphServiceServer(s, &server{})

	s.Serve(lis)
}
