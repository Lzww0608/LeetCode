use std::cmp::Ordering;

impl Solution {
    pub fn earliest_acq(mut logs: Vec<Vec<i32>>, n: i32) -> i32 {
        let mut fa: Vec<i32> = (0..n).collect();
        let mut cnt = 1;

        fn find(fa: &mut Vec<i32>, x: i32) -> i32 {
            if fa[x as usize] != x {
                fa[x as usize] = find(fa, fa[x as usize]);
            }

            fa[x as usize]
        }

        fn merge(fa: &mut Vec<i32>, cnt: &mut i32, x: i32, y: i32) {
            let rx = find(fa, x) as usize;
            let ry = find(fa, y) as usize;
            if rx != ry {
                fa[rx] = ry as i32;
                *cnt += 1;
            }
        }

        logs.sort_unstable_by(|a, b| a[0].cmp(&b[0]));
        
        for v in &logs {
            merge(&mut fa, &mut cnt, v[1], v[2]);
            if cnt == n {
                return v[0];
            }
        }

        -1
    }
}