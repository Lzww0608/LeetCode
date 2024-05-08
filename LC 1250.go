func isGoodArray(nums []int) bool {
    for i := 1; i < len(nums); i++ {
        nums[i] = gcd(nums[i-1], nums[i])
        if nums[i] == 1 {
            return true
        }
    }
    return nums[0] == 1
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}