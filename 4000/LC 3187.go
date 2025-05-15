func countOfPeaks(nums []int, queries [][]int) []int {
    n := len(nums)
    f := make([]int, n - 1)

    update := func(i, v int) {
        for t := i; t < len(f); t += t & (-t) {
            f[t] += v
        }
    }    

    pre := func(i int) (res int) {
        for x := i; x > 0; x &= x - 1 {
            res += f[x]
        }
        return
    }

    query := func(l, r int) int {
        if r < l {
            return 0
        }
        return pre(r) - pre(l - 1)
    }

    solve := func(l, r, v int) {
        for i := l; i <= r; i++ {
            if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
                update(i, v)
            } 
        }
    }
    
    solve(1, n - 2, 1)
    ans := []int{}
    for _, v := range queries {
        if v[0] == 1 {
            ans = append(ans, query(v[1] + 1, v[2] - 1))
            continue
        }
        i := v[1]
        for j := max(i - 1, 1); j <= min(i + 1, n - 2); j++ {
            solve(j, j, -1)
        }
        nums[i] = v[2]
        for j := max(i - 1, 1); j <= min(i + 1, n - 2); j++ {
            solve(j, j, 1)
        }
    }

    return ans
}