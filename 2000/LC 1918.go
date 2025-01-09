func kthSmallestSubarraySum(nums []int, k int) int {
    n := len(nums)
    sum := 0
    for _, x := range nums {
        sum += x
    }

    check := func(target int) bool {
        cnt := 0
        cur := 0
        for l, r := 0, 0; r < n; r++ {
            cur += nums[r]
            for cur >= target {
                cur -= nums[l]
                l++
            }
            cnt += r - l + 1
        }

        return cnt + 1 <= k 
    }

    l, r := 1, sum + 1
    for l < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            l = mid + 1
        } else {
            r = mid
        }
    }

    return r - 1
}