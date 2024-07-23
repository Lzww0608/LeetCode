func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
    for i, g := range group {
        if g == -1 {
            group[i] = m
            m++
        }
    }

    groupAdj := make([][]int, m)
    itemAdj := make([][]int, n)
    groupIn := make([]int, m)
    itemIn := make([]int, n)

    for i, curGroup := range group {
        for _, item := range beforeItems[i] {
            g := group[item]
            if g != curGroup {
                groupAdj[g] = append(groupAdj[g], curGroup)
                groupIn[curGroup]++
            }
            itemAdj[item] = append(itemAdj[item], i)
            itemIn[i]++
        }
    }

    groupList := topologicalSort(groupAdj, groupIn, m)
    if len(groupList) == 0 {
        return nil
    }

    itemList := topologicalSort(itemAdj, itemIn, n)
    if len(itemList) == 0 {
        return nil
    }

    group2Item := make([][]int, m)
    for _, item := range itemList {
        group2Item[group[item]] = append(group2Item[group[item]], item)
    }

    ans := []int{}
    for _, g := range groupList {
        ans = append(ans, group2Item[g]...)
    }

    return ans
}

func topologicalSort(Adj [][]int, In []int, n int) []int {
    ans := []int{}
    q := []int{}
    for i, x := range In {
        if x == 0 {
            q = append(q, i)
        }
    }

    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        ans = append(ans, u)
        for _, v := range Adj[u] {
            if In[v]--; In[v] == 0 {
                q = append(q, v)
            }
            
        }
    }

    if len(ans) < n {
        return nil
    }

    return ans
}