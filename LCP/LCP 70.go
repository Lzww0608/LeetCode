func sandyLandManagement(size int) (ans [][]int) {
    if size == 1 {
        return [][]int{{1, 1}}
    } else if size == 2 {
        return [][]int{{1, 1}, {2, 1}, {2, 3}}
    } else if size == 3 {
        return [][]int{{1, 1}, {2, 1}, {2, 3}, {3, 1}, {3, 5}}
    }

    ans = append(ans, []int{1, 1})
    for i := 2; i <= size; i++ {
        x := (size - i) % 4
        if x == 0 {
            for j := 1; j < i * 2; j += 2 {
                ans = append(ans, []int{i, j})
            }
        } else if x == 1 {
            ans = append(ans, []int{i, 2})
        } else if x == 2 {
            for j := 3; j < i * 2; j += 2 {
                ans = append(ans, []int{i, j})
            }
        } else {
            ans = append(ans, []int{i, 1})
        }
    }

    return
}