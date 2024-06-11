func findDiagonalOrder(nums [][]int) (ans []int) {
    m := make(map[int][]int)
    maxKey := 0
    for i := range nums {
        for j, v := range nums[i] {
            key := i + j
            maxKey = max(maxKey, key)
            m[key] = append(m[key], v)
        }
    }

    for k := 0; k <= maxKey; k++ {
        for i := len(m[k]) - 1; i >= 0; i-- {
            ans = append(ans, m[k][i])
        }
    }

    return 
}
