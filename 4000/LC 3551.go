func minSwaps(nums []int) int {
    n := len(nums)
    a := make([][]int, 0, n)
    for i, x := range nums {
        cur := 0
        for t := x; t > 0; t /= 10 {
            cur += t % 10
        }
        a = append(a, []int{cur, x, i})
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0] || a[i][0] == a[j][0] && a[i][1] < a[j][1]
    })

    cnt := n
    fa := make([]int, n)
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

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            fa[rx] = ry 
            cnt--
        }
        return
    }

    for i, v := range a {
        merge(i, v[2])
    }

    return n - cnt
}