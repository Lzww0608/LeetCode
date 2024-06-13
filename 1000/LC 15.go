func threeSum(nums []int) [][]int {
    ans, n := [][]int{}, len(nums)
    sort.Ints(nums)
    for i := 0; i < n - 2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        } else if nums[i] > 0 {
            break
        }
        l, r, target := i + 1, n - 1, -nums[i]
        for l < r {
            if l > i + 1 && nums[l] == nums[l-1] {
                l++
                continue
            }
            if nums[l] + nums[r] == target {
                ans = append(ans, []int{nums[i],nums[l],nums[r]})
                l++
            } else if nums[l] + nums[r] < target {
                l++
            } else {
                r--
            }
        }
    }
    return ans
}
