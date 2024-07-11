func largestComponentSize(nums []int) (ans int) {
    n := slices.Max(nums)
    parent := make([]int, n + 1)
    for i := range parent {
        parent[i] = i
    }

    var find func(int) int 
    find = func(x int) int {
        if x != parent[x] {
            parent[x] = find(parent[x])
        }
        return parent[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            parent[rx] = ry
        }
    }

    for _, x := range nums {
        for i := 2; i * i <= x; i++ {
            if x % i == 0 {
                merge(x, i)
                merge(x, x / i)
            }
        }
    }

    cnt := map[int]int{}
    for _, x := range nums {
        root := find(x)
        cnt[root]++
        ans = max(ans, cnt[root])
    }

    return
}