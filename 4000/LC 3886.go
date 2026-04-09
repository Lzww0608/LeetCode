func sortableIntegers(nums []int) (ans int) {
    n := len(nums)
    nxt := make([]int, n)
    nxt[n - 1] = n
    last := n 

    for i := n - 2; i >= 0; i-- {
        if nums[i] > nums[i + 1] {
            last = i 
        }
        nxt[i] = last
    }

    solve := func(k int) {
        p := 0
        for r := k - 1; r < n; r += k {
            l := r - k + 1
            m := nxt[l]
            if m >= r {
                if nums[l] < p {
                    return 
                }
                p = nums[r]
            } else {
                if nxt[m + 1] < r || nums[m + 1] < p || nums[r] > nums[l] {
                    return
                }
                p = nums[m]
            }
        }

        ans += k 
        return
    }

    for k := 1; k * k <= n; k++ {
        if n % k == 0 {
            solve(k)
            if k * k < n {
                solve(n / k)
            }
        }
    }
    return
}