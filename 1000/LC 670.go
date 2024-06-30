func maximumSwap(num int) int {
    s := strconv.Itoa(num)
    bytes := []byte(s) // Convert string to slice of bytes for mutability
    maxIdx := len(s) - 1

    x, y := 0, 0
    for i := len(s) - 1; i >= 0; i-- {
        if bytes[i] > bytes[maxIdx] {
            maxIdx = i
        } else if bytes[i] < bytes[maxIdx] {
            x, y = i, maxIdx // Find the first pair to swap
        }
    }

    // Swap
    bytes[x], bytes[y] = bytes[y], bytes[x]

    // Convert back to int
    result, _ := strconv.Atoi(string(bytes))
    return result
}
