impl Solution {
    pub fn maximum_bob_points(num_arrows: i32, alice_arrows: Vec<i32>) -> Vec<i32> {
        let n = alice_arrows.len();
        let mut max_score = 0;
        let mut ans = vec![0;n];

        for i in 1..(1<<n) {
            let mut tmp = vec![0;n];
            let mut score_Bob = 0;
            let mut cur = 0;
            for (j, &x) in alice_arrows.iter().enumerate() {
                if i & (1 << j) != 0 {
                    cur += x + 1;
                    tmp[j] += x + 1;
                    score_Bob += j as i32;
                }
            }

            if cur <= num_arrows && score_Bob > max_score {
                max_score = score_Bob;
                tmp[0] += num_arrows - cur;
                ans.copy_from_slice(&tmp);
            }
        }

        ans
    }
}