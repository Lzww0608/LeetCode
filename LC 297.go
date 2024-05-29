/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    s string
}

func Constructor() Codec {
    return Codec{""}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    q := []*TreeNode{root}
    builder := strings.Builder{}
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
    this.s = builder.String()
    return this.s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {  
    
    if data == "" {
        return nil
    }
    s := strings.Split(data, " ")
    //fmt.Println(s)  
    i, n := 1, len(s)
    t, _ := strconv.Atoi(s[0]) 
    head := &TreeNode{t, nil, nil}
    q := []*TreeNode{head}
    for ;i < n && len(q) > 0; i++ {
        node := q[0]
        q = q[1:]
        if s[i] != "#" {
            t, _ = strconv.Atoi(s[i]) 
            node.Left = &TreeNode{t, nil, nil}
            q = append(q, node.Left)
        }
        i++
        if i < n && s[i] != "#" {
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