impl Solution {
    pub fn sum_of_distances_in_tree(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        let n = n as usize;
        let mut g = vec![vec![]; n];
        for e in &edges {
            let a = e[0] as usize;
            let b = e[1] as usize;
            g[a].push(b);
            g[b].push(a);
        }

        let mut ans = vec![0; n];
        let mut f = vec![0; n];

        fn dfs1(v: usize, fa: isize, depth: i32, g: &Vec<Vec<usize>>, ans: &mut Vec<i32>, f: &mut Vec<i32>) {
            ans[0] += depth;
            f[v] = 1;
            for &u in &g[v] {
                if u as isize != fa {
                    dfs1(u, v as isize, depth + 1, g, ans, f);
                    f[v] += f[u];
                }
            }
        }

        fn dfs2(v: usize, fa: isize, n: usize, g: &Vec<Vec<usize>>, ans: &mut Vec<i32>, f: &mut Vec<i32>) {

            for &u in &g[v] {
                if u as isize != fa {
                    ans[u] = ans[v] + n as i32 - 2 * f[u];
                    dfs2(u, v as isize, n, g, ans, f);
                }
            }
        }

        dfs1(0, -1, 0, &g, &mut ans, &mut f);
        dfs2(0, -1, n, &g, &mut ans, &mut f);

        ans
    }
}