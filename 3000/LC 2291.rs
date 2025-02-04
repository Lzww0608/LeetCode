impl Solution {
    pub fn maximum_profit(present: Vec<i32>, future: Vec<i32>, budget: i32) -> i32 {
        let budget = budget as usize;
        let mut f = vec![0; budget + 1];
        for i in 0..present.len() {
            let x = present[i] as usize;
            for j in (x..=budget).rev() {
                f[j] = f[j].max(f[j-x] + future[i] - present[i]);
            }
        }

        f[budget]
    }
}