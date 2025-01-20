func smallestSubsequence(s string, k int, letter byte, repetition int) string {
    ans := make([]byte, 0, k)
    n := len(s)
    q := []byte{}
    cnt := strings.Count(s, string(letter))
    in := 0
    for i := 0; i < n; i++ {
        for len(q) > 0 && s[i] < q[len(q)-1] {
            if q[len(q)-1] == letter {
                if cnt + in == repetition {
                    break
                } 
                in--
            } 
            q = q[:len(q)-1]
        }

        if i < n {
            if s[i] == letter {
                cnt--
                in++
            }
            q = append(q, s[i])
        }

        if i >= n - k {
            if q[0] == letter {
                in--
                repetition--
                ans = append(ans, q[0])
            } else if len(ans) + repetition < k {
                ans = append(ans, q[0])
            }
            q = q[1:]
        }
    }

    for len(ans) < k {
        if q[0] == letter {
            in--
            repetition--
            ans = append(ans, q[0])
        } else if len(ans) + repetition < k {
            ans = append(ans, q[0])
        }
        q = q[1:]
    }

    return string(ans)
}