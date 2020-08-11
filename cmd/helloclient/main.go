package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/vasu81in/simple-grpc/hellopb"
	"google.golang.org/grpc"
)

func createHelp() {
	var help string = `
usage: helloclient create <size> <clientID>

Create graph.
       eg, helloclient create --size=10
Arguments:
	<size> ... size of the graph (Optional).
	<clientID> ... hostname by default (Optional).
`
	fmt.Printf("%s", help)
}

func deleteHelp() {
	var help string = `
usage: helloclient delete <graphID> <clientID>

Delete graph.
       eg, helloclient delete --graphID=abcd
Arguments:
	<graphID>  ... uuid of the graph (Required).
	<clientID> ... hostname by default (Required).
`
	fmt.Printf("%s", help)
}

func addEdgeHelp() {
	var help string = `
usage: helloclient add_edge <graphID> <vertexA> <vertexB>

Add a Edge to a graph.
       eg, helloclient add_edge --graphID=abcd --vertexA=1 --vertexB=5
Arguments:
	<clientID> ... hostname by default (Optional).
	<graphID>  ... uuid of the graphID.
	<vertexA>  ... vertexA (Mandatory.
	<vertexB>  ... vertexB (Mandatory.
`
	fmt.Printf("%s", help)
}

func spfHelp() {
	var help string = `
usage: helloclient spf <graphID> <vertexA> <vertexB>

Add a Edge to a graph.
       eg, helloclient spf --graphID=abcd --vertexA=1 --vertexB=5
Arguments:
	<clientID> ... hostname by default (Optional).
	<graphID>  ... uuid of the graphID (Required).
	<vertexA>  ... vertexA (Required).
	<vertexB>  ... vertexB (Required).
`
	fmt.Printf("%s", help)
}

func help() {
	var help string = `
available commands:
	helloclient create
	helloclient add_edge
	helloclient delete
	helloclient spf
`
	fmt.Printf("%s", help)
}

// Every request will carry a clientID that server will
// use to authorize. If the clientID doesn't match with the
// serverside ClientID, the request Fails.

// CreateRequest ... Request to create a graph in the server. Returns unique GraphID.
func CreateRequest(cc *grpc.ClientConn, size int32, clientID string) {
	client := hellopb.NewCreateGraphServiceClient(cc)
	request := &hellopb.CreateGraphRequest{
		Size: int32(size), ClientId: clientID}

	resp, _ := client.CreateGraph(context.Background(), request)
	log.Printf("Receive response => [GraphID: %v]\n", resp.GraphId)

}

// Every request will carry a clientID that server will
// use to authorize. If the clientID doesn't match with the
// serverside ClientID, the request Fails.
// If the graphID doesn't match, Success is false.

// DeleteRequest ... Request to delete a graph from the server for graphID.
func DeleteRequest(cc *grpc.ClientConn, graphID string, clientID string) {
	client := hellopb.NewDeleteGraphServiceClient(cc)
	request := &hellopb.DeleteGraphRequest{
		GraphId: graphID, ClientId: clientID}
	resp, _ := client.DeleteGraph(context.Background(), request)
	log.Printf("Receive response => [Deleted: %v]\n", resp.Success)
}

// Every request will carry a clientID that server will
// use to authorize. If the clientID doesn't match with the
// serverside ClientID, the request Fails.
// If the graphID doesn't match, Success is false.

// AddEdgeRequest ... Request to add a edge for graphID.
func AddEdgeRequest(cc *grpc.ClientConn, graphID string,
	clientID string, vertexA, vertexB int32) {
	client := hellopb.NewAddEdgeGraphServiceClient(cc)
	request := &hellopb.AddEdgeRequest{
		GraphId: graphID, ClientId: clientID,
		VertexA: vertexA, VertexB: vertexB}
	resp, _ := client.AddEdgeToGraph(context.Background(), request)
	log.Printf("Receive response => [Added: %v]\n", resp.Success)
}

// Every request will carry a clientID that server will
// use to authorize. If the clientID doesn't match with the
// serverside ClientID, the request Fails.
// If the graphID doesn't match, Distance is 0.

// GetSPFRequest ... Request to find the SPF for a matching graphID. Returns 0 on Failure.
func GetSPFRequest(cc *grpc.ClientConn, graphID string,
	clientID string, vertexA, vertexB int32) {
	client := hellopb.NewGetSPFGraphServiceClient(cc)
	request := &hellopb.GetSPFRequest{
		GraphId: graphID, ClientId: clientID,
		VertexA: vertexA, VertexB: vertexB}
	resp, _ := client.GetSPFFromGraph(context.Background(), request)
	log.Printf("Receive response => [Distance: %v]\n", resp.Distance)

}

