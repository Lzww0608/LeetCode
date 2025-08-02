
import "math"
func maximumValueSum(board [][]int) int64 {
    m := len(board)
    type pair struct {
        x, j int 
    }

    p := [3]pair{}
    suf := make([][3]pair, m)
    for i := range p {
        p[i].x = math.MinInt
    }

    update := func(v []int) {
        for j, x := range v {
            if x > p[0].x {
                if p[0].j != j {
                    if p[1].j != j {
                        p[2] = p[1]
                    }
                    p[1] = p[0]
                }
                p[0] = pair{x, j}
            } else if x > p[1].x && j != p[0].j {
                if p[1].j != j {
                    p[2] = p[1]
                }
                p[1] = pair{x, j}
            } else if x > p[2].x && j != p[0].j && j != p[1].j {
                p[2] = pair{x, j}
            }
        }
    }

    for i := m - 1; i > 1; i-- {
        update(board[i])
        suf[i] = p
    }

    ans := math.MinInt
    for i := range p {
        p[i].x = math.MinInt
    }

    for i, v := range board[:m - 2] {
        update(v)
        for j, x := range board[i + 1] {
            for _, l := range p {
                for _, r := range suf[i + 2] {
                    if l.j != r.j && l.j != j && r.j != j {
                        ans = max(ans, x + l.x + r.x)
                    }
                }
            }
        }
    }

    return int64(ans)
}