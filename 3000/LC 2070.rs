use std::cmp::max;

impl Solution {
    pub fn maximum_beauty(mut items: Vec<Vec<i32>>, queries: Vec<i32>) -> Vec<i32> {
        let n = queries.len();
        let mut ans = vec![0; n];

        items.sort_by(|a, b| a[0].cmp(&b[0]));

        let mut id: Vec<usize> = (0..n).collect();

        id.sort_by(|&a, &b| queries[a].cmp(&queries[b]));

        let mut j = 0;
        let mut cur = 0;
        let m = items.len();
        for &i in &id {
            while j < m && items[j][0] <= queries[i] {
                cur = max(cur, items[j][1]);
                j += 1;
            }
            ans[i] = cur;
        }

        ans
    }
}