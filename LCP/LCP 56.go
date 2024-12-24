
import (
	"container/list"
	"math"
)
type node struct {
    i, j, d int
}
func conveyorBelt(matrix []string, start []int, end []int) int {
    dirs := map[byte][]int{
        '>': {0, 1},
        '<': {0, -1},
        '^': {-1, 0},
        'v': {1, 0},
    }
    m, n := len(matrix), len(matrix[0])
    dis := make([][]int, m)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt32
        }
    }

    q := list.New()
    q.PushFront(node{start[0], start[1], 0})
    for q.Len() != 0 {
        cur := q.Front().Value.(node)
        q.Remove(q.Front())
        i, j, d := cur.i, cur.j, cur.d
        if i == end[0] && j == end[1] {
            return d
        }
        for k, v := range dirs {
            x, y := i + v[0], j + v[1]
            if x < 0 || x >= m || y < 0 || y >= n {
                continue
            }
            if k == matrix[i][j] {
                if dis[x][y] > d {
                    dis[x][y] = d 
                    q.PushFront(node{x, y, d})
                }
            } else {
                if dis[x][y] > d + 1 {
                    dis[x][y] = d + 1 
                    q.PushBack(node{x, y, d + 1})
                }
            }
        }
    }
    
    return -1
}