func minAllOneMultiple(k int) int {
    if k & 1 == 0 || k % 5 == 0 {
        return -1
    }

    ans := 1
    for n := 1; ; n = n * 10 + 1 {
        if n % k == 0 {
            break
        } else {
            n %= k
            ans++
        }
    } 

    return ans
}