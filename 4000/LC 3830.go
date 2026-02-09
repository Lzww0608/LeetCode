func longestAlternating(a []int) int {
    ans := 1
    n := len(a)

    preAsc := make([]int, n + 1)
    preDes := make([]int, n + 1)
    sufAsc := make([]int, n + 1)
    sufDes := make([]int, n + 1)
    preAsc[0], preDes[0] = 1, 1 
    sufAsc[n - 1], sufDes[n - 1] = 1, 1

    for i := 1; i < n; i++ {
        if a[i] > a[i - 1] {
            preAsc[i] = preDes[i - 1] + 1
        } else {
            preAsc[i] = 1
        }

        if a[i] < a[i - 1] {
            preDes[i] = preAsc[i - 1] + 1
        } else {
            preDes[i] = 1
        }

        j := n - i - 1
        if a[j] > a[j + 1] {
            sufAsc[j] = sufDes[j + 1] + 1
        } else {
            sufAsc[j] = 1
        }

        if a[j] < a[j + 1] {
            sufDes[j] = sufAsc[j + 1] + 1
        } else {
            sufDes[j] = 1
        }

        ans = max(ans, preAsc[i + 1], preDes[i + 1], sufAsc[j], sufDes[j])
    }

    for i := range n - 2 {
        if a[i + 2] > a[i] {
            ans = max(ans, preDes[i] + sufAsc[i + 2])
        } else if a[i + 2] < a[i] {
            ans = max(ans, preAsc[i] + sufDes[i + 2])
        }
    }

    return ans
}

