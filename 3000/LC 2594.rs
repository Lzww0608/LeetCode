impl Solution {
    pub fn repair_cars(ranks: Vec<i32>, cars: i32) -> i64 {
        let (mut l, mut r) = (1 as i64, 100 * cars as i64 * cars as i64);

        while l < r {
            let mid = l + ((r - l) >> 1);
            let mut sum: i64 = 0;
            
            for &x in &ranks {
                let t = mid as f64 / x as f64;
                sum += t.sqrt() as i64;
            }

            if sum >= cars as i64 {
                r = mid;
            } else {
                l = mid + 1;
            }
        }

        l
    }
}