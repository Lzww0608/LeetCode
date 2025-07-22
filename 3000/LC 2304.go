
import (
	"math"
	"slices"
)
func minPathCost(grid [][]int, moveCost [][]int) int {
    n := len(grid)
    for i := n - 2; i >= 0; i-- {
        for j, x := range grid[i] {
            t := math.MaxInt32
            for k, y := range moveCost[x] {
                t = min(t, y + grid[i+1][k])
            }
            grid[i][j] += t 
        }
    }

    return slices.Min(grid[0])
}