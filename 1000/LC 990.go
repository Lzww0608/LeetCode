func equationsPossible(equations []string) bool {
    n := len(equations)
    f := make([]int, n << 1)
    m := map[byte]int{}

    t := 0
    for _, e := range equations {
        if _, ok := m[e[0]]; !ok {
            m[e[0]] = t
            t++
        }
        if _, ok := m[e[3]]; !ok {
            m[e[3]] = t
            t++
        }
    }

    for i := 0; i < t; i++ {
        f[i] = i
    }
    
    var find func(int) int
    find = func(x int) int {
        if x != f[x] {
            f[x] = find(f[x])
        }
        return f[x]
    }

    merge := func(x, y int) {
        rx, ry := find(x), find(y)
        if rx != ry {
            f[rx] = ry
        }
    }

    for _, e := range equations {
        a, b := find(m[e[0]]), find(m[e[3]])
        if e[1] == '=' {
            merge(a, b)
        }
    }

    for _, e := range equations {
        a, b := find(m[e[0]]), find(m[e[3]])
        if e[1] == '!' {
            if a == b {
                return false
            }
        } 
    }


    return true
}