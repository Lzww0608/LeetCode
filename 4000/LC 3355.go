func isZeroArray(nums []int, queries [][]int) bool {
    n := len(nums)
    dif := make([]int, n + 1)
    for _, v := range queries {
        dif[v[0]]++
        dif[v[1] + 1]--
    }

    sum := 0
    for i := 0; i < n; i++ {
        sum += dif[i]
        if sum < nums[i] {
            return false
        }
    }

    return true
}