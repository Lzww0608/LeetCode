func minAbsoluteDifference(nums []int) int {
    ans := len(nums) + 1
    a, b := -1, -1
    for i, x := range nums {
        if x == 1 {
            if b != -1 {
                ans = min(ans, i - b)
            }
            a = i
        } else if x == 2 {
            if a != -1 {
                ans = min(ans, i - a)
            }
            b = i 
        }
    }
    if ans > len(nums) {
        return -1
    }
    return ans
}