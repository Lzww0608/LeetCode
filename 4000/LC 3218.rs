impl Solution {
    pub fn minimum_cost(m: i32, n: i32, mut horizontal_cut: Vec<i32>, mut vertical_cut: Vec<i32>) -> i32 {
        horizontal_cut.sort_unstable_by(|a, b| b.cmp(a));
        vertical_cut.sort_unstable_by(|a, b| b.cmp(a));  
        let (mut i, mut j) = (0, 0);
        let mut ans = 0;
        let m = m as usize;
        let n = n as usize;

        while i < m - 1 || j < n - 1 {
            if j >= n - 1 || i < m - 1 && horizontal_cut[i] >= vertical_cut[j] {
                ans += horizontal_cut[i] * (j + 1) as i32;
                i += 1;
            } else {
                ans += vertical_cut[j] * (i + 1) as i32;
                j += 1;
            }
        }       

        ans
    }
}