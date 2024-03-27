func longestArithSeqLength(nums []int) (ans int) {
    n := len(nums)
    m := make([]map[int]int, n)

    for i := range nums {
        m[i] = make(map[int]int)
        for j := 0; j < i; j++ {
            d := nums[j] - nums[i]
            if _, ok := m[j][d]; ok {
                m[i][d] = m[j][d] + 1 
            } else {
                m[i][d] = 2
            }
            ans = max(ans, m[i][d])
        }
    }

    return ans

}