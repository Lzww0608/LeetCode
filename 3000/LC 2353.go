
import (
    "container/heap"

)

type FoodRatings struct {
    food2Rating map[string]int
    food2Cuisine map[string]string
    cuisineHeaps map[string]*FoodHeap
}

type Food struct {
    name   string
    rating int
}

type FoodHeap []Food

func (fh FoodHeap) Len() int { return len(fh) }
func (fh FoodHeap) Less(i, j int) bool {
    if fh[i].rating == fh[j].rating {
        return fh[i].name < fh[j].name
    }
    return fh[i].rating > fh[j].rating
}
func (fh FoodHeap) Swap(i, j int) { fh[i], fh[j] = fh[j], fh[i] }

func (fh *FoodHeap) Push(x interface{}) {
    *fh = append(*fh, x.(Food))
}

func (fh *FoodHeap) Pop() interface{} {
    old := *fh
    n := len(old)
    item := old[n-1]
    *fh = old[0 : n-1]
    return item
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    n := len(foods)
    food2Rating := make(map[string]int)
    food2Cuisine := make(map[string]string)
    cuisineHeaps := make(map[string]*FoodHeap)

    for i := 0; i < n; i++ {
        food2Rating[foods[i]] = ratings[i]
        food2Cuisine[foods[i]] = cuisines[i]
        if _, exists := cuisineHeaps[cuisines[i]]; !exists {
            cuisineHeaps[cuisines[i]] = &FoodHeap{}
            heap.Init(cuisineHeaps[cuisines[i]])
        }
        heap.Push(cuisineHeaps[cuisines[i]], Food{name: foods[i], rating: ratings[i]})
    }

    return FoodRatings{food2Rating, food2Cuisine, cuisineHeaps}
}

func (c *FoodRatings) ChangeRating(food string, newRating int) {
    cuisine := c.food2Cuisine[food]
    c.food2Rating[food] = newRating

    // Push the updated food back into the heap
    heap.Push(c.cuisineHeaps[cuisine], Food{name: food, rating: newRating})
}

func (c *FoodRatings) HighestRated(cuisine string) string {
    // Remove invalid top elements until the top is correct
    for {
        top := (*c.cuisineHeaps[cuisine])[0]
        if top.rating == c.food2Rating[top.name] {
            return top.name
        }
        heap.Pop(c.cuisineHeaps[cuisine])
    }
}
