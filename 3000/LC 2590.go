import "sort"

type TodoList struct {
    id               int
    id2Task          map[int]map[int]bool
    task2Id          map[int]int
    task2Description map[int]string
    task2Due         map[int]int
    task2Tag         map[int][]string
    tag2Task         map[string]map[int]bool
    taskCompleted    map[int]bool
}

func Constructor() TodoList {
    return TodoList{
        id:               1,
        id2Task:          make(map[int]map[int]bool),
        task2Id:          make(map[int]int),
        task2Description: make(map[int]string),
        task2Due:         make(map[int]int),
        task2Tag:         make(map[int][]string),
        tag2Task:         make(map[string]map[int]bool),
        taskCompleted:    make(map[int]bool),
    }
}

func (c *TodoList) AddTask(userId int, taskDescription string, dueDate int, tags []string) int {
    currentTaskId := c.id

    if _, ok := c.id2Task[userId]; !ok {
        c.id2Task[userId] = make(map[int]bool)
    }
    c.id2Task[userId][currentTaskId] = true
    c.task2Id[currentTaskId] = userId
    c.task2Description[currentTaskId] = taskDescription
    c.task2Due[currentTaskId] = dueDate
    c.task2Tag[currentTaskId] = tags

    for _, tag := range tags {
        if _, ok := c.tag2Task[tag]; !ok {
            c.tag2Task[tag] = make(map[int]bool)
        }
        c.tag2Task[tag][currentTaskId] = true
    }

    c.taskCompleted[currentTaskId] = false

    c.id++
    return currentTaskId
}

func (c *TodoList) GetAllTasks(userId int) []string {
    userTasks, ok := c.id2Task[userId]
    if !ok {
        return []string{}
    }

    tmp := []int{}
    for taskId := range userTasks {
        if !c.taskCompleted[taskId] {
            tmp = append(tmp, taskId)
        }
    }

    sort.Slice(tmp, func(i, j int) bool {
        return c.task2Due[tmp[i]] < c.task2Due[tmp[j]]
    })

    ans := make([]string, 0, len(tmp))
    for _, id := range tmp {
        ans = append(ans, c.task2Description[id])
    }
    return ans
}

func (c *TodoList) GetTasksForTag(userId int, tag string) []string {
    tasksForTag, ok := c.tag2Task[tag]
    if !ok {
        return []string{}
    }

    tmp := []int{}
    for taskId := range tasksForTag {
        if c.task2Id[taskId] == userId && !c.taskCompleted[taskId] {
            tmp = append(tmp, taskId)
        }
    }

    sort.Slice(tmp, func(i, j int) bool {
        return c.task2Due[tmp[i]] < c.task2Due[tmp[j]]
    })

    ans := make([]string, 0, len(tmp))
    for _, id := range tmp {
        ans = append(ans, c.task2Description[id])
    }

    return ans
}

func (c *TodoList) CompleteTask(userId int, taskId int) {
    ownerId, taskExists := c.task2Id[taskId]
    if !taskExists || ownerId != userId {
        return 
    }

    if !c.taskCompleted[taskId] {
        c.taskCompleted[taskId] = true
    }
}