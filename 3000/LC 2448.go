func minCost(nums []int, cost []int) int64 {
    cnt, sum := 0, 0
    n := len(nums)
    p := make([]int, n)
    for i := range nums {
        p[i] = i
        sum += nums[i] * cost[i]
        cnt += cost[i]
    }

    cur, mid := 0, 0
    sort.Slice(p, func(i, j int) bool {
        return nums[p[i]] < nums[p[j]]
    })
    for _, i := range p {
        cur += cost[i]
        if cur > cnt / 2 {
            mid = i
            break
        }
    }

    ans := 0
    for i := range nums {
        ans += abs(nums[i] - nums[mid]) * cost[i];
    }

    return int64(ans)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}