func subarraySum(nums []int, k int) int {
    m := map[int]int{}
    ans, pre := 0, 0
    m[0] = 1
    for _, x := range nums {
        pre += x
        if v, ok := m[pre-k]; ok {
            ans += v
        }
        m[pre]++
    }
    return ans
}
