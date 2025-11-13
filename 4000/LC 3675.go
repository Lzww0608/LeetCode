const N int = 26
func minOperations(s string) (ans int) {
    for i := range s {
        x := N - int(s[i] - 'a')
        ans = max(ans, x % N)
    }

    return 
}