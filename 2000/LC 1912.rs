use std::collections::{HashMap, HashSet, BTreeSet};

type Pii = (i32, i32);
type Tii = (i32, i32, i32);

struct MovieRentingSystem {
    shop_to_movie: HashMap<i32, HashMap<i32, i32>>,
    un_rented: HashMap<i32, BTreeSet<Pii>>,
    rented: BTreeSet<Tii>,
}

impl MovieRentingSystem {
    fn new(n: i32, entries: Vec<Vec<i32>>) -> Self {
        let mut shop_to_movie = HashMap::new();
        let mut un_rented = HashMap::new();

        for entry in entries {
            let shop = entry[0];
            let movie = entry[1];
            let price = entry[2];

            shop_to_movie
                .entry(shop)
                .or_insert_with(HashMap::new)
                .insert(movie, price);

            un_rented
                .entry(movie)
                .or_insert_with(BTreeSet::new)
                .insert((price, shop));
        }

        MovieRentingSystem {
            shop_to_movie,
            un_rented,
            rented: BTreeSet::new(),
        }
    }

    fn search(&self, movie: i32) -> Vec<i32> {
        let mut ans = Vec::new();
        if let Some(shops) = self.un_rented.get(&movie) {
            for &(price, shop) in shops.iter().take(5) {
                ans.push(shop);
            }
        }
        ans
    }

    fn rent(&mut self, shop: i32, movie: i32) {
        if let Some(price) = self.shop_to_movie.get(&shop).and_then(|m| m.get(&movie)) {
            if let Some(shops) = self.un_rented.get_mut(&movie) {
                shops.remove(&(price.clone(), shop));
            }
            self.rented.insert((*price, shop, movie));
        }
    }

    fn drop(&mut self, shop: i32, movie: i32) {
        if let Some(price) = self.shop_to_movie.get(&shop).and_then(|m| m.get(&movie)) {
            if let Some(shops) = self.un_rented.get_mut(&movie) {
                shops.insert((*price, shop));
            }
            self.rented.remove(&(*price, shop, movie));
        }
    }

    fn report(&self) -> Vec<Vec<i32>> {
        let mut ans = Vec::new();
        for &(_, shop, movie) in self.rented.iter().take(5) {
            ans.push(vec![shop, movie]);
        }
        ans
    }
}

