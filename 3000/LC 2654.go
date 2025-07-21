
import "math"
func minOperations(nums []int) int {
    g := 0
    n := len(nums)
    cnt := 0

    for _, x := range nums {
        g = gcd(x, g)
        if x == 1 {
            cnt++
        }
    }

    if g != 1 {
        return -1
    } else if cnt > 0 {
        return n - cnt
    }

    type pair struct {
        x, i int 
    }

    f := []pair{}
    ans := math.MaxInt32
    for i, x := range nums {
        for j := range f {
            f[j].x = gcd(x, f[j].x)
        }
        f = append(f, pair{x, i})

        j := 0
        for _, p := range f[1:] {
            if f[j].x != p.x {
                j++
                f[j] = p 
            } else {
                f[j].i = p.i
            }
        }

        f = f[:j+1]
        if f[0].x == 1 {
            ans = min(ans, n - 1 + i - f[0].i)
        }
    }

    return ans
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b 
    }

    return a
}