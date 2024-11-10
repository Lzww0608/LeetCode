impl Solution {
    pub fn largest_combination(candidates: Vec<i32>) -> i32 {
        let mx = *candidates.iter().max().unwrap();
        let mut i = 1;
        let mut ans = 0;
        while i <= mx {
            let mut cur = 0;
            for x in &candidates {
                if x & i != 0 {
                    cur += 1;
                }
            }
            ans = ans.max(cur);
            i <<= 1;
        }

        ans
    }
}