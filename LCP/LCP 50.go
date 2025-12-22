func giveGem(gem []int, operations [][]int) int {
    for _, v := range operations {
        i, j := v[0], v[1]
        gem[j] += gem[i] / 2 
        gem[i] -= gem[i] / 2
    }

    mx := slices.Max(gem)
    mn := slices.Min(gem)
    return mx - mn
}