func missingMultiple(nums []int, k int) int {
    n := len(nums)
    exists := make([]bool, 101)
    for _, x := range nums {
        if x % k == 0 {
            exists[x / k - 1] = true
        }
    }

    for i, x := range exists {
        if !x {
            return (i + 1) * k
        }
    }

    return (n + 1) * k
}