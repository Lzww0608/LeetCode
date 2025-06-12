const N int = 50_001
func amountPainted(paint [][]int) []int {
    n := len(paint)
    ans := make([]int, n)

    fa := make([]int, N)
    for i := range fa {
        fa[i] = i
    }

    var find func(x int) int 
    find = func(x int) int {
        if fa[x] != x {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    for i, v := range paint {
        l, r := v[0], v[1]
        for j := r; j > l; j-- {
            t := find(j)
            if t == j {
                ans[i]++
                fa[j] = min(fa[j], l)
            } else {
                j = t + 1
            }
        }
    }

    return ans
}