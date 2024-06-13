func minSubarray(nums []int, p int) int {
    n := len(nums)
    var sum int64 = 0
    for _, x := range nums {
        sum += int64(x)
    }
    rem := sum % int64(p)
    if rem == 0 {
        return 0
    }
    var pre int64 = 0
    m := map[int64]int{}
    m[0] = -1
    ans := n
    for i, x := range nums {
        pre += int64(x)
        mod := pre % int64(p)
        target := (mod - rem + int64(p)) % int64(p)
        if idx, ok := m[target]; ok {
            ans = min(ans, i - idx)
        }
        m[mod] = i
    }

    if ans == n {
        return -1
    }
    return ans
}