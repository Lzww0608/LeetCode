use std::cmp::{min, max};
use std::collections::VecDeque;
impl Solution {
    pub fn min_score(n: i32, roads: Vec<Vec<i32>>) -> i32 {
        let mut ans = i32::MAX;

        let mut g = vec![Vec::new(); n as usize];
        for e in roads.iter() {
            let (u, v, w) = (e[0] as usize - 1, e[1] as usize - 1, e[2]);
            g[u].push((v, w));
            g[v].push((u, w));
        }

        let mut vis = vec![false; n as usize];

        let mut q = VecDeque::new();
        q.push_back(0);
        vis[0] = true;

        while let Some(u) = q.pop_front() {
            for &(v, w) in g[u].iter() {
                ans = min(ans, w);
                if !vis[v] {
                    vis[v] = true;
                    q.push_back(v);
                }
            }
        }

        ans
    }
}