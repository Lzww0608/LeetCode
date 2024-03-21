func countOdds(low int, high int) int {
    ans := (high - low) >> 1

    if high & 1 == 1 || low & 1 == 1 {
        ans++
    }


    return ans
}