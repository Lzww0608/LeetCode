func maximalRectangle(matrix []string) (ans int) {
    if len(matrix) == 0 {
        return
    }
    n := len(matrix[0])
    f := make([]int, n + 1)

    largestRectangle := func() (res int) {

        st := []int{}
        for i := 0; i <= n; i++ {
            if len(st) == 0 || f[i] >= f[st[len(st)-1]] {
                st = append(st, i)
            } else {
                l, h := 0, f[st[len(st)-1]]
                st = st[:len(st)-1]
                if len(st) == 0 {
                    l = i
                } else {
                    l = i - st[len(st)-1] - 1
                }
                res = max(res, l * h)
                i--
            }
        }
        return 
    }

    for i := range matrix {
        for j, x := range matrix[i] {
            if x == '1' {
                f[j]++
            } else {
                f[j] = 0
            }
        }
        ans = max(ans, largestRectangle())
    }

    return
}