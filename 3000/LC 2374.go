impl Solution {
    pub fn edge_score(edges: Vec<i32>) -> i32 {
        let mut ans = 0;
        let n = edges.len();
        let mut score = vec![0i64; n];

        for (i, &x) in edges.iter().enumerate() {
            let x = x as usize;
            score[x] += i as i64;
            if score[x] > score[ans] || (score[x] == score[ans] && x < ans) {
                ans = x;
            }
        }

        ans as i32
    }
}