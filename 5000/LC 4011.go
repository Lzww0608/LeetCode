func countRatioSubarrays(nums []int, a int, b int) (ans int) {
    // x * b <= a * y 
    n := len(nums)
    for l := range n {
        x, y := 0, 0
        for r := l; r < n; r++ {
            if nums[r] & 1 == 0 {
                x++
            } else {
                y++
            }
            if y > 0 && x * b <= a * y {
                ans++
            }
        }
    }

    return
}