
var directions = [][2]int{{0, 1}, {1, 0}}

func isValid(x, y, n, m int, board [][]bool) bool {
	return x >= 0 && x < n && y >= 0 && y < m && !board[x][y]
}

func dfs(x int, adjList [][]int, visited []bool, match []int) bool {
	for _, neighbor := range adjList[x] {
		if visited[neighbor] {
			continue
		}
		visited[neighbor] = true
		if match[neighbor] == -1 || dfs(match[neighbor], adjList, visited, match) {
			match[neighbor] = x
			return true
		}
	}
	return false
}

func domino(n, m int, broken [][]int) int {
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, m)
	}

	for _, cell := range broken {
		board[cell[0]][cell[1]] = true
	}

	totalCells := n * m
	adjList := make([][]int, totalCells)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] {
				continue
			}
			current := i*m + j
			for _, d := range directions {
				nx, ny := i+d[0], j+d[1]
				if isValid(nx, ny, n, m, board) {
					neighbor := nx*m + ny
					adjList[current] = append(adjList[current], neighbor)
					adjList[neighbor] = append(adjList[neighbor], current)
				}
			}
		}
	}

	match := make([]int, totalCells)
	for i := range match {
		match[i] = -1
	}

	maxDominoCount := 0
	for i := 0; i < totalCells; i++ {
		visited := make([]bool, totalCells)
		if dfs(i, adjList, visited, match) {
			maxDominoCount++
		}
	}

	return maxDominoCount / 2
}

