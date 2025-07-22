const N int = 26
type WordDictionary struct {
    cnt int 
    isEnd bool 
    children [N]*WordDictionary
}


func Constructor() WordDictionary {
    return WordDictionary{}
}


func (c *WordDictionary) AddWord(word string)  {
    cur := c 
    for _, c := range word {
        x := int(c - 'a')
        if cur.children[x] == nil {
            cur.children[x] = &WordDictionary{}
        }
        cur = cur.children[x]
        cur.cnt++
    }
    cur.isEnd = true 
    return
}


func (c *WordDictionary) Search(word string) bool {
    if c == nil {
        return false
    }
    if word == "" {
        return c.isEnd
    } 

    ans := false
    if word[0] == '.' {
        for _, v := range c.children {
            ans = ans || v.Search(word[1:])
        }
    } else {
        x := int(word[0] - 'a')
        ans = c.children[x].Search(word[1:])
    }

    return ans
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */