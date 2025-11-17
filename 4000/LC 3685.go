func subsequenceSumAfterCapping(nums []int, k int) []bool {
    n := len(nums)
    ans := make([]bool, n)
    sort.Ints(nums)
    f := make([]bool, k + 1)
    f[0] = true 

    p := 0
    for i := range n {
        x := i + 1
        for p < n && nums[p] <= x {
            for j := k; j >= nums[p]; j-- {
                f[j] = f[j] || f[j - nums[p]]
            }
            p++
        }

        cnt := n - p
        for j := k; j >= 0; j-- {
            if f[j] && (k - j) % x == 0 && (k - j) / x <= cnt {
                ans[i] = true
                break
            } 
        } 
    }

    return ans
}