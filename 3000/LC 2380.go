func secondsToRemoveOccurrences(s string) (ans int) {
    cnt := 0
    for _, c := range s {
        if c == '0' {
            cnt++
        } else if cnt > 0 {
            ans = max(ans + 1, cnt)
        }
    }

    return
}