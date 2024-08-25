func maxCompatibilitySum(students [][]int, mentors [][]int) int {
    m := len(students)

    cnt := make([][]int, m)
    for i := range cnt {
        cnt[i] = make([]int, m)
    }
    for i := range students {
        for j := range mentors {
            sum := 0
            for k := range students[i] {
                if students[i][k] == mentors[j][k] {
                    sum++
                }
            }
            cnt[i][j] = sum
        }
    }

    f := make([]int, 1 << m)
    for i := range students {
        for j := 1 << m - 1; j > 0; j-- {
            for k := 0; k < m; k++ {
                if (j >> k) & 1 == 1 {
                    f[j] = max(f[j], f[j ^ (1 << k)] + cnt[i][k])
                }
            }
        } 
    }

    return f[len(f)-1]
}