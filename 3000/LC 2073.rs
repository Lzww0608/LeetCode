impl Solution {
    pub fn time_required_to_buy(tickets: Vec<i32>, k: i32) -> i32 {
        let k = k as usize;
        let t = tickets[k];

        tickets.iter()
            .enumerate()
            .map(|(i, &x)| x.min(if i <= k {t} else {t - 1}))
            .sum()
    }
}