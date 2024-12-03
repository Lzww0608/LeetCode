import "strconv"
type ValidWordAbbr struct {
    abbre map[string]int
    str map[string]bool
}

func translate(s string) string {
    if len(s) <= 2 {
        return s
    }
    tmp := []byte{}
    tmp = append(tmp, s[0])
    t := strconv.Itoa(len(s)-2)
    tmp = append(tmp, []byte(t)...)
    tmp = append(tmp, s[len(s)-1])
    return string(tmp)
}

func Constructor(dictionary []string) ValidWordAbbr {
    abbre := make(map[string]int)
    str := make(map[string]bool)
    for _, s := range dictionary {
        if str[s] {
            continue
        }
        abbre[translate(s)]++
        str[s] = true
    }
    //fmt.Println(abbre)
    return ValidWordAbbr{abbre, str}
}


func (c *ValidWordAbbr) IsUnique(word string) bool {
    s := translate(word)
    return c.str[word] && c.abbre[s] == 1 || c.abbre[s] == 0
}


/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */