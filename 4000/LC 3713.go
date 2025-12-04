const N int = 26
func longestBalanced(s string) int {
    n := len(s)

    for d := n; d > 0; d-- {
        cnt := [N]int{}

        check := func() bool {
            pre := -1
            for j := range N {
                if cnt[j] != 0 {
                    if pre == -1 {
                        pre = cnt[j]
                    } else if cnt[j] != pre {
                        return false
                    }
                }
            }

            return true
        }

        for i := range n {
            x := int(s[i] - 'a')
            cnt[x]++
            if i >= d - 1 {
                if check() {
                    return d
                }

                y := int(s[i - d + 1] - 'a')
                cnt[y]--
            }
        }
    }

    return 0
}