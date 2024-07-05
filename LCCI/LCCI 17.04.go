func missingNumber(nums []int) int {
    n := len(nums)
    x := 0
    for i, t := range nums {
        x ^= i ^ t
    }

    return x ^ n
}