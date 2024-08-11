type MagicDictionary struct {
    m map[string]bool
}


func Constructor() MagicDictionary {
    return MagicDictionary{map[string]bool{}}
}


func (c *MagicDictionary) BuildDict(dictionary []string)  {
    for _, s := range dictionary {
        c.m[s] = true
    }
}


func (c *MagicDictionary) Search(searchWord string) bool {
    for s := range c.m {
        if len(s) != len(searchWord) {
            continue
        }
        dif := 0
        for i := range s {
            if s[i] != searchWord[i] {
                dif++
                if dif > 1 {
                    break
                }
            }
        }
        if dif == 1 {
            return true
        }
    }


    return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */