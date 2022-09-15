package main

import (
	"sort"
)

type Twitter struct {
	follow     map[int]map[int]int
	userToPost map[int][]int
	posts      map[int]int
}

func Constructor() Twitter {
	return Twitter{make(map[int]map[int]int), make(map[int][]int), make(map[int]int)}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.userToPost[userId] = append(t.userToPost[userId], tweetId)
	t.posts[tweetId] = len(t.posts)
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	var list []int
	list = append(list, t.userToPost[userId]...)
	for k := range t.follow[userId] {
		list = append(list, t.userToPost[k]...)
	}
	sort.Slice(list, func(i, j int) bool {
		return t.posts[list[i]] > t.posts[list[j]]
	})
	l := len(list)
	l = min(l, 10)

	return list[0:l]
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if t.follow[followerId] == nil {
		t.follow[followerId] = make(map[int]int)
	}
	m := t.follow[followerId]
	if _, ok := m[followeeId]; !ok {
		m[followeeId] = len(m)
		t.follow[followerId] = m
	}
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	m := t.follow[followerId]
	delete(m, followeeId)
}
func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
func main() {
	t := Constructor()
	t.PostTweet(1, 5)
	t.GetNewsFeed(1)
	t.Follow(1, 2)
	t.PostTweet(2, 6)
	t.GetNewsFeed(1)
	t.Unfollow(1, 2)
	t.GetNewsFeed(1)
}
