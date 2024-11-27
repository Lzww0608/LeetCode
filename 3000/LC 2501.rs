use std::collections::HashSet;

impl Solution {
    pub fn longest_square_streak(nums: Vec<i32>) -> i32 {
        let mut ans = -1;

        let mut m: HashSet<i64> = HashSet::new();
        for &x in &nums {
            m.insert(x as i64);
        }

        for &x in &nums {
            let mut cnt = 1;
            let mut cur = x as i64 * x as i64;
            while m.contains(&cur) {
                cnt += 1;
                cur *= cur;
            }

            if cnt > 1 && ans < cnt {
                ans = cnt;
            }
        }

        ans
    }
}