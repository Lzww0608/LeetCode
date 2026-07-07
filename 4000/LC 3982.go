func maxDigitRange(nums []int) (ans int) {
    d := -1
    for _, x := range nums {
        y := x
        mx, mn := x % 10, x % 10
        for x > 0 {
            t := x % 10
            x /= 10
            mx = max(mx, t)
            mn = min(mn, t)
        }

        if d < mx - mn {
            d = mx - mn 
            ans = y
        } else if d == mx - mn {
            ans += y
        }
    }

    return
}