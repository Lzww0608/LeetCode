impl Solution {
    pub fn valid_path(n: i32, edges: Vec<Vec<i32>>, source: i32, destination: i32) -> bool {
        let n = n as usize;
        let mut ans = false;
        let mut g = vec![vec![]; n];
        let mut vis = vec![false; n];
        for edge in edges {
            let (u, v) = (edge[0], edge[1]);
            g[u as usize].push(v);
            g[v as usize].push(u);
        }

        fn dfs(g: &Vec<Vec<i32>>, vis: &mut Vec<bool>, u: i32, dest: i32, ans: &mut bool) {
            if *ans {
                return
            }
            if u == dest {
                *ans = true;
                return
            }
            
            for &v in &g[u as usize] {
                if !vis[v as usize] {
                    vis[v as usize] = true;
                    dfs(g, vis, v, dest, ans);
                }
            }
        } 

        vis[source as usize] = true;
        dfs(&g, &mut vis, source, destination, &mut ans);
        ans
    }
}