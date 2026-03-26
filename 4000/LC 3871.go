func countCommas(n int64) int64 {
    ans := 0

    sum, base := 999_000, 1
    for i := 1000; i <= int(n); i *= 1000 {
        x := min(sum, int(n) - i + 1)
        ans += x * base 
        base++
        sum = sum * 1000
    }

    return int64(ans)
}