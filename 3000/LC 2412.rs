use std::cmp::{min, max};

impl Solution {
    pub fn minimum_money(transactions: Vec<Vec<i32>>) -> i64 {
        let (mut cur, mut mx) = (0, 0);

        for v in transactions {
            cur += max(v[0] - v[1], 0) as i64;
            mx = max(min(v[0], v[1]) as i64, mx);
        }

        cur + mx
    }
}