func main() {
	hostname, _ := os.Hostname()
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createSizePtr := createCmd.Int("size", 10, "Default size of the graph. (Number of vertices)")
	createClientIDPtr := createCmd.String("clientID", hostname, "Default hostname.")

	addEdgeCmd := flag.NewFlagSet("add_edge", flag.ExitOnError)
	addEdgeGraphIDPtr := addEdgeCmd.String("graphID", "", "graphID is Mandatory")
	addEdgeVertexAPtr := addEdgeCmd.String("vertexA", "", "vertexA is Mandatory")
	addEdgeVertexBPtr := addEdgeCmd.String("vertexB", "", "vertexB is Mandatory")
	addEdgeClientIDPtr := addEdgeCmd.String("clientID", hostname, "Default hostname.")

	spfCmd := flag.NewFlagSet("spf", flag.ExitOnError)
	spfGraphIDPtr := spfCmd.String("graphID", "", "graphID is Mandatory")
	spfVertexAPtr := spfCmd.String("vertexA", "", "vertexA is Mandatory")
	spfVertexBPtr := spfCmd.String("vertexB", "", "vertexB is Mandatory")
	spfClientIDPtr := spfCmd.String("clientID", hostname, "Default hostname.")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	deleteGraphIDPtr := deleteCmd.String("graphID", "", "graphID is Mandatory")
	deleteClientIDPtr := deleteCmd.String("clientID", hostname, "Default hostname.")

	flag.Parse()
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
	case "add_edge":
		addEdgeCmd.Parse(os.Args[2:])
	case "spf":
		spfCmd.Parse(os.Args[2:])
	case "delete":
		deleteCmd.Parse(os.Args[2:])
	default:
		help()
		os.Exit(1)
	}

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("simple-grpc-server:50051", opts)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	if createCmd.Parsed() {
		log.Printf("Hi Client: %s, Creating Graph: Size:%d",
			*createClientIDPtr, *createSizePtr)
		size := *createSizePtr
		clientID := *createClientIDPtr
		CreateRequest(cc, int32(size), clientID)
	}

	if deleteCmd.Parsed() {
		clientID := *deleteClientIDPtr
		if *deleteGraphIDPtr == "" {
			deleteHelp()
			os.Exit(1)
		}
		graphID := *deleteGraphIDPtr
		log.Printf("Hi Client: %s, Deleting Graph: %s",
			clientID, graphID)
		DeleteRequest(cc, graphID, clientID)
	}

	if addEdgeCmd.Parsed() {
		clientID := *addEdgeClientIDPtr

		if *addEdgeGraphIDPtr == "" {
			addEdgeHelp()
			os.Exit(1)
		}
		graphID := *addEdgeGraphIDPtr

		if *addEdgeVertexAPtr == "" {
			addEdgeHelp()
			os.Exit(1)
		}

		vertexA, err := strconv.Atoi(*addEdgeVertexAPtr)
		if err != nil {
			log.Fatalf("bad vertexA")
		}

		if *addEdgeVertexBPtr == "" {
			addEdgeHelp()
			os.Exit(1)
		}

		vertexB, err := strconv.Atoi(*addEdgeVertexBPtr)
		if err != nil {
			log.Fatalf("bad vertexB")
		}
		log.Printf("Hi Client: %s, Graph %s, Adding Edge: [%d] -- [%d]",
			clientID, graphID, vertexA, vertexB)
		AddEdgeRequest(cc, graphID, clientID, int32(vertexA), int32(vertexB))
	}

	if spfCmd.Parsed() {
		clientID := *spfClientIDPtr

		if *spfGraphIDPtr == "" {
			spfHelp()
			os.Exit(1)
		}
		graphID := *spfGraphIDPtr

		if *spfVertexAPtr == "" {
			spfHelp()
			os.Exit(1)
		}

		vertexA, err := strconv.Atoi(*spfVertexAPtr)
		if err != nil {
			log.Fatalf("bad vertexA")
		}

		if *spfVertexBPtr == "" {
			spfHelp()
			os.Exit(1)
		}

		vertexB, err := strconv.Atoi(*spfVertexBPtr)
		if err != nil {
			log.Fatalf("bad vertexB")
		}
		log.Printf("Hi Client: %s, Using GraphID: %s, get SPF between Edges: [%d] -- [%d]",
			clientID, graphID, vertexA, vertexB)
		GetSPFRequest(cc, graphID, clientID, int32(vertexA), int32(vertexB))
	}
}
