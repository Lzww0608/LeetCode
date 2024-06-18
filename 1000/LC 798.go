func bestRotation(nums []int) int {
    n := len(nums)
    cnt := make([]int, n)
    sum := 0
    for i, x := range nums {
        if x <= i {
            sum++
            cnt[i-x]++
        }   
    }
    // 2, 3, 1, 4, 0
    // -2, -2, 1, -1, 4
    k, mx := 0, sum
    for i := range nums {
        sum -= cnt[i]
        d := i - nums[i]
        if d + n - 1 - i >= 0 {
            sum++
        }
        if d >= 0 {
            cnt[d]--
        } else if d + n >= 0 && d < 0 {
            cnt[d + n]++
        }
        
        if sum > mx {
            mx, k = sum, i + 1
        }
    }

    return k
}