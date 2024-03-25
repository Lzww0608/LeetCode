var MOD int = int(1e9 + 7)
func threeSumMulti(arr []int, target int) (ans int) {
    m := map[int]int{}
    for _, x := range arr {
        m[x]++
    }

    if target % 3 == 0 && m[target / 3] >= 3 {
        t := int64(m[target / 3])
        ans += int(t * (t - 1) * (t - 2) / 6) % MOD
    }

    for i := 0; i < target; i++ {
        j, k := i, target - 2 * i
        if k > j {
            ans = (ans + m[i] * (m[i] - 1) / 2 * m[k]) % MOD
        }
    }

    for i := 0; i < target; i++ {
        if m[i] == 0 {
            continue
        }
        for j := i + 1; j <= target; j++ {
            k := target - i - j
            if m[j] == 0 || m[k] == 0 || k < j {
                continue
            }
            if k == j {
                ans = (ans + m[i] * m[j] * (m[j] - 1) / 2) % MOD
            } else {
                ans = (ans + m[i] * m[j] * m[k]) % MOD
            }
        } 
        //fmt.Println(i, ans)
    }

    return 
}