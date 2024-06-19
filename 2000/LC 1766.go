const N = 55
func getCoprimes(nums []int, edges [][]int) []int {
    n := len(nums)
    prime := make([][]bool, N)
    for i := range prime {
        prime[i] = make([]bool, N)
    }
    for i := 1; i <= 50; i++ {
        for j := 1; j <= i; j++ {
            if gcd(i, j) == 1 {
                prime[i][j] = true
                prime[j][i] = true
            }
        }
    }

    g := make([][]int, n)
    ans := make([]int, n)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }



    type pair struct {
        x, y int
    }

    m := make([]pair, N)
    for i := range m {
        m[i] = pair{-1, -1}
    }
    var dfs func(int, int, int) 
    dfs = func(root, fa, depth int) {
        //fmt.Println(root, path)
        val := nums[root]
        
        idx, closet := -1, -1
        for x, v := range m {
            if v.x != -1 && v.y > closet && prime[val][x] {
                closet = v.y
                idx = v.x
            }
        }

        ans[root] = idx
        old := m[val]
        m[val] = pair{root, depth}
        for _, node := range g[root] {
            if node != fa {
                dfs(node, root, depth+1)
            }
        } 
        m[val] = old
    }
    dfs(0, -1, 0)

    return ans
}

func gcd(x, y int) int {
    for y > 0 {
        x, y = y, x % y
    }
    return x
}
