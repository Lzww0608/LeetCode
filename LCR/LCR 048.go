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

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    builder := strings.Builder{}

    q := []*TreeNode{root}
    for len(q) > 0 {
        for k := len(q); k > 0; k-- {
            node := q[0]
            q = q[1:]
            if node == nil {
                builder.WriteString("# ")
                continue
            }
            builder.WriteString(strconv.Itoa(node.Val))
            builder.WriteString(" ")
            q = append(q, node.Left)
            q = append(q, node.Right)
        }
    }

    return builder.String()
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {    
    if data == "" {
        return nil
    }
    s := strings.Split(data, " ")
    n := len(s)
    t, _ := strconv.Atoi(s[0])
    head := &TreeNode{t, nil, nil}
    q := []*TreeNode{head}
    for i := 1; i < n && len(q) > 0; i++ {
        node := q[0]
        q = q[1:]
        
        if s[i] != "#" {
            t, _ = strconv.Atoi(s[i])
            node.Left = &TreeNode{t, nil, nil}
            q = append(q, node.Left)
        }

        if i++; i < n && s[i] != "#" {
            t, _ = strconv.Atoi(s[i])
            node.Right = &TreeNode{t, nil, nil}
            q = append(q, node.Right)
        }
        
    }

    return head
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */