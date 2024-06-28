/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
    root *TreeNode
    leaf []*TreeNode
    idx int
}


func Constructor(root *TreeNode) CBTInserter {
    leaf, root := []*TreeNode{}, root
    n := 0
    var dfs func(*TreeNode, int)
    dfs = func(node *TreeNode, d int) {
        if node == nil{
            n = max(n, d)
            return
        }
        dfs(node.Left,d+1)
        dfs(node.Right,d+1)
    }
    dfs(root, 0)
    var dfs1 func(*TreeNode, int, int)
    dfs1 = func(node *TreeNode, d int, n int) {
        if node == nil {
            return
        }
        if (node.Left == nil || node.Right == nil) && d == n{
            leaf = append(leaf, node)
            return
        }
        dfs1(node.Left,d+1,n)
        dfs1(node.Right,d+1,n)
    }
    dfs1(root,0,n-2)
    dfs1(root,0,n-1)
    return CBTInserter{root,leaf,0}
}


func (this *CBTInserter) Insert(val int) int {
    if this.leaf[this.idx].Left == nil {
        this.leaf[this.idx].Left = &TreeNode{Val:val}
        this.leaf = append(this.leaf, this.leaf[this.idx].Left)
        return this.leaf[this.idx].Val
    }
    this.leaf[this.idx].Right = &TreeNode{Val:val}
    this.leaf = append(this.leaf, this.leaf[this.idx].Right)
    this.idx++
    return this.leaf[this.idx-1].Val
}


func (this *CBTInserter) Get_root() *TreeNode {
    return this.root
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
