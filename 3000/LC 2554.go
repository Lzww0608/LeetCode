func maxCount(banned []int, n int, maxSum int) (ans int) {
    m := make(map[int]bool)
    for _, x := range banned {
        m[x] = true
    }

    for x := 1; x <= n; x++ {
        if m[x] {
            continue
        }
        if maxSum -= x; maxSum < 0 {
            break
        }
        ans++
    }

    return
}