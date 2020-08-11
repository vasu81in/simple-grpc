// Author:       Vasu Mahalingam
// Email:        vasu.uky@gmail.com
// Date:         2020-10-09

package graphlib_test

import (
	"fmt"
	myGraph "github.com/vasu81in/simple-grpc/internal/graphlib"
	"reflect"
	"strconv"
	"sync"
	"testing"
	//"strings"
	//"sync"
	"time"
)

var (
	edges = [][]int{
		{0, 1},
		{0, 3},
		{1, 2},
		{1, 4},
		{2, 5},
		{3, 4},
		{4, 5},
	}

	expectedPaths = [][]int{
		{0, 1, 4, 5},
		{0, 3, 4, 5},
		{0, 1, 2, 5},
	}

	expectedDistance int64 = 3
)

func TestCreateGraph(t *testing.T) {
	g := myGraph.CreateGraph("1", 6, edges)

	if g == nil {
		t.Errorf("Graph not created")
	}

	if g.GetID() == "" {
		t.Errorf("Graph ID not matched")
	}

	if g.GetClient() != "1" {
		t.Errorf("Client ID not matched")
	}

	if g.GetGraph() == nil {
		t.Errorf("Mutable object not created")
	}
}

func TestSPFPath(t *testing.T) {
	g := myGraph.CreateGraph("1", 6, edges)

	gotPath, _ := myGraph.ShortestPath(g.GetID(), 0, 5)
	equal := false
	for _, path := range expectedPaths {
		if reflect.DeepEqual(path, gotPath) {
			equal = true
			break
		}
	}
	if !equal {
		t.Errorf("Expected %+v, Got %+v", expectedPaths, gotPath)
	}
}

func TestSPFPathDistance(t *testing.T) {
	g := myGraph.CreateGraph("1", 6, edges)

	_, gotDistance := myGraph.ShortestPath(g.GetID(), 0, 5)
	if gotDistance != expectedDistance {
		t.Errorf("Expected %+v, Got %+v", expectedDistance, gotDistance)
	}
}

func TestDeleteGraph(t *testing.T) {
	g := myGraph.CreateGraph("1", 6, edges)
	myGraph.DeleteGraph(g.GetID(), "1")
	_, expected := myGraph.GetGraph(g.GetID())

	if expected != nil {
		t.Errorf("Graph not deleted")
	}
}

func TestDeleteGraphDB(t *testing.T) {
	expected := 0
	db := myGraph.GetGraphDB()
	count := db.GraphCount()
	fmt.Printf("Graphs Count: %d", count)
	myGraph.DeleteGraphDB()
	gotDb := myGraph.GetGraphDB()
	got := gotDb.GraphCount()

	if got != 0 {
		t.Errorf("Expected %d, Got %d", expected, got)
	}
}

// CreateGraphWorker ...
func CreateGraphWorker(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	g := myGraph.CreateGraph(strconv.Itoa(id), 6, edges)
	fmt.Printf("Worker %v: Started\n", id)
	time.Sleep(time.Second)
	myGraph.ShortestPath(g.GetID(), 0, 5)
	fmt.Printf("Worker %v: Finished\n", id)
}

func TestConcurrencyCreateGraph(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		fmt.Println("Main: Starting worker", i)
		wg.Add(1)
		go CreateGraphWorker(&wg, i)
	}
	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	expected := 10
	g := myGraph.GetGraphDB()
	got := g.GraphCount()
	if got != uint32(expected) {
		t.Errorf("Expected %d Got %d", expected, got)
	}
}
