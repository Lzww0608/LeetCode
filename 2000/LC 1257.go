func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
    fa := make(map[string]string)
    
    for _, v := range regions {
        for i := 1; i < len(v); i++ {
            fa[v[i]] = v[0]
        }
    }

    vis := make(map[string]bool)
    s := region1
    for {
        vis[s] = true
        if _, ok := fa[s]; !ok {
            break
        }
        s = fa[s]
    } 

    s = region2
    for !vis[s] {
        s = fa[s]
    }

    return s
}