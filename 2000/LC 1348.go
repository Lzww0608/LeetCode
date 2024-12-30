type TweetCounts struct {
    m map[string][]int
}


func Constructor() TweetCounts {
    return TweetCounts {
        m: make(map[string][]int),
    }
}


func (c *TweetCounts) RecordTweet(tweetName string, time int)  {
    c.m[tweetName] = append(c.m[tweetName], time)
}


func (c *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
    endTime++
    d := 60 * 60 * 24
    if freq == "minute" {
        d = 60
    } else if freq == "hour" {
        d = 60 * 60
    }

    ans := make([]int, (endTime - startTime + d - 1) / d)
    for _, x := range c.m[tweetName] {
        if x >= startTime && x < endTime {
            ans[(x - startTime) / d]++
        }
        
    }

    return ans
}


/**
 * Your TweetCounts object will be instantiated and called as such:
 * obj := Constructor();
 * obj.RecordTweet(tweetName,time);
 * param_2 := obj.GetTweetCountsPerFrequency(freq,tweetName,startTime,endTime);
 */