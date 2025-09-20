const N int = 26
type Spreadsheet struct {
    a [N][]int 
    n int
}


func Constructor(rows int) Spreadsheet {
    a := [N][]int{}
    for i := range a {
        a[i] = make([]int, rows)
    }

    return Spreadsheet{a, rows}
}


func (c *Spreadsheet) SetCell(cell string, value int)  {
    x := int(cell[0] - 'A')
    y, _ := strconv.Atoi(cell[1:])
    c.a[x][y - 1] = value
}


func (c *Spreadsheet) ResetCell(cell string)  {
    c.SetCell(cell, 0)
}


func (c *Spreadsheet) GetValue(formula string) (ans int) {
    formula = formula[1:]
    s := strings.Split(formula, "+")
    for _, t := range s {
        if t[0] >= 'A' && t[0] <= 'Z' {
            x := int(t[0] - 'A')
            y, _ := strconv.Atoi(t[1:])
            ans += c.a[x][y - 1]
        } else {
            x, _ := strconv.Atoi(t)
            ans += x
        }
    }

    return
}


/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */