func minSensors(n int, m int, k int) int {
    return max(1, (n + k * 2) / (k * 2 + 1)) * max(1, (m + k * 2) / (k * 2 + 1))
}