func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, x := range nums {
        sum += x
    }

    if abs(target) > sum || (sum - target) & 1 == 1 {
        return 0
    }

    t := min((sum - target) >> 1, (sum + target) >> 1)
    f := make([]int, t + 1)
    f[0] = 1
    for _, x := range nums {
        for j := t; j >= x; j-- {
            f[j] += f[j-x]
        }
    }
    
    return f[t]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}