func findDisappearedNumbers(nums []int, lower int, upper int) (ans [][]int) {
    nums = append(nums, lower - 1, upper + 1)
    sort.Ints(nums)
    l := sort.SearchInts(nums, lower)
    r := sort.SearchInts(nums, upper + 1)

    for i := l; i <= r; i++ {
        if nums[i] - nums[i - 1] > 1 {
            ans = append(ans, []int{nums[i - 1] + 1, nums[i] - 1})
        }
    }

    return
}