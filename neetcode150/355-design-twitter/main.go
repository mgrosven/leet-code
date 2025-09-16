package main

import "container/heap"

func main() {

}

type Twitter struct {
	followerMap map[int]map[int]struct{}
	tweets      map[int][]tweet
	time        int
}

type tweet struct {
	id   int
	time int
}

type heapTweet struct {
	idx    int
	userId int
	tweet  tweet
}

type TweetMaxHeap struct {
	items []heapTweet
}

func (h *TweetMaxHeap) Len() int {
	return len(h.items)
}

func (h *TweetMaxHeap) Less(i, j int) bool {
	return h.items[i].tweet.time > h.items[j].tweet.time
}

func (h *TweetMaxHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *TweetMaxHeap) Push(x interface{}) {
	h.items = append(h.items, x.(heapTweet))
}

func (h *TweetMaxHeap) Pop() interface{} {
	last := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	return last
}

func Constructor() Twitter {
	return Twitter{
		followerMap: make(map[int]map[int]struct{}),
		tweets:      make(map[int][]tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	userTweets, ok := this.tweets[userId]
	if !ok {
		userTweets = make([]tweet, 0)
	}

	this.time++

	userTweets = append(userTweets, tweet{
		id:   tweetId,
		time: this.time,
	})
	this.tweets[userId] = userTweets
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	followees, ok := this.followerMap[userId]
	if !ok {
		followees = make(map[int]struct{})
	}

	// Need to follow self to get own tweets
	if _, ok := followees[userId]; !ok {
		followees[userId] = struct{}{}
	}

	tweetHeap := &TweetMaxHeap{items: make([]heapTweet, 0)}

	for followee := range followees {
		tweets := this.tweets[followee]
		if len(tweets) == 0 {
			continue
		}

		heap.Push(tweetHeap, heapTweet{
			tweet:  tweets[len(tweets)-1],
			idx:    len(tweets) - 1,
			userId: followee,
		})
	}

	var result []int
	for tweetHeap.Len() > 0 && len(result) < 10 {
		popped := heap.Pop(tweetHeap).(heapTweet)
		result = append(result, popped.tweet.id)

		nextIdx := popped.idx - 1
		if nextIdx >= 0 {
			tweets := this.tweets[popped.userId]

			heap.Push(tweetHeap, heapTweet{
				tweet:  tweets[nextIdx],
				idx:    nextIdx,
				userId: popped.userId,
			})
		}
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	followees, ok := this.followerMap[followerId]
	if !ok {
		followees = make(map[int]struct{})
		this.followerMap[followerId] = followees
	}

	followees[followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	followees, ok := this.followerMap[followerId]
	if !ok {
		return
	}
	delete(followees, followeeId)
}
