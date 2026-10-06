func maxSubarray(nums []int) (ans int) {
    n := len(nums)
    cnt := make(map[int]int)

    check := func(x int) bool {
        for k := range cnt {
            if k == x {
                if cnt[k] > 1 && cnt[k + x] > 0 {
                    return false
                }
            } else if cnt[k + x] > 0 {
                return false
            }
            
            if k > x {
                y := k - x 
                if y == x && cnt[y] > 1 || y != x && cnt[y] > 0 {
                    return false
                }
            } else if k < x {
                y := x - k
                if y == k && cnt[y] > 1 || y != k && cnt[y] > 0 {
                    return false
                }
            }
        }

        return true
    }

    for l, r := 0, 0; r < n; r++ {
        cnt[nums[r]]++
        for !check(nums[r]) {
            if cnt[nums[l]]--; cnt[nums[l]] == 0 {
                delete(cnt, nums[l])
            }
            l++
        }
        ans = max(ans, r - l + 1)
    }

    return ans
}