func sortByReflection(nums []int) []int {
    sort.Slice(nums, func(i, j int) bool {
        a := int(bits.Reverse(uint(nums[i])) >> bits.LeadingZeros(uint(nums[i])))
        b := int(bits.Reverse(uint(nums[j])) >> bits.LeadingZeros(uint(nums[j])))
        return a < b || a == b && nums[i] < nums[j]
    })

    return nums
}