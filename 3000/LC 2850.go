func minimumMoves(grid [][]int) int {
    
    var from, to []pair
    for i, row := range grid {
        for j, x := range row {
            if x > 1 {
                for k := 1; k < x; k++{
                    from = append(from, pair{i, j})
                }
            } else if x == 0 {
                to = append(to, pair{i, j})
            }
        }
    }

    ans := math.MaxInt
    f := func() {
        sum := 0
        for i, v := range from {
            sum += abs(v.x - to[i].x) + abs(v.y - to[i].y)
        }
        ans = min(ans, sum)
    }

    permute(from, 0, f)

    return ans
}

type pair struct {
    x, y int
}

func permute(a []pair, i int, f func()) {
    if i == len(a) {
        f()
        return
    }
    
    permute(a, i + 1, f)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        permute(a, i + 1, f)
        a[i], a[j] = a[j], a[i]
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}