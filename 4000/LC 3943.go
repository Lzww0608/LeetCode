type pair struct {
    l, r, add int 
    cnt map[int]int
}

func numberOfPairs(a []int, b []int, queries [][]int) []int {
    n, m := len(a), len(b)
    B := int(math.Sqrt(float64(n) * float64(m)))
    c := []pair{}
    for l := 0; l < m; l += B {
        r := min(l + B, m)
        cnt := make(map[int]int)
        for k := l; k < r; k++ {
            cnt[b[k]]++
        }
        c = append(c, pair{l, r, 0, cnt})
    }

    ans := make([]int, 0, len(queries))
    for _, q := range queries {
        if q[0] == 1 {
            x, y, val := q[1], q[2] + 1, q[3]
            for k, v := range c {
                if v.r <= x {
                    continue
                } else if v.l >= y {
                    break
                }

                if v.l >= x && v.r <= y {
                    c[k].add += val
                    continue 
                }
                for j := max(x, v.l); j < min(y, v.r); j++ {
                    c[k].cnt[b[j]]--
                    b[j] += val
                    c[k].cnt[b[j]]++
                }
            }
        } else {
            cur := 0
            for _, v := range c {
                target := q[1] - v.add 
                for _, x := range a {
                    cur += v.cnt[target - x]
                }
            }

            ans = append(ans, cur)
        }
    }

    return ans
}