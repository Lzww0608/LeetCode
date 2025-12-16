func minMoves(nums []int) int64 {
    ans, sum := 0, 0
    p := -1
    for i, x := range nums {
        if x < 0 {
            p = i 
        } else {
            sum += x
        }
    } 

    if p == -1 {
        return 0
    } else if sum < -nums[p] {
        return -1
    }

    n := len(nums)
    l, r := (p - 1 + n) % n, (p + 1) % n 
    for nums[p] < 0 {
        if -nums[p] > nums[l] {
            nums[p] += nums[l]
            ans += nums[l] * ((p - l + n) % n)
            l = (l - 1 + n) % n
        } else {
            ans += -nums[p] * ((p - l + n) % n)
            return int64(ans)
        }

        if -nums[p] > nums[r] {
            nums[p] += nums[r]
            ans += nums[r] * ((r - p + n) % n) 
            r = (r + 1) % n
        } else {
            ans += -nums[p] * ((r - p + n) % n)
            return int64(ans)
        }
    }

    return int64(ans)
}