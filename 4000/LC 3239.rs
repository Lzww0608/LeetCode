impl Solution {
    pub fn min_flips(grid: Vec<Vec<i32>>) -> i32 {
        let mut a = 0;
        let mut b = 0;
        let m = grid.len();
        let n = grid[0].len();

        for i in 0..m {
            let mut j = 0;
            let mut k = n - 1;
            while j < k {
                if grid[i][j] != grid[i][k] {
                    a += 1;
                }
                j += 1;
                k -= 1;
            }
        }

        for j in 0..n {
            let mut i = 0;
            let mut k = m - 1;
            while i < k {
                if grid[i][j] != grid[k][j] {
                    b += 1;
                }
                i += 1;
                k -= 1;
            }
        }

        a.min(b)
    }
}