
const N = 1_000_00
var prime [N+1]int

func init() {
    for i := 2; i <= N; i++ {
        if prime[i] != 0 {
            prime[i] = prime[i-1]
            continue
        }
        prime[i] = prime[i-1] + 1
        for j := i * i; j <= N; j += i {
            prime[j] = -1;
        }
    }
}

func nonSpecialCount(l int, r int) int {
    x := int(math.Sqrt(float64(l-1)))
    y := int(math.Sqrt(float64(r)))

    return r - l + 1 - (prime[y] - prime[x])
}