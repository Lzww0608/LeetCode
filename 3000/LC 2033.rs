
impl Solution {
    pub fn min_operations(grid: Vec<Vec<i32>>, x: i32) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut a: Vec<i32> = Vec::with_capacity(m * n);
        for i in 0..m {
            for j in 0..n {
                if (grid[i][j] - grid[0][0]).abs() % x != 0 {
                    return -1;
                }
                a.push(grid[i][j]);
            }
        }

        a.sort();
        let mut ans = 0;
        let mid = (m * n) / 2;
        for &k in &a {
            ans += (k - a[mid]).abs() / x;
        }

        ans
    }
}