type FileSharing struct {
    m int
    id map[int]bool
    id2File map[int][]int
    file2Id []map[int]bool
}


func Constructor(m int) FileSharing {
    file2Id := make([]map[int]bool, m + 1)
    for i := range m + 1 {
        file2Id[i] = make(map[int]bool)
    }
    return FileSharing{
        m,
        map[int]bool{},
        map[int][]int{},
        file2Id,
    }
}


func (c *FileSharing) Join(ownedChunks []int) int {
    id := 1
    for _, ok := c.id[id]; ok; _, ok = c.id[id] {
        id++
    }
    c.id[id] = true

    for _, v := range ownedChunks {
        c.file2Id[v][id] = true
    }

    c.id2File[id] = ownedChunks
    return id
}


func (c *FileSharing) Leave(userID int)  {
    delete(c.id, userID)

    for _, x := range c.id2File[userID] {
        delete(c.file2Id[x], userID)
    }
    
    delete(c.id2File, userID)
}


func (c *FileSharing) Request(userID int, chunkID int) (ans []int) {
    for x := range c.file2Id[chunkID] {
        ans = append(ans, x)
    }
    sort.Ints(ans)
    if _, ok := c.file2Id[chunkID][userID];len(ans) > 0 && !ok {
        c.file2Id[chunkID][userID] = true
        c.id2File[userID] = append(c.id2File[userID], chunkID)
    }
    return
}


/**
 * Your FileSharing object will be instantiated and called as such:
 * obj := Constructor(m);
 * param_1 := obj.Join(ownedChunks);
 * obj.Leave(userID);
 * param_3 := obj.Request(userID,chunkID);
 */