func minAbbreviation(target string, dictionary []string) (ans string) {
    n := len(target)
    dict := []int{}
    for _, s := range dictionary {
        if len(s) != n {
            continue
        }
        mask := 0
        for i := 0; i < n; i++ {
            if s[i] == target[i] {
                mask |= 1 << i 
            }
        }
        dict = append(dict, mask)
    }
    if len(dict) == 0 {
        return strconv.Itoa(len(target))
    }

    solve := func(x int) (ans int, s []byte) {
        cnt := 0
        for i := 0; i < n; i++ {
            mask := 1 << i 
            if x & mask != 0 {
                if cnt != 0 {
                    s = append(s, []byte(string(strconv.Itoa(cnt)))...)
                    cnt = 0
                }
                ans++
                s = append(s, target[i])
            } else {
                if cnt == 0 {
                    ans++
                }
                cnt++
            }
        }

        if cnt != 0 {
            s = append(s, []byte(string(strconv.Itoa(cnt)))...)
            cnt = 0
        }

        return
    }

    mx := n
    ans = target
    for s := 1; s < (1 << n); s++ {
        f := true
        for _, mask := range dict {
            if s & mask == s {
                f = false
                break
            }
        }
        if f {
            x, t := solve(s)
            if x < mx {
                mx = x 
                ans = string(t)
            }
        }
    }
    
    return
}