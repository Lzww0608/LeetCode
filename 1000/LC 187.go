const N int = 10
func findRepeatedDnaSequences(s string) (ans []string) {
    m := map[byte]int {
        'A': 0x00,
        'C': 0x01,
        'G': 0x02,
        'T': 0x03,
    }

    mask := 0
    cnt := make(map[int]int)
    n := len(s)
    for i := 0; i < min(n, N - 1); i++ {
        mask = (mask << 2) | m[s[i]]
    }

    
    for i := N - 1; i < n; i++ {
        mask = ((mask << 2) | m[s[i]]) & ((1 << (N * 2)) - 1)
        if cnt[mask] == 1 {
            ans = append(ans, string(s[i-N+1:i+1]))
        }
        cnt[mask]++ 
    }

    return
}