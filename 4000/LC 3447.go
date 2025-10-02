func assignElements(groups []int, elements []int) []int {
    n := len(groups)
    ans := make([]int, n)
    mx := slices.Max(groups)

    a := make([]int, mx + 1)
    for i, x := range elements {
        if x > mx || a[x] > 0 {
            continue
        }
        for y := x; y <= mx; y += x {
            if a[y] == 0 {
                a[y] = i + 1
            }
        }
    }

    for i, x := range groups {
        ans[i] = a[x] - 1
    }

    return ans
}