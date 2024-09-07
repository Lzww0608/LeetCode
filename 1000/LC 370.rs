impl Solution {
    pub fn get_modified_array(n: i32, updates: Vec<Vec<i32>>) -> Vec<i32> {
        let n = n as usize;
        let mut pre = vec![0; n + 1];
        let mut ans = vec![0; n + 1];

        for v in updates {
            pre[v[0] as usize] += v[2];
            pre[v[1] as usize + 1] -= v[2];
        }

        for i in 0..n {
            ans[i+1] = ans[i] + pre[i];
        }

        ans[1..].to_vec()
    }
}