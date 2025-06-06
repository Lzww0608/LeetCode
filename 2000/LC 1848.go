func getMinDistance(nums []int, target int, start int) int {
    ans := math.MaxInt32 
    for i, x := range nums {
        if x == target {
            ans = min(ans, abs(i - start))
        }
    }

    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}