type Encrypter struct {
    keyToVal map[byte]string
    valToKey map[string]int
    exist map[string]int
}


func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
    a := map[byte]string{}
    b := map[string]int{}
    exist := map[string]int{}
    for i := range keys {
        a[keys[i]] = values[i]
        b[values[i]]++
    }
    res := &Encrypter{a, b, exist}
    for _, s := range dictionary {
        exist[res.Encrypt(s)]++
    }
    return *res
}


func (this *Encrypter) Encrypt(word1 string) string {
    builder := strings.Builder{}
    for i := range word1 {
        if s, ok := this.keyToVal[word1[i]]; ok {
            builder.WriteString(s)
        } else {
            return ""
        }
    }
    return builder.String()
}


func (this *Encrypter) Decrypt(word2 string) int {
    return this.exist[word2]
}


/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */
