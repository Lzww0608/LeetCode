func bestTeamScore(scores []int, ages []int) int {
    n, m := len(scores), 0
    type pair struct {
        score, age int
    }
    a := make([]pair, n)
    for i := range a {
        m = max(m, ages[i])
        a[i] = pair{scores[i], ages[i]}
    }
    sort.Slice(a, func(i, j int) bool {
        if a[i].score == a[j].score {
            return a[i].age < a[j].age
        }
        return a[i].score < a[j].score
    })

    s := make([]int, m + 1)

    update := func(idx, val int) {
        for i := idx; i < len(s); i += i & (-i) {
            s[i] = max(s[i], val)
        }
    }

    query := func(idx int) (res int) {
        for i := idx; i > 0; i -= i & (-i) {
            res = max(res, s[i])
        }
        return
    }

    
    for _, v := range a {
        update(v.age, query(v.age) + v.score)
    }

    return query(m)
}
