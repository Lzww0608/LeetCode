
import "math"
func closestDivisors(num int) []int {
    a, b := solve(num + 1), solve(num + 2)
    if a[1] - a[0] < b[1] - b[0] {
        return a
    }
    return b
}

func solve(n int) (ans []int) {
    for i := int(math.Sqrt(float64(n))); i > 0; i-- {
        if n % i == 0 {
            k := n / i;
            ans = []int{i, k}
            break
        }
    }

    return 
}