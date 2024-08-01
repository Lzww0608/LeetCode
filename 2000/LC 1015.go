func smallestRepunitDivByK(k int) int {
    if k & 1 == 0 || k % 5 == 0 {
        return -1
    }

    x := 0
    for i := 1; i <= k; i++ {
        x = (x * 10 + 1) % k
        if x == 0 {
            return i
        }
    }

    return -1
}