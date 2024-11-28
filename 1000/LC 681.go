
import (
	"strconv"
	"unicode"
)
func nextClosestTime(s string) string {
    cnt := map[int]struct{}{}
    m := []int{}
    t := []int{}
    for i := range s {
        if unicode.IsDigit(rune(s[i])) {
            if _, ok := cnt[int(s[i] - '0')]; !ok {
                m = append(m, int(s[i] - '0'))
            } 
            cnt[int(s[i] - '0')] = struct{}{}
            t = append(t, int(s[i] - '0'))
        }
    }
    //fmt.Println(m)
    time := [][]int{}
    cur := []int{}
    var dfs func(idx int) 
    dfs = func(idx int) {
        if idx == 4 {
            tmp := make([]int, len(cur))
            copy(tmp, cur)
            time = append(time, tmp)
            return
        }

        for _, i := range m {
            cur = append(cur, i) 
            dfs(idx + 1)
            cur = cur[:len(cur)-1]
        }
    }
    dfs(0)

    mn, ans := math.MaxInt32, t
    for _, cur := range time {
        if tmp := solve(t, cur); tmp > 0 && tmp < mn {
            mn = tmp 
            ans = cur
            //fmt.Println(mn, ans)
        }
    }

    l := strconv.Itoa(ans[0] * 10 + ans[1])
    for len(l) < 2 {
        l = "0" + l
    }
    r := strconv.Itoa(ans[2] * 10 + ans[3])
    for len(r) < 2 {
        r = "0" + r
    }
    return l + ":" + r
}

func solve(a, b []int) int {
    if b[0] > 2 || b[2] > 5 || (b[0] == 2 && b[1] > 3) {
        return math.MaxInt32 
    } 

    if b[0] > a[0] || (b[0] == a[0] && b[1] > a[1] || (b[1] == a[1] && b[2] > a[2] || (b[2] == a[2] && b[3] > a[3]))) {
        return (23 - a[0] * 10 - a[1]) * 60 + (60 - a[2] * 10 - a[3]) - (23 - b[0] * 10 - b[1]) * 60 - (60 - b[2] * 10 - b[3])
    }

    return (b[0] * 10 + b[1]) * 60 + (b[2] * 10 + b[3]) + (23 - a[0] * 10 - a[1]) * 60 + (60 - a[2] * 10 - a[3])
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}

