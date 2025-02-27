func findIndices(nums []int, indexDifference int, valueDifference int) []int {
    n := len(nums)
    mx := make([][2]int, n)
    mn := make([][2]int, n)
    mx[0], mn[0] = [2]int{0, nums[0]}, [2]int{0, nums[0]}
    for i := 1; i < n; i++ {
        if nums[i] > mx[i-1][1] {
            mx[i] = [2]int{i, nums[i]}
        } else {
            mx[i] = mx[i-1]
        }
        if nums[i] < mn[i-1][1] {
            mn[i] = [2]int{i, nums[i]}
        } else {
            mn[i] = mn[i-1]
        }
    }

    for l, r := 0, indexDifference; r < n; l, r = l + 1, r + 1 {
        a := nums[r] + valueDifference
        b := nums[r] - valueDifference
        if mx[l][1] >= a {
            return []int{mx[l][0], r}
        } else if mn[l][1] <= b {
            return []int{mn[l][0], r}
        }
    }

    return []int{-1, -1}
}