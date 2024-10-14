const MAX: usize = 5005;
impl Solution {
    pub fn divide_players(skill: Vec<i32>) -> i64 {
        let mut tot = 0i64;
        let n = skill.len() as i64 / 2;
        let mut cnt = vec![0; MAX];
        for &x in &skill {
            tot += x as i64;
            cnt[x as usize] += 1;
        }
        
        if tot % n != 0 {
            return -1;
        }
        let sum = tot / n;

        let mut ans = 0;
        for i in 1..MAX {
            let x = cnt[i];
            if i as i64 >= sum {
                break
            }
            if x > 0 {
                if cnt[sum as usize - i] != x {
                    return -1;
                } else {
                    ans += x * (sum - i as i64) * i as i64;
                }
            } 
        }

        ans / 2 
    }
}