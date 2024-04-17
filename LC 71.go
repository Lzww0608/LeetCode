func simplifyPath(path string) string {
    st := []string{}

    for _, s := range strings.Split(path, "/") {
        if s == ".." {
            if len(st) > 0 {
                st = st[:len(st)-1]
            }
        } else if s != "" && s != "." {
            st = append(st, s)
        }
        
    }

    return "/" + strings.Join(st, "/")
}