func generatePalindromes(s string) (ans []string) {
    cnt := make([]int, 26)
    for i := range s {
        cnt[int(s[i] & 31)-1]++
    }
    odd, pos := 0, -1
    for i, x := range cnt {
        if x & 1 == 1 {
            odd++
            pos = i
        }
    }
    if odd > 1 {
        return 
    }
    n := len(s)

    cur := make([]byte, n)

    var dfs func(int)
    dfs = func(idx int) {
        if idx * 2 == n {
            ans = append(ans, string(cur))
            return
        } else if idx * 2 + 1 == n {
            cur[idx] = byte('a' + pos)
            ans = append(ans, string(cur))
            return
        }

        for i := range cnt {
            if cnt[i] >= 2 {
                cnt[i] -= 2
                c := byte('a' + i)
                cur[idx], cur[n-idx-1] = c, c
                dfs(idx + 1)
                cnt[i] += 2 
            }
        }
    }
    dfs(0)

    return
}