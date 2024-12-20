const N int = 26
func minAnagramLength(s string) (ans int) {
    cnt := make([]int, N)
    for i := range s {
        x := int(s[i] & 31) - 1
        cnt[x]++
        ans = cnt[x]
    }
    for _, x := range cnt {
        if x > 0 {
            ans = gcd(ans, x)
        }
    }
    sz, k := len(s) / ans, 1
    for i := range cnt {
        cnt[i] /= ans 
    }

    cur := make([]int, N)
    for i := range s {
        x := int(s[i] & 31) - 1
        cur[x]++
        if i + 1 == sz * k {
            f := true
            for j := 0; j < N; j++ {
                if cur[j] != cnt[j] * (sz / (len(s) / ans)) * k {
                    f = false
                    break
                }
            }
            if f {
                k++
            } else {
                sz++
                for sz < len(s) && len(s) % sz != 0 {
                    sz++
                }
            }
        }
    }

    return min(len(s), sz)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y 
    }
    return x
}