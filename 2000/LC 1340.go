func maxJumps(arr []int, d int) (ans int) {
    n := len(arr)
    pos := make([]int, n)
    for i := range pos {
        pos[i] = i
    }
    
    sort.Slice(pos, func(i, j int) bool {
        if arr[pos[i]] == arr[pos[j]] {
            return pos[i] < pos[j]
        }
        return arr[pos[i]] < arr[pos[j]]
    })

    f := make([]int, n)
    for _, x := range pos {
        for j := x - 1; j >= 0 && j >= x - d; j-- {
            if arr[j] < arr[x] {
                f[x] = max(f[x], f[j] + 1)
            } else {
                break
            }
        }

        for j := x + 1; j < n && j <= x + d; j++ {
            if arr[j] < arr[x] {
                f[x] = max(f[x], f[j] + 1)
            } else {
                break
            }
        }
    }

    for _, x := range f {
        ans = max(ans, x + 1)
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
