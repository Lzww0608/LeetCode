func countRotations(s string, k int) (ans int) {
    n := len(s)
    s = s + s 

    cnt := 0
    for i := range n - 1 {
        if s[i] == s[i + 1] {
            cnt++
        }
    }

    if cnt == k {
        ans++
    }

    for j := n - 1; j < n * 2 - 2; j++ {
        if s[j] == s[j + 1] {
            cnt++
        }
        i := j - (n - 1)
        if s[i] == s[i + 1] {
            cnt--
        }

        if cnt == k {
            ans++
        }
    }

    return
}