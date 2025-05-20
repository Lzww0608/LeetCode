func minDifference(nums []int) (ans int) {
    n := len(nums)
    mn, mx := math.MaxInt32, math.MinInt32 

    for i, x := range nums {
        if x != -1 && (i > 0 && nums[i-1] == -1 || i < n - 1 && nums[i+1] == -1) {
            mn = min(mn, x)
            mx = max(mx, x)
        }
    }

    solve := func(l, r int, f bool) {
        d := (min(r - mn, mx - l) + 1) / 2 
        if f {
            d = min(d, (mx - mn + 2) / 3)
        }
        ans = max(ans, d)
    }

    pre := -1
    for i, x := range nums {
        if x == -1 {
            continue
        }

        if pre >= 0 {
            if i - pre == 1 {
                ans = max(ans, abs(x - nums[pre]))
            } else {
                solve(min(nums[pre], x), max(nums[pre], x), i - pre > 2)
            }
        } else if i > 0 {
            solve(x, x, false)
        }

        pre = i
    }
    if pre >= 0 && pre < n - 1 {
        solve(nums[pre], nums[pre], false)
    }
    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}