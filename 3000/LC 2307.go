
import "math"
const EPS = 1e-5
func checkContradictions(equations [][]string, values []float64) bool {
    cnt := 0
    m := make(map[string]int)
    
    search := func(s string) int {
        if _, ok := m[s]; !ok {
            m[s] = cnt 
            cnt++
        }
        return m[s]
    }

    for _, v := range equations {
        search(v[0])
        search(v[1])
    }

    p := make([]int, cnt)
    d := make([]float64, cnt)
    for i := range p {
        p[i] = i
        d[i] = 1.0
    }

    var find func(int) int 
    find = func(x int) int {
        if p[x] != x {
            root := find(p[x])
            d[x] *= d[p[x]]
            p[x] = root
        }
        
        return p[x]
    }
    
    // a / b = w 
    // a = b * w
    for i, v := range equations {
        sa, sb := v[0], v[1]
        a, b := search(sa), search(sb)
        w := values[i]
        fa, fb := find(a), find(b)
        if fa != fb {
            p[fb] = fa
            d[fb] = w * d[a] / d[b]
        } else {
            if math.Abs(w * d[a] - d[b]) >= EPS {
                return true
            } 
        }
    }

    return false
}