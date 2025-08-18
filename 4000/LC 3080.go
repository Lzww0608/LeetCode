func unmarkedSumArray(nums []int, queries [][]int) []int64 {
    n := len(nums)
    vis := make([]bool, n)
    id := make([]int, n)
    sum := 0
    for i := 0; i < n; i++ {
        id[i] = i
        sum += nums[i]
    }
    sort.Slice(id, func(i, j int) bool {
        if nums[id[i]] == nums[id[j]] {
            return id[i] < id[j]
        }
        return nums[id[i]] < nums[id[j]]
    })

    ans := make([]int64, len(queries))
    j := 0
    for i, v := range queries {
        if j >= n {
            break
        }
        if !vis[v[0]] {
            vis[v[0]] = true 
            sum -= nums[v[0]]
        }
        for j < n && v[1] > 0 {
            if !vis[id[j]] {
                vis[id[j]] = true 
                sum -= nums[id[j]]
                v[1]--
            }
            j++
        }

        ans[i] = int64(sum)
    }

    return ans
}