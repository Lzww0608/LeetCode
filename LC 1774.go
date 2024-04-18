func closestCost(baseCosts []int, toppingCosts []int, target int) int {
    f := make([]bool, target + 1)
    ans := math.MaxInt32

    for _, x := range baseCosts {
        if x > target {
            ans = min(ans, x)
        } else {
            f[x] = true
        }
    }

    for _, x := range toppingCosts {
        for i := 0; i < 2; i++ {
            for j := target; j >= 0; j-- {
                if f[j] && j + x > target {
                    ans = min(ans, j + x)
                }
                if j > x {
                    f[j] = f[j] || f[j-x]
                }
            }
        }
    }

    for i := 0; i <= ans - target && i <= target; i++ {
        if f[target - i] {
            return target - i
        }
    }

    return ans
}