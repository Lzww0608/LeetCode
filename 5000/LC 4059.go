func largestPower(nums []int) []int {
    ans := make([]int, 15)
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })

    d := bits.Len(uint(nums[0]))
    for i := d - 1; i >= 0; i-- {
        j := 0
        for j < len(nums) && (nums[j] >> i) & 1 == 1 {
            j++
        }

        ans[14 - i] = j
        for j < len(nums) {
            nums[j] &^= 1 << i 
            j++
        }

        sort.Slice(nums, func(p, q int) bool {
            return nums[p] > nums[q]
        })
    }

    return ans
}