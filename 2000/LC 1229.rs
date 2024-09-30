use std::cmp::{max, min};

impl Solution {
    pub fn min_available_duration(mut a: Vec<Vec<i32>>, mut b: Vec<Vec<i32>>, duration: i32) -> Vec<i32> {
        let m = a.len();
        let n = b.len();
        let (mut i, mut j) = (0, 0);

        a.sort_by(|x, y| x[0].cmp(&y[0]));
        b.sort_by(|x, y| x[0].cmp(&y[0]));

        while i < m && j < n {
            let l_a = a[i][0];
            let r_a = a[i][1];
            let l_b = b[j][0];
            let r_b = b[j][1];

            if min(r_a, r_b) - max(l_a, l_b) >= duration {
                return vec![max(l_a, l_b), max(l_a, l_b) + duration];
            }

            if r_b > r_a {
                i += 1;
            } else {
                j += 1;
            }
        }

        vec![]
    }
}