type Graph struct {
    g [][]int
    n int
}


func Constructor(n int, edges [][]int) Graph {
    g := make([][]int, n)
    for i := range g {
        g[i] = make([]int, n)
    }
    for _, e := range edges {
        a, b, c := e[0], e[1], e[2]
        g[a][b] = c 
    }
    for k := 0; k < n; k++ {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                if g[i][k] != 0 && g[k][j] != 0 && i != j && j != k && (g[i][j] == 0 || g[i][k] + g[k][j] < g[i][j]) {
                    g[i][j] = g[i][k] + g[k][j]
                }
            }
        }
    }
    return Graph{g, n}
}


func (this *Graph) AddEdge(e []int)  {
    a, b, c := e[0], e[1], e[2]
    this.g[a][b] = c 
    for k := 0; k < this.n; k++ {
        for i := 0; i < this.n; i++ {
            for j := 0; j < this.n; j++ {
                if this.g[i][k] != 0 && this.g[k][j] != 0 && i != j && j != k && (this.g[i][j] == 0 || this.g[i][k] + this.g[k][j] < this.g[i][j]) {
                    this.g[i][j] = this.g[i][k] + this.g[k][j]
                }
            }
        }
    }
}


func (this *Graph) ShortestPath(node1 int, node2 int) int {
    if node1 == node2 {
        return 0
    }
   
    if this.g[node1][node2] == 0 {
        return -1
    }
    return this.g[node1][node2]
}


/**
 * Your Graph object will be instantiated and called as such:
 * obj := Constructor(n, edges);
 * obj.AddEdge(edge);
 * param_2 := obj.ShortestPath(node1,node2);
 */