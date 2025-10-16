const N int = 50_001

var isPrimes [N]bool 

func init() {
    isPrimes[1] = true
    for i := 2; i * i < N; i++ {
        if isPrimes[i] {
            continue
        }
        for j := i * i; j < N; j += i {
            isPrimes[j] = true
        }
    }
}

func primeSubarray(nums []int, k int) (ans int) {
    left, pre, ppre := 0, -1, -1
    mx := []int{}
    mn := []int{}

    for i, x := range nums {
        if !isPrimes[x] {
            ppre, pre = pre, i 

            for len(mx) > 0 && nums[mx[len(mx) - 1]] <= x {
                mx = mx[:len(mx) - 1]
            }
            mx = append(mx, i)
            for len(mn) > 0 && nums[mn[len(mn) - 1]] >= x {
                mn = mn[:len(mn) - 1]
            }
            mn = append(mn, i)

            for len(mn) > 0 && nums[mx[0]] - nums[mn[0]] > k {
                if left == mx[0] {
                    mx = mx[1:]
                }
                if left == mn[0] {
                    mn = mn[1:]
                }
                left++
            }
        }

        ans += ppre - left + 1
    }

    return
}