func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    mn, mx := 0, 0
    for i := indexDifference; i < len(nums); i++ {
        if nums[i-indexDifference] > nums[mx] {
            mx = i - indexDifference
        } else if nums[i-indexDifference] < nums[mn] {
            mn = i - indexDifference
        }

        if nums[mx] - nums[i] >= valueDifference {
            return []int{mx, i}
        } else if nums[i] - nums[mn] >= valueDifference {
            return []int{mn, i}
        }
    }

    return []int{-1, -1}
}