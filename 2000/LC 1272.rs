use std::cmp;

impl Solution {
    pub fn remove_interval(intervals: Vec<Vec<i32>>, to_be_removed: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ans: Vec<Vec<i32>> = Vec::new();
        let (i, j) = (to_be_removed[0], to_be_removed[1]);

        for v in &intervals {
            let (l, r) = (v[0], v[1]);

            if l >= j || r <= i {
                ans.push(vec![l, r]);
            } else {
                if l < i {
                    if r > j {
                        ans.push(vec![l, i]);
                        ans.push(vec![j, r]);
                    } else {
                        ans.push(vec![l, i]);
                    }
                } else if r > j {
                    ans.push(vec![j, r]);
                }
            }
        }

        ans.sort_by(|a, b| a[0].cmp(&b[0]));

        ans
    }
}