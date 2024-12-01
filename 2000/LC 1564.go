func maxBoxesInWarehouse(boxes []int, warehouse []int) (ans int) {
    m, n := len(boxes), len(warehouse)
    h := make([]int, n)

    cur := warehouse[0]
    h[0] = warehouse[0]
    for i := 1; i < n; i++ {
        cur = min(cur, warehouse[i])
        h[i] = cur
    }

    sort.Ints(boxes)
    for i, j := n - 1, 0; i >= 0; i-- {
        if j >= m {
            break
        } else if boxes[j] <= h[i] {
            ans++
            j++
        } 
    }

    return ans
}