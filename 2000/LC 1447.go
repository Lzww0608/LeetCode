
import "strconv"
func simplifiedFractions(n int) (ans []string) {
    for i := 2; i <= n; i++ {
        for j := 1; j < i; j++ {
            if gcd(i, j) {
                ans = append(ans, strconv.Itoa(j) + "/" + strconv.Itoa(i))
            }
        }
    }

    return ans
}

func gcd(a, b int) bool {
    for b != 0 {
        a, b = b, a % b
    }

    return a == 1
}