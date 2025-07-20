
import "container/heap"
func minimumCost(nums []int, k int, dist int) int64 {
    h1 := &hp{}
    h2 := &hp{}
    m1 := make(map[int]int)
    m2 := make(map[int]int)

    sum := nums[0]
    for i := 1; i <= dist + 1; i++ {
        heap.Push(h1, nums[i])
        sum += nums[i]
        m1[nums[i]]++
        if h1.Len() > k - 1 {
            x := heap.Pop(h1).(int)
            heap.Push(h2, -x)
            sum -= x
            m2[x]++
            m1[x]--
        }
    }

    ans := sum 
    n := len(nums)
    cnt := k - 1
    for i := dist + 2; i < n; i++ {
        j := i - dist - 1 
        if m2[nums[j]] > 0 {
            m2[nums[j]]--
        } else {
            m1[nums[j]]--
            sum -= nums[j]
            cnt--
        }
        for h2.Len() > 0 && m2[-(*h2)[0]] == 0 {
            heap.Pop(h2)
        }

        if nums[i] >= (*h1)[0] {
            heap.Push(h2, -nums[i])
            m2[nums[i]]++
            if cnt < k - 1 {
                x := -heap.Pop(h2).(int)
                m2[x]--
                m1[x]++
                sum += x
                heap.Push(h1, x)
                cnt++
            }
        } else {
            heap.Push(h1, nums[i])
            m1[nums[i]]++
            sum += nums[i]
            cnt++

            for m1[(*h1)[0]] == 0 {
                heap.Pop(h1)
            }
            if cnt > k - 1 {
                cnt--
                x := heap.Pop(h1).(int)
                m1[x]--
                heap.Push(h2, -x)
                m2[x]++
                sum -= x
            }
        }
        ans = min(ans, sum)
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
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n-1]
    *h = old[:n-1]
    return 
}