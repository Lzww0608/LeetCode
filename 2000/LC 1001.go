type pair struct {
    x, y int
}
func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
    m := len(queries)
    ans := make([]int, m)

    mp := map[int]int{}
    mp1 := map[int]int{}
    mp2 := map[int]int{}
    mp3 := map[int]int{}
    mp_point := map[pair]bool{}
    
    for _, v := range lamps {
        if mp_point[pair{v[0], v[1]}] == true {
            continue
        }
        mp[v[0]]++
        mp1[v[1]]++
        mp2[v[0] - v[1]]++
        mp3[v[0] + v[1]]++
        mp_point[pair{v[0], v[1]}] = true
    }

    for k, q := range queries {
        if mp[q[0]] > 0 || mp1[q[1]] > 0 || mp2[q[0]-q[1]] > 0 || mp3[q[0]+q[1]] > 0 {
            ans[k] = 1
        }
        
        for i := -1; i <= 1; i++ {
            for j := -1; j <= 1; j++ {
                a, b := q[0] + i, q[1] + j
                if a < 0 || a >= n || b < 0 || b >= n || !mp_point[pair{a, b}] {
                    continue
                }
                delete(mp_point, pair{a, b})
                mp[a]--
                mp1[b]--
                mp2[a - b]--
                mp3[a + b]--
            }
        }   
    }
    return ans
}
