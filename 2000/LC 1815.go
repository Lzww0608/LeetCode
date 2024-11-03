func maxHappyGroups(m int, groups []int) (ans int) {
    cnt := make([]int, m)
    for _, x := range groups {
        x = x % m 
        if x == 0 {
            ans++
        } else if cnt[m-x] > 0 {
            cnt[m-x]--
            ans++
        } else {
            cnt[x]++
        }
    }

    v := []int{}
    mask := 0
    for i, x := range cnt {
        if x > 0 {
            v = append(v, i)
            mask = mask << 5 | x 
        }
    }

    if len(v) == 0 {
        return
    }
    for i, j := 0, len(v) - 1; i < j; i, j = i + 1, j - 1 {
        v[i], v[j] = v[j], v[i]
    }

    type pair struct {
        left, mask int
    }
    memo := make(map[pair]int)
    var dfs func(int, int) int 
    dfs = func(left, mask int) (res int) {
        p := pair{left, mask}
        if v, ok := memo[p]; ok {
            return v 
        }
        defer func() {memo[p] = res} ()

        for i, x := range v {
            i *= 5
            if (mask >> i) & 31 != 0 {
                t := dfs((left + x) % m, mask - (1 << i))
                if left == 0 {
                    t++
                }
                res = max(res, t)
            }
        }

        return
    }

    return ans + dfs(0, mask)
}