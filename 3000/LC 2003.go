func smallestMissingValueSubtree(parents []int, nums []int) []int {
    n := len(parents)
    ans := make([]int, n)
    for i := range ans {
        ans[i] = 1
    }
    mn := 0
    for i, x := range nums {
        if x < nums[mn] {
            mn = i
        }
    }
    if nums[mn] > 1 {
        return ans
    }

    g := make([][]int, n)
    for i := 1; i < n; i++ {
        g[parents[i]] = append(g[parents[i]], i)
    }

    cur := mn 
    mex := 1
    st := []int{}
    vis := make(map[int]bool)
    for cur != -1 {
        vis[nums[cur]] = true
        for _, x := range g[cur] {
            if !vis[nums[x]] {
                vis[nums[x]] = true
                st = append(st, x)
            }
        }

        for len(st) > 0 {
            x := st[len(st)-1]
            st = st[:len(st)-1]
            for _, y := range g[x] {
                if !vis[nums[y]] {
                    vis[nums[y]] = true
                    st = append(st, y)
                }
            }
        }

        for vis[mex] {
            mex++
        }

        ans[cur] = mex 
        cur = parents[cur]
    }

    return ans
}