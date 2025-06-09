func videoStitching(clips [][]int, time int) (ans int) {
    maxDis := make([]int, time)
    for _, v := range clips {
        if v[0] < time {
            maxDis[v[0]] = max(maxDis[v[0]], v[1])
        }
    }

    nxt, cur := 0, 0
    for i := 0; i < time; i++ {
        nxt = max(nxt, maxDis[i])
        if i == nxt {
            return -1
        }
        if i == cur {
            ans++
            cur = nxt
        }
    }

    return 
}