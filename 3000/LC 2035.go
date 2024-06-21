func minimumDifference(nums []int) int {
    n := len(nums) >> 1
    a := nums[:n]

    res := make([][]int, n + 1)
    for i := 0; i < (1 << n); i++ {
        sum, cnt := 0, 0
        for j, x := range a {
            if (i >> j) & 1 == 1 {
                sum += x 
                cnt++
            } else {
                sum -= x
            }
        }
        res[cnt] = append(res[cnt], sum)
    }

    for i := range res {
        sort.Ints(res[i])
    }

    a = nums[n:]
    ans := math.MaxInt32

    for i := 0; i < (1 << n); i++ {
        sum, cnt := 0, 0
        for j, x := range a {
            if (i >> j) & 1 == 1 {
                sum += x 
                cnt++
            } else {
                sum -= x
            }
        }

        b := res[cnt]
        j := sort.SearchInts(b, sum)

        if j < len(b) {
            ans = min(ans, b[j] - sum)
        }

        if j > 0 {
            ans = min(ans, sum - b[j - 1])
        }

    }
    return ans
}
