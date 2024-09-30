use std::cmp;

impl Solution {
    pub fn count_different_subsequence_gc_ds(nums: Vec<i32>) -> i32 {
        let mx = *nums.iter().max().unwrap();
        let mut vis = vec![false; (mx + 1) as usize];
        for &x in &nums {
            vis[x as usize] = true;
        }

        let mut ans = 0;
        for i in 1..=mx {
            let mut t = 0;
            let mut j = i;
            while j <= mx {
                if vis[j as usize] {
                    t = Solution::gcd(j, t);
                    if t == i {
                        break
                    }
                }
                j += i;
            }
            if t == i {
                ans += 1;
            }
        }

        ans
    }

    pub fn gcd(mut x: i32, mut y: i32) -> i32 {
        while y != 0 {
            let tmp = y;
            y = x % y;
            x = tmp;
        }

        x
    }
}