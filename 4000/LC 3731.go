func findMissingElements(nums []int) (ans []int) {
    sort.Ints(nums)
    n := len(nums)
    for x, i := nums[0] + 1, 1; x <= nums[n - 1]; x++ {
        if x == nums[i] {
            i++
        } else {
            ans = append(ans, x)
        }
    }

    return 
}