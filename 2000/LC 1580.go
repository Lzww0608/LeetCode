func maxBoxesInWarehouse(boxes []int, warehouse []int) (ans int) {
    n := len(warehouse)
    pre := make([]int, n)
    suf := make([]int, n)
    a, b := math.MaxInt32, math.MaxInt32 
    for i := 0; i < n; i++ {
        j := n - i - 1
        a = min(a, warehouse[i])
        b = min(b, warehouse[j])
        pre[i] = a 
        suf[j] = b 
    }
    for i := 0; i < n; i++ {
        warehouse[i] = max(pre[i], suf[i])
    }

    sort.Ints(boxes)
    sort.Ints(warehouse)
    for i, j := 0, 0; i < len(boxes) && j < n; {
        if boxes[i] <= warehouse[j] {
            i++
            ans++
        }
        j++
    }

    return
}