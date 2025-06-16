func totalFruit(fruits []int) (ans int) {
    types := make(map[int]int)
    n := len(fruits)

    for l, r := 0, 0; r < n; r++ {
        types[fruits[r]]++
        for len(types) > 2 {
            x := fruits[l]
            if types[x]--; types[x] == 0 {
                delete(types, x)
            }
            l++
        }
        ans = max(ans, r - l + 1)
    }

    return
}