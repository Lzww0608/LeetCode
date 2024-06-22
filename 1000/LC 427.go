/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    n := len(grid)
    pre := make([][]int, n + 1)
    pre[0] = make([]int, n + 1)
    for i, row := range grid {
        pre[i+1] = make([]int, n + 1)
        for j, v := range row {
            pre[i+1][j+1] = -pre[i][j] + pre[i+1][j] + pre[i][j+1] + v
        }
    }

    var dfs func(r1, c1, r2, c2 int) * Node 
    dfs = func(r1, c1, r2, c2 int) * Node {
        sum := pre[r2][c2] + pre[r1][c1] - pre[r1][c2] - pre[r2][c1]
        if sum == 0 {
            return &Node{Val: false, IsLeaf: true}
        } else if sum == (r2 - r1) * (c2 - c1) {
            return &Node{Val: true, IsLeaf: true}
        }
        return &Node{
            Val: true,
            IsLeaf: false,
            TopLeft: dfs(r1, c1, (r1 + r2) / 2, (c1 +c2) / 2),
            TopRight: dfs(r1, (c1 +c2) / 2, (r1 + r2) / 2, c2),
            BottomLeft: dfs((r1 + r2) / 2, c1, r2, (c1 +c2) / 2),
            BottomRight: dfs((r1 + r2) / 2, (c1 +c2) / 2, r2, c2),
        }
    }
    return dfs(0, 0, n, n);
}
