func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    n := len(nums)
    f := make([]int, n)
    pos, mx := 0, 0
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if nums[i] % nums[j] == 0 && f[i] < f[j] + 1 {
                f[i] = f[j] + 1
                if f[i] > mx {
                    mx = f[i]
                    pos = i
                }
            }
        }
    }

    ans := make([]int, 0, f[pos] + 1)
    for i := pos; i >= 0 && mx >= 0; i-- {
        if nums[pos] % nums[i] == 0 && f[i] == mx {
            ans = append(ans, nums[i])
            pos = i 
            mx--
        }
    }

    return ans
}