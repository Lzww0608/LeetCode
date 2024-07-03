func divingBoard(shorter int, longer int, k int) (ans []int) {

    if k == 0 {
        return 
    } else if shorter == longer {
        ans = append(ans, shorter * k)
        return
    }

    for i := 0; i <= k; i++ {
        t := i * shorter + (k - i) * longer
        ans = append(ans, t)
    }
    sort.Ints(ans)
    
    return 
}