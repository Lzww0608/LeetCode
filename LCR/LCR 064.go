type MagicDictionary struct {
    isEnd bool
    child [26]*MagicDictionary
}

var root *MagicDictionary

/** Initialize your data structure here. */
func Constructor() MagicDictionary {
    root = &MagicDictionary{}
    return *root
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
    for _, s := range dictionary {
        cur := root 
        for _, c := range s {
            x := int(c - 'a')
            if cur.child[x] == nil {
                cur.child[x] = &MagicDictionary{}
            }
            cur = cur.child[x]
        }
        cur.isEnd = true
    }
}

func (this *MagicDictionary) Search(searchWord string) bool {
    n := len(searchWord)
    var dfs func(*MagicDictionary, int, int) bool
    dfs = func(node *MagicDictionary, i int, diff int) bool {
        if node == nil {
            return false
        }
        if i == n {
            return node.isEnd && diff == 1
        }
        x := int(searchWord[i] - 'a')
        if diff == 0 {
            for j := 0; j < 26; j++ {
                if node.child[j] != nil {
                    t := 1
                    if j == x {
                        t = 0
                    }
                    if dfs(node.child[j], i+1, t) {
                        return true
                    }
                }
            }
        } else if node.child[x] != nil {
            return dfs(node.child[x], i+1, diff)
        }
        return false
    }
    return dfs(root, 0, 0)
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
