func pathSum(nums []int) (ans int) {
    n := len(nums)
    a := make([][3]int, n)
    leaf := make([][20]bool, 6)
    for i, x := range nums {
        a[i][2] = x % 10
        x /= 10
        a[i][1] = x % 10
        a[i][0] = x / 10
        leaf[a[i][0]][a[i][1]] = true
    }
    

    sort.Slice(a, func(i, j int) bool {
        return a[i][0] < a[j][0] || (a[i][0] == a[j][0] && a[i][1] > a[j][1])
    })
    sum := make([]int, 9)
    for _, v := range a {
        x := (v[1] + 1) / 2
        sum[v[1]] = v[2] + sum[x]
        if !leaf[v[0]+1][v[1]*2] && !leaf[v[0]+1][v[1]*2-1] {
            ans += sum[v[1]]
        }
    }

    return
}