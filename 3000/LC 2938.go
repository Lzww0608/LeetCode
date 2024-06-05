func minimumSteps(s string) (ans int64) {
    var cnt int64 = 0
    for _, c := range s {
        if c == '1' {
            cnt++
        } else {
            ans += cnt
        }
    }
    
    return
}