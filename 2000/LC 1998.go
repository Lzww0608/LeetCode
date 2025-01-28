const N int = 100_005
func gcdSort(nums []int) bool {
    fa := make([]int, N)
    for i := range fa {
        fa[i] = i
    }

    var find func(int) int
    find = func(x int) int {
        if x != fa[x] {
            fa[x] = find(fa[x])
        }
        return fa[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[rx] = ry
        }
    }

    for _, x := range nums {
        t := x
        for i := 2; i * i <= x; i++ {
            if x % i == 0 {
                merge(i, t)
                for x % i == 0 {
                    x /= i 
                }
            }
        }
        if x > 1 {
            merge(x, t)
        }
    }

    tmp := make([]int, len(nums))
    copy(tmp, nums)
    sort.Ints(tmp)
    for i := range tmp {
        if find(tmp[i]) != find(nums[i]) {
            return false
        }
    }
    return true
}