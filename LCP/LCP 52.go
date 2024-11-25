

import "github.com/emirpasic/gods/v2/trees/redblacktree"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getNumber(root *TreeNode, ops [][]int) (ans int) {
    rbt := redblacktree.New[int, struct{}]()

    var dfs func(*TreeNode) 
    dfs = func(node *TreeNode) {
        if node == nil {
            return 
        }

        rbt.Put(node.Val, struct{}{})
        dfs(node.Left)
        dfs(node.Right)
    }    
    dfs(root)

    for i := len(ops) - 1; i >= 0; i-- {
        for {
            if it, ok := rbt.Ceiling(ops[i][1]); ok && it.Key <= ops[i][2] {
                rbt.Remove(it.Key)
                if ops[i][0] == 1 {
                    ans++
                } 
            } else {
                break
            }
        }
    }

    return 
}

