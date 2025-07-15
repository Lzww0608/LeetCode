
const N int = 26
func leastInterval(tasks []byte, n int) int {
    cnt := [N]int{}
    for _, c := range tasks {
        cnt[int(c - 'A')]++
    }

    m, c := 0, 0
    for _, x := range cnt {
        if x > m {
            m, c = x, 1
        } else if x == m {
            c++
        }
    }

    return max(len(tasks), (n + 1) * (m - 1) + c)
}

