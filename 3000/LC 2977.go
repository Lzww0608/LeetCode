
import (
	"math"
)
const N int = 26
func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
    root := &node{}
    id := 0

    insert := func(s string) int {
        cur := root
        for _, c := range s {
            c -= 'a'
            if cur.children[c] == nil {
                cur.children[c] = &node{id:-1}
            }
            cur = cur.children[c]
        }

        if cur.id < 0 {
            cur.id = id
            id++
        }
        return cur.id
    }

    m := len(cost)
    dis := make([][]int, m * 2)
    for i := range dis {
        dis[i] = make([]int, m * 2)
        for j := range dis[i] {
            if j != i {
                dis[i][j] = math.MaxInt / 2 
            }
        }
    }

    for i, t := range cost {
        x := insert(original[i])
        y := insert(changed[i])
        dis[x][y] = min(dis[x][y], t)
    }

    for k := 0; k < id; k++ {
        for i := 0; i < id; i++ {
            if dis[i][k] == math.MaxInt / 2 {
                continue
            }
            for j := 0; j < id; j++ {
                dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j])
            }
        }
    }

    n := len(source)
    memo := make([]int, n)
    for i := range memo {
        memo[i] = -1
    }
    var dfs func(int) int
    dfs = func(i int) int {
        if i >= n {
            return 0
        }
        p := &memo[i]
        if *p != -1 {
            return *p
        }

        ans := math.MaxInt / 2
        if source[i] == target[i]{
            ans = dfs(i + 1)
        }
        x, y := root, root
        for j := i; j < n; j++ {
            x = x.children[source[j]-'a']
            y = y.children[target[j]-'a']
            if x == nil || y == nil {
                break
            }
            if x.id >= 0 && y.id >= 0 {
                ans = min(ans, dis[x.id][y.id] + dfs(j + 1))
            }
        }
        *p = ans
        return ans
    }
    ans := dfs(0)
    if ans == math.MaxInt / 2 {
        return -1
    }
    return int64(ans)
}

type node struct {
    children [N]*node
    id int
}