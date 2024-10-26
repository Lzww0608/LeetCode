const N int = 26
func maxNumOfSubstrings(s string) (ans []string) {
    dis := make([][2]int, N)
    n := len(s)
    for i := range dis {
        dis[i] = [2]int{n, -1}
    }
    for i := 0; i < n; i++ {
		c := s[i] - 'a'
		if dis[c][0] == n {
			dis[c][0] = i
		}
		dis[c][1] = i
	}
    
    for i := 0; i < N; i++ {
        l, r := dis[i][0], dis[i][1]
        if r == -1 {
            continue
        }
        for j := l; j <= r; j++ {
            x := s[j] - 'a'
            dis[i][0] = min(dis[i][0], dis[x][0])
            dis[i][1] = max(dis[i][1], dis[x][1])
        }
    }

    for i := 0; i < N; i++ {
        l, r := dis[i][0], dis[i][1]
        if r == -1 {
            continue
        }
        for j := l; j <= r; j++ {
            x := s[j] - 'a'
            dis[i][0] = min(dis[i][0], dis[x][0])
            dis[i][1] = max(dis[i][1], dis[x][1])
        }
    }

    sort.Slice(dis, func(i, j int) bool {
        if dis[i][1] == -1 {
            return false
        } else if dis[j][1] == -1 {
            return true
        }
        return dis[i][1] < dis[j][1] || (dis[i][1] == dis[j][1] && dis[i][1] - dis[i][0] < dis[j][1] - dis[j][0])
    })

    
    l, r := -1, -1
    for _, v := range dis {
        if v[1] == -1 {
            break
        }
        if v[0] > r {
            l, r = v[0], v[1]
            ans = append(ans, s[l:r+1])
        }
    }

    return
}