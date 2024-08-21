func maxDistance(arrays [][]int) (ans int) {
    mn, mx := arrays[0][0], arrays[0][len(arrays[0])-1]
    
    for i := 1; i < len(arrays); i++ {
        v := arrays[i]
        ans = max(ans, v[len(v)-1] - mn, mx - v[0])
        mn = min(mn, v[0])
        mx = max(mx, v[len(v)-1])
    }

    return
}