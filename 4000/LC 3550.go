func smallestIndex(nums []int) int {
    for i := 0; i < min(len(nums), 28); i++ {
        x := nums[i]
        cur := 0 
        for x > 0 {
            cur += x % 10 
            x /= 10
        }
        if cur == i {
            return i
        }
    }

    return -1
}