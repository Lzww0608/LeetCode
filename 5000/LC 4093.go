func maximizeSum(nums []int, k int) int64 {
    ans := 0
    a := make([]int, len(nums)) 
    for i, x := range nums {
        ans += x
        a[i] = solve(x) - x
    }

    sort.Slice(a, func(i, j int) bool {
        return a[i] > a[j]
    })

    for _, x := range a[:k] {
        ans += x
    }

    return int64(ans)
}

func solve(x int) int {
    cnt := [10]int{}
    for x > 0 {
        cnt[x % 10]++
        x /= 10
    }

    res := 0
    for x := 9; x >= 0; x-- {
        for cnt[x] > 0 {
            cnt[x]--
            res = res * 10 + x 
        }
    }

    return res
}