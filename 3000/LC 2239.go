func findClosestNumber(nums []int) (ans int) {
    ans = math.MaxInt32
    for _, x := range nums {
        if abs(ans) > abs(x) || abs(ans) == abs(x) && ans < x {
            ans = x
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