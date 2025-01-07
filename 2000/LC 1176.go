func dietPlanPerformance(calories []int, k int, lower int, upper int) (ans int) {
    n := len(calories)
    sum := 0
    for l, r := 0, 0; r < n; r++ {
        sum += calories[r]
        if r - l + 1 == k {
            if sum > upper {
                ans++
            } else if sum < lower {
                ans--
            }
            sum -= calories[l]
            l++
        }
    }

    return 
}