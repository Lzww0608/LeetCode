func maxSumDivThree(nums []int) int {
    f := [3]int{0, math.MinInt, math.MinInt}
    
    for _, x := range nums {
        f[x%3], f[(x+1)%3], f[(x+2)%3] = max(f[x%3], x + f[0]), max(f[(x+1)%3], x + f[1]), max(f[(x+2)%3], x + f[2])
    }

    return f[0]
}