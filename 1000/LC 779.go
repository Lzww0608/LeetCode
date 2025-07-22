func kthGrammar(n int, k int) int {
    if k == 1 {
        return 0
    } else if k == 2 {
        return 1
    }

    if k & 1 == 0 {
        return kthGrammar(n - 1, (k + 1) / 2) ^ 1
    }

    return kthGrammar(n - 1, (k + 1) / 2)
}