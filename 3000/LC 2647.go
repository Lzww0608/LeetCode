func colorRed(n int) (ans [][]int) {
    if n == 1 {
        return [][]int{{1, 1}}
    } else if n == 2 {
        return [][]int{{1, 1}, {2, 1}, {2, 3}}
    } else if n == 3 {
        return [][]int{{1, 1}, {2, 1}, {2, 3}, {3, 1}, {3, 5}}
    }
    
    ans = append(ans, []int{1, 1})
    for i := 2; i <= n; i++ {
        k := (n - i) % 4
        if k == 0 {
            for j := 1; j < i * 2; j += 2 {
                ans = append(ans, []int{i, j})
            }
        } else if k == 1 {
            ans = append(ans, []int{i, 2})
        } else if k == 2 {
            for j := 3; j < i * 2; j += 2 {
                ans = append(ans, []int{i, j})
            }
        } else {
            ans = append(ans, []int{i, 1})
        }
    }

    return
}