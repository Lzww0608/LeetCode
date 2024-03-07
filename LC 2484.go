var MOD int = int(1e9 + 7)
func countPalindromes(s string) (ans int) {
    n := len(s)
    
    a := [10]int{}
    for i, c := range s {
        t := 0
        for j := n - 1; j > i + 1; j-- {
            if s[i] == s[j] {
                ans = (ans + (j - i - 1) * t % MOD) % MOD
            }
            t = (t + a[int(s[j]-'0')]) % MOD
        }

        a[int(c-'0')]++
    }
    return
}