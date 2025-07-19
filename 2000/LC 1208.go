func equalSubstring(s string, t string, maxCost int) (ans int) {
    n := len(s)
    cur := 0
    for l, r := 0, 0; r < n; r++{
        cur += abs(int(s[r] - 'a') - int(t[r] - 'a'))
        for cur > maxCost {
            cur -= abs(int(s[l] - 'a') - int(t[l] - 'a'))
            l++
        }

        ans = max(ans, r - l + 1)
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
