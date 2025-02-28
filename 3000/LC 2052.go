
import "math"
func minimumCost(sentence string, k int) int {
    s := strings.Split(sentence, " ")
    n := len(s)
    a := make([]int, n)
    for i, t := range s {
        a[i] = len(t) + 1
    }

    f := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        f[i] = math.MaxInt32
    }
    ans := math.MaxInt32
    for i := 1; i <= n; i++ {
        sum := -1
        for j := i; j >= 1; j-- {
            sum += a[j-1]
            if sum > k {
                break
            }
            if f[j-1] == math.MaxInt32 {
                continue
            }
            if i == n { 
                ans = min(ans, f[j-1])
            } else {
                f[i] = min(f[i], calc(sum, k) + f[j-1])
            }
        }
    }

    return ans
}


func calc(x, k int) int {
    return (x - k) * (x - k)
}