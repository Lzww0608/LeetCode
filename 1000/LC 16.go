func threeSumClosest(nums []int, target int) int {
    ans, dif := math.MaxInt32, math.MaxInt32
    sort.Ints(nums)
    n := len(nums)
    for i, x := range nums {
        for j, k := i + 1, n - 1; j < k; {
            if dif > abs(target -(nums[j] + nums[k] + x)) {
                ans = nums[j] + nums[k] + x
                dif = abs(target -(nums[j] + nums[k] + x))
            }
            
            if nums[j] + nums[k] > target - x {
                k--
            } else if nums[j] + nums[k] < target - x {
                j++
            } else {
                return target
            }
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
