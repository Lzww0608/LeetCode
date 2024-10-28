use std::collections::HashMap;

impl Solution {
    pub fn count_distinct_integers(nums: Vec<i32>) -> i32 {
        let mut m: HashMap<i32, bool> = HashMap::new();
        for &x in &nums {
            m.insert(x, true);
            let mut x = x;
            let mut cur = 0;
            while x > 0 {
                cur = cur * 10 + x % 10;
                x /= 10;
            }
            m.insert(cur, true);
        }

        m.len() as i32
    }
}