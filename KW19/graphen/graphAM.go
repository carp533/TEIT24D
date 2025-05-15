package main

import "fmt"

//Augabe: vergleiche die beiden Graphen Implementierungen
//erstelle einen Graph mit den Bsipiel Werten aus graph.go
//und vergleiche die beiden Implementierungen.

// Graph Datentyp mit einer Adjazenz-Matrix
type GraphAM struct {
	nodes   int         // Anzahl der Knoten im Graph (Iindex A-Matrix)
	matrix  [][]int     // Adjazenz-Matrix
	nodeMap map[int]int // Map zwischen Knoten des Graph und Index Adjazenz-Matrix
	idList  []int       // Liste der Knoten
}

func NewMatrixGraph() *GraphAM {
	return &GraphAM{
		nodes:   0,
		matrix:  make([][]int, 0),
		nodeMap: make(map[int]int),
		idList:  make([]int, 0),
	}
}

// Knoten einfügen
func (g *GraphAM) AddNode(id int) {
	if _, exists := g.nodeMap[id]; exists {
		return
	}

	g.nodeMap[id] = g.nodes
	g.idList = append(g.idList, id)

	// Matrix vergrößern
	for i := 0; i < g.nodes; i++ {
		g.matrix[i] = append(g.matrix[i], 0)
	}

	newRow := make([]int, g.nodes+1)
	g.matrix = append(g.matrix, newRow)

	g.nodes++
}

// Kante einfügen
func (g *GraphAM) AddEdge(u, v int) {
	// Sicher stellen, dass beide Knoten existieren
	g.AddNode(u)
	g.AddNode(v)

	idxU := g.nodeMap[u]
	idxV := g.nodeMap[v]

	// Adjazenz-Matrix
	g.matrix[idxU][idxV] = 1
	g.matrix[idxV][idxU] = 1
}

// BFS Adjazenz-Matrix
func (g *GraphAM) BFSAM(start int) []int {
	if _, exists := g.nodeMap[start]; !exists {
		return []int{}
	}

	var result []int
	visited := make([]bool, g.nodes)
	queue := []int{g.nodeMap[start]}
	visited[g.nodeMap[start]] = true

	for len(queue) > 0 {
		nodeIdx := queue[0]
		queue = queue[1:]
		result = append(result, g.idList[nodeIdx])

		for i := 0; i < g.nodes; i++ {
			if g.matrix[nodeIdx][i] == 1 && !visited[i] {
				queue = append(queue, i)
				visited[i] = true
			}
		}
	}

	return result
}

// DFS Adjazenz-Matrix
func (g *GraphAM) DFSAM(start int) []int {
	if _, exists := g.nodeMap[start]; !exists {
		return []int{}
	}

	var result []int
	visited := make([]bool, g.nodes)
	g.dfsAMHelper(g.nodeMap[start], visited, &result)
	return result
}

// helper function
func (g *GraphAM) dfsAMHelper(nodeIdx int, visited []bool, result *[]int) {
	visited[nodeIdx] = true
	*result = append(*result, g.idList[nodeIdx])

	for i := 0; i < g.nodes; i++ {
		if g.matrix[nodeIdx][i] == 1 && !visited[i] {
			g.dfsAMHelper(i, visited, result)
		}
	}
}

func (g *GraphAM) PrintMatrix() {
	fmt.Println("Adjacency Matrix:")
	fmt.Print("  ")
	for _, id := range g.idList {
		fmt.Printf("%d ", id)
	}
	fmt.Println()

	for i := 0; i < g.nodes; i++ {
		fmt.Printf("%d ", g.idList[i])
		for j := 0; j < g.nodes; j++ {
			fmt.Printf("%d ", g.matrix[i][j])
		}
		fmt.Println()
	}
}

// Beispiel für die Nutzung:

func main() {
	// Beispiel für die Nutzung der MatrixGraph-Implementierung
	matrixGraph := NewMatrixGraph()

	matrixGraph.AddEdge(1, 2)
	matrixGraph.AddEdge(1, 3)
	matrixGraph.AddEdge(2, 4)
	matrixGraph.AddEdge(3, 4)

	matrixGraph.PrintMatrix()

	bfsResult := matrixGraph.BFSAM(1)
	fmt.Print("Breitensuche: ")
	for _, node := range bfsResult {
		fmt.Printf("%d ", node)
	}
	fmt.Println()

	dfsResult := matrixGraph.DFSAM(1)
	fmt.Print("Tiefensuche: ")
	for _, node := range dfsResult {
		fmt.Printf("%d ", node)
	}
	fmt.Println()
}
