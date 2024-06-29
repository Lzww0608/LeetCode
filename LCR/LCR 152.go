func verifyTreeOrder(postorder []int) bool {
    return check(postorder, 0, len(postorder) - 1)
}

func check(postorder []int, i, j int) bool {
    if i >= j {
        return true
    }
    k := i
    for postorder[k] < postorder[j] {
        k++
    }
    p := k
    for postorder[p] > postorder[j] {
        p++
    }
    return p == j && check(postorder, i, k - 1) && check(postorder, k, j - 1)
}
