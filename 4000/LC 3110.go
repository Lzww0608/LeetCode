func scoreOfString(s string) (ans int) {
    n := len(s)
    for i := 1; i < n; i++ {
        ans += abs(int(s[i] - 'a') - int(s[i-1] - 'a'))
    }

    return
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}