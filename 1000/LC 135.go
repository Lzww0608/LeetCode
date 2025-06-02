func candy(ratings []int) (ans int) {
    n := len(ratings)
    l := make([]int, n)
    r := make([]int, n)
    for i := 1; i < n; i++ {
        if ratings[i] > ratings[i-1] {
            l[i] = l[i-1] + 1
        }
        j := n - i - 1
        if ratings[j] > ratings[j+1] {
            r[j] = r[j+1] + 1
        }
    }

    for i := range n {
        ans += max(l[i], r[i]) + 1
    }

    return
}