/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) (ans int) {
    m := map[int]int{}
    for i, v := range employees {
        m[v.Id] = i
    }
    q := []int{id}
    for len(q) > 0 {
        u := q[0]
        q = q[1:]
        ans += employees[m[u]].Importance
        for _, v := range employees[m[u]].Subordinates {
            q = append(q, v)
        }
    }

    return
}