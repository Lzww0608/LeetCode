use std::cmp::{min, max};

const MAX: usize = 5000;
impl Solution {
    pub fn minimize_the_difference(mat: Vec<Vec<i32>>, target: i32) -> i32 {
        let mut f = vec![false; MAX];
        
        let (mut mx, mut mn) = (0, 0);
        f[0] = true;

        for i in 0..mat.len() {
            let cur_mx = *mat[i].iter().max().unwrap() as usize;
            let cur_mn = *mat[i].iter().min().unwrap() as usize;
            mn = mn + cur_mn;
            mx = min(cur_mx + mx, target as usize * 2);

            for j in ((mn-cur_mn) as usize..=mx).rev() {
                f[j] = false;
                for &x in &mat[i] {
                    let x = x as usize;
                    if j >= x && f[(j-x) as usize] {
                        f[j] = true;
                        break;
                    }
                }
            }
        }

        let mut ans = (target - mn as i32).abs();
        for (i, &x) in f.iter().enumerate() {
            if x {
                ans = min(ans, (i as i32 - target).abs());
            }
        }

        ans
    }
}