func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    g := make([][]bool, numCourses)
    degree := make([]int, numCourses)
    for i := range g {
        g[i] = make([]bool, numCourses)
    }
    for _, e := range prerequisites {
        a, b := e[0], e[1]
        g[a][b] = true
        degree[b]++
    }

    for k := 0; k < numCourses; k++ {
        for i := 0; i < numCourses; i++ {
            for j := 0; j < numCourses; j++ {
                if i != j && g[i][k] == true && g[k][j] == true {
                    g[i][j] = true
                }
            }
        }
    }

    ans := make([]bool, len(queries))
    for i, q := range queries {
        a, b := q[0], q[1]
        ans[i] = g[a][b]
    }

    return ans
}


func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
    g := make([][]bool, numCourses)
    degree := make([]int, numCourses)
    f := make([][]int, numCourses)
    for i := range g {
        g[i] = make([]bool, numCourses)
    }
    for _, e := range prerequisites {
        a, b := e[0], e[1]
        g[a][b] = true
        degree[b]++
        f[a] = append(f[a], b)
    }

    q := []int{}
    for i, x := range degree {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        i := q[0]
        q = q[1:]
        for _, j := range f[i] {
            g[i][j] = true
            for k := 0; k < numCourses; k++ {
                g[k][j] = g[k][i] || g[k][j]
            }
            degree[j]--
            if degree[j] == 0 {
                q = append(q, j)
            }
        }
        
    }

    ans := make([]bool, len(queries))
    for i, q := range queries {
        a, b := q[0], q[1]
        ans[i] = g[a][b]
    }

    return ans
}
