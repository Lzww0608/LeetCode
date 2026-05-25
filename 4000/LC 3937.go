func minOperations(nums []int, k int) (ans int) {
    ans = math.MaxInt32
    for x := range k {
        for y := range k {
            if x == y {
                continue
            }

            cur := 0
            for i, t := range nums {
                t %= k 
                if i & 1 == 0 {
                    cur += min(abs(y - t), t + k - y, k - t + y)
                } else {
                    cur += min(abs(x - t), t + k - x, k - t + x)
                }
            }

            ans = min(ans, cur)
        } 
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}