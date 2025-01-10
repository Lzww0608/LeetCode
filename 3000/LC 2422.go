func minimumOperations(nums []int) (ans int) {
    n := len(nums)
    l, r := 0, n - 1
    for l < r {
        if nums[l] == nums[r] {
            l++
            r--
            continue
        } else if nums[l] > nums[r] {
            nums[r-1] += nums[r]
            r--
            ans++
        } else {
            nums[l+1] += nums[l]
            l++
            ans++
        }
    }
    
    return
}