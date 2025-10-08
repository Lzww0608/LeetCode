func maxAlternatingSum(nums []int, swaps [][]int) int64 {
    ans := 0
    n := len(nums)

    fa := make([]int, n)
    for i := range n {
        fa[i] = i
        ans -= nums[i]
    }

    var find func(int) int 
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
        }
    }

    for _, v := range swaps {
        merge(v[0], v[1])
    }

    cnt := make([]int, n)
    a := make([][]int, n)
    for i := range n {
        x := find(i)
        if i & 1 == 0 {
            cnt[x]++
        }
        a[x] = append(a[x], nums[i])
    }

    for i := range n {
        x := find(i)
        if x == i && cnt[i] > 0 {
            sort.Ints(a[i])
            m := len(a[i])
            for j := 0; j < cnt[i]; j++ {
                ans += 2 * a[i][m - j - 1]
            }
        }
    }

    return int64(ans)
}