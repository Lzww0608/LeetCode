
import "strconv"
func sumOfLargestPrimes(s string) int64 {
    a := []int{}
    m := make(map[int]bool)
    n := len(s)
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            x, _ := strconv.Atoi(s[i:j+1])
            if !m[x] && check(x) {
                m[x] = true
                a = append(a, x)
            }
        }
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i] > a[j]
    })
    
    ans := 0
    for i := 0; i < min(3, len(a)); i++ {
        ans += a[i]
    }

    return int64(ans)
}

func check(x int) bool {
    if x == 1 {
        return false
    }
    for i := 2; i * i <= x; i++ {
        if x % i == 0 {
            return false
        }
    }

    return true
}