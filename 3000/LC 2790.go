func maxIncreasingGroups(usageLimits []int) int {
    sort.Ints(usageLimits)
    ans, rem := 1, 0

    for _, x := range usageLimits {
        rem += x 
        if rem >= ans {
            rem -= ans
            ans++
        }
    }

    return ans - 1
}