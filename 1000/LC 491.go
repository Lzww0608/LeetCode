func findSubsequences(nums []int) (ans [][]int) {
    tmp := []int{}
    n := len(nums)

    var dfs func(int, int)
    dfs = func(idx, pre int) {
        if idx == n {
            if len(tmp) > 1 {
                ans = append(ans, append([]int(nil), tmp...))
            }
            return
        }  

        if nums[idx] >= pre {
            tmp = append(tmp, nums[idx])
            dfs(idx + 1, nums[idx])
            tmp = tmp[:len(tmp)-1]
        }

        if nums[idx] != pre {
            dfs(idx + 1, pre)
        }
    }
    dfs(0, math.MinInt32)

    return 
}