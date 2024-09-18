use std::collections::HashMap;
impl Solution {
    pub fn find_strobogrammatic(n: i32) -> Vec<String> {
        let n = n as usize;
        let m = vec![
            ('9', '6'),
            ('6', '9'),
            ('1', '1'),
            ('0', '0'),
            ('8', '8'),
        ].into_iter().collect::<std::collections::HashMap<char, char>>();

        let mut ans = Vec::new();
        let mut s = vec![' '; n];

        fn dfs(cnt: usize, n: usize, s: &mut Vec<char>, m: &HashMap<char, char>, ans: &mut Vec<String>) {
            if cnt * 2 == n {
                ans.push(s.iter().collect());
                return;
            } else if cnt * 2 + 1 == n {
                s[n/2] = '1';
                ans.push(s.iter().collect());
                s[n/2] = '8';
                ans.push(s.iter().collect());
                s[n/2] = '0';
                ans.push(s.iter().collect());
                return;
            }

            for (&k, &v) in m.iter() {
                if cnt == 0 && k == '0' {
                    continue;
                }

                s[cnt] = k;
                s[n-cnt-1] = v;
                dfs(cnt+1, n, s, m, ans);
            }
        }
        dfs(0, n, &mut s, &m, &mut ans);

        ans
    }
    
}