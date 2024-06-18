func minAbsDifference(nums []int, goal int) int {
    n := len(nums) / 2
    ans := math.MaxInt32
    a := nums[:n]
    b := []int{}
    for i := 0; i < (1 << n); i++ {
        sum := 0
        for j, x := range a {
            if (i >> j) & 1 == 1 {
                sum += x
            }
        }
        ans = min(ans, abs(sum - goal))
        b = append(b, sum)
    }
    sort.Ints(b)

    a = nums[n:]
    for i := 0; i < (1 << (len(nums) - n)); i++ {
        sum := 0
        for j, x := range a {
            if (i >> j) & 1 == 1 {
                sum += x
            }
        } 

        ans = min(ans, abs(sum - goal))
        t := goal - sum
        j := sort.SearchInts(b, t)
        if j < len(b) {
            ans = min(ans, abs(t - b[j]))
        }

        if j > 0 {
            ans = min(ans, abs(b[j-1] - t))
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
