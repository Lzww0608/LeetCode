impl Solution {
    pub fn satisfies_conditions(grid: Vec<Vec<i32>>) -> bool {
        let m = grid.len();
        let n = grid[0].len();
        for i in 0..m {
            for j in 0..n {
                if i + 1 < m && grid[i][j] != grid[i+1][j] {
                    return false;
                }

                if j + 1 < n && grid[i][j] == grid[i][j+1] {
                    return false;
                }
            }
        }

        true
    }
}