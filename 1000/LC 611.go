func triangleNumber(nums []int) (ans int) {
    n := len(nums)
    sort.Ints(nums)
    for i := 0; i < n - 2; i++ {
        if nums[i] == 0 {
            continue
        }
        k := i + 2 
        for j := i + 1; j < n - 1; j++ {
            for k < n && nums[i] + nums[j] > nums[k] {
                k++
            }
            ans += k - j - 1
        }
    }

    return
}