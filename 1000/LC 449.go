/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

func encode(root *TreeNode) string {
    if root == nil {
        return "#"
    }

    l := encode(root.Left)
    r := encode(root.Right)
    return strconv.Itoa(root.Val) + " " + l + " " + r;
}

func decode(s []string, pos *int) *TreeNode {
    if s[*pos] == "#" {
        *pos++
        return nil
    }
    x, _ := strconv.Atoi(s[*pos])
    *pos++
    cur := &TreeNode{x, nil, nil}
    cur.Left = decode(s, pos)
    cur.Right = decode(s, pos)
    return cur
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
    return encode(root)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {   
    if data == " #" {
        return nil
    } 
    s := strings.Split(data, " ")
    pos := 0
    return decode(s, &pos)
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */