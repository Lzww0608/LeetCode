func longestCommonSubsequence(a [][]int) (ans []int) {
    n := len(a)
    mx, cnt := 0, 0 
    p := make([]int, n)

    for i := 0; p[i] < len(a[i]); i = (i + 1) % n {
        if a[i][p[i]] > mx {
            mx = a[i][p[i]]
            cnt = 1
        } else if a[i][p[i]] == mx {
            cnt++
        } else {
            for p[i] < len(a[i]) && a[i][p[i]] < mx {
                p[i]++
            }
            if p[i] >= len(a[i]) {
                return
            }
            if a[i][p[i]] == mx {
                cnt++
            } else {
                mx = a[i][p[i]] 
                cnt = 1
            }
        }

        if cnt == n {
            ans = append(ans, mx)
            for j := 0; j < n; j++ {
                p[j]++
            }
        }
    }

    return 
}