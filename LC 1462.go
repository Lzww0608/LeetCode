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