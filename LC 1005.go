func largestSumAfterKNegations(a []int, k int) (ans int) {
    nums := map[int]int{}
    for _, x := range a {
        nums[x]++
        ans += x
    }
    for i := -100; i < 0 && k > 0; i++{
        if nums[i] > 0 {
            t := min(k, nums[i])
            ans -= i * 2 * t
            k -= t
            nums[-i] += t
        } 
    }
    if k & 1 == 1 && nums[0] == 0 {
        for i := 1; i <= 100; i++ {
            if nums[i] > 0 {
                ans -= i * 2
                break
            }
        }
    }
    

    return 
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}