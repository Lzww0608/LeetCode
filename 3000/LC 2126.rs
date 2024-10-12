impl Solution {
    pub fn asteroids_destroyed(mass: i32, asteroids: Vec<i32>) -> bool {
        let mut mn = vec![-1; 17];
        let mut sum = vec![0i64; 17];

        for i in 0..asteroids.len() {
            let h = 31 - (asteroids[i].leading_zeros() as usize);
            if mn[h] == -1 || asteroids[i] < mn[h] {
                mn[h] = asteroids[i];
            }
            sum[h] += asteroids[i] as i64;
        }

        let mut cur = mass as i64;
        for i in 0..17 {
            if cur < mn[i] as i64 {
                return false;
            }
            cur += sum[i];
        }

        true
    }
}