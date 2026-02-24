func maxCaloriesBurnt(heights []int) int64 {
    sort.Ints(heights)
    n := len(heights)
    ans := heights[n - 1] * heights[n - 1]

    for i := 0; i < n / 2; i++ {
        ans += (heights[i] - heights[n - i - 1]) * (heights[i] - heights[n - i - 1])
        if n - i - 2 >= 0 {
            ans += (heights[i] - heights[n - i - 2]) * (heights[i] - heights[n - i - 2])
        }
    }

    return int64(ans)
}