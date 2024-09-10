
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    arr []int
    p int
    st []*TreeNode
    last *TreeNode
}


func Constructor(root *TreeNode) BSTIterator {
    return BSTIterator{
        arr: []int{},
        p: -1,
        st: []*TreeNode{},
        last: root,
    }
}


func (c *BSTIterator) HasNext() bool {
    return c.last != nil || c.p < len(c.arr) - 1 || len(c.st) > 0
}


func (c *BSTIterator) Next() int {
    c.p++
    if c.p == len(c.arr) {
        for c.last != nil {
            c.st = append(c.st, c.last)
            c.last = c.last.Left
        }
        cur := c.st[len(c.st)-1]
        c.st = c.st[:len(c.st)-1]
        c.arr = append(c.arr, cur.Val)
        c.last = cur.Right
    }

    return c.arr[c.p]
}


func (c *BSTIterator) HasPrev() bool {
    return c.p > 0
}


func (c *BSTIterator) Prev() int {
    c.p--
    return c.arr[max(0, c.p)]
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.HasNext();
 * param_2 := obj.Next();
 * param_3 := obj.HasPrev();
 * param_4 := obj.Prev();
 */