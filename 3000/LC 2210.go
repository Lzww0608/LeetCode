func countHillValley(nums []int) (ans int) {
    n := len(nums)
    for i := 1; i < n - 1; i++ {
        x := nums[i]
        l, r := i - 1, i + 1 
        for l >= 0 && nums[l] == x {
            l--
        }
        for r < n && nums[r] == x {
            r++
        }

        i = r - 1
        if l < 0 || r >= n {
            continue
        }

        if nums[l] < x && nums[r] < x || nums[l] > x && nums[r] > x {
            ans++
        }
    }

    return
}