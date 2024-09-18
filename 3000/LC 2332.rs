impl Solution {
    pub fn latest_time_catch_the_bus(mut buses: Vec<i32>, mut passengers: Vec<i32>, capacity: i32) -> i32 {
        let m = buses.len();
        let n = passengers.len();
        buses.sort_unstable();
        passengers.sort_unstable();

        let mut j = 0;
        let mut cnt = 0;
        for i in 0..m {
            cnt = capacity;
            while cnt > 0 && j < n && passengers[j] <= buses[i] {
                j += 1;
                cnt -= 1;
            }
        }

        j -= 1;
        let mut ans = if cnt > 0 {
            *buses.last().unwrap()
        } else {
            passengers[j]
        };
        
        while j < n && ans == passengers[j] {
            j -= 1;
            ans -= 1;
        }

        ans
    }
}