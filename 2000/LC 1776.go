func getCollisionTimes(cars [][]int) []float64 {
    n := len(cars)
    ans := make([]float64, n)
    st := []int{}

    for i := n - 1; i >= 0; i-- {
        for len(st) > 0 {
            j := st[len(st)-1]
            if cars[i][1] <= cars[j][1] {
                st = st[:len(st)-1]
            } else {
                if ans[j] < 0 {
                    break
                }
                time := float64(cars[j][0] - cars[i][0]) / float64(cars[i][1] - cars[j][1])
                if time < ans[j] {
                    break
                } else {
                    st = st[:len(st)-1]
                }
            }
        }

        if len(st) == 0 {
            ans[i] = -1.0
        } else {
            j := st[len(st)-1]
            time := float64(cars[j][0] - cars[i][0]) / float64(cars[i][1] - cars[j][1])
            ans[i] = time
        }

        st = append(st, i)
    }

    return ans
}