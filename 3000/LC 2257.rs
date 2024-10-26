use std::collections::HashMap;

impl Solution {
    pub fn count_unguarded(m: i32, n: i32, guards: Vec<Vec<i32>>, walls: Vec<Vec<i32>>) -> i32 {
        let dirs = vec![vec![1, 0], vec![-1, 0], vec![0, 1], vec![0, -1]];
        let mut vis = vec![false; (m * n) as usize];
        let mut w = HashMap::new();
        let mut ans = m * n - walls.len() as i32 - guards.len() as i32;

        for v in &guards {
            w.insert(v[0] * n + v[1], true);
        }
        for v in &walls {
            w.insert(v[0] * n + v[1], true);
        }

        for v in &guards {
            let (i, j) = (v[0], v[1]);
            for dir in &dirs {
                let (mut x, mut y) = (i + dir[0], j + dir[1]);
                while x >= 0 && x < m && y >= 0 && y < n && !w.contains_key(&(x * n + y)) {
                    if !vis[(x * n + y) as usize] {
                        vis[(x * n + y) as usize] = true;
                        ans -= 1;
                    }
                    x += dir[0];
                    y += dir[1];
                }
            }
        }

        ans
        
    }
}