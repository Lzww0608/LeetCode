func wiggleSort(nums []int)  {
    for i := 1; i < len(nums); i++ {
        if i & 1 == 1 {
            if nums[i] > nums[i-1] {
                nums[i], nums[i-1] = nums[i-1], nums[i]
            }
        } else if nums[i] < nums[i-1] {
            nums[i], nums[i-1] = nums[i-1], nums[i]
        }
    }
}