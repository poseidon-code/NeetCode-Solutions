// 355: Design Twitter
// https://leetcode.com/problems/design-twitter

package main

import (
	"container/heap"
	"fmt"
)

type Tweet struct {
	count       int
	tweet_id    int
	followee_id int
	index       int
}

type Heap []*Tweet

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i].count < h[j].count }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// SOLUTION
type Twitter struct {
	count      int
	tweet_map  map[int][]*Tweet
	follow_map map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{
		tweet_map:  make(map[int][]*Tweet),
		follow_map: make(map[int]map[int]bool),
	}
}

func (this *Twitter) PostTweet(user_id int, tweet_id int) {
	if _, ok := this.tweet_map[user_id]; !ok {
		this.tweet_map[user_id] = make([]*Tweet, 0)
	}
	this.tweet_map[user_id] = append(this.tweet_map[user_id], &Tweet{count: this.count, tweet_id: tweet_id})
	this.count -= 1
}

func (this *Twitter) GetNewsFeed(user_id int) []int {
	result := make([]int, 0)
	min_heap := Heap{}

	if _, ok := this.follow_map[user_id]; !ok {
		this.follow_map[user_id] = make(map[int]bool)
	}
	this.follow_map[user_id][user_id] = true

	for followee_id, _ := range this.follow_map[user_id] {
		if _, ok := this.tweet_map[followee_id]; ok {
			index := len(this.tweet_map[followee_id]) - 1
			tweet := this.tweet_map[followee_id][index]

			heap.Push(
				&min_heap,
				&Tweet{
					count:       tweet.count,
					tweet_id:    tweet.tweet_id,
					followee_id: followee_id,
					index:       index - 1,
				},
			)
		}
	}

	for len(min_heap) > 0 && len(result) < 10 {
		tweet := heap.Pop(&min_heap).(*Tweet)
		result = append(result, tweet.tweet_id)

		if tweet.index >= 0 {
			next_tweet := this.tweet_map[tweet.followee_id][tweet.index]

			heap.Push(
				&min_heap,
				&Tweet{
					count:       next_tweet.count,
					tweet_id:    next_tweet.tweet_id,
					followee_id: tweet.followee_id,
					index:       tweet.index - 1,
				},
			)
		}
	}

	return result
}

func (this *Twitter) Follow(follower_id int, followee_id int) {
	if _, ok := this.follow_map[follower_id]; !ok {
		this.follow_map[follower_id] = make(map[int]bool)
	}
	this.follow_map[follower_id][followee_id] = true
}

func (this *Twitter) Unfollow(follower_id int, followee_id int) {
	if _, ok := this.follow_map[follower_id]; ok {
		delete(this.follow_map[follower_id], followee_id)
	}
}

func main() {
	o := Constructor()

	// OPERATIONS
	o.PostTweet(1, 5)
	fmt.Println(o.GetNewsFeed(1))
	o.Follow(1, 2)
	o.PostTweet(2, 6)
	fmt.Println(o.GetNewsFeed(1))
	o.Unfollow(1, 2)
	fmt.Println(o.GetNewsFeed(1))
}
