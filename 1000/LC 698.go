func canPartitionKSubsets(nums []int, k int) bool {
    n := len(nums)
    sum := 0
    for _, x := range nums {
        sum += x
    }

    if sum % k != 0 {
        return false
    }
    target := sum / k

    sort.Ints(nums)
    if nums[n-1] > target {
        return false
    }

    f := make([]bool, 1 << n)
    for i := range f {
        f[i] = true
    }

    var dfs func(int, int) bool
    dfs = func(s, t int) bool {
        if s == 0 {
            return true
        }

        if !f[s] {
            return f[s]
        } 

        f[s] = false
        for i, x := range nums {
            if x + t > target {
                break
            }

            if (s >> i) & 1 == 1 {
                if dfs(s ^ (1 << i), (t + x) % target) {
                    return true
                }
            }
        }
        
        return false
    }

    return dfs(1 << n - 1, 0)
}