type folder struct {
    son map[string]*folder
    name string 
    isDel bool
}
func deleteDuplicateFolder(paths [][]string) (ans [][]string) {
    root := &folder{}
    for _, path := range paths {
        cur := root
        for _, s := range path {
            if cur.son == nil {
                cur.son = map[string]*folder{}
            }
            if cur.son[s] == nil {
                cur.son[s] = &folder{}
            }
            cur = cur.son[s]
            cur.name = s 
        }
    }

    folders := make(map[string][]*folder)
    var dfs func(*folder) string 
    dfs = func(f *folder) string {
        if f.son == nil {
            return "(" + f.name + ")"
        }
        m := make([]string, 0, len(f.son))
        for _, son := range f.son {
            m = append(m, dfs(son))
        }
        sort.Strings(m)
        sub := strings.Join(m, "")
        folders[sub] = append(folders[sub], f)
        return "(" + f.name + sub + ")"
    }
    dfs(root)

    for _, f := range folders {
        if len(f) > 1 {
            for _, v := range f {
                v.isDel = true
            } 
        }
    }

    path := []string{}
    var dfs2 func(*folder)
    dfs2 = func(f *folder) {
        if f.isDel {
            return
        }
        path = append(path, f.name)
        ans = append(ans, append([]string(nil), path...))
        for _, son := range f.son {
            dfs2(son)
        }
        path = path[:len(path)-1]
    }

    for _, son := range root.son {
        dfs2(son)
    }

    return 
}