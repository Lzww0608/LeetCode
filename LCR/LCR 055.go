/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    st []*TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    st := []*TreeNode{}
    for root != nil {
        st = append(st, root)
        root = root.Left
    }
    return BSTIterator{st}
}


func (this *BSTIterator) Next() int {
    cur := this.st[len(this.st)-1]
    this.st = this.st[:len(this.st)-1]
    res := cur.Val

    for cur = cur.Right; cur != nil; cur = cur.Left {
        this.st = append(this.st, cur)
    }

    return res
}


func (this *BSTIterator) HasNext() bool {
    return len(this.st) > 0 
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
