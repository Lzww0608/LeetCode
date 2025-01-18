
import "math/bits"
func findKthSmallest(coins []int, k int) int64 {
    sort.Ints(coins)
    n := len(coins)
    // 刷表法
    s := make([]int, 1 << n)
    s[0] = 1
    for i := 0; i < n; i++ {
        j := 1 << i 
        for mask := 0; mask < j; mask++ {
            s[j | mask] = lcm(s[mask], coins[i])
        }
    }

    check := func(mid int) bool {
        cnt := 0
        for i := 1; i < (1 << n); i++ {
            if bits.OnesCount(uint(i)) & 1 == 0 {
                cnt += - mid / s[i]
            } else {
                cnt += mid / s[i]
            }
        }

        return cnt >= k
    }

    l, r := k - 1, coins[0] * k + 1
    for l + 1 < r {
        mid := l + ((r - l) >> 1)
        if check(mid) {
            r = mid
        } else {
            l = mid
        }
    }

    return int64(r)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}

func lcm(x, y int) int {
    return x * y / gcd(x, y)
}