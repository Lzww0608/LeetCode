use std::cmp::{min, max};
impl Solution {
    pub fn grid_game(grid: Vec<Vec<i32>>) -> i64 {
        let n = grid[0].len();
        let mut pre1 = vec![0i64; n+1];
        let mut pre2 = vec![0i64; n+1];

        for i in 0..n {
            pre1[i+1] = pre1[i] + grid[0][i] as i64;
            pre2[i+1] = pre2[i] + grid[1][i] as i64;
        }

        let mut ans = i64::MAX;

        for i in 0..n {
            ans = min(ans, max(pre2[i], pre1[n] - pre1[i+1]));
        }

        ans
    }
}