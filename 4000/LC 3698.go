func minBitwiseArray(nums []int) []int {
    
}func splitArray(nums []int) int64 {
    n := len(nums)
    sum := 0
    up := make([]bool, n)
    down := make([]bool, n) 
    for i, x := range nums {
        sum += x
        if i == 0 || up[i - 1] && x > nums[i - 1] {
            up[i] = true
        }
    }

    down[n - 1] = true 
    for i := n - 2; i >= 0; i-- {
        if nums[i] > nums[i + 1] {
            down[i] = true 
        } else {
            break
        }
    }

    ans := math.MaxInt / 2
    cur := 0
    for i := range n - 1 {
        cur += nums[i]
        if up[i] {
            if down[i + 1] {
                ans = min(ans, abs(sum - cur * 2))
            }
        } else {
            break
        }
    }

    if ans >= math.MaxInt / 2 {
        return -1
    }

    return int64(ans)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}