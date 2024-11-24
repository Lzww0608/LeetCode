const MOD int = 1_000_000_007
func numWays(s string) (ans int) {
    cnt := 0
    n := len(s)
    for i := range s {
        if s[i] == '1' {
            cnt++
        }
    }

    if cnt % 3 != 0 {
        return
    }

    cnt /= 3
    if cnt == 0 {
        return (n - 2) * (n - 1) / 2  %  MOD
    }
    a, b, k := 0, 0, 0
    for i := range s {
        if s[i] == '1' {
            k++
        } 

        if k == cnt {
            a++
        } else if k == cnt * 2 {
            b++
        } 
    }

    ans += (a * b) % MOD
    return
}