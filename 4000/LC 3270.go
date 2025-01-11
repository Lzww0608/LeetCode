const N int = 1_000
func generateKey(num1 int, num2 int, num3 int) (ans int) {
    for k := N; k > 0; k /= 10 {
        ans += min(num1 / k, num2 / k, num3 / k) * k
        num1 %= k 
        num2 %= k
        num3 %= k
    }

    return
}