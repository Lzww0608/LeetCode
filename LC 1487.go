func getFolderNames(names []string) []string {
    n := len(names)
    ans := make([]string, n)

    m := map[string]int{}
    for i, s := range names {
        if m[s] > 0 {
            for j := m[s]; true; j++ {
                t := fmt.Sprintf("%s(%d)", s, j)
                if m[t] == 0 {
                    ans[i] = t 
                    m[t] = 1
                    m[s] = j + 1
                    break
                } 
            }
        } else {
            ans[i] = s
            m[s] = 1
        }
        
    }

    return ans
}