impl Solution {
    pub fn maximize_sweetness(sweetness: Vec<i32>, k: i32) -> i32 {
        let mut sum = 0;
        let mut mn = i32::MAX;

        for &x in &sweetness {
            sum += x;
            mn = mn.min(x);
        }

        let check = |target: i32| -> bool {
            let mut cnt = 0;
            let mut cur = 0;

            for &x in &sweetness {
                cur += x;
                if cur >= target {
                    cnt += 1;
                    cur = 0;
                }
            }

            cnt > k
        };

        let mut l = mn;
        let mut ans = 0;
        let mut r = sum / (k + 1) + 1;

        while l < r {
            let mid = l + ((r - l) >> 1);
            if check(mid) {
                ans = mid;
                l = mid + 1;
            } else {
                r = mid;
            }
        }

        ans
    }
}