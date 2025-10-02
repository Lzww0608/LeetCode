func countArrays(original []int, bounds [][]int) int {
    type pair struct {
        l, r int 
    }

    n := len(original)
    cur := pair{bounds[0][0], bounds[0][1]}
    for i := 1; i < n; i++ {
        d := original[i] - original[i - 1]
        cur.l = max(cur.l + d, bounds[i][0])
        cur.r = min(cur.r + d, bounds[i][1])
        if cur.r < cur.l {
            return 0
        }
    }

    return cur.r - cur.l + 1
}