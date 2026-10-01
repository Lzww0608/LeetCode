func longestSubarray(nums []int, k int) int {
    n := len(nums)

    for d := n; d > 0; d-- {
        sum := 0
        m := make(map[int]int)
        for i := range d {
            sum += nums[i]
            m[(nums[i] * 2 % k + k) % k]++
        }
        
        if t := (sum % k + k) % k; t == 0 || m[t] > 0 {
            return d
        }

        for i := d; i < n; i++ {
            sum += nums[i] - nums[i - d]
            m[(nums[i - d] * 2 % k + k) % k]--
            m[(nums[i] * 2 % k + k) % k]++
            if t := (sum % k + k) % k; t == 0 || m[t] > 0 {
                return d
            }
        }
    }

    return 0
}