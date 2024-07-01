func convert(s string, numRows int) (ans string) {
    if numRows == 1 {
        return s
    }
    
    m := make([]string, numRows)
    flag, i := -1, 0
    for _, c := range s {
        m[i] += string(c)
        if i == 0 || i == numRows - 1 {
            flag = -flag
        }
        i += flag
    }
    for _, str := range m {
        ans += str
    }
    return 
}
