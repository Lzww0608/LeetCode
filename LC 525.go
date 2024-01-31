func findMaxLength(nums []int) int {
    a, b, cnt := 0, 0, map[int]int{}
    cnt[0] = -1
    ans := 0
    for i, x := range nums {
        if x == 0 {
            a++
        } else {
            b++
        }
        if v, ok := cnt[a-b]; ok {
            ans = max(ans, i - v)
        } else {
            cnt[a-b] = i
        }
    }
    return ans
}