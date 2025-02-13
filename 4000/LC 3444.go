
import "math"
func minimumIncrements(nums []int, target []int) int {
    n := len(target)
    m := 1 << n
    sub := make([]int, m) 
    sub[0] = 1
    for i := 0; i < n; i++ {
        mask := 1 << i 
        for s := 0; s < mask; s++ {
            sub[s|mask] = lcm(sub[s], target[i]) 
        }
    }

    f := make([]int, m)
    for i := range f {
        f[i] = math.MaxInt32
    }
    f[0] = 0
    for _, x := range nums {
        for mask := m - 1; mask > 0; mask-- {
            for s := mask; s > 0; s = (s - 1) & mask {
                y := sub[s]
                f[mask] = min(f[mask], f[mask ^ s] + (y - x % y) % y)
            }
        }
    }

    return f[m-1]
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x 
}

func lcm(x, y int) int {
    return x / gcd(x, y) * y
}