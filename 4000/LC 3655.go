const MOD int = 1_000_000_007
func xorAfterQueries(nums []int, queries [][]int) (ans int) {
    n := len(nums)
    B := int(math.Sqrt(float64(len(queries))))

    dif := make([][]int, B)
    has := make([][]bool, B)

    for _, q := range queries {
        l, r, k, v := q[0], q[1], q[2], q[3]
        if k >= B {
            for i := l; i <= r; i += k {
                nums[i] = nums[i] * v % MOD
            }
        } else {
            if len(dif[k]) == 0 {
                dif[k] = make([]int, n + k)
                has[k] = make([]bool, k)
                for i := range dif[k] {
                    dif[k][i] = 1
                }
            }
            has[k][l % k] = true
            dif[k][l] = dif[k][l] * v % MOD
            r = r - (r - l) % k + k
            dif[k][r] = dif[k][r] * quickPow(v, MOD -2) % MOD
        }
    }

    for k := 1; k < B; k++ {
        if len(dif[k]) == 0 {
            continue
        }

        for i := 0; i < k; i++ {
            if !has[k][i] {
                continue
            }

            x := 1
            for j := i; j < n; j += k {
                x = x * dif[k][j] % MOD
                nums[j] = nums[j] * x % MOD
            }
        }
    }

    for _, x := range nums {
        ans ^= x
    }

    return
}

func quickPow(a, r int) int {
    res := 1
    for r > 0 {
        if r & 1 == 1 {
            res = res * a % MOD
        }
        a = a * a % MOD
        r >>= 1
    }

    return res
}