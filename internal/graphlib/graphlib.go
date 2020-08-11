package graphlib

import (
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"
	graph "github.com/yourbasic/graph"

	//"runtime"
	//"strconv"
	//"strings"
	"sync"
	//"time"
)

var gGraphDB *GraphDB

// Client is a struct used to identify a host
type Client struct {
	Name string
}

func (c *Client) String() string {
	return c.Name
}

// InitGraphDB ... init sets up the Global Graph DB
func initGraphDB() *GraphDB {
	log.Info("Initializing GraphDB()")
	return &GraphDB{
		graphDB: make(map[string]*Graph),
		count:   0,
	}
}

// GetGraphDB ... getter for Graph DB
func GetGraphDB() *GraphDB {
	gGraphDB.mux.Lock()
	defer gGraphDB.mux.Unlock()
	log.Info(fmt.Sprintf("%s :%+v", WhereAmI(), gGraphDB))
	return gGraphDB
}

// DeleteGraphDB ... Reinit Graph DB
func DeleteGraphDB() {
	gGraphDB = initGraphDB()
}

// Sets up the global graph DB
func init() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	gGraphDB = initGraphDB()
}

// GraphDB is a struct that keeps a
// a map of Graphs and the size of the
// DB.
type GraphDB struct {
	mux     sync.Mutex        // Safe mutex
	graphDB map[string]*Graph // graphID is the key
	count   uint32            // Number of  graphs
}

// get ... Getter for the graph DB based on graphID
func (db *GraphDB) get(graphID string) (bool, *Graph) {
	db.mux.Lock()
	defer db.mux.Unlock()
	if g, ok := db.graphDB[graphID]; ok {
		log.Info(fmt.Sprintf("%s :%+v", WhereAmI(), g))
		return ok, g
	}
	log.Info(fmt.Sprintf("%s :%s ", WhereAmI(), graphID))
	return false, nil
}

// add ... Inserts the graph into graph DB
func (db *GraphDB) add(g *Graph) {
	db.mux.Lock()
	defer db.mux.Unlock()
	log.Info(fmt.Sprintf("%s :%+v", WhereAmI(), g))
	db.graphDB[g.graphID] = g
	db.count++
	log.Info(fmt.Sprintf("DB: %+v", db))
}

// Delete ... Deletes the graph from graph DB
func (db *GraphDB) delete(graphID string) {
	db.mux.Lock()
	defer db.mux.Unlock()
	log.Info(fmt.Sprintf("%s :%+v", WhereAmI(), db.graphDB[graphID]))
	delete(db.graphDB, graphID)
	db.count--
}

// GraphCount ... Counts the number of graphs tracked by graph DB
func (db *GraphDB) GraphCount() uint32 {
	db.mux.Lock()
	defer db.mux.Unlock()
	return db.count
}

// GetNextID ... returns the ID to be used for graph
func (db *GraphDB) getGraphID() string {
	db.mux.Lock()
	defer db.mux.Unlock()
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// Graph is a struct to provide
// an abstraction over graph
// library
type Graph struct {
	client  *Client
	graph   *graph.Mutable
	graphID string
}

// GetID ... Returns graph ID
func (g *Graph) GetID() string {
	return g.graphID
}

// GetClient ... Returns client ID
func (g *Graph) GetClient() string {
	return g.client.String()
}

// GetGraph .. Returns the Mutable graph handle
func (g *Graph) GetGraph() *graph.Mutable {
	return g.graph
}

// NewCreateGraph ... returns a new graph object
func NewCreateGraph(clientID string, size int) string {
	g := &Graph{}
	g.graphID = gGraphDB.getGraphID()
	g.client = &Client{Name: clientID}
	g.graph = graph.New(size)
	gGraphDB.add(g)
	return g.GetID()
}

// AddEdgeToGraph ... Add edge between vertextA and vertextB
func AddEdgeToGraph(graphID string, clientID string,
	vertexA int, vertexB int) bool {
	found, current := gGraphDB.get(graphID)
	if !found {
		log.Error(fmt.Sprintf("%s :%s", WhereAmI(), graphID))
		return false
	}
	log.Info(fmt.Sprintf("%s :%+v", WhereAmI(), current))
	same := current.client.String() == clientID
	if !same {
		log.Error(fmt.Sprintf("Client: %s not the owner of +%v",
			clientID, current.graphID))
		return false
	}
	current.graph.AddBothCost(vertexA, vertexB, 1)
	return true
}

// CreateGraph ... Callers call this method to get a new graph object
func CreateGraph(clientID string, size int, edges [][]int) *Graph {
	g := &Graph{}
	g.graphID = gGraphDB.getGraphID()
	g.client = &Client{Name: clientID}
	g.graph = graph.New(size)
	for _, edge := range edges {
		g.graph.AddBothCost(edge[0], edge[1], 1)
	}
	gGraphDB.add(g)
	return g
}

// DeleteGraph ... Deletes a graph with ID from the graph DB
func DeleteGraph(graphID string, clientID string) bool {
	found, current := gGraphDB.get(graphID)
	if !found {
		log.Error(fmt.Sprintf("%s :%s", WhereAmI(), graphID))
		return false
	}
	log.Info(fmt.Sprintf("%s :%+v", WhereAmI(), current))
	same := current.client.String() == clientID
	if !same {
		log.Error(fmt.Sprintf("Client: %s not the owner of +%v",
			clientID, current.graphID))
		return false
	}
	gGraphDB.delete(graphID)
	return true
}

// GetGraph ... Returns the graph matching the graphID
func GetGraph(graphID string) (bool, *Graph) {
	return gGraphDB.get(graphID)
}

// NewShortestPath ... Returns a shortest path between vertices
func NewShortestPath(graphID string, clientID string, v1 int, v2 int) int32 {
	found, current := gGraphDB.get(graphID)
	if !found {
		log.Error(fmt.Sprintf("%s :%s", WhereAmI(), graphID))
		return 0
	}
	log.Info(fmt.Sprintf("%s :%+v", WhereAmI(), current))
	same := current.client.String() == clientID
	if !same {
		log.Error(fmt.Sprintf("Client: %s not the owner of +%v",
			clientID, current.graphID))
		return 0
	}
	_, distance := graph.ShortestPath(current.graph, v1, v2)
	return int32(distance)
}

// ShortestPath ... Returns a shortest path between vertices
func ShortestPath(graphID string, v1 int, v2 int) ([]int, int64) {
	found, g := gGraphDB.get(graphID)
	if !found {
		log.Error(fmt.Sprintf("%s :%s", WhereAmI(), graphID))
		return nil, 0
	}
	return graph.ShortestPath(g.graph, v1, v2)
}
