func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
    vis := make([]int, n)
    for i := range vis {
        vis[i] = -2
    }
    
    var dfs func(int, int) bool
    dfs = func(id int, fa int) bool {
        if vis[id] != -2 {
            if vis[id] == fa {
                return true
            }
            return false
        } 
        vis[id] = -1
        if leftChild[id] != -1 {
            if !dfs(leftChild[id], id) {
                
                return false
            }
        }
        if rightChild[id] != -1 {
            if !dfs(rightChild[id], id) {
                
                return false
            }
        }
        vis[id] = fa
        return true
    }

    for i := range leftChild {
        if vis[i] != -2 {
            continue
        }
        if leftChild[i] != -1 {
            if !dfs(leftChild[i], i) {
                return false
            }
        }
        if rightChild[i] != -1 {
            if !dfs(rightChild[i], i) {
                return false
            }
        }
    }

   
    cnt := 0
    for _, x := range vis {
        if x == -2 {
            cnt++
        }
    }
    fmt.Println(vis)
    return cnt == 1
}