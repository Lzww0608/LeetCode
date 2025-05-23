
import "math"
func mostSimilar(n int, roads [][]int, names []string, targetPath []string) []int {
    g := make([][]int, n)
    for _, v := range roads {
        a, b := v[0], v[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }

    f := make([]int, n)
    path := make([][]int, n)
    for _, s := range targetPath {
        ff := make([]int, n)
        for i := range ff {
            ff[i] = math.MaxInt32
        }
        tmp := make([][]int, n)
        for i := 0; i < n; i++ {
            for _, j := range g[i] {
                add := 0
                if names[j] != s {
                    add = 1
                } 
                if ff[j] > f[i] + add {
                    ff[j] = f[i] + add
                    tmp[j] = make([]int, len(path[i]))
                    copy(tmp[j], path[i])
                    tmp[j] = append(tmp[j], j)
                }
            }
        }
        f = ff 
        path = tmp
    }

    k := 0
    for i, x := range f {
        if f[k] > x {
            k = i
        }
    }

    return path[k]
}