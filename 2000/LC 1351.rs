impl Solution {
    pub fn count_negatives(grid: Vec<Vec<i32>>) -> i32 {
        let mut ans: i32 = 0;
        for v in &grid {
            for &x in v {
                if x < 0 {
                    ans += 1;
                }
            }
        }

        ans
    }
}