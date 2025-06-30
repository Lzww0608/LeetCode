func possibleStringCount(word string) int {
    ans := 1
    n := len(word)
    cnt := 1
    for i := 1; i <= n; i++ {
        if i == n || word[i] != word[i-1] {
            ans += cnt - 1
            cnt = 0
        }
        cnt++
    }

    return ans
}