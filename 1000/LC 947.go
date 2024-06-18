const N int = 10_001
func removeStones(stones [][]int) (ans int) {
    f := map[int]int{}

    var find func(int) int
    find = func(x int) int {
        if _, ok := f[x]; !ok {
            f[x] = x
        }
        if x != f[x] {
            f[x] = find(f[x])
        }
        return f[x]
    } 

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            f[rx] = ry
        }
    }

    for _, v := range stones {
        merge(v[0], v[1] + N)
    }

    ans = len(stones)
    for x, fa := range f {
        if x == fa {
            ans--
        }
    }

    return ans
}