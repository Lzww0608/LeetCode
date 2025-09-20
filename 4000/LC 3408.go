type TaskManager struct {
    user map[int]int
    task map[int]int
    h hp
}


func Constructor(tasks [][]int) TaskManager {
    user := make(map[int]int)
    task := make(map[int]int)
    h := hp{}

    for _, v := range tasks {
        user[v[1]] = v[0]
        task[v[1]] = v[2]

        heap.Push(&h, pair{v[2], v[1], v[0]})
    }

    return TaskManager{user, task, h}
}


func (c *TaskManager) Add(userId int, taskId int, priority int)  {
    c.user[taskId] = userId
    c.task[taskId] = priority
    heap.Push(&c.h, pair{priority, taskId, userId})
}


func (c *TaskManager) Edit(taskId int, newPriority int)  {
    c.task[taskId] = newPriority
    heap.Push(&c.h, pair{newPriority, taskId, c.user[taskId]})
}


func (c *TaskManager) Rmv(taskId int) {
    delete(c.task, taskId)
    delete(c.user, taskId)
}


func (c *TaskManager) ExecTop() int {
    for c.h.Len() > 0 {
        cur := heap.Pop(&c.h).(pair)
        if v, ok := c.task[cur.i]; !ok || v != cur.p || c.user[cur.i] != cur.u {
            continue
        }

        delete(c.user, cur.i)
        delete(c.task, cur.i)
        return cur.u
    }

    return -1
}

type pair struct {
    p, i, u int
}
type hp []pair 
func (h hp) Len() int {return len(h)}
func (h hp) Less(i, j int) bool {return h[i].p > h[j].p || h[i].p == h[j].p && h[i].i > h[j].i}
func (h hp) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *hp) Push(x any) {
    *h = append(*h, x.(pair))
}
func (h *hp) Pop() (x any) {
    old := *h 
    n := len(old)
    x = old[n - 1]
    *h = old[:n - 1]
    return
}



/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */