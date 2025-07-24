type Tweet struct {
    tid, time int
}

type User struct {
    uid int 
    tweet []*Tweet
    followee map[int]bool
}

type Twitter struct {
    m map[int]*User
    time int
}


func Constructor() Twitter {
    return Twitter{map[int]*User{},  0}
}


func (c *Twitter) PostTweet(userId int, tweetId int)  {
    if _, ok := c.m[userId]; !ok {
        c.m[userId] = &User{}
        c.m[userId].followee = make(map[int]bool)
    }

    c.m[userId].uid = userId
    c.m[userId].tweet = append(c.m[userId].tweet, &Tweet{tweetId, c.time})
    c.time++
}


func (c *Twitter) GetNewsFeed(userId int) (ans []int) {
    if _, ok := c.m[userId]; !ok {
        c.m[userId] = &User{}
        c.m[userId].followee = make(map[int]bool)
    }
    n := len(c.m[userId].followee) + 1
    a := make([]*Tweet, 0, n * 10)
    for _, v := range c.m[userId].tweet {
        a = append(a, v)
    }
    for k := range c.m[userId].followee {
        for _, v := range c.m[k].tweet {
            a = append(a, v)
        }
    }
    sort.Slice(a, func(i, j int) bool {
        return a[i].time > a[j].time
    })
    
    for i := 0; i < min(10, len(a)); i++ {
        ans = append(ans, a[i].tid)
    }

    return 
}


func (c *Twitter) Follow(followerId int, followeeId int)  {
    if _, ok := c.m[followerId]; !ok {
        c.m[followerId] = &User{}
        c.m[followerId].followee = make(map[int]bool)
    }
    if _, ok := c.m[followeeId]; !ok {
        c.m[followeeId] = &User{}
        c.m[followeeId].followee = make(map[int]bool)
    }
    c.m[followerId].followee[followeeId] = true
}


func (c *Twitter) Unfollow(followerId int, followeeId int)  {
    if _, ok := c.m[followerId]; !ok {
        c.m[followerId] = &User{}
        c.m[followerId].followee = make(map[int]bool)
    }
    delete(c.m[followerId].followee, followeeId)
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */