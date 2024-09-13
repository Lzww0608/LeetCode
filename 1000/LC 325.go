func maxSubArrayLen(nums []int, k int) (ans int) {
    pre := 0
    m := map[int]int{}
    m[0] = -1
    for i, x := range nums {
        pre += x
        if v, ok := m[pre-k]; ok {
            ans = max(ans, i - v)
        }
        if _, ok := m[pre]; !ok {
            m[pre] = i
        } 
        
    }

    return 
}