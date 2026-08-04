func maximumWidth(a []int) (ans int) {
    cnt := make(map[int]int)
    for _, x := range a {
        cnt[x]++
    }

    p := make(map[int]int)
    for k, v := range cnt {
        p[k] += v 
        p[k * 2] += v / 2 
        for kk, vv := range cnt {
            if kk > k {
                p[kk + k] += min(v, vv)
            }
        }
    }    

    for _, x := range p {
        ans = max(ans, x)
    }
    
    return
}