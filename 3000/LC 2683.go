func doesValidArrayExist(derived []int) bool {
    sum := 0
    for _, x := range derived {
        sum ^= x
    }
    return sum == 0
}