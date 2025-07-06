
import "strconv"
type SQL struct {
    m map[string]map[int][]string
    cnt map[string]int
    id map[string]int
}


func Constructor(names []string, columns []int) SQL {
    id := make(map[string]int)
    m := make(map[string]map[int][]string)
    cnt := make(map[string]int)
    for i, name := range names {
        id[name] = 1
        cnt[name] = columns[i]
        m[name] = make(map[int][]string)
    }

    return SQL{m, cnt, id}
}


func (c *SQL) Ins(name string, row []string) bool {
    if len(row) != c.cnt[name] {
        return false
    }

    c.m[name][c.id[name]] = row
    c.id[name]++
    return true
}


func (c *SQL) Rmv(name string, rowId int)  {
    delete(c.m[name], rowId)
}


func (c *SQL) Sel(name string, rowId int, columnId int) string {
    columnId -= 1
    if v, ok1 := c.m[name]; !ok1 {
        return "<null>"
    } else if _, ok2 := v[rowId]; !ok2 || c.cnt[name] <= columnId {
        return "<null>"
    }

    return c.m[name][rowId][columnId]
}


func (c *SQL) Exp(name string) (ans []string) {
    if _, ok := c.m[name]; !ok {
        return nil
    }
    
    for i, s := range c.m[name] {
        tmp := strconv.Itoa(i) 
        tmp += "," + strings.Join(s, ",")
        ans = append(ans, tmp)
    }

    return 
}


/**
 * Your SQL object will be instantiated and called as such:
 * obj := Constructor(names, columns);
 * param_1 := obj.Ins(name,row);
 * obj.Rmv(name,rowId);
 * param_3 := obj.Sel(name,rowId,columnId);
 * param_4 := obj.Exp(name);
 */