func fourSum(nums []int, target int) (ans [][]int) {
    sort.Ints(nums)
    n := len(nums)
    for i := range nums {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        for j := i + 1; j < n; j++ {
            if j > i + 1 && nums[j] == nums[j-1] {
                continue
            }
            for p, q := j + 1, n - 1; p < q; {
                t := nums[i] + nums[j] + nums[p] + nums[q] 
                if t > target {
                    q--
                } else if t < target {
                    p++
                } else {
                    ans = append(ans, []int{nums[i], nums[j], nums[p], nums[q]})
                    t = nums[p]
                    for p < q && nums[p] == t {
                        p++
                    }
                    t = nums[q]
                    for p < q && nums[q] == t {
                        q--
                    }
                }
            }
        }
    }

    return
}
