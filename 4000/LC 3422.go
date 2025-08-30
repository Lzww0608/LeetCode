func minOperations(nums []int, k int) int64 {
    pq1 := hp{}
    pq2 := hp{}

    pre, suf := 0, 0

    for i := 0; i < k; i++ {
        heap.Push(&pq1, nums[i])
        pre += nums[i]
    }
    for i := 0; i < k / 2; i++ {
        t := heap.Pop(&pq1).(int)
        heap.Push(&pq2, -t)
        pre -= t
        suf += t
    }

    calc := func() int {
        if k & 1 == 1 {
            return suf - pre + pq1[0]
        }
        return suf - pre 
    }

    m := make(map[int]int)
    ans := calc()
    for i := k; i < len(nums); i++ {
        balance := 0
        x, y := nums[i - k], nums[i]
        m[x]++
        
        if pq1.Len() > 0 && x <= pq1[0] {
            balance--
            pre -= x
        } else {
            balance++
            suf -= x
        }

        if pq1.Len() > 0 && y <= pq1[0] {
            heap.Push(&pq1, y)
            pre += y
            balance++
        } else {
            heap.Push(&pq2, -y)
            suf += y
            balance--
        }

        for pq1.Len() > 0 && m[pq1[0]] > 0 {
            m[pq1[0]]--
            heap.Pop(&pq1)
        }
        for pq2.Len() > 0 && m[-pq2[0]] > 0 {
            m[-pq2[0]]--
            heap.Pop(&pq2)
        }

        if balance > 0 {
            t := heap.Pop(&pq1).(int)
            heap.Push(&pq2, -t)
            if m[t] == 0 {
                pre -= t 
                suf += t
            }
            
        } else if balance < 0 {
            t := -heap.Pop(&pq2).(int)
            heap.Push(&pq1, t)
            if m[t] == 0 {
                pre += t 
                suf -= t
            }
            
        }

        for pq1.Len() > 0 && m[pq1[0]] > 0 {
            m[pq1[0]]--
            heap.Pop(&pq1)
        }
        for pq2.Len() > 0 && m[-pq2[0]] > 0 {
            m[-pq2[0]]--
            heap.Pop(&pq2)
        }

        //fmt.Println(pre, suf)
        ans = min(ans, calc())
    }

    return int64(ans)
}

type hp []int
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i] > h[j]}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(int))
}
func (h *hp) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func abs(x int) int {
    if x < 0 {
        return -x
    }

    return x
}