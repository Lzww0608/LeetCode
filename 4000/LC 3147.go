func maximumEnergy(energy []int, k int) int {
    n := len(energy)
    f := make([]int, n + 1)
    for i := n - 1; i >= 0; i-- {
        f[i] = energy[i] + f[min(n, i+k)]
    }

    return slices.Max(f[:n])
}