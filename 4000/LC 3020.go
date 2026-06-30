func maximumLength(nums []int) int {
    cnt := make(map[int]int)
    for _, x := range nums {
        cnt[x]++
    }

    ans := cnt[1]
    if ans & 1 == 0 {
        ans--
    }
    delete(cnt, 1)
    for _, x := range nums {
        y := int(math.Sqrt(float64(x)))
        if y * y == x && cnt[y] >= 2 {
            continue
        }

        res := 0
        for cnt[x] >= 2 {
            res += 2
            x *= x 
        }
        if cnt[x] >= 1 {
            res++
        } else {
            res--
        }

        ans = max(ans, res)
    }

    return ans
}