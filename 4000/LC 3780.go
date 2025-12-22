func maximumSum(nums []int) (ans int) {
    a := [3]int{}
    b := [3]int{}
    for _, x := range nums {
        y := x % 3
        t := (3 - x % 3) % 3
        if b[t] != 0 {
            ans = max(ans, b[t] + x)
        }
        for i := range 3 {
            if a[i] != 0 {
                b[(i + y) % 3] = max(b[(i + y) % 3], a[i] + x)
            }
        }
        a[y] = max(a[y], x)
    }

    return
}