func uniqueMorseRepresentations(words []string) int {
    m := [26]string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    ans := map[string]bool{}
    for _, s := range words {
        builder := strings.Builder{}
        for i := range s {
            builder.WriteString(m[int(s[i] - 'a')])
        } 
        ans[builder.String()] = true
    } 

    return len(ans)
}