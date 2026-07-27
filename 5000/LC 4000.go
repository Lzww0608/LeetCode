func largestInteger(n int, s int) (ans int) {
    if n * 9 < s {
        return -1
    }

    for range n {
        x := min(s, 9)
        ans = ans * 10 + x
        s -= x
    }

    return 
}