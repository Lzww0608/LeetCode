func areNumbersAscending(s string) bool {
    str := strings.Split(s, " ")
    pre := -1
    for _, t := range str {
        x, ok := strconv.Atoi(t)
        if ok != nil {
            continue
        } else {
            if x <= pre {
                return false
            }
            pre = x 
        }
    }
    
    return true
}
