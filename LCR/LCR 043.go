/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
    q []*TreeNode
    root *TreeNode
}


func Constructor(root *TreeNode) CBTInserter {
    que := []*TreeNode{}
    q := []*TreeNode{root}
    for len(q) > 0 {
        node := q[0]
        q = q[1:]
        if node.Left == nil || node.Right == nil {
            que = append(que, node)
        } 
        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
    }
    return CBTInserter{que, root}
}


func (this *CBTInserter) Insert(v int) int {
    node := &TreeNode{v, nil, nil}
    if this.q[0].Left == nil {
        this.q[0].Left = node
        this.q = append(this.q, node)
        return this.q[0].Val
    }
    res := this.q[0].Val
    this.q[0].Right = node
    this.q = append(this.q, node)
    this.q = this.q[1:]
    return res
}


func (this *CBTInserter) Get_root() *TreeNode {
    return this.root
}


/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
