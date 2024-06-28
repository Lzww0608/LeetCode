func iceBreakingGame(num int, target int) int {
    ans := 0
    for i := 2; i <= num; i++ {
        ans = (target + ans) % i
    }

    return ans
}
