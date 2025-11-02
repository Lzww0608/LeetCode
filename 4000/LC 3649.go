func perfectPairs(nums []int) int64 {
    // a * b >= 0 
    // 0 < a < b, |a + b| <= |a|, a + b <= a, b <= 0
    // 0 < b < a, |a + b| <= |b|, a + b <= b, a <= 0
    // a < b < 0, -a - b <= -b,  a >= 0
    // b < a < 0, b >= 0

    // a * b < 0 
    // a < 0 < b, |a + b| <= min(-a, b), a + b <= -a,  
    // b < 0 < a, 

    n := len(nums)
    for i := range n {
        if nums[i] < 0 {
            nums[i] = -nums[i]
        }
    }
    sort.Ints(nums)

    ans, j := 0, 0
    for i := range n {
        for j < i && nums[i] > nums[j] * 2 {
            j++
        }

        ans += i - j
    }

    return int64(ans)
}